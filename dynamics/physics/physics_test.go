package physix

import (
	"testing"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
)

func TestApplyForce(t *testing.T) {
	body := &rigidbody.RigidBody{
		Position:  vector.Vector{X: 10, Y: 10},
		Velocity:  vector.Vector{X: 0, Y: 0},
		Mass:      2,
		IsMovable: true,
	}

	// Apply 10N force over 1 second
	force := vector.Vector{X: 10, Y: 0}
	dt := 1.0

	// a = F / m = 10 / 2 = 5
	// v = v0 + a * dt = 0 + 5 * 1.0 = 5
	// p = p0 + v * dt = 10 + 5 * 1.0 = 15
	ApplyForce(body, force, dt)

	if body.Velocity.X != 5 {
		t.Errorf("Velocity expected 5, got %f", body.Velocity.X)
	}
	if body.Position.X != 15 {
		t.Errorf("Position expected 15, got %f", body.Position.X)
	}

	// Test immutability
	staticBody := &rigidbody.RigidBody{
		Position:  vector.Vector{X: 10, Y: 10},
		Velocity:  vector.Vector{X: 0, Y: 0},
		Mass:      2,
		IsMovable: false,
	}
	ApplyForce(staticBody, force, dt)
	if staticBody.Velocity.X != 0 || staticBody.Position.X != 10 {
		t.Errorf("Unmovable object altered by ApplyForce()")
	}
}
