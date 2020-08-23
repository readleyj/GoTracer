package internal

import "sort"

type Intersection struct {
	T      float64
	Object *Sphere
}

type Intersections []Intersection

func MakeIntersection(T float64, object *Sphere) Intersection {
	return Intersection{T, object}
}

func MakeIntersections(intersects ...Intersection) Intersections {
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
