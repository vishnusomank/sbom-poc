package models

type SBOM struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	ImageName     string `json:"image"`
	Version       string `json:"version"`
	Value         string `json:"sbom"`
	Vulnerability string `json:"vulnerability"`
}
