package models

import "time"

type Lead struct {
	Id           uint      `json:"id"`
	LeadId       string    `json:"leadId"`
	Name         string    `json:"name"`
	Email        string    `json:"email" gorm:"unique"`
	Comapany     string    `json:"company"`
	Domain       string    `json:"domain"`
	ReportTitle  string    `json:"reportTitle"`
	Comment      string    `json:"comment"`
	Designations string    `json:"designation"`
	Mobile       string    `json:"mobile"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}
