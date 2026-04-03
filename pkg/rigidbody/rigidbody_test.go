package rigidbody

import (
	"math"
	"testing"

	"github.com/rudransh61/Physix-go/pkg/vector"
)

const epsilon = 1e-9

func TestApplyTorque(t *testing.T) {
	rb := &RigidBody{}
	rb.ApplyTorque(5)
	rb.ApplyTorque(3)
	if rb.Torque != 8 {
		t.Errorf("Torque expected 8, got %v", rb.Torque)
	}
}

func TestApplyImpulse(t *testing.T) {
	rb := &RigidBody{
		Mass:     4,
		Velocity: vector.Vector{X: 0, Y: 0},
	}
	rb.ApplyImpulse(vector.Vector{X: 12, Y: 8})

	// change = impulse / mass = (12/4, 8/4) = (3, 2)
	if math.Abs(rb.Velocity.X-3) > epsilon || math.Abs(rb.Velocity.Y-2) > epsilon {
		t.Errorf("Velocity expected (3, 2), got (%v, %v)", rb.Velocity.X, rb.Velocity.Y)
	}
}

func TestApplyImpulseZeroMass(t *testing.T) {
	rb := &RigidBody{
		Mass:     0,
		Velocity: vector.Vector{X: 1, Y: 1},
	}
	// Should not panic
	rb.ApplyImpulse(vector.Vector{X: 10, Y: 10})
	if rb.Velocity.X != 1 || rb.Velocity.Y != 1 {
		t.Error("Zero-mass body velocity should be unchanged")
	}
}

func TestUpdateRotation(t *testing.T) {
	// Place body at (1, 0) with angular velocity such that pi/2 rotation in dt=1
	rb := &RigidBody{
		Position:        vector.Vector{X: 1, Y: 0},
		AngularVelocity: math.Pi / 2,
	}
	rb.UpdateRotation(1.0)

	// After 90-degree rotation, (1, 0) -> (0, 1)
	if math.Abs(rb.Position.X) > epsilon || math.Abs(rb.Position.Y-1) > epsilon {
		t.Errorf("UpdateRotation: position expected (0, 1), got (%v, %v)", rb.Position.X, rb.Position.Y)
	}
}

func TestUpdateRotationZeroAngularVelocity(t *testing.T) {
	rb := &RigidBody{
		Position:        vector.Vector{X: 5, Y: 3},
		AngularVelocity: 0,
	}
	rb.UpdateRotation(1.0)

	if rb.Position.X != 5 || rb.Position.Y != 3 {
		t.Errorf("Zero angular velocity should not change position, got (%v, %v)", rb.Position.X, rb.Position.Y)
	}
}

func TestRotateCoordinates(t *testing.T) {
	rb := &RigidBody{
		Position: vector.Vector{X: 1, Y: 0},
	}
	// Rotate 90 degrees
	result := rb.rotateCoordinates(90)

	if math.Abs(result.X) > epsilon || math.Abs(result.Y-1) > epsilon {
		t.Errorf("rotateCoordinates(90) expected (0, 1), got (%v, %v)", result.X, result.Y)
	}
}

func TestInfiniteMass(t *testing.T) {
	if Infinite_mass != 1e10 {
		t.Errorf("Infinite_mass expected 1e10, got %v", Infinite_mass)
	}
}
