package model

import "time"

type Game struct {
	GameID      string    `json:"_id" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"title,omitempty"`
	Description string    `json:"description" bson:"description,omitempty"`
	ReleaseDate time.Time `json:"release_date" bson:"release_date,omitempty"`
	Version     string    `json:"version" bson:"version,omitempty"`
	Platform    string    `json:"platform" bson:"platform,omitempty"`
	UpdatedDate time.Time `json:"updated_date" bson:"updated_date,omitempty"`
	GoToUpdate  *bool     `json:"go_to_update" bson:"go_to_update,omitempty"`
}

type GameCreateRequest struct {
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	ReleaseDate time.Time `json:"release_date" bson:"release_date"`
	Version     string    `json:"version" bson:"version"`
	Platform    string    `json:"platform" bson:"platform"`
	GoToUpdate  bool      `json:"go_to_update" bson:"go_to_update"`
}

type GameUpdateRequest struct {
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Platform    string    `json:"platform" bson:"platform"`
	GoToUpdate  *bool     `json:"go_to_update" bson:"go_to_update"`
}
