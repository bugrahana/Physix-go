package vector

import "math"

// Vector3 represents a 3D vector.
type Vector3 struct {
	X, Y, Z float64
}

// NewVector3 creates a new 3D vector with the given x, y, and z components.
func NewVector3(x, y, z float64) Vector3 {
	return Vector3{X: x, Y: y, Z: z}
}

// Add performs 3D vector addition.
func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{v.X + other.X, v.Y + other.Y, v.Z + other.Z}
}

// Sub performs 3D vector subtraction.
func (v Vector3) Sub(other Vector3) Vector3 {
	return Vector3{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

// Scale multiplies the 3D vector by a scalar.
func (v Vector3) Scale(scalar float64) Vector3 {
	return Vector3{v.X * scalar, v.Y * scalar, v.Z * scalar}
}

// Magnitude calculates the magnitude (length) of the 3D vector.
func (v Vector3) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalize normalizes the 3D vector to have a magnitude of 1.
func (v Vector3) Normalize() Vector3 {
	magnitude := v.Magnitude()
	if magnitude == 0 {
		return Vector3{}
	}
	return v.Scale(1 / magnitude)
}

// InnerProduct performs the dot product of two 3D vectors.
func (v Vector3) InnerProduct(other Vector3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// Cross returns the cross product of two 3D vectors.
func (v Vector3) Cross(other Vector3) Vector3 {
	return Vector3{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

// Distance3 calculates the distance between the tips of two 3D vectors.
func Distance3(v1, v2 Vector3) float64 {
	return v1.Sub(v2).Magnitude()
}
