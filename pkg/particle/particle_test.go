package particle

import (
	"math"
	"testing"

	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
)

const epsilon = 1e-9

func newTestBody(x, y, mass float64, movable bool) *PVEBody {
	return &PVEBody{
		RigidBody: &rigidbody.RigidBody{
			Position:  vector.Vector{X: x, Y: y},
			Velocity:  vector.Vector{X: 0, Y: 0},
			Force:     vector.Vector{X: 0, Y: 0},
			Mass:      mass,
			Shape:     "Circle",
			Radius:    5,
			IsMovable: movable,
		},
	}
}

func TestApplyForceUpdatesVelocityAndPosition(t *testing.T) {
	body := newTestBody(0, 0, 2, true)
	force := vector.Vector{X: 10, Y: 0}
	dt := 1.0

	ApplyForce(body, force, dt)

	// a = F/m = 10/2 = 5, v = 0 + 5*1 = 5, p = 0 + 5*1 = 5
	if math.Abs(body.Velocity.X-5) > epsilon {
		t.Errorf("Velocity.X expected 5, got %v", body.Velocity.X)
	}
	if math.Abs(body.Position.X-5) > epsilon {
		t.Errorf("Position.X expected 5, got %v", body.Position.X)
	}
}

func TestApplyForceCalculatesHeat(t *testing.T) {
	body := newTestBody(0, 0, 2, true)
	force := vector.Vector{X: 10, Y: 0}
	ApplyForce(body, force, 1.0)

	// KE = 0.5 * m * v^2 = 0.5 * 2 * 25 = 25
	if math.Abs(body.Heat-25) > epsilon {
		t.Errorf("Heat expected 25, got %v", body.Heat)
	}
}

func TestApplyForceStaticBody(t *testing.T) {
	body := newTestBody(10, 10, 2, false)
	force := vector.Vector{X: 100, Y: 100}
	ApplyForce(body, force, 1.0)

	if body.Velocity.X != 0 || body.Velocity.Y != 0 {
		t.Error("Static body should not move")
	}
	if body.Position.X != 10 || body.Position.Y != 10 {
		t.Error("Static body position should not change")
	}
}

func TestApplyForceZeroMass(t *testing.T) {
	body := newTestBody(0, 0, 0, true)
	force := vector.Vector{X: 10, Y: 0}

	// Should not panic
	ApplyForce(body, force, 1.0)

	if body.Velocity.X != 0 {
		t.Error("Zero-mass body velocity should stay 0")
	}
}

func TestResolveCollisionSeparatesOverlappingBalls(t *testing.T) {
	ball1 := &rigidbody.RigidBody{
		Position: vector.Vector{X: 0, Y: 0}, Velocity: vector.Vector{},
		Mass: 1, Radius: 5, Shape: "Circle", IsMovable: true,
	}
	ball2 := &rigidbody.RigidBody{
		Position: vector.Vector{X: 6, Y: 0}, Velocity: vector.Vector{},
		Mass: 1, Radius: 5, Shape: "Circle", IsMovable: true,
	}

	ResolveCollision(ball1, ball2, 0.1)

	// After collision, balls should be pushed apart: ball1 leftward, ball2 rightward
	if ball1.Velocity.X >= 0 {
		t.Errorf("ball1 should have negative X velocity, got %v", ball1.Velocity.X)
	}
	if ball2.Velocity.X <= 0 {
		t.Errorf("ball2 should have positive X velocity, got %v", ball2.Velocity.X)
	}

	// Positions should also be corrected apart
	dist := ball1.Position.Sub(ball2.Position).Magnitude()
	if dist <= 6 {
		t.Errorf("balls should be pushed apart, distance is %v", dist)
	}
}

func TestResolveCollisionNoOverlap(t *testing.T) {
	ball1 := &rigidbody.RigidBody{
		Position: vector.Vector{X: 0, Y: 0}, Velocity: vector.Vector{},
		Mass: 1, Radius: 5, Shape: "Circle", IsMovable: true,
	}
	ball2 := &rigidbody.RigidBody{
		Position: vector.Vector{X: 20, Y: 0}, Velocity: vector.Vector{},
		Mass: 1, Radius: 5, Shape: "Circle", IsMovable: true,
	}

	ResolveCollision(ball1, ball2, 0.1)

	// No overlap, velocities should stay zero
	if ball1.Velocity.X != 0 || ball2.Velocity.X != 0 {
		t.Error("No overlap, velocities should remain zero")
	}
}

func TestResolveCollisionZeroMass(t *testing.T) {
	ball1 := &rigidbody.RigidBody{
		Position: vector.Vector{X: 0, Y: 0}, Velocity: vector.Vector{},
		Mass: 0, Radius: 5, Shape: "Circle", IsMovable: true,
	}
	ball2 := &rigidbody.RigidBody{
		Position: vector.Vector{X: 6, Y: 0}, Velocity: vector.Vector{},
		Mass: 0, Radius: 5, Shape: "Circle", IsMovable: true,
	}

	// Should not panic
	ResolveCollision(ball1, ball2, 0.1)
}
