package internal

import (
	"math"
	"sort"
)

type Intersection struct {
	T      float64
	U      float64
	V      float64
	Object Shape
}

type Intersections []Intersection

func NewIntersection(T float64, object Shape) Intersection {
	return Intersection{T, 0, 0, object}
}

func NewIntersectionUV(T float64, object Shape, U, V float64) Intersection {
	return Intersection{T, U, V, object}
}

func IntersectEquals(i1, i2 Intersection) bool {
	return i1.T == i2.T && i1.Object.GetID() == i2.Object.GetID()
}

func NewIntersections(intersects ...Intersection) Intersections {
	results := make([]Intersection, len(intersects))
	copy(results[:], intersects)
	sort.Slice(results, func(i, j int) bool {
		return results[i].T < results[j].T
	})
	return results
}

func Hit(intersects Intersections) Intersection {
	var result Intersection

	for _, v := range intersects {
		if v.T > 0 {
			result = v
			break
		}
	}

	return result
}

func ShadeHit(w World, comps Computation, remaining int) Color {
	var surfaceColor Color

	shadowed := IsShadowed(w, comps.OverPoint)

	for _, light := range w.Lights {
		surfaceColor = AddColors(
			surfaceColor,
			Lighting(comps.Object.GetMaterial(), comps.Object, light, comps.OverPoint, comps.EyeV, comps.NormalV, shadowed),
		)
	}

	reflected := ReflectedColor(w, comps, remaining)
	refracted := RefractedColor(w, comps, remaining)

	material := comps.Object.GetMaterial()

	if material.Reflective > 0 && material.Transparency > 0 {
		reflectance := Schlick(comps)

		return AddColors(
			surfaceColor,
			AddColors(
				ColorScalarMultiply(reflected, reflectance),
				ColorScalarMultiply(refracted, (1-reflectance)),
			),
		)
	}

	return AddColors(AddColors(surfaceColor, reflected), refracted)
}

func ColorAt(w World, r Ray, remaining int) Color {
	intersections := IntersectWorld(w, r)
	hit := Hit(intersections)
	empty := Intersection{}

	if hit == empty {
		return NewColor(0, 0, 0)
	}

	comps := PrepareComputations(hit, r, intersections)

	return ShadeHit(w, comps, remaining)
}

func ReflectedColor(w World, comps Computation, remaining int) Color {
	objectMaterial := comps.Object.GetMaterial()

	if remaining <= 0 || objectMaterial.Reflective == 0 {
		return black
	}

	reflectRay := NewRay(comps.OverPoint, comps.ReflectV)
	color := ColorAt(w, reflectRay, remaining-1)

	return ColorScalarMultiply(color, objectMaterial.Reflective)
}

func RefractedColor(w World, comps Computation, remaining int) Color {
	indexRatio := comps.N1 / comps.N2
	cosI := Dot(comps.EyeV, comps.NormalV)
	sin2T := (indexRatio * indexRatio) * (1 - cosI*cosI)

	if remaining <= 0 || sin2T > 1 || comps.Object.GetMaterial().Transparency == 0 {
		return black
	}

	cosT := math.Sqrt(1.0 - sin2T)
	direction := SubTuples(
		TupleScalarMultiply(comps.NormalV, indexRatio*cosI-cosT),
		TupleScalarMultiply(comps.EyeV, indexRatio),
	)
	refractRay := NewRay(comps.UnderPoint, direction)
	transparency := comps.Object.GetMaterial().Transparency
	color := ColorScalarMultiply(ColorAt(w, refractRay, remaining-1), transparency)

	return color
}

func Schlick(comps Computation) float64 {
	cos := Dot(comps.EyeV, comps.NormalV)

	if comps.N1 > comps.N2 {
		ratio := comps.N1 / comps.N2
		sin2T := (ratio * ratio) * (1 - cos*cos)

		if sin2T > 1.0 {
			return 1.0
		}

		cos = math.Sqrt(1.0 - sin2T)
	}

	r0 := math.Pow(((comps.N1 - comps.N2) / (comps.N1 + comps.N2)), 2)

	return r0 + (1-r0)*math.Pow(1-cos, 5)
}
