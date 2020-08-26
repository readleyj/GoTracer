package internal

import "math"

type Camera struct {
	Hsize      int
	Vsize      int
	HalfWidth  float64
	HalfHeight float64
	FOV        float64
	PixelSize  float64
	Transform  Matrix
}

func NewCamera(hsize, vsize int, fov float64) Camera {
	var halfWidth, halfHeight float64
	halfView := math.Tan(fov / 2)
	aspect := float64(hsize) / float64(vsize)

	if aspect >= 1 {
		halfWidth = halfView
		halfHeight = halfView / aspect
	} else {
		halfWidth = halfView * aspect
		halfHeight = halfView
	}

	pixelSize := (halfWidth * 2) / float64(hsize)

	return Camera{
		Hsize:      hsize,
		Vsize:      vsize,
		HalfWidth:  halfWidth,
		HalfHeight: halfHeight,
		FOV:        fov,
		PixelSize:  pixelSize,
		Transform:  NewIdentity4(),
	}
}

func RayForPixel(c Camera, px, py int) Ray {
	xOffset := (float64(px) + 0.5) * c.PixelSize
	yOffset := (float64(py) + 0.5) * c.PixelSize

	worldX := c.HalfWidth - xOffset
	worldY := c.HalfHeight - yOffset

	transformInverse := MatrixInverse(c.Transform)
	pixel := MatrixTupleMultiply(transformInverse, NewPoint(worldX, worldY, -1))
	origin := MatrixTupleMultiply(transformInverse, NewPoint(0, 0, 0))
	direction := Normalize(SubTuples(pixel, origin))

	return NewRay(origin, direction)
}

func Render(c Camera, w World) *Canvas {
	image := NewCanvas(c.Hsize, c.Vsize)

	for y := 0; y < c.Vsize-1; y++ {
		for x := 0; x < c.Hsize-1; x++ {
			ray := RayForPixel(c, x, y)
			color := ColorAt(w, ray)
			image.WritePixelAtCoord(x, y, color)
		}
	}

	return image
}
