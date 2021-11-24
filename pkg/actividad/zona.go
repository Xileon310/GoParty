package actividad

// Zona representa una localidad en la que puede organizarse
// una actividad
type Zona struct {
	// Localidad contiene el nombre de la localidad representada por la Zona
	localidad string

	// Provincia en la que se encuentra la Zona
	provincia string

	// País en el que se encuentra la Zona
	pais string
}

// NewZona inicializa y devuelve un objeto Zona
func NewZona(localidad string, provincia string, pais string) Zona {
	return Zona{localidad, provincia, pais}
}
