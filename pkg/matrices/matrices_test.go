package matrices

import (
	"math"
	"testing"
)

const epsilon = 1e-9

func TestIdentityMatrices(t *testing.T) {
	if Identity3[0][0] != 1 || Identity3[1][1] != 1 || Identity3[2][2] != 1 {
		t.Error("Identity3 diagonal should be 1")
	}
	if Identity3[0][1] != 0 || Identity3[1][0] != 0 {
		t.Error("Identity3 off-diagonal should be 0")
	}
	if Identity2[0][0] != 1 || Identity2[1][1] != 1 || Identity2[0][1] != 0 || Identity2[1][0] != 0 {
		t.Error("Identity2 is incorrect")
	}
}

func TestAdd(t *testing.T) {
	a := [][]float64{{1, 2}, {3, 4}}
	b := [][]float64{{5, 6}, {7, 8}}
	result, err := Add(a, b)
	if err != nil {
		t.Fatalf("Add() returned error: %v", err)
	}
	expected := [][]float64{{6, 8}, {10, 12}}
	for i := range expected {
		for j := range expected[i] {
			if result[i][j] != expected[i][j] {
				t.Errorf("Add()[%d][%d] = %v, want %v", i, j, result[i][j], expected[i][j])
			}
		}
	}
}

func TestAddDimensionMismatch(t *testing.T) {
	a := [][]float64{{1, 2}}
	b := [][]float64{{1, 2}, {3, 4}}
	_, err := Add(a, b)
	if err == nil {
		t.Error("Add() should return error for mismatched dimensions")
	}
}

func TestSubtract(t *testing.T) {
	a := [][]float64{{5, 6}, {7, 8}}
	b := [][]float64{{1, 2}, {3, 4}}
	result, err := Subtract(a, b)
	if err != nil {
		t.Fatalf("Subtract() returned error: %v", err)
	}
	expected := [][]float64{{4, 4}, {4, 4}}
	for i := range expected {
		for j := range expected[i] {
			if result[i][j] != expected[i][j] {
				t.Errorf("Subtract()[%d][%d] = %v, want %v", i, j, result[i][j], expected[i][j])
			}
		}
	}
}

func TestMultiply(t *testing.T) {
	a := [][]float64{{1, 2}, {3, 4}}
	b := [][]float64{{5, 6}, {7, 8}}
	result := Multiply(a, b)
	// [1*5+2*7, 1*6+2*8] = [19, 22]
	// [3*5+4*7, 3*6+4*8] = [43, 50]
	expected := [][]float64{{19, 22}, {43, 50}}
	for i := range expected {
		for j := range expected[i] {
			if result[i][j] != expected[i][j] {
				t.Errorf("Multiply()[%d][%d] = %v, want %v", i, j, result[i][j], expected[i][j])
			}
		}
	}
}

func TestMultiplyIdentity(t *testing.T) {
	a := [][]float64{{2, 3, 4}, {5, 6, 7}, {8, 9, 10}}
	result := Multiply(a, Identity3)
	for i := range a {
		for j := range a[i] {
			if math.Abs(result[i][j]-a[i][j]) > epsilon {
				t.Errorf("Multiply by identity changed value at [%d][%d]: got %v, want %v", i, j, result[i][j], a[i][j])
			}
		}
	}
}

func TestMultiplyDimensionMismatch(t *testing.T) {
	a := [][]float64{{1, 2, 3}}
	b := [][]float64{{1, 2}, {3, 4}}
	result := Multiply(a, b)
	if len(result) != 0 {
		t.Error("Multiply() should return empty for mismatched inner dimensions")
	}
}

func TestRotationX(t *testing.T) {
	r := RotationX(0)
	// Rotation by 0 should be identity
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			expected := 0.0
			if i == j {
				expected = 1.0
			}
			if math.Abs(r[i][j]-expected) > epsilon {
				t.Errorf("RotationX(0)[%d][%d] = %v, want %v", i, j, r[i][j], expected)
			}
		}
	}
}

func TestRotationY(t *testing.T) {
	r := RotationY(math.Pi / 2)
	// Rotating (1,0,0) by pi/2 around Y should give (0,0,-1)
	x := r[0][0]*1 + r[0][1]*0 + r[0][2]*0
	z := r[2][0]*1 + r[2][1]*0 + r[2][2]*0
	if math.Abs(x) > epsilon || math.Abs(z+1) > epsilon {
		t.Errorf("RotationY(pi/2) * (1,0,0) expected (0,0,-1), got (%v, _, %v)", x, z)
	}
}

func TestRotationZ(t *testing.T) {
	r := RotationZ(math.Pi / 2)
	// Rotating (1,0,0) by pi/2 around Z should give (0,1,0)
	x := r[0][0]*1 + r[0][1]*0 + r[0][2]*0
	y := r[1][0]*1 + r[1][1]*0 + r[1][2]*0
	if math.Abs(x) > epsilon || math.Abs(y-1) > epsilon {
		t.Errorf("RotationZ(pi/2) * (1,0,0) expected (0,1,0), got (%v, %v, _)", x, y)
	}
}
