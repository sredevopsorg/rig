package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	databaseIDIndex = "database_id_idx"
	nameIndex       = "database_name_idx"
)

type MongoRepository struct {
	DatabaseCollection *mongo.Collection
}

func (r *MongoRepository) BuildIndexes(ctx context.Context) error {
	databaseIdIndexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "database_id", Value: 1},
		},
		Options: options.Index().SetName(databaseIDIndex).SetUnique(true),
	}
	if _, err := r.DatabaseCollection.Indexes().CreateOne(ctx, databaseIdIndexModel); err != nil {
		return err
	}
	nameIndexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "project_id", Value: 1},
			{Key: "name", Value: 1},
		},
		Options: options.Index().SetName(nameIndex).SetUnique(true),
	}
	if _, err := r.DatabaseCollection.Indexes().CreateOne(ctx, nameIndexModel); err != nil {
		return err
	}
	return nil
}

func NewRepository(c *mongo.Client) (*MongoRepository, error) {
	repo := &MongoRepository{
		DatabaseCollection: c.Database("rig").Collection("databases"),
	}
	err := repo.BuildIndexes(context.Background())
	if err != nil {
		return nil, err
	}
	return repo, nil
}
