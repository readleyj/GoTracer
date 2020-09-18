package internal

import (
	"math"
	"math/rand"
)

type LightSource interface {
	GetPosition() Tuple
	GetIntensity() Color
}

type PointLight struct {
	Position  Tuple
	Intensity Color
}

func (light PointLight) GetPosition() Tuple {
	return light.Position
}

func (light PointLight) GetIntensity() Color {
	return light.Intensity
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

type AreaLight struct {
	Position  Tuple
	Intensity Color
	CornerPos Tuple
	UVec      Tuple
	USteps    int
	VVec      Tuple
	VSteps    int
	Samples   int
	Jitter    bool
}

func (light AreaLight) GetPosition() Tuple {
	return light.Position
}

func (light AreaLight) GetIntensity() Color {
	return light.Intensity
}

func (light AreaLight) PointOnLight(u, v int, jitter bool) Tuple {
	var offset Tuple

	if jitter {
		offset = AddTuples(
			TupleScalarMultiply(light.UVec, float64(u)+0.5*rand.Float64()),
			TupleScalarMultiply(light.VVec, float64(v)+0.5*rand.Float64()),
		)
	} else {
		offset = AddTuples(
			TupleScalarMultiply(light.UVec, float64(u)+0.5),
			TupleScalarMultiply(light.VVec, float64(v)+0.5),
		)
	}

	return AddTuples(light.CornerPos, offset)
}

func NewAreaLight(corner Tuple, uvec Tuple, usteps int, vvec Tuple, vsteps int, intensity Color) AreaLight {
	return AreaLight{
		Position: AddTuples(
			corner,
			NewVector(
				(uvec.X+vvec.X)/2,
				(uvec.Y+vvec.Y)/2,
				(uvec.Z+vvec.Z)/2,
			),
		),
		Intensity: intensity,
		CornerPos: corner,
		UVec:      TupleScalarDivide(uvec, float64(usteps)),
		USteps:    usteps,
		VVec:      TupleScalarDivide(vvec, float64(vsteps)),
		VSteps:    vsteps,
		Samples:   usteps * vsteps,
		Jitter:    false,
	}
}

func Lighting(m Material, object Shape, light LightSource, point, eyeV, normalV Tuple, intensity float64) Color {
	var color Color

	if m.Pattern != nil {
		color = PatternAtShape(m.Pattern, object, point)
	} else {
		color = m.Color
	}

	switch light.(type) {
	case AreaLight:
		l := light.(AreaLight)

		var ambient, diffuse, specular, total Color

		effectiveColor := HadamardProduct(color, light.GetIntensity())
		ambient = ColorScalarMultiply(effectiveColor, m.Ambient)

		for u := 0; u < l.USteps; u++ {
			for v := 0; v < l.VSteps; v++ {
				lightV := Normalize(SubTuples(l.PointOnLight(u, v, l.Jitter), point))
				lightDotNormal := Dot(lightV, normalV)

				if lightDotNormal < 0 {
					diffuse = NewColor(0, 0, 0)
					specular = NewColor(0, 0, 0)
				} else {
					diffuse = ColorScalarMultiply(effectiveColor, m.Diffuse*lightDotNormal*intensity)

					reflectV := Reflect(Negate(lightV), normalV)
					reflectDotEye := Dot(reflectV, eyeV)

					if reflectDotEye <= 0 {
						specular = NewColor(0, 0, 0)
					} else {
						factor := math.Pow(reflectDotEye, m.Shininess)
						specular = ColorScalarMultiply(light.GetIntensity(), m.Specular*factor*intensity)
					}
				}

				total = AddColors(total, AddColors(diffuse, specular))
			}
		}

		return AddColors(ambient, ColorScalarMultiply(ColorScalarDivide(total, float64(l.Samples)), intensity))

	case PointLight:
		l := light.(PointLight)

		var ambient, diffuse, specular Color

		effectiveColor := HadamardProduct(color, light.GetIntensity())
		lightV := Normalize(SubTuples(l.GetPosition(), point))
		ambient = ColorScalarMultiply(effectiveColor, m.Ambient)

		lightDotNormal := Dot(lightV, normalV)

		if lightDotNormal < 0 {
			diffuse = NewColor(0, 0, 0)
			specular = NewColor(0, 0, 0)
		} else {
			diffuse = ColorScalarMultiply(effectiveColor, m.Diffuse*lightDotNormal*intensity)

			reflectV := Reflect(Negate(lightV), normalV)
			reflectDotEye := Dot(reflectV, eyeV)

			if reflectDotEye <= 0 {
				specular = NewColor(0, 0, 0)
			} else {
				factor := math.Pow(reflectDotEye, m.Shininess)
				specular = ColorScalarMultiply(light.GetIntensity(), m.Specular*factor*intensity)
			}
		}
		return AddColors(ambient, AddColors(diffuse, specular))
	default:
		panic("Provided light source of unknown type")
	}
}

func IntensityAt(light LightSource, point Tuple, w World) float64 {
	switch light.(type) {
	case PointLight:
		if IsShadowed(w, light.GetPosition(), point) {
			return 0.0
		}

		return 1.0

	case AreaLight:
		light := light.(AreaLight)
		var total float64

		for v := 0; v < light.VSteps; v++ {
			for u := 0; u < light.USteps; u++ {
				lightPos := light.PointOnLight(u, v, light.Jitter)

				if !IsShadowed(w, lightPos, point) {
					total += 1.0
				}
			}
		}

		return total / float64(light.Samples)

	default:
		return 0.0
	}
}
