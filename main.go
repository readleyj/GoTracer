package main

import (
	"gotracer/internal"
	"io/ioutil"
)

func main() {
	rayOrigin := internal.Point(0.0, 0.0, -5.0)
	wallZ := 10.0

	wallSize := 7.0
	canvasPixels := 100
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2

	canvas := internal.CreateCanvas(canvasPixels, canvasPixels)
	color := internal.Color{1, 0, 0}
	shape := internal.MakeSphere()

	for y := 0; y < canvasPixels-1; y++ {
		worldY := half - pixelSize*float64(y)

		for x := 0; x < canvasPixels-1; x++ {
			worldX := -half + pixelSize*float64(x)

			position := internal.Point(worldX, worldY, wallZ)
			r := internal.Ray{rayOrigin, internal.SubTuples(position, rayOrigin).Normalize()}
			xs := internal.Intersect(shape, r)
			empty := internal.Intersection{}

			if internal.Hit(xs) != empty {
				canvas.WritePixelAtCoord(x, y, color)
			}
		}
	}

	ioutil.WriteFile("circle.ppm", []byte(canvas.ToPPM()), 0)
}
