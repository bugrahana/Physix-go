# Documentation 

- [Introduction](#introduction)
- [Installation](#installation)
- [Project Architecture & `main.go`](#project-architecture--maingo)
- [Vectors](#vectors)
- [Rigid Body](#rigid-body)
- [Collision Detection](#collision-detection)
- [Springs](#springs)

---

## Introduction 
Physix-go is a simple, highly-optimized, and flexible 2D physics engine written in Go. It provides comprehensive functions to perform physics calculations efficiently, ranging from standard rigid body collision bounds to massive particle-based simulations.

*Note: Polygon-based physics functions are currently experimental. Active development and performance assurances are focused heavily on particle, rectangular, and circular methodologies.*

---

## Installation

First, ensure you have installed [Go](https://go.dev/doc/install).

```bash
go get github.com/rudransh61/Physix-go@v1.0.0  
```

Alternatively, you can clone the repository to test out the internal library examples directly:

```bash
git clone https://github.com/rudransh61/Physix-go.git
```

If everything is configured correctly, running the `main.go` root file will execute the heavily optimized graphical physics showcase:
```bash
go run main.go
```

---

## Project Architecture & `main.go`

Physix-go strongly strictly separates the mathematical **Core Engine** from the visual **Frontend Implementation**. For a comprehensive breakdown of the engine directories, refer to [CODE_STRUCTURE.md](./CODE_STRUCTURE.md).

### The Role of `main.go`
The root `main.go` file acts as the ultimate reference implementation of how to integrate the internal engine. 
- It intentionally contains **zero** explicit physics formulas. 
- Instead, it acts purely as a graphical presentation layer powered by [ebiten/v2](https://github.com/hajimehoshi/ebiten).
- It handles rendering, input tracking, array data allocation (up to 100,000 entities), and calls the engine packages `pkg/particle` and `pkg/broadphase` strictly to resolve the complex grid mathematics.
- New contributors should analyze `main.go` to learn how to safely hook external UI frameworks into Physix-go.

---

## Vectors

A `Vector` is a mathematical object defining both scalar properties `X` and `Y`. It is the operational foundation of the entire library, used continuously to store positions, bounds, forces, and kinetic acceleration limits.

```go
import "github.com/rudransh61/Physix-go/pkg/vector"

func main() {
    // Standard initialization
    vec1 := vector.Vector{X: 1, Y: 2}

    // Constructor initialization
    vec2 := vector.NewVector(1, 2)
}
```

### Operations on Vectors
```go
vec_add       = vec1.Add(vec2)
vec_sub       = vec1.Sub(vec2)
vec_dot       = vec1.Dot(vec2)            // Inner Product
vec_scale     = vec1.Scale(2)             // Scale by float
vec_magnitude = vec1.Magnitude()          // Length limit
vec_norm      = vec1.Normalize()
vec_dist      = vector.Distance(vec1, vec2)
```

---

## Rigid Body

A `RigidBody` is an entity that inherently possesses physical traits (Mass, Form) and actively reacts to external forces via Velocity and Position fields.

```go
import "github.com/rudransh61/Physix-go/pkg/rigidbody"

// Instantiate Body
body := &rigidbody.RigidBody{
    Position: vector.Vector{X: 50, Y: 50},
    Velocity: vector.Vector{X: 30, Y: 20},
    Mass:     1,
    
    Shape:    "Circle",      // Accepts "Circle" or "Rectangle"
    Radius:   10,            // Required if using "Circle"
    Width:    20,            // Required if using "Rectangle"
    Height:   30,            // Required if using "Rectangle"

    IsMovable: true,         // Toggle immutability (e.g. floors/static platforms)
}
```

### Applying Dynamics
To advance the internal velocity parameters based on active forces, utilize the integration methodology governed by the `dynamics/physics` engine core. This applies physics stepwise across `dt` seconds per frame.

```go
import "github.com/rudransh61/Physix-go/dynamics/physics"

dt := 0.1
forceVector := vector.Vector{X: 0, Y: 9.8} // Gravity
physix.ApplyForce(body, forceVector, dt) 
```

---

## Collision Detection
Collision functions determine bounding intersections deterministically over overlapping ranges. Mathematical solutions for overlap prevention are housed centrally in `github.com/rudransh61/Physix-go/dynamics/collision`.

### Querying Overlaps
```go
// Returns boolean resolutions
is_hit1 := collision.CircleCollided(ball1, ball2)
is_hit2 := collision.RectangleCollided(rect1, rect2)

// Strict order enforcement: (circle, rectangle)
is_hit3 := collision.CircleRectangleCollided(circle, rect)
```

### Collision Responses (Narrow-phase)
If an overlap actively occurs, invoke structural constraints to prevent body piercing.

```go
collision.PreventCircleOverlap(ball1, ball2)
collision.PreventRectangleOverlap(rect1, rect2)
collision.PreventCircleRectangleOverlap(circle, rect) // Order strict
```

**Note**: `Prevent*Overlap` algorithms actively separate entangled properties cleanly but *do not manipulate final kinetic limits*. Immediately chain `BounceOnCollision` to natively execute inverse momentum.

```go
restitution := 1.0 // Velocity reflection threshold
collision.BounceOnCollision(ball1, ball2, restitution)
```

---

## Springs

The engine supports physical soft-body dynamics using Hooke's internal methodologies. Linking two elements via Spring naturally limits tension and controls structural elasticity over time.

```go
import "github.com/rudransh61/Physix-go/pkg/spring"

stiffness := 10.0
damping := 0.5
restLength := 150.0

// RestLength parameter is optional; default calculations resolve length implicitly
tether := spring.NewSpring(ballA, ballB, stiffness, damping, restLength)

// Frame execution applying mutual structural tension
tether.ApplyForce() 
```

---
*For a comprehensively deployed suite of working simulations and architectural physics edge cases, refer externally to the code structures located fundamentally in the `/examples` repository space.*
