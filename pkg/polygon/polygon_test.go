package polygon

import (
	"math"
	"testing"

	"github.com/rudransh61/Physix-go/pkg/vector"
)

const epsilon = 1e-9

func squareVertices() []vector.Vector {
	return []vector.Vector{
		{X: 0, Y: 0},
		{X: 4, Y: 0},
		{X: 4, Y: 4},
		{X: 0, Y: 4},
	}
}

func TestCalculateCentroid(t *testing.T) {
	c := CalculateCentroid(squareVertices())
	if math.Abs(c.X-2) > epsilon || math.Abs(c.Y-2) > epsilon {
		t.Errorf("Centroid expected (2, 2), got (%v, %v)", c.X, c.Y)
	}
}

func TestCalculateCentroidEmpty(t *testing.T) {
	c := CalculateCentroid([]vector.Vector{})
	if c.X != 0 || c.Y != 0 {
		t.Errorf("Centroid of empty vertices should be (0, 0), got (%v, %v)", c.X, c.Y)
	}
}

func TestCalculateCentroidSingle(t *testing.T) {
	c := CalculateCentroid([]vector.Vector{{X: 5, Y: 7}})
	if math.Abs(c.X-5) > epsilon || math.Abs(c.Y-7) > epsilon {
		t.Errorf("Centroid of single vertex should be that vertex, got (%v, %v)", c.X, c.Y)
	}
}

func TestNewPolygon(t *testing.T) {
	verts := squareVertices()
	p := NewPolygon(verts, 10, true)

	if p.Mass != 10 {
		t.Errorf("Mass expected 10, got %v", p.Mass)
	}
	if !p.IsMovable {
		t.Error("Polygon should be movable")
	}
	if p.Shape != "polygon" {
		t.Errorf("Shape expected 'polygon', got %v", p.Shape)
	}
	if math.Abs(p.Position.X-2) > epsilon || math.Abs(p.Position.Y-2) > epsilon {
		t.Errorf("Position should be centroid (2, 2), got (%v, %v)", p.Position.X, p.Position.Y)
	}
	if p.Rotation != 0 || p.Torque != 0 {
		t.Error("Initial rotation and torque should be 0")
	}
	if p.Restitution != 1.0 {
		t.Errorf("Default restitution expected 1.0, got %v", p.Restitution)
	}
}

func TestUpdatePosition(t *testing.T) {
	p := NewPolygon(squareVertices(), 10, true)
	// Move all vertices right by 10
	for i := range p.Vertices {
		p.Vertices[i].X += 10
	}
	p.Torque = 5
	p.UpdatePosition()

	if math.Abs(p.Position.X-12) > epsilon || math.Abs(p.Position.Y-2) > epsilon {
		t.Errorf("Updated position expected (12, 2), got (%v, %v)", p.Position.X, p.Position.Y)
	}
	// Rotation = Torque / Mass = 5 / 10 = 0.5
	if math.Abs(p.Rotation-0.5) > epsilon {
		t.Errorf("Rotation expected 0.5, got %v", p.Rotation)
	}
}

func TestUpdatePositionZeroMass(t *testing.T) {
	p := NewPolygon(squareVertices(), 0, true)
	p.Torque = 5
	// Should not panic
	p.UpdatePosition()
}

func TestApplyImpulse(t *testing.T) {
	p := NewPolygon(squareVertices(), 4, true)
	impulse := vector.Vector{X: 8, Y: 0}
	p.ApplyImpulse(impulse)

	// change_velocity = impulse / mass = 8/4 = 2
	if math.Abs(p.Velocity.X-2) > epsilon {
		t.Errorf("Velocity.X expected 2, got %v", p.Velocity.X)
	}
}

func TestApplyImpulseZeroMass(t *testing.T) {
	p := NewPolygon(squareVertices(), 0, true)
	// Should not panic
	p.ApplyImpulse(vector.Vector{X: 10, Y: 10})
	if p.Velocity.X != 0 || p.Velocity.Y != 0 {
		t.Error("Zero-mass polygon velocity should stay 0")
	}
}

func TestProject(t *testing.T) {
	p := NewPolygon(squareVertices(), 1, true)
	// Project onto X-axis
	axis := vector.Vector{X: 1, Y: 0}
	min, max := Project(*p, axis)
	if math.Abs(min-0) > epsilon || math.Abs(max-4) > epsilon {
		t.Errorf("Project onto X expected (0, 4), got (%v, %v)", min, max)
	}

	// Project onto Y-axis
	axis = vector.Vector{X: 0, Y: 1}
	min, max = Project(*p, axis)
	if math.Abs(min-0) > epsilon || math.Abs(max-4) > epsilon {
		t.Errorf("Project onto Y expected (0, 4), got (%v, %v)", min, max)
	}
}

func TestProjectEmptyVertices(t *testing.T) {
	p := Polygon{}
	min, max := Project(p, vector.Vector{X: 1, Y: 0})
	if min != 0 || max != 0 {
		t.Errorf("Project on empty polygon expected (0, 0), got (%v, %v)", min, max)
	}
}

func TestMove(t *testing.T) {
	p := NewPolygon(squareVertices(), 1, true)
	p.Move(vector.Vector{X: 10, Y: 5})

	if math.Abs(p.Vertices[0].X-10) > epsilon || math.Abs(p.Vertices[0].Y-5) > epsilon {
		t.Errorf("Vertex 0 expected (10, 5), got (%v, %v)", p.Vertices[0].X, p.Vertices[0].Y)
	}
	if math.Abs(p.Vertices[2].X-14) > epsilon || math.Abs(p.Vertices[2].Y-9) > epsilon {
		t.Errorf("Vertex 2 expected (14, 9), got (%v, %v)", p.Vertices[2].X, p.Vertices[2].Y)
	}
}

func TestClosestPoint(t *testing.T) {
	p := NewPolygon(squareVertices(), 1, true)
	// Reset position to origin so vertex offsets work as expected
	p.Position = vector.Vector{X: 0, Y: 0}

	cx, cy := p.ClosestPoint(5, 0)
	// Closest vertex to (5, 0) is (4, 0)
	if math.Abs(cx-4) > epsilon || math.Abs(cy-0) > epsilon {
		t.Errorf("ClosestPoint(5, 0) expected (4, 0), got (%v, %v)", cx, cy)
	}
}
