package internal

import "math"

type PointLight struct {
	Position  Tuple
	Intensity Color
}

func NewPointLight(position Tuple, intensity Color) PointLight {
	return PointLight{
		Position:  position,
		Intensity: intensity,
	}
}

func PointLightEquals(l1, l2 PointLight) bool {
	return TupleEquals(l1.Position, l2.Position) && ColorEquals(l1.Intensity, l2.Intensity)
}

func Lighting(m Material, object Shape, light PointLight, point, eyeV, normalV Tuple, inShadow bool) Color {
	var color Color

	if m.Pattern != nil {
		color = PatternAtShape(m.Pattern, object, point)
	} else {
		color = m.Color
	}

	var ambient, diffuse, specular Color

	effectiveColor := HadamardProduct(color, light.Intensity)
	lightV := Normalize(SubTuples(light.Position, point))
	ambient = ColorScalarMultiply(effectiveColor, m.Ambient)

	lightDotNormal := Dot(lightV, normalV)

	if lightDotNormal < 0 || inShadow {
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
