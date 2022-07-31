package database

import (
	"gorm.io/gorm"
)

type Artist struct {
	gorm.Model
	Name  string  `json:"name"`
	Songs []*Song `json:"songs"`
}

type Song struct {
	gorm.Model 
	Name     string `json:"name"`
	ArtistID int
	Artist *Artist `json:"artist"`
}

type User struct {
	gorm.Model  
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	FavoriteSongs []*Song `gorm:"many2many:user_favorite_songs;" json:"favoriteSongs"`
}
