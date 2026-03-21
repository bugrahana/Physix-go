# Physix-go Code Structure

Physix-go is structured into clearly separated layers: **Core Engine** and **Implementation/Frontend**. This enforces modularity, ensuring that physical mathematical algorithms do not bleed directly into the visual presentation loop.

---

## Directory Map

### 1. Engine Libraries (`pkg/` and `dynamics/`)
These directories contain the pure, rendering-agnostic logic of the engine. They expose APIs for manipulating abstract physical properties across discrete steps in time (`dt`).

- **`pkg/vector`**: Provides foundational 2D mathematical operations (`Add`, `Sub`, `Scale`, `Dot`, `Normalize`). Almost every physical property (Position, Velocity, Force) is built upon this `Vector` struct.
- **`pkg/rigidbody`**: Defines the central `RigidBody` object. This structure maintains the state memory of an entity (`Mass`, `Radius`, `Shape`, `Velocity`, `Force`, `IsMovable`).
- **`pkg/particle`**: Extends standard rigid bodies for particle-based thermodynamics. It defines the `PVEBody` struct to manage kinetic `Heat` translation, alongside utilizing optimized `ApplyForce` and `ResolveCollision` methodologies tailored strictly to massive particle swarms.
- **`pkg/broadphase`**: Implements O(1) grid-based Spatial Hashing (`spatial_hash.go`). It cleverly partitions the 2D plane into indexed cells to drastically reduce the volume of $O(N^2)$ collision intersection checks.
- **`pkg/spring`**: Provides soft-body tension constraints, utilizing Hooke's Law and numerical damping factors to physically chain bodies together.
- **`dynamics/physics`**: Contains external force integrators (`ApplyForce`). These functions translate aggregated mathematical vectors into accelerations, seamlessly resolving velocity and position.
- **`dynamics/collision`**: Handles narrow-phase resolution formulas. Calculates true deterministic overlaps (`RectangleCollided`, `PreventOverlap`, `BounceOnCollision`).

### 2. Implementation Frontend (`examples/` and Base)
These files act strictly as presentation layers. They orchestrate cameras, UI bindings, and UI render loops using the Ebiten graphics library. They strictly reference the Engine to compute logic.

- **`main.go`**: The primary entry point and high-performance visual showcase (the Star Simulation). This application handles array indexing, the infinite update UI loop, inputs, and frame-rate presentation. It structurally offloads spatial hashing and gravity math directly to the internal engine packages, serving as the benchmark example of building a pure implementation frontend.
- **`examples/`**: Contains highly varied physical testing environments. Each file is designed to demonstrate an explicit physics methodology (e.g., `box_mover.go` for collision resolution, `softbody.go` for tension springs, `flappybird.go` for continuous impulses).