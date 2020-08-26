package internal

type World struct {
	Light   PointLight
	Objects []*Sphere
}

func NewWorld() World {
	return World{}
}

func NewDefaultWorld() World {
	light := NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1))

	s1 := NewSphere()
	s1.SetMaterial(NewMaterial(
		NewColor(0.8, 1.0, 0.6),
		DefaultMaterial.Ambient,
		0.7,
		0.2,
		DefaultMaterial.Shininess,
	))

	s2 := NewSphere()
	s2.SetTransform(Scale(0.5, 0.5, 0.5))

	return World{
		light,
		[]*Sphere{
			s1,
			s2,
		},
	}
}

func IntersectWorld(w World, r Ray) Intersections {
	var intersects []Intersection

	for _, obj := range w.Objects {
		i := Intersect(obj, r)

		for _, innerIntersect := range i {
			intersects = append(intersects, innerIntersect)
		}
	}

	return NewIntersections(intersects...)
}

func (w World) ContainsObject(s *Sphere) bool {
	for _, obj := range w.Objects {
		if SphereEquals(obj, s) {
			return true
		}
	}

	return false
}
