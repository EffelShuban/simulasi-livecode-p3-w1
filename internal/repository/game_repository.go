package repository

import (
	"context"
	"errors"
	"sim-livecodep3w1/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GameRepository interface {
	FindAll(ctx context.Context) ([]model.Game, error)
	FindByGameID(ctx context.Context, gameID primitive.ObjectID) (model.Game, error)
	Create(ctx context.Context, game *model.Game) error
	Update(ctx context.Context, game model.Game) error
	Delete(ctx context.Context, gameID primitive.ObjectID) error
	UpdateVersion(ctx context.Context) error
}

type mongodbGameRepository struct {
	gameCollection *mongo.Collection
}

func NewMongodbGameRepository(client *mongo.Client) GameRepository {
	return &mongodbGameRepository{
		gameCollection: client.Database("livecode-sim").Collection("games"),
	}
}

func (r *mongodbGameRepository) FindAll(ctx context.Context) ([]model.Game, error) {
	opts := options.Find().SetSort(bson.D{{Key: "_id", Value: 1}})
	res, err := r.gameCollection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer res.Close(ctx)

	var games []model.Game
	if err = res.All(ctx, &games); err != nil {
		return nil, err
	}

	return games, nil
}

func (r *mongodbGameRepository) FindByGameID(ctx context.Context, gameID primitive.ObjectID) (model.Game, error) {
	var game model.Game
	err := r.gameCollection.FindOne(ctx, bson.M{"_id": gameID}).Decode(&game)
	if err != nil {
		return model.Game{}, err
	}
	if game == (model.Game{}) {
		return model.Game{}, errors.New("game not found")
	}

	return game, nil
}

func (r *mongodbGameRepository) Create(ctx context.Context, game *model.Game) error {
	insertResult, err := r.gameCollection.InsertOne(ctx, game)
	if err != nil {
		return err
	}

	game.GameID = insertResult.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r *mongodbGameRepository) Update(ctx context.Context, game model.Game) error {
	objID, err := primitive.ObjectIDFromHex(game.GameID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	
	game.GameID = "" // Clear ID so it is not included in the $set update (requires omitempty in model)
	update := bson.M{"$set": game}
	res, err := r.gameCollection.UpdateOne(ctx, filter, update)
	if err != nil{
		return err
	}

	if res.MatchedCount == 0{
		return errors.New("game not found")
	}

	return nil
}

func (r *mongodbGameRepository) Delete(ctx context.Context, gameID primitive.ObjectID) error {
	opts := options.Delete().SetCollation(&options.Collation{})
	res, err := r.gameCollection.DeleteOne(ctx, bson.M{"_id": gameID}, opts)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("Game not found")
	}
	return nil
}

func (r *mongodbGameRepository) UpdateVersion(ctx context.Context) error{
	games, err := r.FindAll(ctx)
	if err != nil{
		return err
	}

	for k := range games{
		if games[k].GoToUpdate != nil && *games[k].GoToUpdate == true{
			games[k].UpdatedDate = time.Now()
			games[k].Version = "V.1.0." + string(games[k].Version[6]+1)
			f := false
			games[k].GoToUpdate = &f
			r.Update(ctx, games[k])
		}
	}

	return nil
}