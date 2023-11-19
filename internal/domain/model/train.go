package model

import "gorm.io/gorm"

type Train struct {
	gorm.Model

	From     string
	To       string
	Sections []*Section
}

func NewTrain(from string, to string) *Train {
	train := &Train{From: from, To: to}
	sections := []*Section{
		NewSection("Section 1", train),
		NewSection("Section 2", train),
	}
	train.Sections = sections

	return train
}
