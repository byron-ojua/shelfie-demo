package db

import (
	"context"
)

// Book is a struct representing a book document
type Book struct {
	ISBN        string `json:"isbn"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishYear string `json:"publish_year"`
	Publisher   string `json:"publisher"`
	ImageURLS   string `json:"image_url_s"`
	ImageURLM   string `json:"image_url_m"`
	ImageURLL   string `json:"image_url_l"`
	ID          string `json:"id"`
}

type GetBooksInput struct {
	Limit int
	Page  int
}

func (env *env) GetBooks(ctx context.Context) ([]Book, error) {
	env.logger.Info("Getting books")

	// Query for books
	// query := "SELECT * FROM c"
	// iter := env.client.QueryItems(ctx, query, nil)

	return nil, nil
}
