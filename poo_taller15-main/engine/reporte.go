package engine

import (
	"fmt"

	"poo_taller15-main/model"
)

func GenerarReporteJustificado(
	evaluacion model.Evaluacion,
	docente model.Docente,
	competencia model.Competencia,
) string {

	puntajeFinal, reglas :=
		CalcularPuntajeConReglas(
			evaluacion,
			docente,
			competencia,
		)

	reporte := "===== REPORTE DE EVALUACIÓN =====\n"

	reporte += fmt.Sprintf(
		"Docente: %s\n",
		docente.Nombre)

	reporte += fmt.Sprintf(
		"Departamento: %s\n",
		docente.Departamento)

	reporte += fmt.Sprintf(
		"Cargo: %s\n",
		docente.Cargo)

	reporte += fmt.Sprintf(
		"Puntaje Base: %.2f\n",
		evaluacion.PuntajeBase)

	reporte += "\nReglas Aplicadas:\n"

	if len(reglas) == 0 {
		reporte += "- Ninguna\n"
	}

	for _, r := range reglas {
		reporte += "- " + r + "\n"
	}

	reporte += fmt.Sprintf(
		"\nPuntaje Final: %.2f\n",
		puntajeFinal)

	return reporte
}