package database

import (
	"fmt"
	"log"
	"spotify/graph/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func Connect() *Database {
	host := "localhost"
	user := "postgres"
	password := "postgres"
	name := "spotify"
	port := 5432
	time_zone := "Asia/Jakarta"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", host, user, password, name, port, time_zone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	log.Print("successfully connected to database")

	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&model.Artist{}, &model.Song{}, &model.User{})

	return &Database{
		DB: db,
	}
}

func (database *Database) GetUsers() []*model.User  {
	users := []*model.User{}
	database.DB.Find(&users)
	return users
}

func (database *Database) AddUser(input model.NewUser) *model.User {
	user := &model.User{
		Name: input.Name,
	}
	database.DB.Create(user)
	return user
}

func (database *Database) GetArtist(id int) *model.Artist {
	artist := &model.Artist{}
	database.DB.Model(&model.Artist{}).Preload("Songs").First(artist, id)
	return artist
}

func (database *Database) GetArtists() []*model.Artist {
	artists := []*model.Artist{}
	database.DB.Model(&model.Artist{}).Preload("Songs").Find(&artists)
	return artists
}

func (database *Database) AddArtist(input model.NewArtist) *model.Artist {
	artist := &model.Artist{
		Name:  input.Name,
		Songs: []*model.Song{},
	}
	database.DB.Create(artist)
	return artist
}

func (database *Database) GetSongs() []*model.Song {
	songs := []*model.Song{}
	database.DB.Model(&model.Song{}).Preload("Artist").Find(&songs)
	return songs
}

func (database *Database) AddSong(input model.NewSong) *model.Song {
	song := &model.Song{
		Name:     input.Name,
		ArtistID: input.ArtistID,
		Artist: database.GetArtist(input.ArtistID),
	}
	database.DB.Create(song)
	return song
}