---
title: Reference Invalidation
slug: guide/memory
sidebar:
  order: 21
---

Thanks to the garbage collector in Go, you shouldn't need to worry about memory management.

`graphics.gd` will track any Godot references that you access each frame and will invalidate them once
they've remained unused for two entire frames. As such, if you need to keep any Objects
or high-performance reference types (`Array.Any`, `Packed.Array[T]`, `Dictionary.Any`, `Signal.Any`,
`Callable.Function`, `String.Readable`, `String.Name`, `Path.ToNode`) around in Go for longer than a frame, you should either do something with the
reference every frame to keep it alive, or, in the case of Objects, store the ID of the object
instead of the `Object.Instance` itself.

If you are sticking with convienence types and `Object.ID` (each classdb object has it's own typed ID available), you won't run into any problems here.

In the worst case, `graphics.gd` will produce a panic if you try to use an Engine reference
that has been invalidated, (this may terminate your application). `graphics.gd` has been designed
to provide the best possible memory safety protections when working with the engine, if you run into
a memory safety issue, this means either `graphics.gd` or the engine is at fault, please [open an issue](https://github.com/quaadgras/graphics.gd/issues/new/choose).

### Globals
If you initalise an object as a global, the underlying reference will not be invalidated until shutdown. These objects
are only valid after:

  * [startup.LoadingScene] is called, or
  * [startup.Rendering] is called, or
  * [startup.Scene] is called, or
  * they are accessed inside a method called by the engine

If accessed earlier, these objects will panic-on-use.

```go
package main

import "graphics.gd/classdb/Sprite2D"

var global = Sprite2D.New()
```

If you assign newly created objects to these globals, they still need to be used every frame and unless you made a copy first,
the original reference will be lost. Therefore, unless you really know what you are doing, do not assign new objects to any
global variables after the engine has started up.

### Goroutines

Memory safety protections only apply within a single thread, please be careful when using goroutines and make sure only to use
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
