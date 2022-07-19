package models

type SBOM struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	ImageName     string `json:"image"`
	Version       string `json:"version"`
	Value         string `json:"sbom"`
	Vulnerability string `json:"vulnerability"`
}

type PolicyDB struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	CVEId      string `json:"cve"`
	PolicyData string `json:"policy"`
}

type SBOMPolicy struct {
	SbomID   int `json:"sbomID"`
	PolicyID int `json:"policyID"`
}
