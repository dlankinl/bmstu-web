package web

import "github.com/google/uuid"

type EntrepreneursResponse struct {
	Pages         int    `json:"numPages"`
	Entrepreneurs []User `json:"entrepreneurs"`
}

type EntrepreneurResponse struct {
	Entrepreneur User `json:"entrepreneur"`
}

type ContactResponse struct {
	Contact Contact `json:"contact"`
}

type ContactsResponse struct {
	EntrepreneurId uuid.UUID `json:"entrepreneurId"`
	Contacts       []Contact `json:"contacts"`
}

type ActFieldResponse struct {
	ActField ActivityField `json:"activityField"`
}

type ActFieldsResponse struct {
	Pages     int             `json:"numPages"`
	ActFields []ActivityField `json:"activityFields"`
}

type CompanyResponse struct {
	Company Company `json:"company"`
}

type CompaniesResponse struct {
	Pages          int       `json:"numPages"`
	EntrepreneurId uuid.UUID `json:"entrepreneurId"`
	Companies      []Company `json:"companies"`
}

type FinReportResponse struct {
	FinReport FinancialReport `json:"financialReport"`
}

type FinReportByPeriodResponse struct {
	CompanyId uuid.UUID         `json:"companyId"`
	Period    Period            `json:"period"`
	Revenue   float32           `json:"revenue"`
	Costs     float32           `json:"costs"`
	Profit    float32           `json:"profit"`
	Reports   []FinancialReport `json:"reports"`
}

type RatingResponse struct {
	Rating float32 `json:"rating"`
}

type EntrepreneurReportResponse struct {
	Revenue float32 `json:"revenue"`
	Costs   float32 `json:"costs"`
	Profit  float32 `json:"profit"`
	Taxes   float32 `json:"taxes"`
	TaxLoad float32 `json:"taxLoad"`
}
