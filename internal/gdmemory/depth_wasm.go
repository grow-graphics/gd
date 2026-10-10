//go:build wasm

package gdmemory

import "graphics.gd/internal/threadcheck"

// depth is the number of engine->Go callbacks that are running.
func depth() int { return threadcheck.FrameDepth() }
