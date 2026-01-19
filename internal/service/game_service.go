package service

import (
	"context"
	"sim-livecodep3w1/internal/model"
	"sim-livecodep3w1/internal/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameService interface{
	FindAll(ctx context.Context) ([]model.Game, error)
	FindByID(ctx context.Context, gameID primitive.ObjectID) (model.Game, error)
	Create(ctx context.Context, req model.GameCreateRequest) (model.Game, error)
	Update(ctx context.Context, req model.GameUpdateRequest, gameId primitive.ObjectID) (error)
	Delete(ctx context.Context, gameID primitive.ObjectID)error
	UpdateVersion(ctx context.Context) error
}

type gameService struct{
	repo repository.GameRepository
}

func NewGameService(repo repository.GameRepository) GameService{
	return &gameService{
		repo: repo,
	}
}

func (s *gameService) FindAll(ctx context.Context) ([]model.Game, error){
	return s.repo.FindAll(ctx)
}

func (s *gameService) FindByID(ctx context.Context, gameID primitive.ObjectID) (model.Game, error){
	return s.repo.FindByGameID(ctx, gameID)
}

func (s *gameService) Create(ctx context.Context, req model.GameCreateRequest) (model.Game, error){
	game := model.Game{
		Title: req.Title,
		Description: req.Description,
		ReleaseDate: req.ReleaseDate,
		Version: req.Version,
		Platform: req.Platform,
		GoToUpdate: &req.GoToUpdate,
	}
	err := s.repo.Create(ctx, &game)
	if err != nil{
		return model.Game{}, err
	}

	return game, nil
}

func (s *gameService) Update(ctx context.Context, req model.GameUpdateRequest, gameId primitive.ObjectID) (error){
	game := model.Game{
		Title: req.Title,
		Description: req.Description,
		Platform: req.Platform,
		GoToUpdate: req.GoToUpdate,
		GameID: gameId.Hex(),
		UpdatedDate: time.Now(),
	}

	return s.repo.Update(ctx, game)
}

func (s *gameService) Delete(ctx context.Context, gameID primitive.ObjectID)error{
	return s.repo.Delete(ctx, gameID)
}

func (s *gameService) UpdateVersion(ctx context.Context) error{
	return s.repo.UpdateVersion(ctx)
}