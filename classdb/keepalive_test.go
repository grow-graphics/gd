package classdb

import (
	"context"
	"net/http"
	"reflect"
	"strconv"
	"sync"
	"testing"
	"time"

	"graphics.gd/classdb/MeshInstance3D"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node3D"
)

// TestKeepAlive — existing baseline: unexported Instance inside a nested
// non-extension struct should be picked up.
func TestKeepAlive(t *testing.T) {
	type Child struct {
		node Node.Instance
	}
	type TestStruct struct {
		child Child
	}
	fn := compile_keepalive(reflect.TypeFor[TestStruct]())
	if fn == nil {
		t.Fatal("not detected")
	}
	fn(reflect.ValueOf(&TestStruct{}).Elem())
}

// TestKeepAlive_ExtensionUnexportedInstance — matches CitizenEditor.scene:
// an unexported Node3D.Instance field directly on an extension class.
//
// The walker has an early `continue` for fields that implement Node.Any and
// are exported on extension classes (auto-bound scene nodes). The condition
// requires IsExported() == true, so unexported fields should NOT be skipped
// and must be walked.
func TestKeepAlive_ExtensionUnexportedInstance(t *testing.T) {
	type CitizenLike struct {
		Node3D.Extension[CitizenLike]
		scene Node3D.Instance
	}
	fn := compile_keepalive(reflect.TypeFor[CitizenLike]())
	if fn == nil {
		t.Fatal("compile_keepalive returned nil for *CitizenLike — the unexported Node3D.Instance field on an extension class was not picked up as a keepalive root")
	}
	fn(reflect.ValueOf(&CitizenLike{}).Elem())
}

// TestKeepAlive_PointerToStructWithInstance — matches CitizenEditor.body
// (*CitizenBody) where CitizenBody contains an unexported Instance field.
// The walker should: pointer-deref → struct fields → Instance match.
func TestKeepAlive_PointerToStructWithInstance(t *testing.T) {
	type Inner struct {
		mesh MeshInstance3D.Instance
	}
	type Outer struct {
		Node3D.Extension[Outer]
		body *Inner
	}
	fn := compile_keepalive(reflect.TypeFor[Outer]())
	if fn == nil {
		t.Fatal("compile_keepalive returned nil — the *Inner.mesh chain wasn't walked")
	}
	fn(reflect.ValueOf(&Outer{}).Elem())
}

// TestKeepAlive_FullCitizenShape — the exact nesting the citizen editor
// uses: extension class with both unexported Instance and unexported
// pointer-to-struct-with-Instance fields.
func TestKeepAlive_FullCitizenShape(t *testing.T) {
	type Body struct {
		mesh MeshInstance3D.Instance
	}
	type Editor struct {
		Node3D.Extension[Editor]
		scene Node3D.Instance
		body  *Body
	}
	fn := compile_keepalive(reflect.TypeFor[Editor]())
	if fn == nil {
		t.Fatal("compile_keepalive returned nil for the CitizenEditor-shaped struct")
	}
	// Smoke-call the closure on a zero value — must not panic on nil
	// pointers (body is nil here).
	val := reflect.New(reflect.TypeFor[Editor]()).Elem()
	fn(val)
}

// TestKeepAlive_ValueTypedBody — matches the current citizen editor
// after the user reverted body to a value type instead of a pointer.
// Verifies compile_keepalive still threads through the nested struct.
func TestKeepAlive_ValueTypedBody(t *testing.T) {
	type Body struct {
		mesh MeshInstance3D.Instance
	}
	type Editor struct {
		Node3D.Extension[Editor]
		scene Node3D.Instance
		body  Body
	}
	fn := compile_keepalive(reflect.TypeFor[Editor]())
	if fn == nil {
		t.Fatal("compile_keepalive returned nil for editor with value-typed body field")
	}
	val := reflect.New(reflect.TypeFor[Editor]()).Elem()
	fn(val)
}

// TestKeepAlive_MapOfPointers: a map walk hands each entry over in one
// reused scratch value, so the pointer walker must dedupe on the target
// rather than on the value it is given, or only the first entry is kept
// alive and handles held by the rest go stale after a frame.
func TestKeepAlive_MapOfPointers(t *testing.T) {
	type Inner struct {
		held map[int]Node.Instance // nil: walked without touching the engine
	}
	type Outer struct {
		entries map[int]*Inner
	}
	fn := compile_keepalive(reflect.TypeFor[Outer]())
	if fn == nil {
		t.Fatal("compile_keepalive returned nil for a map of pointers to structs with Instances")
	}
	outer := Outer{entries: map[int]*Inner{1: {}, 2: {}, 3: {}}}
	clear(skips)
	fn(reflect.ValueOf(&outer).Elem())
	walked := 0
	for val := range skips {
		if val.Type() == reflect.TypeFor[Inner]() {
			walked++
		}
	}
	if walked != len(outer.entries) {
		t.Fatalf("walked %d of %d map entries", walked, len(outer.entries))
	}
}

// TestKeepAlive_StandardLibrary: values of the standard library types that
// other goroutines change under their own locks are not walked. Walking an
// extension's context.Context iterated cancelCtx.children while an HTTP
// goroutine added and removed child contexts, a fatal "concurrent map
// iteration and map write" on the main thread.
func TestKeepAlive_StandardLibrary(t *testing.T) {
	type Held struct {
		ctx    context.Context
		server *http.Server
		wg     sync.WaitGroup
		at     time.Time
	}
	if fn := compile_keepalive(reflect.TypeFor[Held]()); fn != nil {
		t.Fatal("compile_keepalive walks into standard library types")
	}
}

// TestKeepAlive_SkipList: the built-in list names packages and their
// subpackages, not other packages that only share a prefix.
func TestKeepAlive_SkipList(t *testing.T) {
	for pkg, want := range map[string]bool{
		"context": true, "net": true, "net/http": true, "crypto/tls": true,
		"sync": true, "sync/atomic": true, "time": true, "os": true, "io": true, "bufio": true,
		"": false, "main": false, "container/list": false, "netx": false, "graphics.gd/classdb": false,
	} {
		if got := keepalive_skips(pkg); got != want {
			t.Errorf("keepalive_skips(%q) = %v, want %v", pkg, got, want)
		}
	}
}

// TestKeepAlive_SkipKeepAlive: a package outside the built-in list is walked
// until it is added with [SkipKeepAlive], then its subpackages are skipped too.
func TestKeepAlive_SkipKeepAlive(t *testing.T) {
	t.Cleanup(func() {
		compiled_keepalives_mu.Lock()
		defer compiled_keepalives_mu.Unlock()
		keepalive_skipped = nil
	})
	type Before struct{ node Node.Instance }
	type After struct{ node Node.Instance }
	if compile_keepalive(reflect.TypeFor[Before]()) == nil {
		t.Fatal("compile_keepalive does not walk a type outside the built-in list")
	}
	SkipKeepAlive("graphics.gd/classdb")
	if compile_keepalive(reflect.TypeFor[Before]()) == nil {
		t.Fatal("SkipKeepAlive changed a type that was already walked")
	}
	if compile_keepalive(reflect.TypeFor[After]()) != nil {
		t.Fatal("compile_keepalive walks into a package given to SkipKeepAlive")
	}
	for pkg, want := range map[string]bool{
		"graphics.gd/classdb":     true,
		"graphics.gd/classdb/sub": true,
		"graphics.gd/classdbx":    false,
		"graphics.gd":             false,
	} {
		if got := keepalive_skips(pkg); got != want {
			t.Errorf("keepalive_skips(%q) = %v, want %v", pkg, got, want)
		}
	}
}

// TestKeepAlive_SkipKeepAliveConcurrent: SkipKeepAlive is safe next to the
// compile and cache lookups that the walk and adopted classes make.
func TestKeepAlive_SkipKeepAliveConcurrent(t *testing.T) {
	t.Cleanup(func() {
		compiled_keepalives_mu.Lock()
		defer compiled_keepalives_mu.Unlock()
		keepalive_skipped = nil
	})
	type Held struct{ node Node.Instance }
	var wg sync.WaitGroup
	for i := range 8 {
		wg.Go(func() {
			for range 100 {
				SkipKeepAlive("example.com/skip" + strconv.Itoa(i))
				compile_keepalive(reflect.TypeFor[Held]())
			}
		})
	}
	wg.Wait()
}

// TestKeepAlive_ContextChildren walks a context.Context while another
// goroutine adds and cancels children of it, which is how a request context
// is used. Unpatched it ends in a fatal error within a second.
func TestKeepAlive_ContextChildren(t *testing.T) {
	type Held struct {
		ctx context.Context
		any any
	}
	parent, cancel := context.WithCancel(context.Background())
	defer cancel()
	held := Held{ctx: parent, any: parent}
	stop := make(chan struct{})
	var wg sync.WaitGroup
	wg.Go(func() {
		for {
			select {
			case <-stop:
				return
			default:
				_, cancel := context.WithCancel(parent)
				cancel()
			}
		}
	})
	defer wg.Wait()
	defer close(stop)
	for deadline := time.Now().Add(time.Second); time.Now().Before(deadline); {
		clear(skips)
		if fn := compile_keepalive(reflect.TypeFor[Held]()); fn != nil {
			fn(reflect.ValueOf(&held).Elem())
		}
	}
}
