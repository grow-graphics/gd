package gd_test

import (
	"testing"

	"graphics.gd/variant/Object"
)

// TestSignalArgsNestedCall emits a signal with Object.call("emit_signal")
// to two Go handlers. The first handler emits another signal with another
// argument. The second handler must still receive the first signal's
// argument: emit_signal passes each handler pointers into the arguments of
// the outer call, so a nested engine call must not reuse that memory.
func TestSignalArgsNestedCall(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		sender := new(CustomStringSignals)
		other := new(CustomStringSignals)
		// Make the engine objects first, so that their signals exist.
		sender.AsNode()
		other.AsNode()
		var got []string
		handler := func(s string) {
			got = append(got, s)
			if len(got) == 1 {
				Object.Call(other, "emit_signal", "on_string", "clobber")
			}
		}
		// Two distinct callables: Godot refuses to connect one twice.
		sender.StringSignal.Call(func(s string) { handler(s) })
		sender.StringSignal.Call(func(s string) { handler(s) })
		Object.Call(sender, "emit_signal", "on_string", "sent")
		if len(got) != 2 || got[0] != "sent" || got[1] != "sent" {
			t.Errorf("handlers received %q, want [\"sent\" \"sent\"]", got)
		}
		sender.AsNode().QueueFree()
		other.AsNode().QueueFree()
	})
}
