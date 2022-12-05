package models

import "gorm.io/gorm"

type Gpu struct {
	gorm.Model
	productName string `json:"productName"`
	gpuChip     string `json:"gpuChip"`
	released    string `json:"released`
	memory      string `json:memory`
}
