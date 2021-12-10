package actividad

import (
	"testing"
)

// ---------------------- Tests del Objeto Valor Zona ----------------------
func TestLocalidadNoVacia(t *testing.T) {
	_, err := NewZona("", "Granada", "España")
	if err != nil {
		t.Errorf("Localidad no válida en objeto Zona")
	}
}

func TestProvinciaNoVacia(t *testing.T) {
	_, err := NewZona("Motril", "", "España")
	if err != nil {
		t.Errorf("Provincia no válida en objeto Zona")
	}
}

func TestPaisNoVacio(t *testing.T) {
	_, err := NewZona("Motril", "Granada", "")
	if err != nil {
		t.Errorf("País no válido en objeto Zona")
	}
}
