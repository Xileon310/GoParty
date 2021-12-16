package actividad

import (
	"testing"
)

// ---------------------- Tests del Objeto Valor Zona ----------------------
func TestZonaValida(t *testing.T) {
	_, err := NewZona("Motril", "Granada", "España")
	if err != nil {
		t.Fatalf(err.Error())
	}
}

// ---------------------- Tests de la Entidad Actividad ----------------------
func TestActividadValida(t *testing.T) {
	zona := Zona{"Motril", "Granada", "España"}
	_, err := NewActividad("Ruta de senderismo", zona, Ocio)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
