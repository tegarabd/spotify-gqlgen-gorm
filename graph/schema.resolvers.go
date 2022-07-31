package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"spotify/graph/generated"
	"spotify/graph/model"
)

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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
