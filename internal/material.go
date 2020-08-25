package internal

type Material struct {
	Color     Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

var DefaultMaterial = NewMaterial(
	NewColor(1, 1, 1),
	0.1,
	0.9,
	0.9,
	200.0,
)

func NewDefaultMaterial() Material {
	return Material{
		DefaultMaterial.Color,
		DefaultMaterial.Ambient,
		DefaultMaterial.Diffuse,
		DefaultMaterial.Specular,
		DefaultMaterial.Shininess,
	}
}

func NewMaterial(color Color, amb, dif, spec, shin float64) Material {
	return Material{
		color,
		amb,
		dif,
		spec,
		shin,
	}
}
