package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"spotify/graph/generated"
	"spotify/graph/model"
)

// Songs is the resolver for the songs field.
func (r *artistResolver) Songs(ctx context.Context, obj *model.Artist) ([]*model.Song, error) {
	return r.DB.GetSongsByArtist(obj), nil
}

// AddSong is the resolver for the addSong field.
func (r *mutationResolver) AddSong(ctx context.Context, input model.NewSong) (*model.Song, error) {
	return r.DB.AddSong(input), nil
}

// AddArtist is the resolver for the addArtist field.
func (r *mutationResolver) AddArtist(ctx context.Context, input model.NewArtist) (*model.Artist, error) {
	return r.DB.AddArtist(input), nil
}

// AddUser is the resolver for the addUser field.
func (r *mutationResolver) AddUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.DB.AddUser(input), nil
}

// AddUserFavoriteSong is the resolver for the addUserFavoriteSong field.
func (r *mutationResolver) AddUserFavoriteSong(ctx context.Context, input model.NewUserFavoriteSong) (*model.User, error) {
	return r.DB.AddUserFavoriteSong(input), nil
}

// GetSongs is the resolver for the getSongs field.
func (r *queryResolver) GetSongs(ctx context.Context) ([]*model.Song, error) {
	return r.DB.GetSongs(), nil
}

// GetArtists is the resolver for the getArtists field.
func (r *queryResolver) GetArtists(ctx context.Context) ([]*model.Artist, error) {
	return r.DB.GetArtists(), nil
}

// GetUsers is the resolver for the getUsers field.
func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	return r.DB.GetUsers(), nil
}

// Artist is the resolver for the artist field.
func (r *songResolver) Artist(ctx context.Context, obj *model.Song) (*model.Artist, error) {
	return r.DB.GetArtistBySong(obj), nil
}

// Artist returns generated.ArtistResolver implementation.
func (r *Resolver) Artist() generated.ArtistResolver { return &artistResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Song returns generated.SongResolver implementation.
func (r *Resolver) Song() generated.SongResolver { return &songResolver{r} }

type artistResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type songResolver struct{ *Resolver }
