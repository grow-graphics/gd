package gd_test

import (
	"testing"

	gd "graphics.gd/internal"
	"graphics.gd/variant"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector3"
)

func TestVariants(t *testing.T) {
	var f = variant.New(Float.X(3.14))
	if f.Float32() != 3.14 {
		t.Fatal()
	}
	if f.Interface() != Float.X(3.14) {
		t.Fatal()
	}

	var v3 = variant.New(Vector3.XYZ{1, 2, 3})
	if v3.Vector3().X != 1 || v3.Vector3().Y != 2 || v3.Vector3().Z != 3 {
		t.Fatal()
	}

	var iv3 = gd.InternalVariant(v3)
	if iv3.Vector3().X != 1 || iv3.Vector3().Y != 2 || iv3.Vector3().Z != 3 {
		t.Fatal()
	}

	var if64 = gd.InternalVariant(f)
	if !Float.IsApproximatelyEqual(if64.Float(), 3.14) {
		t.Fatal()
	}
}
