package repositories

import (
	"context"
	"fmt"

	dtos "items/dtos"
	model "items/models"
	e "items/utils/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepositoryMongoDB struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection string
}

func NewMongoDB(host string, port int, collection string) *RepositoryMongoDB {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d/", host, port)))
	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}

	names, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}

	fmt.Println("[MongoDB] Initialized connection")
	fmt.Println(fmt.Printf("[MongoDB] Available databases: %s\n", names))

	return &RepositoryMongoDB{
		Client:     client,
		Database:   client.Database("items"),
		Collection: collection,
	}
}

func (repo *RepositoryMongoDB) Get(ctx context.Context, id string) (dtos.ItemDto, e.ApiError) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dtos.ItemDto{}, e.NewBadRequestApiError(fmt.Sprintf("error getting item %s invalid id", id))
	}
	result := repo.Database.Collection(repo.Collection).FindOne(context.TODO(), bson.M{
		"_id": objectID,
	})
	if result.Err() == mongo.ErrNoDocuments {
		return dtos.ItemDto{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
	}
	var item model.Item
	if err := result.Decode(&item); err != nil {
		return dtos.ItemDto{}, e.NewInternalServerApiError(fmt.Sprintf("error getting item %s", id), err)
	}
	return dtos.ItemDto{
		Id:        id,
		Title:     item.Title,
		Price:     item.Price,
		Image:     item.Image,
		Sale_sate: item.Sale_sate,
		Condition: item.Condition,
		Address:   item.Address,
	}, nil
}

func (repo *RepositoryMongoDB) InsertItem(ctx context.Context, item dtos.ItemDto) (dtos.ItemDto, e.ApiError) {
	result, err := repo.Database.Collection(repo.Collection).InsertOne(context.TODO(), model.Item{
		Title:     item.Title,
		UserId:    item.UserId,
		Price:     item.Price,
		Image:     item.Image,
		Sale_sate: item.Sale_sate,
		Condition: item.Condition,
		Address:   item.Address,
	})
	if err != nil {
		return dtos.ItemDto{}, e.NewInternalServerApiError(fmt.Sprintf("error inserting item %s", item.Id), err)
	}
	item.Id = result.InsertedID.(primitive.ObjectID).Hex()
	return item, nil
}

func (repo *RepositoryMongoDB) InsertItems(ctx context.Context, itemsdto dtos.ItemsDto) (dtos.ItemsDto, e.ApiError) {
	var inserteditems dtos.ItemsDto
	var items model.Items

	// Crear lista de modelos de items
	for _, itemdto := range itemsdto {
		items = append(items, model.Item{
			Title:     itemdto.Title,
			UserId:    itemdto.UserId,
			Price:     itemdto.Price,
			Image:     itemdto.Image,
			Sale_sate: itemdto.Sale_sate,
			Condition: itemdto.Condition,
			Address:   itemdto.Address,
		})
	}

	// InsertMany espera como parametro datos de tipo interface
	var itemsInterface []interface{}
	for _, item := range items {
		itemsInterface = append(itemsInterface, item)
	}

	// Insertar los items en la base de datos
	result, err := repo.Database.Collection(repo.Collection).InsertMany(context.TODO(), itemsInterface)
	if err != nil {
		return dtos.ItemsDto{}, e.NewInternalServerApiError("error inserting items", err)
	}

	// Asignar los IDs generados a los items insertados
	for i, item := range itemsdto {
		item.Id = result.InsertedIDs[i].(primitive.ObjectID).Hex()
		inserteditems = append(inserteditems, item)
	}

	return inserteditems, nil
}

func (repo *RepositoryMongoDB) Update(ctx context.Context, item dtos.ItemDto) (dtos.ItemDto, e.ApiError) {
	_, err := repo.Database.Collection(repo.Collection).UpdateByID(context.TODO(), fmt.Sprintf("%v", item.Id), model.Item{
		Title:     item.Title,
		UserId:    item.UserId,
		Price:     item.Price,
		Image:     item.Image,
		Sale_sate: item.Sale_sate,
		Condition: item.Condition,
		Address:   item.Address,
	})
	if err != nil {
		return dtos.ItemDto{}, e.NewInternalServerApiError(fmt.Sprintf("error updating item %s", item.Id), err)
	}
	return item, nil
}

func (repo *RepositoryMongoDB) Delete(ctx context.Context, id string) e.ApiError {
	_, err := repo.Database.Collection(repo.Collection).DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return e.NewInternalServerApiError(fmt.Sprintf("error deleting item %s", id), err)
	}
	return nil
}

func (repo *RepositoryMongoDB) GetByUserId(ctx context.Context, userid int) (model.Items, e.ApiError) {
	cursor, err := repo.Database.Collection(repo.Collection).Find(ctx, bson.M{"userid": userid})
	if err != nil {
		return model.Items{}, e.NewInternalServerApiError("failed to delete items", err)
	}
	defer cursor.Close(ctx)

	var items model.Items
	for cursor.Next(ctx) {
		var item model.Item
		err := cursor.Decode(&item)
		if err != nil {
			return model.Items{}, e.NewInternalServerApiError("failed to decode item", err)
		}

		items = append(items, item)
	}

	if err := cursor.Err(); err != nil {
		return model.Items{}, e.NewInternalServerApiError("cursor error", err)
	}

	return items, nil
}

func (repo *RepositoryMongoDB) DeleteByUserId(ctx context.Context, userid int) e.ApiError {
	_, err := repo.Database.Collection(repo.Collection).DeleteMany(context.TODO(), bson.M{"userid": userid})
	if err != nil {
		return e.NewInternalServerApiError(fmt.Sprintf("error deleting items by user %v", userid), err)
	}
	return nil
}
