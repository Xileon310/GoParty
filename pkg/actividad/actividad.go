// El paquete actividad contiene las estructuras necesarias para
// representar una actividad organizada por un usuario.
package actividad

import "fmt"

// Representa las distintas categorías de actividades
type Categoria int

const (
	Ocio Categoria = iota
	Cultura
	Deporte
)

// Representación más básica de la información con la que trabajará
// la aplicación. Esta clase representa una actividad.
type Actividad struct {
	// Titulo de la actividad
	titulo string

	// Zona en la que se desarrolla la actividad
	zona Zona

	// Categoría a la que pertenece la actividad
	categoria Categoria
}

type errorActividad struct {
	err string
}

func (e *errorActividad) Error() string {
	return fmt.Sprintf("Error al crear la actividad: %s", e.err)
}

// NewActividad inicializa y devuelve un objeto Actividad
func NewActividad(titulo string, zona Zona, categoria Categoria) (Actividad, error) {
	var act Actividad

	if titulo == "" {
		return act, &errorActividad{"titulo vacío"}
	}

	if categoria < Ocio || categoria > Deporte {
		return act, &errorActividad{"categoría no válida"}
	}

	act = Actividad{titulo, zona, categoria}

	return act, nil
}

func BuscarActividadPorZona(actividades []Actividad, zona Zona) []Actividad {
	var actividadesZona []Actividad = nil

	for _, act := range actividades {
		if CompararZona(act.zona, zona) {
			actividadesZona = append(actividadesZona, act)
		}
	}

	return actividadesZona
}
