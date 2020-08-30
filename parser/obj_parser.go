package parser

import (
	"bufio"
	"gotracer/internal"
	"strconv"
	"strings"
)

type Object struct {
	Vertices []internal.Tuple
	Normals  []internal.Tuple
	Groups   map[string]*internal.Group
}

func NewObject() *Object {
	obj := Object{}
	obj.Vertices = append(obj.Vertices, internal.Tuple{})
	obj.Normals = append(obj.Normals, internal.Tuple{})
	obj.Groups = make(map[string]*internal.Group)
	obj.Groups["DefaultGroup"] = internal.NewGroup()

	return &obj
}

func ParseObjectFile(objData string) (*Object, int) {
	obj := NewObject()
	linesIgnored := 0
	lastGroup := "DefaultGroup"

	scanner := bufio.NewScanner(strings.NewReader(objData))

	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		words := strings.Split(line, " ")

		switch words[0] {
		case "v":
			{
				x, _ := strconv.ParseFloat(words[1], 64)
				y, _ := strconv.ParseFloat(words[2], 64)
				z, _ := strconv.ParseFloat(words[3], 64)
				vertex := internal.NewPoint(x, y, z)

				obj.Vertices = append(obj.Vertices, vertex)
			}

		case "f":
			{
				var vertices []internal.Tuple
				var normals []internal.Tuple
				normalsSet := len(obj.Normals) != 1

				for i := 1; i <= len(words)-1; i++ {
					var vertIndex, normalIndex int
					indices := strings.Split(words[i], "/")
					vertIndex, _ = strconv.Atoi(indices[0])

					if normalsSet {
						normalIndex, _ = strconv.Atoi(indices[2])
					}

					vertices = append(vertices, obj.Vertices[vertIndex])
					normals = append(normals, obj.Normals[normalIndex])
				}

				for i := 1; i < len(vertices)-1; i++ {
					if normalsSet {
						tri := internal.NewSmoothTriangle(
							vertices[0],
							vertices[i],
							vertices[i+1],
							normals[0],
							normals[i],
							normals[i+1],
						)

						obj.Groups[lastGroup].AddChild(tri)
					} else {
						tri := internal.NewTriangle(
							vertices[0],
							vertices[i],
							vertices[i+1],
						)

						obj.Groups[lastGroup].AddChild(tri)
					}

				}
			}

		case "g":
			{
				lastGroup = words[1]
				_, present := obj.Groups[lastGroup]

				if !present {
					obj.Groups[lastGroup] = internal.NewGroup()
				}
			}

		case "vn":
			{
				x, _ := strconv.ParseFloat(words[1], 64)
				y, _ := strconv.ParseFloat(words[2], 64)
				z, _ := strconv.ParseFloat(words[3], 64)
				normal := internal.NewVector(x, y, z)

				obj.Normals = append(obj.Normals, normal)
			}

		default:
			{
				linesIgnored++
			}

		}
	}

	return obj, linesIgnored
}

func (obj *Object) GetGroup(groupName string) *internal.Group {
	return obj.Groups[groupName]
}

func (obj *Object) GetDefaultGroup() *internal.Group {
	return obj.GetGroup("DefaultGroup")
}

func (obj *Object) ToGroup() *internal.Group {
	group := internal.NewGroup()

	for _, g := range obj.Groups {
		group.AddChild(g)
	}

	return group
}
