package internal

type World struct {
	Lights  []PointLight
	Objects []Shape
}

func NewWorld() World {
	return World{}
}

func NewDefaultWorld() World {
	var lights []PointLight
	light := NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1))
	lights = append(lights, light)

	s1 := NewSphere()
	s1.SetMaterial(NewMaterial(
		NewColor(0.8, 1.0, 0.6),
		DefaultMaterial.Pattern,
		DefaultMaterial.Ambient,
		0.7,
		0.2,
		DefaultMaterial.Shininess,
		DefaultMaterial.Reflective,
		DefaultMaterial.Transparency,
		DefaultMaterial.RefractiveIndex,
	))

	s2 := NewSphere()
	s2.SetTransform(Scale(0.5, 0.5, 0.5))

	return World{
		lights,
		[]Shape{
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

// Currently IsShadowed does the calculation
// on the first light in the passed World's Lights
// This will have to be fixed
// The other lights have to be considered as well
func IsShadowed(world World, point Tuple) bool {
	v := SubTuples(world.Lights[0].Position, point)
	distance := Magnitude(v)
	direction := Normalize(v)

	r := NewRay(point, direction)
	intersections := IntersectWorld(world, r)
	h := Hit(intersections)
	empty := Intersection{}

	return h != empty && h.T < distance
}

func (w World) ContainsObject(s Shape) bool {
	for _, obj := range w.Objects {
		if ShapeEquals(obj, s) {
			return true
		}
	}

	return false
}
