package vector

import "math"

// --- 2D Coordinate Transformations ---

// CartesianToPolar converts a 2D Cartesian vector to polar coordinates (r, theta).
// theta is the angle in radians measured counter-clockwise from the positive X-axis.
func CartesianToPolar(v Vector) (r, theta float64) {
	r = v.Magnitude()
	theta = math.Atan2(v.Y, v.X)
	return
}

// PolarToCartesian converts polar coordinates (r, theta) to a 2D Cartesian vector.
// theta is the angle in radians measured counter-clockwise from the positive X-axis.
func PolarToCartesian(r, theta float64) Vector {
	return Vector{
		X: r * math.Cos(theta),
		Y: r * math.Sin(theta),
	}
}

// --- 3D Coordinate Transformations ---

// CartesianToSpherical converts a 3D Cartesian vector to spherical coordinates (r, theta, phi).
// r is the radial distance, theta is the polar angle from the positive Z-axis [0, pi],
// and phi is the azimuthal angle in the XY-plane from the positive X-axis.
func CartesianToSpherical(v Vector3) (r, theta, phi float64) {
	r = v.Magnitude()
	if r == 0 {
		return 0, 0, 0
	}
	theta = math.Acos(v.Z / r)
	phi = math.Atan2(v.Y, v.X)
	return
}

// SphericalToCartesian converts spherical coordinates (r, theta, phi) to a 3D Cartesian vector.
// theta is the polar angle from the positive Z-axis, phi is the azimuthal angle in the XY-plane.
func SphericalToCartesian(r, theta, phi float64) Vector3 {
	return Vector3{
		X: r * math.Sin(theta) * math.Cos(phi),
		Y: r * math.Sin(theta) * math.Sin(phi),
		Z: r * math.Cos(theta),
	}
}

// CartesianToCylindrical converts a 3D Cartesian vector to cylindrical coordinates (rho, phi, z).
// rho is the radial distance in the XY-plane, phi is the azimuthal angle from the positive X-axis,
// and z is the height.
func CartesianToCylindrical(v Vector3) (rho, phi, z float64) {
	rho = math.Sqrt(v.X*v.X + v.Y*v.Y)
	phi = math.Atan2(v.Y, v.X)
	z = v.Z
	return
}

// CylindricalToCartesian converts cylindrical coordinates (rho, phi, z) to a 3D Cartesian vector.
// rho is the radial distance in the XY-plane, phi is the azimuthal angle from the positive X-axis.
func CylindricalToCartesian(rho, phi, z float64) Vector3 {
	return Vector3{
		X: rho * math.Cos(phi),
		Y: rho * math.Sin(phi),
		Z: z,
	}
}
