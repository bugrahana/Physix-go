package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rudransh61/Physix-go/pkg/vector"
)

// Demonstrates coordinate system transformations:
// - 2D: Cartesian <-> Polar (orbiting dots)
// - 3D: Cartesian <-> Spherical (projected orbiting sphere point)
// - 3D: Cartesian <-> Cylindrical (projected helix)

var (
	time_ float64
	dt    = 0.02
)

const (
	cx = 400.0 // screen center X
	cy = 300.0 // screen center Y
)

func update() error {
	time_ += dt
	return nil
}

// project3D does a simple perspective projection from 3D to 2D screen coords.
func project3D(v vector.Vector3, offsetX, offsetY float64) (float64, float64) {
	dist := 6.0
	scale := 120.0 / (v.Z + dist)
	return v.X*scale + offsetX, -v.Y*scale + offsetY
}

func draw(screen *ebiten.Image) {
	// --- 2D: Polar coordinates ---
	// An object orbits using polar coords, converted back to Cartesian for rendering.
	polarCx, polarCy := 200.0, 200.0
	numDots := 8
	for i := 0; i < numDots; i++ {
		// Define position in polar coordinates
		r := 60.0 + 20.0*math.Sin(time_*2+float64(i))
		theta := time_ + 2*math.Pi*float64(i)/float64(numDots)

		// Convert polar -> Cartesian
		pos := vector.PolarToCartesian(r, theta)

		// Verify round-trip
		rBack, thetaBack := vector.CartesianToPolar(pos)
		_ = rBack
		_ = thetaBack

		sx := pos.X + polarCx
		sy := pos.Y + polarCy
		clr := color.RGBA{R: 0xff, G: uint8(180 + i*10), B: 0x40, A: 0xff}
		ebitenutil.DrawCircle(screen, sx, sy, 5, clr)
	}
	ebitenutil.DrawCircle(screen, polarCx, polarCy, 3, color.White)
	drawLabel(screen, polarCx-60, polarCy+90, "Polar (2D)")

	// --- 3D: Spherical coordinates ---
	// Points on a sphere defined in spherical coords, projected to 2D.
	sphereCx, sphereCy := 600.0, 200.0
	sphereR := 2.0
	for lat := 0; lat < 12; lat++ {
		for lon := 0; lon < 12; lon++ {
			theta := math.Pi * float64(lat) / 11          // polar angle [0, pi]
			phi := 2 * math.Pi * float64(lon) / 12        // azimuthal [0, 2pi]
			phi += time_ * 0.5                             // spin the sphere

			// Spherical -> Cartesian
			cart := vector.SphericalToCartesian(sphereR, theta, phi)

			// Verify round-trip
			rB, tB, pB := vector.CartesianToSpherical(cart)
			_ = rB
			_ = tB
			_ = pB

			sx, sy := project3D(cart, sphereCx, sphereCy)
			brightness := uint8(128 + int(127*(cart.Z+sphereR)/(2*sphereR)))
			ebitenutil.DrawCircle(screen, sx, sy, 2, color.RGBA{R: 0x40, G: brightness, B: 0xff, A: 0xff})
		}
	}
	drawLabel(screen, sphereCx-70, sphereCy+90, "Spherical (3D)")

	// --- 3D: Cylindrical coordinates ---
	// A helix defined in cylindrical coords, projected to 2D.
	helixCx, helixCy := 400.0, 450.0
	steps := 80
	var prevSx, prevSy float64
	for i := 0; i < steps; i++ {
		frac := float64(i) / float64(steps-1)
		rho := 1.5
		phi := 4*math.Pi*frac + time_
		z := 3*frac - 1.5

		// Cylindrical -> Cartesian
		cart := vector.CylindricalToCartesian(rho, phi, z)

		// Verify round-trip
		rhoB, phiB, zB := vector.CartesianToCylindrical(cart)
		_ = rhoB
		_ = phiB
		_ = zB

		sx, sy := project3D(cart, helixCx, helixCy)
		if i > 0 {
			ebitenutil.DrawLine(screen, prevSx, prevSy, sx, sy,
				color.RGBA{R: 0xff, G: 0x80, B: uint8(frac * 255), A: 0xff})
		}
		prevSx, prevSy = sx, sy
	}
	drawLabel(screen, helixCx-75, helixCy+60, "Cylindrical (3D)")

	// Info
	ebitenutil.DebugPrint(screen, "Coordinate Transforms Demo")

	// Show live polar values for first dot
	r := 60.0 + 20.0*math.Sin(time_*2)
	theta := time_
	pos := vector.PolarToCartesian(r, theta)
	rBack, thetaBack := vector.CartesianToPolar(pos)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf(
		"Polar dot 0: r=%.1f theta=%.2f rad\nRound-trip:   r=%.1f theta=%.2f rad",
		r, math.Mod(theta, 2*math.Pi), rBack, thetaBack), 10, 30)
}

func drawLabel(screen *ebiten.Image, x, y float64, text string) {
	ebitenutil.DebugPrintAt(screen, text, int(x), int(y))
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Coordinate Transforms Demo")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

type Game struct{}

func (g *Game) Update() error                                    { return update() }
func (g *Game) Draw(screen *ebiten.Image)                        { draw(screen) }
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) { return 800, 600 }
