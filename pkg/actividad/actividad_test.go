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
