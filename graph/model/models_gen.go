// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Artist struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Songs []*Song `json:"songs"`
}

type NewArtist struct {
	Name string `json:"name"`
}

type NewSong struct {
	Name     string `json:"name"`
	ArtistID int    `json:"artistId"`
}

type NewUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type NewUserFavoriteSong struct {
	UserID int `json:"userId"`
	SongID int `json:"songId"`
}

type Song struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Artist *Artist `json:"artist"`
}

type User struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	FavoriteSongs []*Song `json:"favoriteSongs"`
}
