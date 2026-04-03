package vector

import (
	"math"
	"testing"
)

func TestVector3Operations(t *testing.T) {
	v1 := Vector3{X: 1, Y: 2, Z: 3}
	v2 := Vector3{X: 4, Y: 5, Z: 6}

	// Test Add
	vAdd := v1.Add(v2)
	if vAdd.X != 5 || vAdd.Y != 7 || vAdd.Z != 9 {
		t.Errorf("Add() expected {5, 7, 9}, got {%v, %v, %v}", vAdd.X, vAdd.Y, vAdd.Z)
	}

	// Test Sub
	vSub := v2.Sub(v1)
	if vSub.X != 3 || vSub.Y != 3 || vSub.Z != 3 {
		t.Errorf("Sub() expected {3, 3, 3}, got {%v, %v, %v}", vSub.X, vSub.Y, vSub.Z)
	}

	// Test Scale
	vScale := v1.Scale(2)
	if vScale.X != 2 || vScale.Y != 4 || vScale.Z != 6 {
		t.Errorf("Scale() expected {2, 4, 6}, got {%v, %v, %v}", vScale.X, vScale.Y, vScale.Z)
	}

	// Test InnerProduct
	dot := v1.InnerProduct(v2)
	if dot != 32 { // 1*4 + 2*5 + 3*6 = 32
		t.Errorf("InnerProduct() expected 32, got %v", dot)
	}

	// Test Magnitude
	v := Vector3{X: 2, Y: 3, Z: 6}
	mag := v.Magnitude()
	if mag != 7 { // sqrt(4 + 9 + 36) = 7
		t.Errorf("Magnitude() expected 7, got %v", mag)
	}

	// Test Normalize
	vNorm := v.Normalize()
	if math.Abs(vNorm.X-2.0/7) > 1e-9 || math.Abs(vNorm.Y-3.0/7) > 1e-9 || math.Abs(vNorm.Z-6.0/7) > 1e-9 {
		t.Errorf("Normalize() expected {2/7, 3/7, 6/7}, got {%v, %v, %v}", vNorm.X, vNorm.Y, vNorm.Z)
	}

	// Test Distance3
	dist := Distance3(v1, v2)
	expected := math.Sqrt(27) // sqrt(9+9+9)
	if math.Abs(dist-expected) > 1e-9 {
		t.Errorf("Distance3() expected %v, got %v", expected, dist)
	}
}

func TestVector3Cross(t *testing.T) {
	// i x j = k
	i := Vector3{X: 1, Y: 0, Z: 0}
	j := Vector3{X: 0, Y: 1, Z: 0}
	k := i.Cross(j)
	if k.X != 0 || k.Y != 0 || k.Z != 1 {
		t.Errorf("Cross(i, j) expected {0, 0, 1}, got {%v, %v, %v}", k.X, k.Y, k.Z)
	}

	// j x i = -k
	negK := j.Cross(i)
	if negK.X != 0 || negK.Y != 0 || negK.Z != -1 {
		t.Errorf("Cross(j, i) expected {0, 0, -1}, got {%v, %v, %v}", negK.X, negK.Y, negK.Z)
	}

	// Parallel vectors: cross product is zero
	v1 := Vector3{X: 2, Y: 4, Z: 6}
	v2 := Vector3{X: 1, Y: 2, Z: 3}
	zero := v1.Cross(v2)
	if zero.X != 0 || zero.Y != 0 || zero.Z != 0 {
		t.Errorf("Cross of parallel vectors expected {0, 0, 0}, got {%v, %v, %v}", zero.X, zero.Y, zero.Z)
	}
}

func TestVector3NormalizeZero(t *testing.T) {
	v := Vector3{X: 0, Y: 0, Z: 0}
	vNorm := v.Normalize()
	if vNorm.X != 0 || vNorm.Y != 0 || vNorm.Z != 0 {
		t.Errorf("Normalize() of zero vector should be {0, 0, 0}, got {%v, %v, %v}", vNorm.X, vNorm.Y, vNorm.Z)
	}
}
