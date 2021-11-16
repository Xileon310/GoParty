// El paquete actividad contiene las estructuras necesarias para
// representar una actividad organizada por un usuario.
package actividad

// Actividad representa una actividad
type Actividad struct {
	// Titulo de la actividad
	titulo string

	// Zona en la que se desarrolla la actividad
	zona Zona
}

// NewActividad inicializa y devuelve un objeto Actividad
func NewActividad(titulo string, zona Zona) Actividad {
	a := Actividad{titulo, zona}
	return a
}
