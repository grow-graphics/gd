package project

import (
	"path/filepath"
	"testing"
	"time"
)

func TestStaleImports(t *testing.T) {
	dir := t.TempDir()
	old := time.Now().Add(-time.Hour)
	new := time.Now()

	write(t, dir, "icon.svg", "", old)
	write(t, dir, "icon.svg.import", "", new) // imported
	write(t, dir, "main.tscn", "", new)       // not importable
	if stale := staleImports(dir); len(stale) != 0 {
		t.Fatalf("expected nothing to import, got %v", stale)
	}

	write(t, dir, "models/house.glb", "", old)       // newly added
	write(t, dir, "sounds/door.ogg", "", new)        // changed since import
	write(t, dir, "sounds/door.ogg.import", "", old) //
	write(t, dir, ".godot/imported/x.png", "", new)  // cache
	write(t, dir, "raw/texture.png", "", new)        // ignored
	write(t, dir, "raw/.gdignore", "", new)          //
	stale := staleImports(dir)
	want := []string{filepath.Join(dir, "models/house.glb"), filepath.Join(dir, "sounds/door.ogg")}
	if len(stale) != len(want) || stale[0] != want[0] || stale[1] != want[1] {
		t.Fatalf("expected %v, got %v", want, stale)
	}
}
