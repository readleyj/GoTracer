package main

import (
	"gotracer/internal"
	"io/ioutil"
	"math"
)

func RenderCircle() {
	rayOrigin := internal.NewPoint(0.0, 0.0, -5.0)
	wallZ := 10.0

	wallSize := 7.0
	canvasPixels := 500
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2

	canvas := internal.NewCanvas(canvasPixels, canvasPixels)
	color := internal.NewColor(1, 0, 0)
	shape := internal.NewSphere()

	for y := 0; y < canvasPixels-1; y++ {
		worldY := half - pixelSize*float64(y)

		for x := 0; x < canvasPixels-1; x++ {
			worldX := -half + pixelSize*float64(x)

			position := internal.NewPoint(worldX, worldY, wallZ)
			r := internal.NewRay(rayOrigin, internal.SubTuples(position, rayOrigin).Normalize())
			xs := internal.Intersect(shape, r)
			empty := internal.Intersection{}

			if internal.Hit(xs) != empty {
				canvas.WritePixelAtCoord(x, y, color)
			}
		}
	}

	ioutil.WriteFile("circle.ppm", []byte(canvas.ToPPM()), 0)
}

func RenderSphere() {
	rayOrigin := internal.NewPoint(0.0, 0.0, -5.0)
	wallZ := 10.0

	wallSize := 7.0
	canvasPixels := 500
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2

	canvas := internal.NewCanvas(canvasPixels, canvasPixels)
	shape := internal.NewSphere()
	shape.Material.SetColor(internal.NewColor(0.0, 0.2, 1))

	lightPosition := internal.NewPoint(-10, 10, -10)
	lightColor := internal.NewColor(1, 1, 1)
	light := internal.NewPointLight(lightPosition, lightColor)

	for y := 0; y < canvasPixels-1; y++ {
		worldY := half - pixelSize*float64(y)

		for x := 0; x < canvasPixels-1; x++ {
			worldX := -half + pixelSize*float64(x)

			position := internal.NewPoint(worldX, worldY, wallZ)
			r := internal.NewRay(rayOrigin, internal.SubTuples(position, rayOrigin).Normalize())
			xs := internal.Intersect(shape, r)
			empty := internal.Intersection{}

			if hit := internal.Hit(xs); hit != empty {
				point := internal.Position(r, hit.T)
				normal := internal.NormalAt(hit.Object, point)
				eye := internal.Negate(r.Direction)
				color := internal.Lighting(hit.Object.GetMaterial(), light, point, eye, normal, false)

				canvas.WritePixelAtCoord(x, y, color)
			}
		}
	}

	ioutil.WriteFile("sphere.ppm", []byte(canvas.ToPPM()), 0)
}

func RenderScene() {
	world := internal.NewWorld()

	floor := internal.NewSphere()
	floor.SetTransform(internal.Scale(10, 0.1, 10))
	floor.SetMaterial(internal.NewDefaultMaterial())
	floor.Material.SetColor(internal.NewColor(1, 0.9, 0.9))
	floor.Material.Specular = 0
	world.Objects = append(world.Objects, floor)

	leftWall := internal.NewSphere()
	leftTransform := internal.MatrixMultiply(
		internal.MatrixMultiply(
			internal.MatrixMultiply(
				internal.Translate(0, 0, 5),
				internal.RotateY(-math.Pi/4),
			),
			internal.RotateX(math.Pi/2),
		),
		internal.Scale(10, 0.1, 10),
	)
	leftWall.SetTransform(leftTransform)
	world.Objects = append(world.Objects, leftWall)

	rightWall := internal.NewSphere()
	rightTransform := internal.MatrixMultiply(
		internal.MatrixMultiply(
			internal.MatrixMultiply(
				internal.Translate(0, 0, 5),
				internal.RotateY(math.Pi/4),
			),
			internal.RotateX(math.Pi/2),
		),
		internal.Scale(10, 0.1, 10),
	)
	rightWall.SetTransform(rightTransform)
	world.Objects = append(world.Objects, rightWall)

	middle := internal.NewSphere()
	middle.SetTransform(internal.Translate(-0.5, 1, 0.5))
	middle.SetMaterial(internal.NewDefaultMaterial())
	middle.Material.SetColor(internal.NewColor(0.1, 1, 0.5))
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3
	world.Objects = append(world.Objects, middle)

	right := internal.NewSphere()
	right.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(1.5, 0.5, -0.5),
			internal.Scale(0.5, 0.5, 0.5),
		),
	)
	right.SetMaterial(internal.NewDefaultMaterial())
	right.Material.SetColor(internal.NewColor(0.5, 1, 0.1))
	right.Material.Diffuse = 0.7
	right.Material.Specular = 0.3
	world.Objects = append(world.Objects, right)

	left := internal.NewSphere()
	left.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(-1.5, 0.33, -0.75),
			internal.Scale(0.33, 0.33, 0.33),
		),
	)
	left.SetMaterial(internal.NewDefaultMaterial())
	left.Material.SetColor(internal.NewColor(1.0, 0.8, 0.1))
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.3
	world.Objects = append(world.Objects, left)

	world.Lights = append(world.Lights, internal.NewPointLight(internal.NewPoint(-10, 10, -10), internal.NewColor(1, 1, 1)))

	camera := internal.NewCamera(100, 50, math.Pi/3)
	camera.Transform = internal.ViewTransform(
		internal.NewPoint(0, 1.5, -5),
		internal.NewPoint(0, 1, 0),
		internal.NewVector(0, 1, 0),
	)

	canvas := internal.Render(camera, world)
	ioutil.WriteFile("scene.ppm", []byte(canvas.ToPPM()), 0)
}

func main() {
	// RenderCircle()
	// RenderSphere()
	RenderScene()
}
