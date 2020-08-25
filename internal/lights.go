package internal

import "math"

type PointLight struct {
	Position  Tuple
	Intensity Color
}

func NewPointLight(position Tuple, intensity Color) PointLight {
	return PointLight{position, intensity}
}

func Lighting(m Material, light PointLight, point, eyeV, normalV Tuple) Color {
	var ambient, diffuse, specular Color

	effectiveColor := HadamardProduct(m.Color, light.Intensity)
	lightV := Normalize(SubTuples(light.Position, point))
	ambient = ColorScalarMultiply(effectiveColor, m.Ambient)

	lightDotNormal := Dot(lightV, normalV)

	if lightDotNormal < 0 {
		diffuse = NewColor(0, 0, 0)
		specular = NewColor(0, 0, 0)
	} else {
		diffuse = ColorScalarMultiply(effectiveColor, m.Diffuse*lightDotNormal)

		reflectV := Reflect(Negate(lightV), normalV)
		reflectDotEye := Dot(reflectV, eyeV)

		if reflectDotEye <= 0 {
			specular = NewColor(0, 0, 0)
		} else {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = ColorScalarMultiply(light.Intensity, m.Specular*factor)
		}
	}

	return AddColors(AddColors(ambient, diffuse), specular)
}
