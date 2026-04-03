package vector

import (
	"math"
	"testing"
)

const epsilon = 1e-9

func TestCartesianToPolar(t *testing.T) {
	// (1, 0) -> r=1, theta=0
	r, theta := CartesianToPolar(Vector{X: 1, Y: 0})
	if math.Abs(r-1) > epsilon || math.Abs(theta) > epsilon {
		t.Errorf("CartesianToPolar(1,0) expected (1, 0), got (%v, %v)", r, theta)
	}

	// (0, 1) -> r=1, theta=pi/2
	r, theta = CartesianToPolar(Vector{X: 0, Y: 1})
	if math.Abs(r-1) > epsilon || math.Abs(theta-math.Pi/2) > epsilon {
		t.Errorf("CartesianToPolar(0,1) expected (1, pi/2), got (%v, %v)", r, theta)
	}

	// (3, 4) -> r=5, theta=atan2(4,3)
	r, theta = CartesianToPolar(Vector{X: 3, Y: 4})
	if math.Abs(r-5) > epsilon || math.Abs(theta-math.Atan2(4, 3)) > epsilon {
		t.Errorf("CartesianToPolar(3,4) expected (5, %v), got (%v, %v)", math.Atan2(4, 3), r, theta)
	}
}

func TestPolarToCartesian(t *testing.T) {
	// r=1, theta=0 -> (1, 0)
	v := PolarToCartesian(1, 0)
	if math.Abs(v.X-1) > epsilon || math.Abs(v.Y) > epsilon {
		t.Errorf("PolarToCartesian(1, 0) expected (1, 0), got (%v, %v)", v.X, v.Y)
	}

	// r=5, theta=pi/2 -> (0, 5)
	v = PolarToCartesian(5, math.Pi/2)
	if math.Abs(v.X) > epsilon || math.Abs(v.Y-5) > epsilon {
		t.Errorf("PolarToCartesian(5, pi/2) expected (0, 5), got (%v, %v)", v.X, v.Y)
	}
}

func TestPolarRoundTrip(t *testing.T) {
	original := Vector{X: -3, Y: 7}
	r, theta := CartesianToPolar(original)
	result := PolarToCartesian(r, theta)
	if math.Abs(result.X-original.X) > epsilon || math.Abs(result.Y-original.Y) > epsilon {
		t.Errorf("Polar round-trip failed: started with %v, got %v", original, result)
	}
}

func TestCartesianToSpherical(t *testing.T) {
	// Point along Z axis: (0, 0, 5) -> r=5, theta=0, phi=0
	r, theta, phi := CartesianToSpherical(Vector3{X: 0, Y: 0, Z: 5})
	if math.Abs(r-5) > epsilon || math.Abs(theta) > epsilon || math.Abs(phi) > epsilon {
		t.Errorf("CartesianToSpherical(0,0,5) expected (5, 0, 0), got (%v, %v, %v)", r, theta, phi)
	}

	// Point along X axis: (3, 0, 0) -> r=3, theta=pi/2, phi=0
	r, theta, phi = CartesianToSpherical(Vector3{X: 3, Y: 0, Z: 0})
	if math.Abs(r-3) > epsilon || math.Abs(theta-math.Pi/2) > epsilon || math.Abs(phi) > epsilon {
		t.Errorf("CartesianToSpherical(3,0,0) expected (3, pi/2, 0), got (%v, %v, %v)", r, theta, phi)
	}

	// Zero vector
	r, theta, phi = CartesianToSpherical(Vector3{})
	if r != 0 || theta != 0 || phi != 0 {
		t.Errorf("CartesianToSpherical(0,0,0) expected (0, 0, 0), got (%v, %v, %v)", r, theta, phi)
	}
}

func TestSphericalRoundTrip(t *testing.T) {
	original := Vector3{X: 2, Y: -3, Z: 4}
	r, theta, phi := CartesianToSpherical(original)
	result := SphericalToCartesian(r, theta, phi)
	if math.Abs(result.X-original.X) > epsilon || math.Abs(result.Y-original.Y) > epsilon || math.Abs(result.Z-original.Z) > epsilon {
		t.Errorf("Spherical round-trip failed: started with %v, got %v", original, result)
	}
}

func TestCartesianToCylindrical(t *testing.T) {
	// Point along X axis: (3, 0, 7) -> rho=3, phi=0, z=7
	rho, phi, z := CartesianToCylindrical(Vector3{X: 3, Y: 0, Z: 7})
	if math.Abs(rho-3) > epsilon || math.Abs(phi) > epsilon || math.Abs(z-7) > epsilon {
		t.Errorf("CartesianToCylindrical(3,0,7) expected (3, 0, 7), got (%v, %v, %v)", rho, phi, z)
	}

	// Point along Y axis: (0, 4, 2) -> rho=4, phi=pi/2, z=2
	rho, phi, z = CartesianToCylindrical(Vector3{X: 0, Y: 4, Z: 2})
	if math.Abs(rho-4) > epsilon || math.Abs(phi-math.Pi/2) > epsilon || math.Abs(z-2) > epsilon {
		t.Errorf("CartesianToCylindrical(0,4,2) expected (4, pi/2, 2), got (%v, %v, %v)", rho, phi, z)
	}
}

func TestCylindricalRoundTrip(t *testing.T) {
	original := Vector3{X: -1, Y: 5, Z: -2}
	rho, phi, z := CartesianToCylindrical(original)
	result := CylindricalToCartesian(rho, phi, z)
	if math.Abs(result.X-original.X) > epsilon || math.Abs(result.Y-original.Y) > epsilon || math.Abs(result.Z-original.Z) > epsilon {
		t.Errorf("Cylindrical round-trip failed: started with %v, got %v", original, result)
	}
}
