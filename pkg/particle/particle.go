package particle

import (
	"image/color"

	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
)

// PVEBody extends rigidbody.RigidBody with Heat and Color fields for particle simulations.
type PVEBody struct {
	*rigidbody.RigidBody
	Heat  float64
	Color color.RGBA
}

// ApplyForce applies a force to a PVEBody, calculating its acceleration, velocity, position, and heat.
func ApplyForce(body *PVEBody, force vector.Vector, dt float64) {
	if body.IsMovable {
		// F = ma -> a = F/m
		body.Force = force
		if body.Mass == 0 {
			return
		}
		acceleration := body.Force.Scale(1 / body.Mass)

		// Update velocity
		body.Velocity = body.Velocity.Add(acceleration.Scale(dt))

		// Update position
		body.Position = body.Position.Add(body.Velocity.Scale(dt))

		// Recalculate heat based on kinetic energy
		speed := body.Velocity.Magnitude()
		body.Heat = 0.5 * body.Mass * speed * speed
	}
}

// ResolveCollision handles simple overlap resolution between two rigid body representations of particles.
func ResolveCollision(ball1, ball2 *rigidbody.RigidBody, dt float64) {
	distance := ball1.Position.Sub(ball2.Position)
	distanceMagnitude := distance.Magnitude()
	minimumDistance := ball1.Radius + ball2.Radius

	if distanceMagnitude < minimumDistance {
		moveDirection := distance.Normalize()
		overlap := (minimumDistance - distanceMagnitude) * 5

		// Calculate the repulsive force magnitude based on the overlap
		mag := 10.0
		repulsiveForceMagnitude := overlap * mag
		repulsiveForce := moveDirection.Scale(repulsiveForceMagnitude)

		// Apply the repulsive force to the velocities
		if ball1.Mass != 0 {
			ball1.Velocity = ball1.Velocity.Add(repulsiveForce.Scale(dt / ball1.Mass).Scale(0.9))
		}
		if ball2.Mass != 0 {
			ball2.Velocity = ball2.Velocity.Add(repulsiveForce.Scale(-dt / ball2.Mass).Scale(0.9))
		}

		// Adjust positions slightly to avoid sticking
		correctionFactor := 0.5
		positionCorrection := moveDirection.Scale(correctionFactor * overlap * 5)
		ball1.Position = ball1.Position.Add(positionCorrection)
		ball2.Position = ball2.Position.Sub(positionCorrection)
	}
}
