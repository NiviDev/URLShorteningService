package main

import (
	"gorm.io/gorm"
)

type Shorten struct {
	gorm.Model
	URL         string `form:"url"`
	ShortCode   string `form:"shortened"`
	AccessCount uint
}
