package internal

import "sort"

type Intersection struct {
	T      float64
	Object *Sphere
}

type Intersections []Intersection

func NewIntersection(T float64, object *Sphere) Intersection {
	return Intersection{T, object}
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

func ShadeHit(w World, comps Computation) Color {
	return Lighting(
		comps.Object.Material, w.Light, comps.Point, comps.EyeV, comps.NormalV,
	)
}

func ColorAt(w World, r Ray) Color {
	intersections := IntersectWorld(w, r)
	hit := Hit(intersections)
	empty := Intersection{}

	if hit == empty {
		return NewColor(0, 0, 0)
	}

	comps := PrepareComputations(hit, r)

	return ShadeHit(w, comps)
}
