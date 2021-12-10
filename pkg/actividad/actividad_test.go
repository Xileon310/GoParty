package actividad

import (
	"testing"
)

// ---------------------- Tests del Objeto Valor Zona ----------------------
func TestLocalidadNoVacia(t *testing.T) {
	_, err := NewZona("", "Granada", "España")
	if err == nil {
		t.Fatalf("Localidad no válida en objeto Zona")
	}
}

func TestProvinciaNoVacia(t *testing.T) {
	_, err := NewZona("Motril", "", "España")
	if err == nil {
		t.Fatalf("Provincia no válida en objeto Zona")
	}
}

func TestPaisNoVacio(t *testing.T) {
	_, err := NewZona("Motril", "Granada", "")
	if err == nil {
		t.Fatalf("País no válido en objeto Zona")
	}
}

// ---------------------- Tests de la Entidad Actividad ----------------------
func TestTituloNoVacio(t *testing.T) {
	zona := Zona{"Motril", "Granada", "España"}
	_, err := NewActividad("", zona, Ocio)
	if err == nil {
		t.Fatalf("Título no válido en objeto Actividad")
	}
}

func TestCategoriaNoValida(t *testing.T) {
	zona := Zona{"Motril", "Granada", "España"}
	_, err := NewActividad("Ruta de senderismo", zona, 99)
	if err == nil {
		t.Fatalf("Categoría no válida en objeto Actividad")
	}
}
