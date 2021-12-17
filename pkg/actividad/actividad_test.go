package actividad

import (
	"testing"
)

// ---------------------- Tests del Objeto Valor Zona ----------------------
func TestZonaValida(t *testing.T) {
	zona, err := NewZona("Motril", "Granada", "España")
	if zona.localidad != "Motril" || zona.provincia != "Granada" || zona.pais != "España" || err != nil {
		t.Fatalf(err.Error())
	}
}

// ---------------------- Tests de la Entidad Actividad ----------------------
func TestActividadValida(t *testing.T) {
	zona := Zona{"Motril", "Granada", "España"}
	act, err := NewActividad("Ruta de senderismo", zona, Ocio)
	if act.titulo != "Ruta de senderismo" || !CompararZona(act.zona, zona) || act.categoria != Ocio || err != nil {
		t.Fatalf(err.Error())
	}
}

func TestBuscarActividadPorZona(t *testing.T) {
	zona1, _ := NewZona("Motril", "Granada", "España")
	zona2, _ := NewZona("Almería", "Almería", "España")

	act1, _ := NewActividad("Ruta de senderismo", zona1, Ocio)
	act2, _ := NewActividad("Ruta de ciclismo", zona2, Ocio)

	var actividadesTotales []Actividad
	actividadesTotales = append(actividadesTotales, act1, act2)

	actividadesZona := BuscarActividadPorZona(actividadesTotales, zona2)

	if actividadesZona != nil {
		for _, act := range actividadesZona {
			if !CompararZona(act.zona, zona2) {
				t.Fatalf("Se ha devuelto una actividad que no corresponde.")
			}
		}
	} else {
		t.Fatalf("No se ha encontrado ninguna actividad y debería haberlo hecho.")
	}
}
