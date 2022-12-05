package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Gpu struct {
	gorm.Model
	ProductName string `json:"productName" validate:"nonzero"`
	GpuChip     string `json:"gpuChip" validate:"nonzero"`
	Released    string `json:"released" validate:"nonzero"`
	Memory      string `json:"memory" validate:"nonzero"`
}

func ValidateGpuData(gpu *Gpu) error {
	if err := validator.Validate(gpu); err != nil {
		return err
	}
	return nil
}
