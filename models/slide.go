package models

import (
	"gorm.io/gorm"
)

type Slide struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageLink   string `json:"image_link"`
	IsShow      bool   `json:"is_show"`
	Sequence    int    `json:"sequence"`
}

func GetSlides() (slides []Slide) {
	db.Order("sequence").Order("updated_at desc").Find(&slides)
	return
}

func AddSlide(title string, description string, imageLink string, isShow bool, sequence int) bool {
	db.Create(&Slide{
		Title:       title,
		Description: description,
		ImageLink:   imageLink,
		IsShow:      isShow,
		Sequence:    sequence,
	})

	return true
}

func EditSlide(id int, title string, description string, imageLink string, isShow bool, sequence int) bool {

	var slide Slide
	db.Where("id = ?", id).First(&slide)

	slide.Title = title
	slide.Description = description
	slide.ImageLink = imageLink
	slide.IsShow = isShow
	slide.Sequence = sequence

	db.Save(&slide)

	return true
}

func DeleteSlide(id int) bool {
	var slide Slide
	db.Table("slide").Where("id = ?", id).First(&slide)
	db.Table("slide").Delete(&slide)

	return true
}
