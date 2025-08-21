---
title: Memory Safety
slug: guide/memory-safety
sidebar:
  order: 21
---

Thanks to the garbage collector in Go, you shouldn't need to worry about memory safety.

`graphics.gd` will track any Godot pointers that you reference each frame and will invalidate them once
they've remained unused for two whole frames. As such, if you need to keep an Object
or a performance-optimised collection type (`Array.Any`, `Packed.Array[T]`, `Dictionary.Any`, `Signal.Any`,
`Callable.Function`, `String.Readable`, `String.Name`, `Path.ToNode`) around in Go for longer than a frame, you should either do something with the
reference every frame to keep it alive, or, in the case of Objects, store the ID of the object
instead of the `Object.Instance` itself.

If you are sticking with convienence types and `Object.ID` (each object has it's own typed ID too), you won't run into any problems here.

In the worst case, graphics.gd will raise a panic if you try to use an Engine reference
that has been invalidated, which may crash your program. `graphics.gd` has been designed
to provide the maximum level of memory safety when working with the engine, if you run into
a memory safety issue, please [open an issue](https://github.com/quaadgras/graphics.gd/issues/new/choose).

## Goroutines

Please be careful when using goroutines and make sure only to use
[Thread Safe APIs](https://docs.godotengine.org/en/latest/tutorials/performance/thread_safe_apis.html).
Prefer channels or signals wherever possible and [Share Memory by Communicating](https://go.dev/blog/codelab-share)
to prevent data races.

```go
type MyTimer struct {
	Object.Extension[MyTimer]

  Timer Signal.Void
}

func (t *MyTimer) Ready() {
	go func() {
		time.Sleep(2 * time.Second)
		t.Timer.Emit() // safe to call from a goroutine
	}()
}

```
