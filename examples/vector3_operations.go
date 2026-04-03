package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rudransh61/Physix-go/pkg/vector"
)

// A spinning 3D wireframe cube rendered using Vector3 operations.

var (
	angle float64
	dt    = 0.02

	// Unit cube vertices defined with Vector3
	cubeVerts = [8]vector.Vector3{
		vector.NewVector3(-1, -1, -1),
		vector.NewVector3(1, -1, -1),
		vector.NewVector3(1, 1, -1),
		vector.NewVector3(-1, 1, -1),
		vector.NewVector3(-1, -1, 1),
		vector.NewVector3(1, -1, 1),
		vector.NewVector3(1, 1, 1),
		vector.NewVector3(-1, 1, 1),
	}

	// Edges as vertex index pairs
	cubeEdges = [12][2]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 0}, // back face
		{4, 5}, {5, 6}, {6, 7}, {7, 4}, // front face
		{0, 4}, {1, 5}, {2, 6}, {3, 7}, // connecting edges
	}
)

// rotateY rotates a Vector3 around the Y axis using Vector3 operations.
func rotateY(v vector.Vector3, a float64) vector.Vector3 {
	return vector.NewVector3(
		v.X*math.Cos(a)+v.Z*math.Sin(a),
		v.Y,
		-v.X*math.Sin(a)+v.Z*math.Cos(a),
	)
}

// rotateX rotates a Vector3 around the X axis.
func rotateX(v vector.Vector3, a float64) vector.Vector3 {
	return vector.NewVector3(
		v.X,
		v.Y*math.Cos(a)-v.Z*math.Sin(a),
		v.Y*math.Sin(a)+v.Z*math.Cos(a),
	)
}

// project projects a 3D point to 2D screen coordinates using perspective division.
func project(v vector.Vector3) (float64, float64) {
	dist := 4.0
	scale := 150.0 / (v.Z + dist)
	return v.X*scale + 400, v.Y*scale + 300
}

func update() error {
	angle += dt
	return nil
}

func draw(screen *ebiten.Image) {
	// Transform and project all vertices
	var projected [8][2]float64
	for i, v := range cubeVerts {
		r := rotateY(v, angle)
		r = rotateX(r, angle*0.7)
		projected[i][0], projected[i][1] = project(r)
	}

	// Draw edges
	for _, e := range cubeEdges {
		ebitenutil.DrawLine(screen,
			projected[e[0]][0], projected[e[0]][1],
			projected[e[1]][0], projected[e[1]][1],
			color.RGBA{R: 0x00, G: 0xff, B: 0xff, A: 0xff},
		)
	}

	// Show cross product info: compute face normal of front face using Cross
	v0 := rotateX(rotateY(cubeVerts[4], angle), angle*0.7)
	v1 := rotateX(rotateY(cubeVerts[5], angle), angle*0.7)
	v3 := rotateX(rotateY(cubeVerts[7], angle), angle*0.7)
	edge1 := v1.Sub(v0)
	edge2 := v3.Sub(v0)
	normal := edge1.Cross(edge2).Normalize()

	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"Vector3 Demo: Spinning Cube\n\n"+
			"Front face normal (Cross product):\n"+
			"  X=%.2f  Y=%.2f  Z=%.2f\n"+
			"  Magnitude=%.2f",
		normal.X, normal.Y, normal.Z, normal.Magnitude(),
	))
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Vector3 Operations - 3D Cube")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

type Game struct{}

func (g *Game) Update() error                                    { return update() }
func (g *Game) Draw(screen *ebiten.Image)                        { draw(screen) }
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) { return 800, 600 }
