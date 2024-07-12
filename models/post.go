package models

import(
	"go_fiber_restfull/validator"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string `json:"title"`
	Content string `json:"content"`
}

func (p *Post) Validate() error {
	v := validator.NewValidator()
	v.AddRule("title", validator.FieldRule{Required: true, Max: 255})
	v.AddRule("content", validator.FieldRule{Required: true})
	return v.Validate(p)
}