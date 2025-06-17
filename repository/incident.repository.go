package repository

import (
	"fmt"
	"incidence_grade/config"
	"incidence_grade/models"
	"incidence_grade/utils"
	"strconv"
	"time"

	"gorm.io/gorm"
)

const (
	statusIDResolved uint = 2
	statusIDRejected uint = 3
	statusIDDraft    uint = 4
)

const (
	typeIncidentADD = "Adición"
	typeIncidentRET = "Retiro"
)

type IncidentRepository struct {
	db     *gorm.DB
	config *config.Config
}

func NewIncidentRepository(db *gorm.DB, config *config.Config) *IncidentRepository {
	return &IncidentRepository{db: db, config: config}
}

func (r *IncidentRepository) Create(incident *models.Incident) (*models.Incident, error) {
	if incident.Title == typeIncidentADD || incident.Title == typeIncidentRET {
		MonthsAgo := time.Now().AddDate(0, -4, 0)
		existingCount := int64(0)
		err := r.db.Model(&models.Incident{}).
			Where("user_id = ? AND title = ? AND created_at > ?",
				incident.UserID,
				incident.Title,
				MonthsAgo).
			Count(&existingCount).Error

		if err != nil {
			return nil, err
		}

		if existingCount > 0 {
			return nil, fmt.Errorf("%s, Solo una solicitud por semestre", incident.Title)
		}
	}

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
		Updates(incident).
		Error; dbErr != nil {
		return nil, dbErr
	}

	incidentWithDetails, err := r.FindById(incident.ID)
	if err != nil {
		fmt.Printf("Error fetching incident details after update (ID: %d) for email: %v\n", incident.ID, err)
		return incident, nil
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

		var userToken models.UserToken
		if err := r.db.Where("user_id = ?", incidentWithDetails.UserID).First(&userToken).Error; err == nil {
			notification := &utils.ExpoNotification{
				Token: userToken.DeviceToken,
				Title: "Cambio de estatus en incidencia",
				Body:  fmt.Sprintf("Se ha cambiado el estatus de una incidencia a %s", statusStringToEmail),
				Data: map[string]string{
					"incidenceId": strconv.FormatUint(uint64(incidentWithDetails.ID), 10),
					"status":      statusStringToEmail,
				},
			}

			if _, err := utils.SendExpoNotification(notification); err != nil {
				fmt.Printf("Error sending push notification for incidence ID %d: %v\n", incidentWithDetails.ID, err)
			}
		} else {
			fmt.Printf("No device token found for user ID %d, cannot send push notification\n", incidentWithDetails.UserID)
		}

	} else {
		if incidentWithDetails == nil {
			fmt.Printf("Could not fetch incident details (ID: %d) after update for email.\n", incident.ID)
		} else {
			fmt.Printf("User email not found for incidence ID %d, cannot send update email.\n", incidentWithDetails.ID)
		}
	}

	return incident, nil
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
