package database

import (
	"fmt"
	"log"
	"spotify/graph/model"
	"spotify/utility"

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

	db.AutoMigrate(&Artist{}, &Song{}, &User{})

	return &Database{
		DB: db,
	}
}

func (database *Database) GetUsers() []*model.User  {
	usersDB := []*User{}
	database.DB.Model(&User{}).Preload("FavoriteSongs").Find(&usersDB)

	users := []*model.User{}
	utility.Recast(&usersDB, &users)
	return users
}

func (database *Database) AddUser(input model.NewUser) *model.User {
	userDB := &User{
		Name: input.Name,
		Email: input.Email,
	}
	database.DB.Create(userDB)

	user := &model.User{}
	utility.Recast(userDB, user)
	return user
}

func (database *Database) AddUserFavoriteSong(input model.NewUserFavoriteSong) *model.User {
	userDB := &User{}
	database.DB.Model(userDB).Preload("FavoriteSongs").First(userDB, input.UserID)

	songDB := &Song{}
	database.DB.Model(songDB).Preload("Artist").First(songDB, input.SongID)

	database.DB.Model(userDB).Association("FavoriteSongs").Append(songDB)

	user := &model.User{}
	utility.Recast(userDB, user)
	return user
}

func (database *Database) GetArtist(id int) *model.Artist {
	artistDB := &Artist{}
	database.DB.Model(&Artist{}).Preload("Songs").First(artistDB, id)

	artist := &model.Artist{}
	utility.Recast(artistDB, artist)
	return artist
}

func (database *Database) GetArtists() []*model.Artist {
	artistsDB := []*Artist{}
	database.DB.Model(&Artist{}).Preload("Songs").Find(&artistsDB)

	artists := []*model.Artist{}
	utility.Recast(&artistsDB, &artists)
	return artists
}

func (database *Database) AddArtist(input model.NewArtist) *model.Artist {
	artistDB := &Artist{
		Name:  input.Name,
		Songs: []*Song{},
	}
	database.DB.Create(artistDB)

	artist := &model.Artist{}
	utility.Recast(artistDB, artist)
	return artist
}

func (database *Database) GetSongs() []*model.Song {
	songsDB := []*Song{}
	database.DB.Model(&Song{}).Preload("Artist").Find(&songsDB)

	songs := []*model.Song{}
	utility.Recast(&songsDB, &songs)
	return songs
}

func (database *Database) AddSong(input model.NewSong) *model.Song {
	songDB := &Song{
		Name:   input.Name,
		ArtistID: input.ArtistID,
	}
	database.DB.Create(songDB)
	database.DB.Model(&Song{}).Preload("Artist").Find(songDB)

	song := &model.Song{}
	utility.Recast(songDB, song)
	return song
}