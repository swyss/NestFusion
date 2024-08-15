package models

type Setting struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Key   string `json:"key"`
	Value string `json:"value"`
}
