package main

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rudransh61/Physix-go/dynamics/collision"
	"github.com/rudransh61/Physix-go/pkg/broadphase"
	"github.com/rudransh61/Physix-go/pkg/particle"
	"github.com/rudransh61/Physix-go/pkg/rigidbody"
	"github.com/rudransh61/Physix-go/pkg/vector"
)

var (
	balls           []*particle.PVEBody
	dt              = 0.1
	ticker          *time.Ticker
	initialInterval = time.Second / 5
	center          vector.Vector
	limit           = 10000

	spatialHash *broadphase.SpatialHash
)

const (
	Mass       = 1
	Shape      = "Circle"
	Radius     = 2
	Friction   = 0.899
	Gravity    = 50
	InitRadius = 1000.0
)

var (
	particlesAdded = 0
	maxParticles   = 100000
)

func update() error {
	if particlesAdded < maxParticles {
		select {
		case <-ticker.C:
			for i := 0; i < 14; i++ {
				addParticle()
			}
		default:
		}
	}

	spatialHash.Clear()
	for _, ball := range balls {
		spatialHash.Add(ball, ball.Position)
	}

	for _, ball := range balls {
		gravity := center.Sub(ball.Position).Normalize().Scale(Gravity)
		particle.ApplyForce(ball, gravity, dt)
		ball.Velocity = ball.Velocity.Scale(Friction)
	}

	for i := 0; i < len(balls); i++ {
		ball := balls[i]
		nearbyObjects := spatialHash.Query(ball.Position)
		for _, obj := range nearbyObjects {
			other, ok := obj.(*particle.PVEBody)
			if !ok || other == ball {
				continue
			}
			if collision.CircleCollided(ball.RigidBody, other.RigidBody) {
				particle.ResolveCollision(ball.RigidBody, other.RigidBody, dt)
			}
		}
	}

	return nil
}

func draw(screen *ebiten.Image) {
	for _, ball := range balls {
		ebitenutil.DrawCircle(screen, ball.Position.X, ball.Position.Y, ball.Radius, ball.Color)
	}
}

func main() {
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Star Simulation")

	center = vector.Vector{X: 400, Y: 400}
	ticker = time.NewTicker(initialInterval)

	cellSize := 2.0 * Radius
	screenWidth, screenHeight := ebiten.WindowSize()
	spatialHash = broadphase.NewSpatialHash(cellSize, float64(screenWidth), float64(screenHeight))

	initializeBalls(10000)

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

func initializeBalls(n int) {
	balls = make([]*particle.PVEBody, 0, n)
	for i := 0; i < n; i++ {
		angle := 2 * math.Pi * float64(i) / float64(n)
		radius := rand.Float64() * InitRadius
		x := center.X + radius*math.Cos(angle)
		y := center.Y + radius*math.Sin(angle)
		colorValue := uint8(rand.Int())
		colorValue1 := uint8(rand.Int())
		colorValue2 := uint8(rand.Int())
		ball := &particle.PVEBody{
			RigidBody: &rigidbody.RigidBody{
				Position:  vector.Vector{X: x, Y: y},
				Velocity:  vector.Vector{X: 0, Y: 0},
				Mass:      Mass,
				Shape:     Shape,
				Radius:    Radius,
				IsMovable: true,
			},
			Color: color.RGBA{R: colorValue1, G: colorValue2, B: colorValue, A: 0xff},
			Heat:  100.0,
		}
		balls = append(balls, ball)
	}
}

func addParticle() {
	screenWidth, screenHeight := ebiten.WindowSize()
	x := rand.Float64() * float64(screenWidth)
	y := rand.Float64() * float64(screenHeight)
	colorValue := uint8(rand.Int())
	colorValue1 := uint8(rand.Int())
	colorValue2 := uint8(rand.Int())
	ball := &particle.PVEBody{
		RigidBody: &rigidbody.RigidBody{
			Position:  vector.Vector{X: x, Y: y},
			Velocity:  vector.Vector{X: 0, Y: 0},
			Mass:      Mass,
			Shape:     Shape,
			Radius:    Radius,
			IsMovable: true,
		},
		Color: color.RGBA{R: colorValue1, G: colorValue2, B: colorValue, A: 0xff},
		Heat:  100.0,
	}
	balls = append(balls, ball)
	particlesAdded++
}

type Game struct{}

func (g *Game) Update() error {
	return update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 800
}
