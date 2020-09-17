package main

import (
	"gotracer/internal"
	"image"
	"image/png"
	"io/ioutil"
	"math"
	"os"
)

func renderCircle() {
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

	writeToPng(canvas, "circle.png")
}

func renderSphere() {
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
				normal := internal.NormalAt(hit.Object, point, internal.Intersection{})
				eye := internal.Negate(r.Direction)
				color := internal.Lighting(hit.Object.GetMaterial(), internal.NewSphere(), light, point, eye, normal, 1.0)

				canvas.WritePixelAtCoord(x, y, color)
			}
		}
	}

	ioutil.WriteFile("sphere.ppm", []byte(canvas.ToPPM()), 0)
}

func renderScene() {
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

	camera := internal.NewCamera(1920, 1080, math.Pi/3)
	camera.Transform = internal.ViewTransform(
		internal.NewPoint(0, 1.5, -5),
		internal.NewPoint(0, 1, 0),
		internal.NewVector(0, 1, 0),
	)

	canvas := internal.Render(camera, world)
	writeToPng(canvas, "scene.png")
}

func renderRefraction() {
	world := internal.NewWorld()

	camera := internal.NewCamera(1920, 1080, 0.5)
	camera.Transform = internal.ViewTransform(internal.NewPoint(-4.5, 0.85, -4), internal.NewPoint(0, 0.85, 0), internal.NewVector(0, 1, 0))

	wallMaterial := internal.NewDefaultMaterial()
	pattern := internal.NewCheckersPattern(internal.NewColor(0, 0, 0), internal.NewColor(0.75, 0.75, 0.74))
	pattern.SetTransform(internal.Scale(0.5, 0.5, 0.5))
	wallMaterial.SetPattern(pattern)
	wallMaterial.Specular = 0.0

	floor := internal.NewPlane()
	floor.SetTransform(internal.RotateY(0.31415))
	floorMaterial := internal.NewDefaultMaterial()
	floorMaterial.SetPattern(pattern)
	floorMaterial.Ambient = 0.5
	floorMaterial.Diffuse = 0.4
	floorMaterial.Specular = 0.8
	floorMaterial.Reflective = 0.1
	floor.SetMaterial(floorMaterial)

	ceil := internal.NewPlane()
	ceil.SetTransform(internal.Translate(0, 5, 0))
	ceilMaterial := internal.NewDefaultMaterial()
	ceilPattern := internal.NewCheckersPattern(internal.NewColor(0.85, 0.85, 0.85), internal.NewColor(1, 1, 1))
	ceilPattern.SetTransform(internal.Scale(0.2, 0.2, 0.2))
	ceilMaterial.SetPattern(ceilPattern)
	ceilMaterial.Ambient = 0.5
	ceilMaterial.Specular = 0
	ceil.SetMaterial(ceilMaterial)

	westWall := internal.NewPlane()
	westWall.SetTransform(
		internal.MatrixMultiply(
			internal.MatrixMultiply(
				internal.Translate(-5, 0, 0),
				internal.RotateZ(1.5708),
			),
			internal.RotateY(1.5708),
		))

	westWall.SetMaterial(wallMaterial)

	eastWall := internal.NewPlane()
	eastWall.SetTransform(
		internal.MatrixMultiply(
			internal.MatrixMultiply(
				internal.Translate(5, 0, 0),
				internal.RotateZ(1.5708),
			),
			internal.RotateY(1.5708),
		))
	eastWall.SetMaterial(wallMaterial)

	northWall := internal.NewPlane()
	northWall.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(0, 0, 5),
			internal.RotateX(1.5708),
		),
	)
	northWall.SetMaterial(wallMaterial)

	southWall := internal.NewPlane()
	southWall.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(0, 0, -5),
			internal.RotateX(1.5708),
		),
	)
	southWall.SetMaterial(wallMaterial)

	ball1 := internal.NewSphere()
	ball1.SetTransform(internal.Translate(4, 1, 4))
	material1 := internal.NewDefaultMaterial()
	material1.SetColor(internal.NewColor(0.8, 0.1, 0.3))
	material1.Specular = 0
	ball1.SetMaterial(material1)

	ball2 := internal.NewSphere()
	ball2.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(4.6, 0.4, 2.9),
			internal.Scale(0.4, 0.4, 0.4),
		),
	)
	material2 := internal.NewDefaultMaterial()
	material2.SetColor(internal.NewColor(0.1, 0.8, 0.2))
	material2.Shininess = 200
	ball2.SetMaterial(material2)

	ball3 := internal.NewSphere()
	ball3.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(2.6, 0.6, 4.4),
			internal.Scale(0.6, 0.6, 0.6),
		),
	)
	material3 := internal.NewDefaultMaterial()
	material3.SetColor(internal.NewColor(0.2, 0.1, 0.8))
	material3.Shininess = 10
	material3.Specular = 0.4
	ball3.SetMaterial(material3)

	glassBall := internal.NewSphere()
	glassBall.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(0.25, 1, 0),
			internal.Scale(1, 1, 1),
		),
	)

	glassMaterial := internal.NewMaterial(internal.NewColor(0.8, 0.8, 0.9), nil, 0, 0.2, 0.9, 300, 0.0, 0.8, 1.5)
	glassBall.SetMaterial(glassMaterial)

	world.Lights = append(world.Lights, internal.NewPointLight(internal.NewPoint(-4.9, 4.9, 1), internal.NewColor(1, 1, 1)))
	world.Objects = append(world.Objects, ceil, floor, northWall, eastWall, southWall, westWall, ball1, ball2, ball3, glassBall)

	canvas := internal.Render(camera, world)
	writeToPng(canvas, "refraction.png")
}

func renderShadowGlamour() {
	world := internal.NewWorld()

	camera := internal.NewCamera(400, 160, 0.7854)
	camera.Transform = internal.ViewTransform(internal.NewPoint(-3, 1, 2.5), internal.NewPoint(0, 0.5, 0), internal.NewVector(0, 1, 0))

	light := internal.NewAreaLight(
		internal.NewPoint(-1, 2, 4),
		internal.NewVector(2, 0, 0),
		10,
		internal.NewVector(0, 2, 0),
		10,
		internal.NewColor(1.5, 1.5, 1.5),
	)

	cube := internal.NewCube()
	cubeMaterial := internal.Material{}
	cubeMaterial.SetColor(internal.NewColor(1.5, 1.5, 1.5))
	cubeMaterial.Ambient = 1.0
	cubeMaterial.Diffuse = 0.0
	cubeMaterial.Specular = 0.0
	cube.SetMaterial(cubeMaterial)
	cube.SetTransform(
		internal.MatrixMultiply(
			internal.Scale(1, 1, 0.01),
			internal.Translate(0, 3, 4),
		),
	)

	plane := internal.NewPlane()
	planeMaterial := internal.Material{}
	planeMaterial.SetColor(internal.NewColor(1, 1, 1))
	planeMaterial.Ambient = 0.025
	planeMaterial.Diffuse = 0.67
	planeMaterial.Specular = 0.0
	plane.SetMaterial(planeMaterial)

	sphere1 := internal.NewSphere()
	sphereMaterial1 := internal.Material{}
	sphereMaterial1.SetColor(internal.NewColor(1, 0, 0))
	sphereMaterial1.Ambient = 0.1
	sphereMaterial1.Diffuse = 0.6
	sphereMaterial1.Reflective = 0.3
	sphereMaterial1.Specular = 0.0
	sphere1.SetMaterial(cubeMaterial)
	sphere1.SetTransform(
		internal.MatrixMultiply(
			internal.Scale(0.5, 0.5, 0.5),
			internal.Translate(0.5, 0.5, 0),
		),
	)

	sphere2 := internal.NewSphere()
	sphereMaterial2 := internal.NewDefaultMaterial()
	sphereMaterial2.SetColor(internal.NewColor(0.5, 0.5, 1))
	sphereMaterial2.Ambient = 0.1
	sphereMaterial2.Diffuse = 0.6
	sphereMaterial2.Reflective = 0.3
	sphereMaterial2.Specular = 0.0
	sphere2.SetMaterial(cubeMaterial)
	sphere2.SetTransform(
		internal.MatrixMultiply(
			internal.Scale(0.33, 0.33, 0.33),
			internal.Translate(-0.25, 0.33, 0),
		),
	)

	world.Lights = append(world.Lights, light)
	world.Objects = append(world.Objects, cube, plane, sphere1, sphere2)

	canvas := internal.Render(camera, world)
	writeToPng(canvas, "shadow_glamour.png")
}

func renderReflectionRefraction() {
	world := internal.NewWorld()

	camera := internal.NewCamera(1920, 1080, 1.152)
	camera.Transform = internal.ViewTransform(
		internal.NewPoint(-2.6, 1.5, -3.9),
		internal.NewPoint(-0.6, 1, -0.8),
		internal.NewPoint(0, 1, 0),
	)

	light := internal.NewPointLight(internal.NewPoint(-4.9, 4.9, -1), internal.NewColor(1, 1, 1))

	wallMaterial := internal.NewDefaultMaterial()
	wallPattern := internal.NewStripePattern(
		internal.NewColor(0.45, 0.45, 0.45),
		internal.NewColor(0.55, 0.55, 0.55),
	)
	wallPattern.SetTransform(
		internal.MatrixMultiply(
			internal.RotateY(1.5708),
			internal.Scale(0.25, 0.25, 0.25),
		),
	)
	wallMaterial.Ambient = 0.0
	wallMaterial.Diffuse = 0.4
	wallMaterial.Specular = 0.0
	wallMaterial.Reflective = 0.3
	wallMaterial.SetPattern(wallPattern)

	floor := internal.NewPlane()
	floor.SetTransform(internal.RotateY(0.31415))
	floorMaterial := internal.NewDefaultMaterial()
	floorPattern := internal.NewCheckersPattern(
		internal.NewColor(0.35, 0.35, 0.35),
		internal.NewColor(0.65, 0.65, 0.65),
	)
	floorMaterial.SetPattern(floorPattern)
	floorMaterial.Specular = 0.0
	floorMaterial.Reflective = 0.4
	floor.SetMaterial(floorMaterial)

	ceiling := internal.NewPlane()
	ceiling.SetTransform(internal.Translate(0, 5, 0))
	ceilingMaterial := internal.NewDefaultMaterial()
	ceilingMaterial.Ambient = 0.3
	ceilingMaterial.Specular = 0.0
	ceilingMaterial.SetColor(internal.NewColor(0.8, 0.8, 0.8))
	ceiling.SetMaterial(ceilingMaterial)

	westWall := internal.NewPlane()
	westWall.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(-5, 0, 0),
			internal.MatrixMultiply(
				internal.RotateZ(1.5708),
				internal.RotateY(1.5708),
			),
		),
	)
	westWall.SetMaterial(wallMaterial)

	eastWall := internal.NewPlane()
	eastWall.SetTransform(
		internal.MatrixMultiply(
			internal.MatrixMultiply(
				internal.RotateY(1.5708),
				internal.RotateZ(1.5708),
			),
			internal.Translate(5, 0, 0),
		),
	)
	eastWall.SetMaterial(wallMaterial)

	northWall := internal.NewPlane()
	northWall.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(0, 0, 5),
			internal.RotateX(1.5708),
		),
	)
	northWall.SetMaterial(wallMaterial)

	southWall := internal.NewPlane()
	southWall.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(0, 0, -5),
			internal.RotateX(1.5708),
		),
	)
	southWall.SetMaterial(wallMaterial)

	ball1 := internal.NewSphere()
	ball1.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(4.6, 0.4, 1),
			internal.Scale(0.4, 0.4, 0.4),
		),
	)
	ball1Material := internal.NewDefaultMaterial()
	ball1Material.SetColor(internal.NewColor(0.8, 0.5, 0.3))
	ball1Material.Shininess = 50
	ball1.SetMaterial(ball1Material)

	ball2 := internal.NewSphere()
	ball2.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(4.7, 0.3, 0.4),
			internal.Scale(0.3, 0.3, 0.3),
		),
	)
	ball2Material := internal.NewDefaultMaterial()
	ball2Material.SetColor(internal.NewColor(0.9, 0.4, 0.5))
	ball2Material.Shininess = 50
	ball2.SetMaterial(ball2Material)

	ball3 := internal.NewSphere()
	ball3.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(-1, 0.5, 4.5),
			internal.Scale(0.5, 0.5, 0.5),
		),
	)
	ball3Material := internal.NewDefaultMaterial()
	ball3Material.SetColor(internal.NewColor(0.4, 0.9, 0.6))
	ball3Material.Shininess = 50
	ball3.SetMaterial(ball3Material)

	ball4 := internal.NewSphere()
	ball4.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(-1.7, 0.3, 4.7),
			internal.Scale(0.3, 0.3, 0.3),
		),
	)
	ball4Material := internal.NewDefaultMaterial()
	ball4Material.SetColor(internal.NewColor(0.4, 0.6, 0.9))
	ball4Material.Shininess = 50
	ball4.SetMaterial(ball4Material)

	ball5 := internal.NewSphere()
	ball5.SetTransform(
		internal.Translate(-0.6, 1, 0.6),
	)
	ball5Material := internal.NewDefaultMaterial()
	ball5Material.SetColor(internal.NewColor(1, 0.3, 0.2))
	ball5Material.Specular = 0.4
	ball5Material.Shininess = 5
	ball5.SetMaterial(ball5Material)

	ball6 := internal.NewSphere()
	ball6.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(0.6, 0.7, -0.6),
			internal.Scale(0.7, 0.7, 0.7),
		),
	)
	ball6Material := internal.NewDefaultMaterial()
	ball6Material.SetColor(internal.NewColor(0, 0, 0.2))
	ball6Material.Ambient = 0.0
	ball6Material.Diffuse = 0.4
	ball6Material.Specular = 0.9
	ball6Material.Shininess = 300
	ball6Material.Reflective = 0.9
	ball6Material.Transparency = 0.9
	ball6Material.RefractiveIndex = 1.5
	ball6.SetMaterial(ball6Material)

	ball7 := internal.NewSphere()
	ball7.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(-0.7, 0.5, -0.8),
			internal.Scale(0.5, 0.5, 0.5),
		),
	)
	ball7Material := internal.Material{}
	ball7Material.SetColor(internal.NewColor(0, 0.2, 0))
	ball7Material.Ambient = 0.0
	ball7Material.Diffuse = 0.4
	ball7Material.Specular = 0.9
	ball7Material.Shininess = 300
	ball7Material.Reflective = 0.9
	ball7Material.Transparency = 0.9
	ball7Material.RefractiveIndex = 1.5
	ball7.SetMaterial(ball7Material)

	// Walls rendered incorrectly for some reason
	objects := []internal.Shape{
		floor,
		ceiling,
		// northWall,
		// southWall,
		// eastWall,
		// westWall,
		ball1,
		ball2,
		ball3,
		ball4,
		ball5,
		ball6,
		ball7,
	}
	world.Lights = append(world.Lights, light)
	world.Objects = append(world.Objects, objects...)

	canvas := internal.Render(camera, world)
	writeToPng(canvas, "reflection_refraction.png")
}

func main() {
	// renderCircle()
	// renderSphere()
	// renderScene()
	// renderRefraction()
	renderReflectionRefraction()
}

// Adapted from https://github.com/eriklupander/rt/blob/master/main.go
func writeToPng(canvas *internal.Canvas, file string) {
	image := image.NewRGBA(image.Rect(0, 0, canvas.W, canvas.H))
	canvas.ToPNG(image)
	outputFile, _ := os.Create(file)
	defer outputFile.Close()
	png.Encode(outputFile, image)
}
