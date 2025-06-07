package repository

import (
	"fmt"
	"incidence_grade/config"
	"incidence_grade/models"
	"incidence_grade/utils"
	"strconv"

	"gorm.io/gorm"
)

const (
	statusIDResolved uint = 2
	statusIDRejected uint = 3
	statusIDDraft    uint = 4
)

type IncidentRepository struct {
	db     *gorm.DB
	config *config.Config
}

func NewIncidentRepository(db *gorm.DB, config *config.Config) *IncidentRepository {
	return &IncidentRepository{db: db, config: config}
}

func (r *IncidentRepository) Create(incident *models.Incident) (*models.Incident, error) {
	dbErr := r.db.Create(incident).Error
	if dbErr != nil {
		return nil, dbErr
	}

	incidentWithDetails, err := r.FindById(incident.ID)
	if err != nil {
		fmt.Printf("Error fetching incident details after create (ID: %d) for email: %v\n", incident.ID, err)
		return incident, nil
	}

	if incidentWithDetails != nil && incidentWithDetails.User.Email != "" {
		if incidentWithDetails.StatusID == statusIDDraft {
			data := utils.TemplateData{
				"IncidenceID": strconv.FormatUint(uint64(incidentWithDetails.ID), 10),
				"StudentName": incidentWithDetails.User.FirstName + " " + incidentWithDetails.User.LastName,
			}

			emailErr := utils.SendIncidenceUpdateEmail(incidentWithDetails.User.Email, "borrador", data, r.config.ResendAPIKey, r.config.FromEmail)
			if emailErr != nil {
				fmt.Printf("Error sending creation email (status 'borrador') for incidence ID %d: %v\n", incidentWithDetails.ID, emailErr)
			}
		}
	} else {
		if incidentWithDetails == nil {
			fmt.Printf("Could not fetch incident details (ID: %d) after creation for email.\n", incident.ID)
		} else {
			fmt.Printf("User email not found for incidence ID %d, cannot send creation email.\n", incidentWithDetails.ID)
		}
	}

	return incident, nil
}

func (r *IncidentRepository) FindById(id uint) (*models.Incident, error) {
	var incident models.Incident
	error := r.db.Preload("Status").Preload("User").Preload("Attachment").First(&incident, id).Error
	if error != nil {
		return nil, error
	}
	return &incident, nil
}

func (r *IncidentRepository) FindAll() ([]models.Incident, error) {
	var incidents []models.Incident
	error := r.db.Preload("Status").Preload("User").Preload("Attachment").Find(&incidents).Error
	if error != nil {
		return nil, error
	}
	return incidents, nil
}

func (r *IncidentRepository) Update(incident *models.Incident) (*models.Incident, error) {
	if dbErr := r.db.
		Model(&models.Incident{ID: incident.ID}).
		Updates(incident). // This updates fields passed in 'incident'. Ensure StatusID is part of 'incident' if it's changing.
		Error; dbErr != nil {
		return nil, dbErr
	}

	// Fetch the updated incident with User and Status preloaded for email
	incidentWithDetails, err := r.FindById(incident.ID)
	if err != nil {
		fmt.Printf("Error fetching incident details after update (ID: %d) for email: %v\n", incident.ID, err)
		return incident, nil // Still return the incident as potentially modified
	}

	if incidentWithDetails != nil && incidentWithDetails.User.Email != "" {
		var statusStringToEmail string

		switch incidentWithDetails.StatusID {
		case statusIDResolved:
			statusStringToEmail = "resuelto"
		case statusIDRejected:
			statusStringToEmail = "rechazado"
		default:
			return incident, nil
		}

		data := utils.TemplateData{
			"IncidenceID": strconv.FormatUint(uint64(incidentWithDetails.ID), 10),
			"StudentName": incidentWithDetails.User.FirstName + " " + incidentWithDetails.User.LastName,
		}

		if incidentWithDetails.StatusID == statusIDRejected {
			data["RejectionReason"] = incidentWithDetails.Response
		}

		emailErr := utils.SendIncidenceUpdateEmail(incidentWithDetails.User.Email, statusStringToEmail, data, r.config.ResendAPIKey, r.config.FromEmail)
		if emailErr != nil {
			fmt.Printf("Error sending update email for incidence ID %d (status: %s): %v\n", incidentWithDetails.ID, statusStringToEmail, emailErr)
		}

	} else {
		if incidentWithDetails == nil {
			fmt.Printf("Could not fetch incident details (ID: %d) after update for email.\n", incident.ID)
		} else {
			fmt.Printf("User email not found for incidence ID %d, cannot send update email.\n", incidentWithDetails.ID)
		}
	}

	return incident, nil // Return the incident as potentially modified by Updates
}

func (r *IncidentRepository) Delete(id uint) (map[string]interface{}, error) {

	error := r.db.Delete(&models.Incident{}, id).Error
	if error != nil {
		return nil, error
	}
	resutl := map[string]interface{}{
		"ok": true,
		"id": id,
	}
	return resutl, nil
}

func (r *IncidentRepository) FindByIdUser(user_id uint) ([]models.Incident, error) {
	var incident []models.Incident
	error := r.db.Where("user_id = ?", user_id).Preload("Status").Preload("User").Preload("Attachment").Find(&incident).Error
	if error != nil {
		return nil, error
	}
	return incident, nil
}

func (r *IncidentRepository) SaveFile(files []models.Attachment) ([]models.Attachment, error) {
	for i := range files {
		error := r.db.Create(&files[i]).Error
		if error != nil {
			return nil, error
		}
	}
	return files, nil

}
