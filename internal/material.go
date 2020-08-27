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

func MaterialEquals(m1, m2 Material) bool {
	return ColorEquals(m1.Color, m2.Color) && m1.Ambient == m2.Ambient &&
		m1.Diffuse == m2.Diffuse &&
		m1.Specular == m2.Specular &&
		m1.Shininess == m2.Shininess
}

func (m *Material) SetColor(c Color) {
	m.Color = c
}
