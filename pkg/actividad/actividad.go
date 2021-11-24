// El paquete actividad contiene las estructuras necesarias para
// representar una actividad organizada por un usuario.
package actividad

// Representa las distintas categorías de actividades
type Categoria int

const (
	Ocio Categoria = iota
	Cultura
	Deporte
)

// Actividad representa una actividad
type Actividad struct {
	// Titulo de la actividad
	titulo string

	// Zona en la que se desarrolla la actividad
	zona Zona

	// Categoría a la que pertenece la actividad
	categoria Categoria
}

// NewActividad inicializa y devuelve un objeto Actividad
func NewActividad(titulo string, zona Zona, categoria Categoria) Actividad {
	return Actividad{titulo, zona, categoria}
}
