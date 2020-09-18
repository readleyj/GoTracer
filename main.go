package main

import (
	"gotracer/internal"
	"image"
	"image/png"
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
	canvasPixels := 1024
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

	writeToPng(canvas, "sphere.png")
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

	camera := internal.NewCamera(1920, 1080, 0.7854)
	camera.Transform = internal.ViewTransform(internal.NewPoint(-3, 1, 2.5), internal.NewPoint(0, 0.5, 0), internal.NewVector(0, 1, 0))

	light := internal.NewAreaLight(
		internal.NewPoint(-1, 2, 4),
		internal.NewVector(2, 0, 0),
		10,
		internal.NewVector(0, 2, 0),
		10,
		internal.NewColor(1.5, 1.5, 1.5),
	)
	light.Jitter = true

	cube := internal.NewCube()
	cube.HasShadow = false
	cubeMaterial := internal.NewDefaultMaterial()
	cubeMaterial.SetColor(internal.NewColor(1.5, 1.5, 1.5))
	cubeMaterial.Ambient = 1.0
	cubeMaterial.Diffuse = 0.0
	cubeMaterial.Specular = 0.0
	cube.SetMaterial(cubeMaterial)
	cube.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(0, 3, 4),
			internal.Scale(1, 1, 0.01),
		),
	)

	plane := internal.NewPlane()
	planeMaterial := internal.NewDefaultMaterial()
	planeMaterial.SetColor(internal.NewColor(1, 1, 1))
	planeMaterial.Ambient = 0.025
	planeMaterial.Diffuse = 0.67
	planeMaterial.Specular = 0.0
	plane.SetMaterial(planeMaterial)

	sphere1 := internal.NewSphere()
	sphereMaterial1 := internal.NewDefaultMaterial()
	sphereMaterial1.SetColor(internal.NewColor(1, 0, 0))
	sphereMaterial1.Ambient = 0.1
	sphereMaterial1.Diffuse = 0.6
	sphereMaterial1.Reflective = 0.3
	sphereMaterial1.Specular = 0.0
	sphere1.SetMaterial(sphereMaterial1)
	sphere1.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(0.5, 0.5, 0),
			internal.Scale(0.5, 0.5, 0.5),
		),
	)

	sphere2 := internal.NewSphere()
	sphereMaterial2 := internal.NewDefaultMaterial()
	sphereMaterial2.SetColor(internal.NewColor(0.5, 0.5, 1))
	sphereMaterial2.Ambient = 0.1
	sphereMaterial2.Diffuse = 0.6
	sphereMaterial2.Reflective = 0.3
	sphereMaterial2.Specular = 0.0
	sphere2.SetMaterial(sphereMaterial2)
	sphere2.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(-0.25, 0.33, 0),
			internal.Scale(0.33, 0.33, 0.33),
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
			internal.MatrixMultiply(
				internal.Translate(-5, 0, 0),
				internal.RotateZ(1.5708),
			),
			internal.RotateY(1.5708),
		),
	)
	westWall.SetMaterial(wallMaterial)

	eastWall := internal.NewPlane()
	eastWall.SetTransform(
		internal.MatrixMultiply(
			internal.MatrixMultiply(
				internal.Translate(5, 0, 0),
				internal.RotateZ(1.5708),
			),
			internal.RotateY(1.5708),
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

	objects := []internal.Shape{
		floor,
		ceiling,
		northWall,
		southWall,
		eastWall,
		westWall,
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

func renderTable() {
	world := internal.NewWorld()

	camera := internal.NewCamera(1920, 1080, 0.785)
	camera.Transform = internal.ViewTransform(internal.NewPoint(8, 6, -8), internal.NewPoint(0, 3, 0), internal.NewVector(0, 1, 0))

	light := internal.NewPointLight(internal.NewPoint(0, 6.9, -5), internal.NewColor(1, 1, 0.9))

	floor := internal.NewCube()
	floorMaterial := internal.NewDefaultMaterial()
	floorPattern := internal.NewCheckersPattern(
		internal.NewColor(0, 0, 0),
		internal.NewColor(0.25, 0.25, 0.25),
	)
	floorPattern.SetTransform(internal.Scale(0.07, 0.07, 0.07))
	floorMaterial.Ambient = 0.25
	floorMaterial.Diffuse = 0.7
	floorMaterial.Specular = 0.9
	floorMaterial.Shininess = 300
	floorMaterial.Reflective = 0.1
	floorMaterial.SetPattern(floorPattern)
	floor.SetMaterial(floorMaterial)
	floor.SetTransform(
		internal.MatrixMultiply(
			internal.Scale(20, 7, 20),
			internal.Translate(0, 1, 0),
		),
	)

	wall := internal.NewCube()
	wallMaterial := internal.NewDefaultMaterial()
	wallPattern := internal.NewCheckersPattern(
		internal.NewColor(0.4863, 0.3765, 0.2941),
		internal.NewColor(0.3725, 0.2902, 0.2275),
	)
	wallPattern.SetTransform(internal.Scale(0.05, 20, 0.05))
	wallMaterial.Ambient = 0.1
	wallMaterial.Diffuse = 0.7
	wallMaterial.Specular = 0.9
	wallMaterial.Shininess = 300
	wallMaterial.Reflective = 0.1
	wallMaterial.SetPattern(wallPattern)
	wall.SetMaterial(wallMaterial)
	wall.SetTransform(internal.Scale(10, 10, 10))

	tableTop := internal.NewCube()
	tableMaterial := internal.NewDefaultMaterial()
	tablePattern := internal.NewStripePattern(
		internal.NewColor(0.5529, 0.4235, 0.3255),
		internal.NewColor(0.6588, 0.5098, 0.4000),
	)
	tablePattern.SetTransform(
		internal.MatrixMultiply(
			internal.Scale(0.05, 0.05, 0.05),
			internal.RotateY(0.1),
		),
	)
	tableMaterial.Ambient = 0.1
	tableMaterial.Diffuse = 0.7
	tableMaterial.Specular = 0.9
	tableMaterial.Shininess = 300
	tableMaterial.Reflective = 0.2
	tableMaterial.SetPattern(tablePattern)
	tableTop.SetMaterial(tableMaterial)
	tableTop.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(0, 3.1, 0),
			internal.Scale(3, 0.1, 2),
		),
	)

	legMaterial := internal.NewDefaultMaterial()
	legMaterial.SetColor(internal.NewColor(0.5529, 0.4235, 0.3255))
	legMaterial.Ambient = 0.2
	legMaterial.Diffuse = 0.7

	leg1 := internal.NewCube()
	leg1.SetMaterial(legMaterial)
	leg1.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(2.7, 1.5, -1.7),
			internal.Scale(0.1, 1.5, 0.1),
		),
	)

	leg2 := internal.NewCube()
	leg2.SetMaterial(legMaterial)
	leg2.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(2.7, 1.5, 1.7),
			internal.Scale(0.1, 1.5, 0.1),
		),
	)

	leg3 := internal.NewCube()
	leg3.SetMaterial(legMaterial)
	leg3.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(-2.7, 1.5, -1.7),
			internal.Scale(0.1, 1.5, 0.1),
		),
	)

	leg4 := internal.NewCube()
	leg4.SetMaterial(legMaterial)
	leg4.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(-2.7, 1.5, 1.7),
			internal.Scale(0.1, 1.5, 0.1),
		),
	)

	glassCube := internal.NewCube()
	glassCube.HasShadow = false
	glassCube.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(0, 3.45001, 0),
			internal.MatrixMultiply(
				internal.RotateY(0.2),
				internal.Scale(0.25, 0.25, 0.25),
			),
		),
	)
	glassCubeMaterial := internal.NewDefaultMaterial()
	glassCubeMaterial.SetColor(internal.NewColor(1, 1, 0.8))
	glassCubeMaterial.Ambient = 0.0
	glassCubeMaterial.Diffuse = 0.3
	glassCubeMaterial.Specular = 0.9
	glassCubeMaterial.Shininess = 300
	glassCubeMaterial.Reflective = 0.7
	glassCubeMaterial.Transparency = 0.7
	glassCubeMaterial.RefractiveIndex = 1.5
	glassCube.SetMaterial(glassCubeMaterial)

	cube1 := internal.NewCube()
	cube1.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(1, 3.35, -0.9),
			internal.MatrixMultiply(
				internal.RotateY(-0.4),
				internal.Scale(0.15, 0.15, 0.15),
			),
		),
	)
	cube1Material := internal.NewDefaultMaterial()
	cube1Material.SetColor(internal.NewColor(1, 0.5, 0.5))
	cube1Material.Reflective = 0.6
	cube1Material.Diffuse = 0.4
	cube1.SetMaterial(cube1Material)

	cube2 := internal.NewCube()
	cube2.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(-1.5, 3.27, 0.3),
			internal.MatrixMultiply(
				internal.RotateY(0.4),
				internal.Scale(0.15, 0.07, 0.15),
			),
		),
	)
	cube2Material := internal.NewDefaultMaterial()
	cube2Material.SetColor(internal.NewColor(1, 1, 0.5))
	cube2.SetMaterial(cube2Material)

	cube3 := internal.NewCube()
	cube3.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(0, 3.25, 1),
			internal.MatrixMultiply(
				internal.RotateY(0.4),
				internal.Scale(0.2, 0.05, 0.05),
			),
		),
	)
	cube3Material := internal.NewDefaultMaterial()
	cube3Material.SetColor(internal.NewColor(0.5, 1, 0.5))
	cube3.SetMaterial(cube3Material)

	cube4 := internal.NewCube()
	cube4.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(-0.6, 3.4, -1),
			internal.MatrixMultiply(
				internal.RotateY(0.8),
				internal.Scale(0.05, 0.2, 0.05),
			),
		),
	)
	cube4Material := internal.NewDefaultMaterial()
	cube4Material.SetColor(internal.NewColor(0.5, 0.5, 1))
	cube4.SetMaterial(cube4Material)

	cube5 := internal.NewCube()
	cube5.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(2, 3.4, 1),
			internal.MatrixMultiply(
				internal.RotateY(0.8),
				internal.Scale(0.05, 0.2, 0.05),
			),
		),
	)
	cube5Material := internal.NewDefaultMaterial()
	cube5Material.SetColor(internal.NewColor(0.5, 1, 1))
	cube5.SetMaterial(cube5Material)

	frame1 := internal.NewCube()
	frame1.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(-10, 4, 1),
			internal.Scale(0.05, 1, 1),
		),
	)
	frame1Material := internal.NewDefaultMaterial()
	frame1Material.SetColor(internal.NewColor(0.7098, 0.2471, 0.2196))
	frame1Material.Diffuse = 0.6
	frame1.SetMaterial(frame1Material)

	frame2 := internal.NewCube()
	frame2.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(-10, 3.4, 2.7),
			internal.Scale(0.05, 0.4, 0.4),
		),
	)
	frame2Material := internal.NewDefaultMaterial()
	frame2Material.SetColor(internal.NewColor(0.2667, 0.2706, 0.6902))
	frame2Material.Diffuse = 0.6
	frame2.SetMaterial(frame2Material)

	frame3 := internal.NewCube()
	frame3.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(-10, 4.6, 2.7),
			internal.Scale(0.05, 0.4, 0.4),
		),
	)
	frame3Material := internal.NewDefaultMaterial()
	frame3Material.SetColor(internal.NewColor(0.3098, 0.5961, 0.3098))
	frame3Material.Diffuse = 0.6
	frame3.SetMaterial(frame3Material)

	frame4 := internal.NewCube()
	frame4.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(-2, 3.5, 9.95),
			internal.Scale(5, 1.5, 0.05),
		),
	)
	frame4Material := internal.NewDefaultMaterial()
	frame4Material.SetColor(internal.NewColor(0.3882, 0.2627, 0.1882))
	frame4Material.Diffuse = 0.7
	frame4.SetMaterial(frame4Material)

	mirror := internal.NewCube()
	mirror.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(-2, 3.5, 9.95),
			internal.Scale(4.8, 1.4, 0.06),
		),
	)
	mirrorMaterial := internal.NewDefaultMaterial()
	mirrorMaterial.SetColor(internal.NewColor(0, 0, 0))
	mirrorMaterial.Diffuse = 0.0
	mirrorMaterial.Ambient = 0.0
	mirrorMaterial.Specular = 1.0
	mirrorMaterial.Shininess = 300
	mirrorMaterial.Reflective = 1.0
	mirror.SetMaterial(mirrorMaterial)

	objects := []internal.Shape{
		floor,
		wall,
		tableTop,
		leg1,
		leg2,
		leg3,
		leg4,
		frame1,
		frame2,
		frame3,
		frame4,
		glassCube,
		cube1,
		cube2,
		cube3,
		cube4,
		cube5,
		mirror,
	}
	world.Lights = append(world.Lights, light)
	world.Objects = append(world.Objects, objects...)

	canvas := internal.Render(camera, world)
	writeToPng(canvas, "table.png")
}

func renderCylinders() {
	world := internal.NewWorld()

	camera := internal.NewCamera(1920, 1080, 0.314)
	camera.Transform = internal.ViewTransform(internal.NewPoint(8, 3.5, -9), internal.NewPoint(0, 0.3, 0), internal.NewPoint(0, 1, 0))

	light := internal.NewPointLight(internal.NewPoint(1, 6.9, -4.9), internal.NewColor(1, 1, 1))

	floor := internal.NewPlane()
	floorMaterial := internal.NewDefaultMaterial()
	floorPattern := internal.NewCheckersPattern(
		internal.NewColor(0.5, 0.5, 0.5),
		internal.NewColor(0.75, 0.75, 0.75),
	)
	floorPattern.SetTransform(
		internal.MatrixMultiply(
			internal.RotateY(0.3),
			internal.Scale(0.25, 0.25, 0.25),
		),
	)
	floorMaterial.SetPattern(floorPattern)
	floorMaterial.Ambient = 0.2
	floorMaterial.Diffuse = 0.9
	floorMaterial.Specular = 0.0
	floor.SetMaterial(floorMaterial)

	cylinder1 := internal.NewCylinder()
	cylinder1.Minimum = 0
	cylinder1.Maximum = 0.75
	cylinder1.Closed = true
	cylinder1.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(-1, 0, 1),
			internal.Scale(0.5, 1, 0.5),
		),
	)
	cylinder1Material := internal.NewDefaultMaterial()
	cylinder1Material.SetColor(internal.NewColor(0, 0, 0.6))
	cylinder1Material.Diffuse = 0.1
	cylinder1Material.Specular = 0.9
	cylinder1Material.Shininess = 300
	cylinder1Material.Reflective = 0.9
	cylinder1.SetMaterial(cylinder1Material)

	cylinder2 := internal.NewCylinder()
	cylinder2.Minimum = 0
	cylinder2.Maximum = 0.2
	cylinder2.Closed = false
	cylinder2.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(1, 0, 0),
			internal.Scale(0.8, 1, 0.8),
		),
	)
	cylinder2Material := internal.NewDefaultMaterial()
	cylinder2Material.SetColor(internal.NewColor(1, 1, 0.3))
	cylinder2Material.Ambient = 0.1
	cylinder2Material.Diffuse = 0.8
	cylinder2Material.Specular = 0.9
	cylinder2Material.Shininess = 300
	cylinder2.SetMaterial(cylinder2Material)

	cylinder3 := internal.NewCylinder()
	cylinder3.Minimum = 0
	cylinder3.Maximum = 0.3
	cylinder3.Closed = false
	cylinder3.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(1, 0, 0),
			internal.Scale(0.6, 1, 0.6),
		),
	)
	cylinder3Material := internal.NewDefaultMaterial()
	cylinder3Material.SetColor(internal.NewColor(1, 0.9, 0.4))
	cylinder3Material.Ambient = 0.1
	cylinder3Material.Diffuse = 0.8
	cylinder3Material.Specular = 0.9
	cylinder3Material.Shininess = 300
	cylinder3.SetMaterial(cylinder3Material)

	cylinder4 := internal.NewCylinder()
	cylinder4.Minimum = 0
	cylinder4.Maximum = 0.4
	cylinder4.Closed = false
	cylinder4.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(1, 0, 0),
			internal.Scale(0.4, 1, 0.4),
		),
	)
	cylinder4Material := internal.NewDefaultMaterial()
	cylinder4Material.SetColor(internal.NewColor(1, 0.8, 0.5))
	cylinder4Material.Ambient = 0.1
	cylinder4Material.Diffuse = 0.8
	cylinder4Material.Specular = 0.9
	cylinder4Material.Shininess = 300
	cylinder4.SetMaterial(cylinder4Material)

	cylinder5 := internal.NewCylinder()
	cylinder5.Minimum = 0
	cylinder5.Maximum = 0.5
	cylinder5.Closed = true
	cylinder5.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(1, 0, 0),
			internal.Scale(0.2, 1, 0.2),
		),
	)
	cylinder5Material := internal.NewDefaultMaterial()
	cylinder5Material.SetColor(internal.NewColor(1, 0.7, 0.6))
	cylinder5Material.Ambient = 0.1
	cylinder5Material.Diffuse = 0.8
	cylinder5Material.Specular = 0.9
	cylinder5Material.Shininess = 300
	cylinder5.SetMaterial(cylinder5Material)

	cylinder6 := internal.NewCylinder()
	cylinder6.Minimum = 0
	cylinder6.Maximum = 0.3
	cylinder6.Closed = true
	cylinder6.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(0, 0, -0.75),
			internal.Scale(0.05, 1, 0.05),
		),
	)
	cylinder6Material := internal.NewDefaultMaterial()
	cylinder6Material.SetColor(internal.NewColor(1, 0, 0))
	cylinder6Material.Ambient = 0.1
	cylinder6Material.Diffuse = 0.9
	cylinder6Material.Specular = 0.9
	cylinder6Material.Shininess = 300
	cylinder6.SetMaterial(cylinder6Material)

	cylinder7 := internal.NewCylinder()
	cylinder7.Minimum = 0
	cylinder7.Maximum = 0.3
	cylinder7.Closed = true
	cylinder7.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(0, 0, -2.25),
			internal.MatrixMultiply(
				internal.RotateY(-0.15),
				internal.MatrixMultiply(
					internal.Translate(0, 0, 1.5),
					internal.Scale(0.05, 1, 0.05),
				),
			),
		),
	)
	cylinder7Material := internal.NewDefaultMaterial()
	cylinder7Material.SetColor(internal.NewColor(1, 1, 0))
	cylinder7Material.Ambient = 0.1
	cylinder7Material.Diffuse = 0.9
	cylinder7Material.Specular = 0.9
	cylinder7Material.Shininess = 300
	cylinder7.SetMaterial(cylinder7Material)

	cylinder8 := internal.NewCylinder()
	cylinder8.Minimum = 0
	cylinder8.Maximum = 0.3
	cylinder8.Closed = true
	cylinder8.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(0, 0, -2.25),
			internal.MatrixMultiply(
				internal.RotateY(-0.3),
				internal.MatrixMultiply(
					internal.Translate(0, 0, 1.5),
					internal.Scale(0.05, 1, 0.05),
				),
			),
		),
	)
	cylinder8Material := internal.NewDefaultMaterial()
	cylinder8Material.SetColor(internal.NewColor(0, 1, 0))
	cylinder8Material.Ambient = 0.1
	cylinder8Material.Diffuse = 0.9
	cylinder8Material.Specular = 0.9
	cylinder8Material.Shininess = 300
	cylinder8.SetMaterial(cylinder8Material)

	cylinder9 := internal.NewCylinder()
	cylinder9.Minimum = 0
	cylinder9.Maximum = 0.3
	cylinder9.Closed = true
	cylinder9.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(0, 0, -2.25),
			internal.MatrixMultiply(
				internal.RotateY(-0.45),
				internal.MatrixMultiply(
					internal.Translate(0, 0, 1.5),
					internal.Scale(0.05, 1, 0.05),
				),
			),
		),
	)
	cylinder9Material := internal.NewDefaultMaterial()
	cylinder9Material.SetColor(internal.NewColor(0, 1, 1))
	cylinder9Material.Ambient = 0.1
	cylinder9Material.Diffuse = 0.9
	cylinder9Material.Specular = 0.9
	cylinder9Material.Shininess = 300
	cylinder9.SetMaterial(cylinder9Material)

	glassCylinder := internal.NewCylinder()
	glassCylinder.Minimum = 0.0001
	glassCylinder.Maximum = 0.5
	glassCylinder.Closed = true
	glassCylinder.SetTransform(
		internal.MatrixMultiply(
			internal.Translate(0, 0, -1.5),
			internal.Scale(0.33, 1, 0.33),
		),
	)
	glassCylinderMaterial := internal.NewDefaultMaterial()
	glassCylinderMaterial.SetColor(internal.NewColor(0.25, 0, 0))
	glassCylinderMaterial.Diffuse = 0.1
	glassCylinderMaterial.Specular = 0.9
	glassCylinderMaterial.Shininess = 300
	glassCylinderMaterial.Reflective = 0.9
	glassCylinderMaterial.Transparency = 0.9
	glassCylinderMaterial.RefractiveIndex = 1.5
	glassCylinder.SetMaterial(glassCylinderMaterial)

	objects := []internal.Shape{
		floor,
		cylinder1,
		cylinder2,
		cylinder3,
		cylinder4,
		cylinder5,
		cylinder6,
		cylinder7,
		cylinder8,
		cylinder9,
		glassCylinder,
	}
	world.Lights = append(world.Lights, light)
	world.Objects = append(world.Objects, objects...)

	canvas := internal.Render(camera, world)
	writeToPng(canvas, "cylinders.png")
}

func main() {
	// renderCircle()
	// renderSphere()
	// renderScene()
	// renderRefraction()
	// renderReflectionRefraction()
	// renderTable()
	// renderCylinders()
	renderShadowGlamour()
}

// Adapted from https://github.com/eriklupander/rt/blob/master/main.go
func writeToPng(canvas *internal.Canvas, file string) {
	image := image.NewRGBA(image.Rect(0, 0, canvas.W, canvas.H))
	canvas.ToPNG(image)
	outputFile, _ := os.Create(file)
	defer outputFile.Close()
	png.Encode(outputFile, image)
}
