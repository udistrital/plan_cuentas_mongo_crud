// Modelos faltantes para completar Swagger

package models

type Root1 struct {
	Codigo          string
	Nombre          string
	Hijos           []string
	IsLeaf          bool
	UnidadEjecutora string
	ValorInicial    float64
}

type VigenciaNamespace map[string]int

type PorDefinir struct {
}
