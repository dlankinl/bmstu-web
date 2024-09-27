package web

import (
	"ppo/domain"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Username string    `json:"username,omitempty"`
	FullName string    `json:"fullName,omitempty"`
	Gender   string    `json:"gender,omitempty"`
	Birthday time.Time `json:"birthday,omitempty"`
	City     string    `json:"city,omitempty"`
	Role     string    `json:"role,omitempty"`
}

type Contact struct {
	ID      uuid.UUID `json:"id,omitempty"`
	OwnerID uuid.UUID `json:"ownerId,omitempty"`
	Name    string    `json:"name,omitempty"`
	Value   string    `json:"value,omitempty"`
}

type ActivityField struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Cost        float32   `json:"cost,omitempty"`
}

type Company struct {
	ID              uuid.UUID `json:"id,omitempty"`
	OwnerID         uuid.UUID `json:"ownerId,omitempty"`
	ActivityFieldId uuid.UUID `json:"activityFieldId,omitempty"`
	Name            string    `json:"name,omitempty"`
	City            string    `json:"city,omitempty"`
}

type FinancialReport struct {
	ID        uuid.UUID `json:"id,omitempty"`
	CompanyID uuid.UUID `json:"companyId,omitempty"`
	Revenue   float32   `json:"revenue,omitempty"`
	Costs     float32   `json:"costs,omitempty"`
	Year      int       `json:"year,omitempty"`
	Quarter   int       `json:"quarter,omitempty"`
}

type Period struct {
	StartYear    int `json:"startYear"`
	StartQuarter int `json:"startQuarter"`
	EndYear      int `json:"endYear"`
	EndQuarter   int `json:"endQuarter"`
}

func toUserTransport(user *domain.User) User {
	return User{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		Gender:   user.Gender,
		Birthday: user.Birthday,
		City:     user.City,
		Role:     user.Role,
	}
}

func toUserModel(user *User) domain.User {
	return domain.User{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		Gender:   user.Gender,
		Birthday: user.Birthday,
		City:     user.City,
		Role:     user.Role,
	}
}

func toContactTransport(contact *domain.Contact) Contact {
	return Contact{
		ID:      contact.ID,
		OwnerID: contact.OwnerID,
		Name:    contact.Name,
		Value:   contact.Value,
	}
}

func toContactModel(contact *Contact) domain.Contact {
	return domain.Contact{
		ID:      contact.ID,
		OwnerID: contact.OwnerID,
		Name:    contact.Name,
		Value:   contact.Value,
	}
}

func toActFieldTransport(field *domain.ActivityField) ActivityField {
	return ActivityField{
		ID:          field.ID,
		Name:        field.Name,
		Description: field.Description,
		Cost:        field.Cost,
	}
}

func toActFieldModel(field *ActivityField) domain.ActivityField {
	return domain.ActivityField{
		ID:          field.ID,
		Name:        field.Name,
		Description: field.Description,
		Cost:        field.Cost,
	}
}

func toCompanyTransport(company *domain.Company) Company {
	return Company{
		ID:              company.ID,
		OwnerID:         company.OwnerID,
		ActivityFieldId: company.ActivityFieldId,
		Name:            company.Name,
		City:            company.City,
	}
}

func toCompanyModel(company *Company) domain.Company {
	return domain.Company{
		ID:              company.ID,
		OwnerID:         company.OwnerID,
		ActivityFieldId: company.ActivityFieldId,
		Name:            company.Name,
		City:            company.City,
	}
}

func toFinReportTransport(finReport *domain.FinancialReport) FinancialReport {
	return FinancialReport{
		ID:        finReport.ID,
		CompanyID: finReport.CompanyID,
		Revenue:   finReport.Revenue,
		Costs:     finReport.Costs,
		Year:      finReport.Year,
		Quarter:   finReport.Quarter,
	}
}

func toFinReportModel(finReport *FinancialReport) domain.FinancialReport {
	return domain.FinancialReport{
		ID:        finReport.ID,
		CompanyID: finReport.CompanyID,
		Revenue:   finReport.Revenue,
		Costs:     finReport.Costs,
		Year:      finReport.Year,
		Quarter:   finReport.Quarter,
	}
}

func toPeriodTransport(per *domain.Period) Period {
	return Period{
		StartYear:    per.StartYear,
		StartQuarter: per.StartQuarter,
		EndYear:      per.EndYear,
		EndQuarter:   per.EndQuarter,
	}
}
