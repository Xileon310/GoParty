package actividad

import "fmt"

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

type errorZona struct {
	err string
}

func (e *errorZona) Error() string {
	return fmt.Sprintf("Error al crear la zona: %s", e.err)
}

// NewZona inicializa y devuelve un objeto Zona
func NewZona(localidad string, provincia string, pais string) (Zona, error) {
	var zona Zona

	if localidad == "" {
		return zona, &errorZona{"localidad vacía"}
	}

	if provincia == "" {
		return zona, &errorZona{"provincia vacía"}
	}

	if pais == "" {
		return zona, &errorZona{"país vacío"}
	}

	zona = Zona{localidad, provincia, pais}

	return zona, nil
}

func CompararZona(una Zona, otra Zona) bool {
	if una.localidad == otra.localidad && una.provincia == otra.provincia && una.pais == otra.pais {
		return true
	} else {
		return false
	}
}
