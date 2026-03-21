package vector

import (
	"math"
	"testing"
)

func TestVectorOperations(t *testing.T) {
	v1 := Vector{X: 3, Y: 4}
	v2 := Vector{X: 1, Y: 2}

	// Test Vector Add
	vAdd := v1.Add(v2)
	if vAdd.X != 4 || vAdd.Y != 6 {
		t.Errorf("Add() expected {4, 6}, got {%v, %v}", vAdd.X, vAdd.Y)
	}

	// Test Vector Sub
	vSub := v1.Sub(v2)
	if vSub.X != 2 || vSub.Y != 2 {
		t.Errorf("Sub() expected {2, 2}, got {%v, %v}", vSub.X, vSub.Y)
	}

	// Test Vector Scale
	vScale := v1.Scale(2)
	if vScale.X != 6 || vScale.Y != 8 {
		t.Errorf("Scale() expected {6, 8}, got {%v, %v}", vScale.X, vScale.Y)
	}

	// Test InnerProduct
	dot := v1.InnerProduct(v2)
	if dot != 11 {
		t.Errorf("InnerProduct() expected 11, got %v", dot)
	}

	// Test Magnitude
	mag := v1.Magnitude()
	if mag != 5 {
		t.Errorf("Magnitude() expected 5.0, got %v", mag)
	}

	// Test Normalize
	vNorm := v1.Normalize()
	if math.Abs(vNorm.X-0.6) > 1e-9 || math.Abs(vNorm.Y-0.8) > 1e-9 {
		t.Errorf("Normalize() expected {0.6, 0.8}, got {%v, %v}", vNorm.X, vNorm.Y)
	}

	// Test Distance
	dist := Distance(v1, v2)
	expectedDist := math.Sqrt(8)
	if math.Abs(dist-expectedDist) > 1e-9 {
		t.Errorf("Distance() expected %v, got %v", expectedDist, dist)
	}
}

func TestNormalizeZero(t *testing.T) {
	v := Vector{X: 0, Y: 0}
	vNorm := v.Normalize()
	if vNorm.X != 0 || vNorm.Y != 0 {
		t.Errorf("Normalize() of zero vector should be {0, 0}, got {%v, %v}", vNorm.X, vNorm.Y)
	}
}
