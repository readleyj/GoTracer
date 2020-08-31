package internal

import (
	"math/rand"
	"sort"
	"time"
)

type Group struct {
	ID        int64
	Material  Material
	Transform Matrix
	Parent    Shape
	Children  []Shape
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewGroup() *Group {
	return &Group{
		ID:        rand.Int63(),
		Material:  NewDefaultMaterial(),
		Transform: NewIdentity4(),
		Parent:    nil,
		Children:  []Shape{},
	}
}

func (group *Group) GetID() int64 {
	return group.ID
}

func (group *Group) LocalIntersect(ray Ray) Intersections {
	var intersects Intersections

	for _, child := range group.Children {
		childIntersects := Intersect(child, ray)
		intersects = append(intersects, childIntersects...)
	}

	sort.Slice(intersects, func(i, j int) bool {
		return intersects[i].T < intersects[j].T
	})

	return intersects
}

func (group *Group) LocalNormalAt(point Tuple, hit Intersection) Tuple {
	return Tuple{}
}

func (group *Group) GetTransform() Matrix {
	return group.Transform
}

func (group *Group) SetTransform(transform Matrix) {
	group.Transform = transform
}

func (group *Group) GetMaterial() Material {
	return group.Material
}

func (group *Group) SetMaterial(material Material) {
	group.Material = material
}

func (group *Group) GetParent() Shape {
	return group.Parent
}

func (group *Group) SetParent(s Shape) {
	group.Parent = s
}

func (group *Group) AddChild(s Shape) {
	group.Children = append(group.Children, s)
	s.SetParent(group)
}
