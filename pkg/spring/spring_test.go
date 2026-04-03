package spring

import (
	"math"
	"testing"

	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
)

const epsilon = 1e-9

func makeBall(x, y, mass float64, movable bool) *rigidbody.RigidBody {
	return &rigidbody.RigidBody{
		Position:  vector.Vector{X: x, Y: y},
		Velocity:  vector.Vector{X: 0, Y: 0},
		Mass:      mass,
		IsMovable: movable,
	}
}

func TestNewSpringDefaultRestLength(t *testing.T) {
	a := makeBall(0, 0, 1, true)
	b := makeBall(10, 0, 1, true)
	s := NewSpring(a, b, 1.0, 0.1)

	if math.Abs(s.RestLength-10) > epsilon {
		t.Errorf("Default RestLength expected 10, got %v", s.RestLength)
	}
	if s.Stiffness != 1.0 || s.Damping != 0.1 {
		t.Error("Stiffness or Damping not set correctly")
	}
}

func TestNewSpringCustomRestLength(t *testing.T) {
	a := makeBall(0, 0, 1, true)
	b := makeBall(10, 0, 1, true)
	s := NewSpring(a, b, 1.0, 0.1, 5.0)

	if math.Abs(s.RestLength-5) > epsilon {
		t.Errorf("Custom RestLength expected 5, got %v", s.RestLength)
	}
}

func TestApplyForceStretched(t *testing.T) {
	// Spring at rest length 5, balls 10 apart -> stretched
	a := makeBall(0, 0, 1, true)
	b := makeBall(10, 0, 1, true)
	s := NewSpring(a, b, 1.0, 0.0, 5.0)

	s.ApplyForce()

	// Force direction A->B, magnitude = stiffness * (10 - 5) = 5
	// BallA should gain positive X velocity (toward B)
	if a.Velocity.X <= 0 {
		t.Errorf("BallA should move toward B (positive X), got %v", a.Velocity.X)
	}
	// BallB should gain negative X velocity (toward A)
	if b.Velocity.X >= 0 {
		t.Errorf("BallB should move toward A (negative X), got %v", b.Velocity.X)
	}
}

func TestApplyForceCompressed(t *testing.T) {
	// Spring at rest length 10, balls 2 apart -> compressed
	a := makeBall(0, 0, 1, true)
	b := makeBall(2, 0, 1, true)
	s := NewSpring(a, b, 1.0, 0.0, 10.0)

	s.ApplyForce()

	// Compressed: force pushes balls apart
	// BallA should move away from B (negative X)
	if a.Velocity.X >= 0 {
		t.Errorf("BallA should be pushed away (negative X), got %v", a.Velocity.X)
	}
	// BallB should move away from A (positive X)
	if b.Velocity.X <= 0 {
		t.Errorf("BallB should be pushed away (positive X), got %v", b.Velocity.X)
	}
}

func TestApplyForceAtRest(t *testing.T) {
	// Spring at rest length = actual distance -> no force
	a := makeBall(0, 0, 1, true)
	b := makeBall(5, 0, 1, true)
	s := NewSpring(a, b, 1.0, 0.0, 5.0)

	s.ApplyForce()

	if math.Abs(a.Velocity.X) > epsilon || math.Abs(b.Velocity.X) > epsilon {
		t.Errorf("At rest length, velocities should be 0. Got a=%v, b=%v", a.Velocity.X, b.Velocity.X)
	}
}

func TestApplyForceImmovableBall(t *testing.T) {
	a := makeBall(0, 0, 1, false) // immovable
	b := makeBall(10, 0, 1, true)
	s := NewSpring(a, b, 1.0, 0.0, 5.0)

	s.ApplyForce()

	if a.Velocity.X != 0 || a.Velocity.Y != 0 {
		t.Error("Immovable ball A should not move")
	}
	if b.Velocity.X >= 0 {
		t.Errorf("BallB should move toward A, got Vx=%v", b.Velocity.X)
	}
}

func TestApplyForceWithDamping(t *testing.T) {
	// Two balls moving apart; damping should reduce velocity change
	a := makeBall(0, 0, 1, true)
	b := makeBall(10, 0, 1, true)
	a.Velocity = vector.Vector{X: -5, Y: 0} // moving away from B
	b.Velocity = vector.Vector{X: 5, Y: 0}  // moving away from A

	sNoDamp := NewSpring(makeBall(0, 0, 1, true), makeBall(10, 0, 1, true), 1.0, 0.0, 5.0)
	sNoDamp.BallA.Velocity = vector.Vector{X: -5, Y: 0}
	sNoDamp.BallB.Velocity = vector.Vector{X: 5, Y: 0}

	sDamp := NewSpring(a, b, 1.0, 0.5, 5.0)

	sNoDamp.ApplyForce()
	sDamp.ApplyForce()

	// With damping, BallA should gain MORE positive velocity (damping adds to spring force here)
	if sDamp.BallA.Velocity.X <= sNoDamp.BallA.Velocity.X {
		t.Error("Damping should increase corrective force when balls move apart")
	}
}

func TestApplyForceZeroMass(t *testing.T) {
	a := makeBall(0, 0, 0, true)
	b := makeBall(10, 0, 0, true)
	s := NewSpring(a, b, 1.0, 0.0, 5.0)

	// Should not panic
	s.ApplyForce()
}
