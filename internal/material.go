package internal

type Material struct {
	Color           Color
	Pattern         Pattern
	Ambient         float64
	Diffuse         float64
	Specular        float64
	Shininess       float64
	Reflective      float64
	Transparency    float64
	RefractiveIndex float64
}

var DefaultMaterial = NewMaterial(
	NewColor(1, 1, 1),
	nil,
	0.1,
	0.9,
	0.9,
	200.0,
	0.0,
	0.0,
	1.0,
)

func NewDefaultMaterial() Material {
	return Material{
		DefaultMaterial.Color,
		nil,
		DefaultMaterial.Ambient,
		DefaultMaterial.Diffuse,
		DefaultMaterial.Specular,
		DefaultMaterial.Shininess,
		DefaultMaterial.Reflective,
		DefaultMaterial.Transparency,
		DefaultMaterial.RefractiveIndex,
	}
}

func NewMaterial(color Color, pattern Pattern, amb, dif, spec, shin, reflect, trans, refr float64) Material {
	return Material{
		color,
		pattern,
		amb,
		dif,
		spec,
		shin,
		reflect,
		trans,
		refr,
	}
}

func MaterialEquals(m1, m2 Material) bool {
	return ColorEquals(m1.Color, m2.Color) && m1.Ambient == m2.Ambient &&
		m1.Diffuse == m2.Diffuse &&
		m1.Specular == m2.Specular &&
		m1.Shininess == m2.Shininess &&
		m1.Reflective == m2.Reflective
}

func (m *Material) SetColor(c Color) {
	m.Color = c
}

func (m *Material) SetPattern(pattern Pattern) {
	m.Pattern = pattern
}

func (m *Material) HasPattern() bool {
	return m.Pattern != nil
}
