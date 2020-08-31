package parser

import (
	"gotracer/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIgnoreUnrecognizedFiles(t *testing.T) {
	file := `There was a young lady named Bright​
​who traveled much faster than light.​
She set out one day​
in a relative way,​
and came back the previous night.`

	_, linesIgnored := ParseObjectFile(file)

	assert.Equal(t, 5, linesIgnored)
}

func TestParseVertexRecords(t *testing.T) {
	file := `
v -1 1 0​
v -1.0000 0.5000 0.0000​
v 1 0 0​
v 1 1 0`

	obj, _ := ParseObjectFile(file)

	assert.True(t, internal.TupleEquals(internal.NewPoint(-1, 1, 0), obj.Vertices[1]))
	assert.True(t, internal.TupleEquals(internal.NewPoint(-1, 0.5, 0), obj.Vertices[2]))
	assert.True(t, internal.TupleEquals(internal.NewPoint(1, 0, 0), obj.Vertices[3]))
	assert.True(t, internal.TupleEquals(internal.NewPoint(1, 1, 0), obj.Vertices[4]))
}

func TestParseTriangleFaces(t *testing.T) {
	file := `
v -1 1 0​
v -1 0 0​
v 1 0 0​
v 1 1 0​

f 1 2 3
f 1 3 4`

	object, _ := ParseObjectFile(file)
	g := object.GetDefaultGroup()
	t1 := g.Children[0].(*internal.Triangle)
	t2 := g.Children[1].(*internal.Triangle)

	assert.Equal(t, object.Vertices[1], t1.P1)
	assert.Equal(t, object.Vertices[2], t1.P2)
	assert.Equal(t, object.Vertices[3], t1.P3)

	assert.Equal(t, object.Vertices[1], t2.P1)
	assert.Equal(t, object.Vertices[3], t2.P2)
	assert.Equal(t, object.Vertices[4], t2.P3)
}

func TestTriangulatingPolygons(t *testing.T) {
	file := `​
v -1 1 0​
v -1 0 0
v 1 0 0
v 1 1 0​
v 0 2 0

f 1 2 3 4 5`

	parser, _ := ParseObjectFile(file)
	g := parser.GetDefaultGroup()
	t1 := g.Children[0].(*internal.Triangle)
	t2 := g.Children[1].(*internal.Triangle)
	t3 := g.Children[2].(*internal.Triangle)

	assert.Equal(t, parser.Vertices[1], t1.P1)
	assert.Equal(t, parser.Vertices[2], t1.P2)
	assert.Equal(t, parser.Vertices[3], t1.P3)

	assert.Equal(t, parser.Vertices[1], t2.P1)
	assert.Equal(t, parser.Vertices[3], t2.P2)
	assert.Equal(t, parser.Vertices[4], t2.P3)

	assert.Equal(t, parser.Vertices[1], t3.P1)
	assert.Equal(t, parser.Vertices[4], t3.P2)
	assert.Equal(t, parser.Vertices[5], t3.P3)
}

func TestTrianglesInGroups(t *testing.T) {
	file := `
v -1 1 0
v -1 0 0
v 1 0 0
v 1 1 0

g FirstGroup
f 1 2 3
g SecondGroup
f 1 3 4`

	obj, _ := ParseObjectFile(file)
	g1 := obj.GetGroup("FirstGroup")
	g2 := obj.GetGroup("SecondGroup")
	t1 := g1.Children[0].(*internal.Triangle)
	t2 := g2.Children[0].(*internal.Triangle)

	assert.Equal(t, obj.Vertices[1], t1.P1)
	assert.Equal(t, obj.Vertices[2], t1.P2)
	assert.Equal(t, obj.Vertices[3], t1.P3)

	assert.Equal(t, obj.Vertices[1], t2.P1)
	assert.Equal(t, obj.Vertices[3], t2.P2)
	assert.Equal(t, obj.Vertices[4], t2.P3)
}

func TestConvertObjFileToGroup(t *testing.T) {
	file := `
v -1 1 0
v -1 0 0
v 1 0 0
v 1 1 0

g FirstGroup
f 1 2 3
g SecondGroup
f 1 3 4`

	obj, _ := ParseObjectFile(file)
	g := obj.ToGroup()
	g1 := obj.GetGroup("FirstGroup")
	g2 := obj.GetGroup("SecondGroup")

	_, includesG1 := internal.IndexOf(g.Children, g1)
	_, includesG2 := internal.IndexOf(g.Children, g2)

	assert.True(t, includesG1)
	assert.True(t, includesG2)
}

func TestVertexNormalRecords(t *testing.T) {
	file := `
vn 0 0 1
vn 0.707 0 -0.707
vn 1 2 3`

	obj, _ := ParseObjectFile(file)

	assert.True(t, internal.TupleEquals(internal.NewVector(0, 0, 1), obj.Normals[1]))
	assert.True(t, internal.TupleEquals(internal.NewVector(0.707, 0, -0.707), obj.Normals[2]))
	assert.True(t, internal.TupleEquals(internal.NewVector(1, 2, 3), obj.Normals[3]))
}

func TestFacesWithNormals(t *testing.T) {
	file := `
v 0 1 0
v -1 0 0
v 1 0 0
vn -1 0 0
vn 1 0 0
vn 0 1 0
f 1//3 2//1 3//2
f 1/0/3 2/102/1 3/14/2`

	obj, _ := ParseObjectFile(file)
	g := obj.GetDefaultGroup()
	t1 := g.Children[0].(*internal.SmoothTriangle)
	t2 := g.Children[1].(*internal.SmoothTriangle)

	assert.Equal(t, obj.Vertices[1], t1.P1)
	assert.Equal(t, obj.Vertices[2], t1.P2)
	assert.Equal(t, obj.Vertices[3], t1.P3)

	assert.Equal(t, obj.Normals[3], t2.N1)
	assert.Equal(t, obj.Normals[1], t2.N2)
	assert.Equal(t, obj.Normals[2], t2.N3)

	assert.True(t, internal.ShapeEquals(t1, t2))
}
