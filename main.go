package main

import (
	"gotracer/internal"
	"io/ioutil"
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
				color := internal.Lighting(hit.Object.Material, light, point, eye, normal)

				canvas.WritePixelAtCoord(x, y, color)
			}
		}
	}

	ioutil.WriteFile("sphere.ppm", []byte(canvas.ToPPM()), 0)
}

func main() {
	RenderCircle()
	// RenderSphere()
}
