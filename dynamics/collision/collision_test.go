package collision

import (
	"testing"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
)

func TestCircleCollided(t *testing.T) {
	b1 := &rigidbody.RigidBody{Position: vector.Vector{X: 10, Y: 10}, Radius: 5, Shape: "Circle"}
	b2 := &rigidbody.RigidBody{Position: vector.Vector{X: 15, Y: 10}, Radius: 5, Shape: "Circle"} // touching (dist=5 < 10)
	b3 := &rigidbody.RigidBody{Position: vector.Vector{X: 30, Y: 10}, Radius: 5, Shape: "Circle"} // distinct (dist=20 > 10)

	if !CircleCollided(b1, b2) {
		t.Errorf("Expected circle collision to be true")
	}
	if CircleCollided(b1, b3) {
		t.Errorf("Expected distinct circles to NOT collide")
	}
}

func TestRectangleCollided(t *testing.T) {
	r1 := &rigidbody.RigidBody{Position: vector.Vector{X: 10, Y: 10}, Width: 10, Height: 10, Shape: "Rectangle"}
	r2 := &rigidbody.RigidBody{Position: vector.Vector{X: 15, Y: 15}, Width: 10, Height: 10, Shape: "Rectangle"} // overlapping
	r3 := &rigidbody.RigidBody{Position: vector.Vector{X: 30, Y: 30}, Width: 10, Height: 10, Shape: "Rectangle"} // distinct

	if !RectangleCollided(r1, r2) {
		t.Errorf("Expected rectangle collision to be true")
	}
	if RectangleCollided(r1, r3) {
		t.Errorf("Expected distinct rectangles to NOT collide")
	}
}

func TestBounceOnCollision(t *testing.T) {
	b1 := &rigidbody.RigidBody{Position: vector.Vector{X: 10, Y: 10}, Velocity: vector.Vector{X: 5, Y: 0}, Mass: 1, Radius: 5, Shape: "Circle", IsMovable: true}
	b2 := &rigidbody.RigidBody{Position: vector.Vector{X: 15, Y: 10}, Velocity: vector.Vector{X: -5, Y: 0}, Mass: 1, Radius: 5, Shape: "Circle", IsMovable: true}

	restitution := 1.0 // Fully elastic

	// Simulate bounce
	BounceOnCollision(b1, b2, restitution)

	if b1.Velocity.X > 0 {
		t.Errorf("Expected b1.Velocity to reflect leftwards, got %f", b1.Velocity.X)
	}
	if b2.Velocity.X < 0 {
		t.Errorf("Expected b2.Velocity to reflect rightwards, got %f", b2.Velocity.X)
	}
}
