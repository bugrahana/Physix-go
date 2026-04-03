package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	physix "github.com/rudransh61/Physix-go/dynamics/physics"
	"github.com/rudransh61/Physix-go/pkg/polygon"
	"github.com/rudransh61/Physix-go/pkg/vector"
)

var (
	poly1 *polygon.Polygon // A falling triangle
	poly2 *polygon.Polygon // A static pentagon (floor)
	dt    = 0.05
)

const (
	Gravity = 20.0
)

func update() error {
	// Apply gravity to the movable polygon
	gravity := vector.Vector{X: 0, Y: Gravity}
	physix.ApplyForcePolygon(poly1, gravity, dt)

	// Simple floor collision: stop the triangle when its centroid reaches the pentagon
	if poly1.Position.Y > poly2.Position.Y-60 {
		poly1.Velocity = vector.Vector{X: 0, Y: 0}
		// Bounce it back up with some horizontal push and spin
		poly1.ApplyImpulse(vector.Vector{X: 15, Y: -80})
		poly1.Torque = 0.3
	}

	return nil
}

func drawPolygon(screen *ebiten.Image, p *polygon.Polygon, clr color.RGBA) {
	n := len(p.Vertices)
	for i := 0; i < n; i++ {
		v1 := p.Vertices[i]
		v2 := p.Vertices[(i+1)%n]
		ebitenutil.DrawLine(screen, v1.X, v1.Y, v2.X, v2.Y, clr)
	}
	// Draw centroid marker
	ebitenutil.DrawCircle(screen, p.Position.X, p.Position.Y, 3, clr)
}

func draw(screen *ebiten.Image) {
	drawPolygon(screen, poly1, color.RGBA{R: 0x00, G: 0xff, B: 0x80, A: 0xff})
	drawPolygon(screen, poly2, color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff})

	ebitenutil.DebugPrint(screen, "Polygon Physics: triangle falling onto a pentagon")
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Polygon Physics")

	// Create a triangle (movable)
	triVerts := []vector.Vector{
		{X: 380, Y: 50},
		{X: 420, Y: 50},
		{X: 400, Y: 10},
	}
	poly1 = polygon.NewPolygon(triVerts, 5, true)

	// Create a pentagon (static floor)
	pentVerts := make([]vector.Vector, 5)
	cx, cy, r := 400.0, 450.0, 80.0
	for i := 0; i < 5; i++ {
		angle := 2*math.Pi*float64(i)/5 - math.Pi/2
		pentVerts[i] = vector.Vector{
			X: cx + r*math.Cos(angle),
			Y: cy + r*math.Sin(angle),
		}
	}
	poly2 = polygon.NewPolygon(pentVerts, 1000, false)

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

type Game struct{}

func (g *Game) Update() error                                    { return update() }
func (g *Game) Draw(screen *ebiten.Image)                        { draw(screen) }
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) { return 800, 600 }
