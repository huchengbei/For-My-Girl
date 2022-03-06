package models

import (
	"time"

	"gorm.io/gorm"
)

type Moment struct {
	gorm.Model
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Datetime   time.Time `json:"datetime"`
	Link       string    `json:"link"`
	MomentType string    `json:"moment_type"`
	ImageLink  string    `json:"image_link"`
	AlbumLink  string    `json:"album_link"`
	MusicLink  string    `json:"music_link"`
	VideoLink  string    `json:"video_link"`
}

func GetMoments() (moments []Moment) {
	db.Order("datetime desc").Find(&moments)
	return
}

func AddMoment(title string, content string, datetimeStr string, link string, momentType string, imageLink string,
	albumLink string, musicLink string, videoLink string) bool {
	datetime, err := time.Parse(time.RFC3339, datetimeStr)

	if err != nil {
		return false
	}

	db.Create(&Moment{
		Title:      title,
		Content:    content,
		Datetime:   datetime,
		Link:       link,
		MomentType: momentType,
		ImageLink:  imageLink,
		AlbumLink:  albumLink,
		MusicLink:  musicLink,
		VideoLink:  videoLink,
	})

	return true
}

func EditMoment(id int, title string, content string, datetimeStr string, link string, momentType string, imageLink string,
	albumLink string, musicLink string, videoLink string) bool {
	datetime, err := time.Parse(time.RFC3339, datetimeStr)

	if err != nil {
		return false
	}
	var moment Moment
	db.Where("id = ?", id).First(&moment)

	moment.Title = title
	moment.Content = content
	moment.Datetime = datetime
	moment.Link = link
	moment.MomentType = momentType
	moment.ImageLink = imageLink
	moment.AlbumLink = albumLink
	moment.MomentType = musicLink
	moment.VideoLink = videoLink

	db.Save(&moment)

	return true
}

func DeleteMoment(id int) bool {
	var moment Moment
	db.Where("id = ?", id).First(&moment)
	db.Delete(&moment)

	return true
}
