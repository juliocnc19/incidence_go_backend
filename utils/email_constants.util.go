package utils

import (
	"fmt"
	"strings"
)

const (
	SubjectIncidenceCreated  = "Nueva Incidencia Registrada: %s"
	SubjectIncidenceResolved = "Incidencia Resuelta: %s"
	SubjectIncidenceRejected = "Incidencia Rechazada: %s"

	// Placeholders: {{.IncidenceID}}, {{.StudentName}}, {{.Status}}, {{.DetailsLink}}
	BodyIncidenceCreated = `
		<h1>Nueva Incidencia Creada</h1>
		<p>Hola {{.StudentName}},</p>
		<p>Se ha registrado una nueva incidencia con ID: <strong>{{.IncidenceID}}</strong>.</p>
		<p>Estado actual: <strong>Borrador</strong>.</p>
		<p>Saludos,</p>
		<p>Equipo de Soporte de AIS</p>
	`

	BodyIncidenceResolved = `
		<h1>Incidencia Resuelta</h1>
		<p>Hola {{.StudentName}},</p>
		<p>Tu incidencia con ID: <strong>{{.IncidenceID}}</strong> ha sido marcada como <strong>Resuelta</strong>.</p>
		<p>Si consideras que no ha sido resuelta, por favor, contacta con nosotros.</p>
		<p>Saludos,</p>
		<p>Equipo de Soporte de AIS</p>
	`

	BodyIncidenceRejected = `
		<h1>Incidencia Rechazada</h1>
		<p>Hola {{.StudentName}},</p>
		<p>Lamentamos informarte que tu incidencia con ID: <strong>{{.IncidenceID}}</strong> ha sido <strong>Rechazada</strong>.</p>
		<p>Motivo del rechazo: {{.RejectionReason}}</p> // Este placeholder es un ejemplo, puedes añadir más según necesites
		<p>Si tienes alguna duda, por favor, contacta con nosotros.</p>
		<p>Saludos,</p>
		<p>Equipo de Soporte de AIS</p>
	`
)

type TemplateData map[string]string

func GetEmailContent(status string, data TemplateData) (subject, body string, err error) {
	incidenceID, ok := data["IncidenceID"]
	if !ok {
		incidenceID = "N/A"
	}

	switch status {
	case "borrador": // Assuming 'borrador' is the status for a newly created incidence
		subject = fmt.Sprintf(SubjectIncidenceCreated, incidenceID)
		body = populateTemplate(BodyIncidenceCreated, data)
	case "resuelto":
		subject = fmt.Sprintf(SubjectIncidenceResolved, incidenceID)
		body = populateTemplate(BodyIncidenceResolved, data)
	case "rechazado":
		subject = fmt.Sprintf(SubjectIncidenceRejected, incidenceID)
		body = populateTemplate(BodyIncidenceRejected, data)
	default:
		return "", "", fmt.Errorf("estado de incidencia no válido para notificación: %s", status)
	}
	return subject, body, nil
}

func populateTemplate(template string, data TemplateData) string {
	result := template
	for key, value := range data {
		placeholder := fmt.Sprintf("{{.%s}}", key)
		result = strings.ReplaceAll(result, placeholder, value)
	}
	return result
}
