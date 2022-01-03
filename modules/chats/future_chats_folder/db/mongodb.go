package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/zhashkevych/todo-app/modules/chats/chat_naw"
	"github.com/zhashkevych/todo-app/modules/chats/future_chats_folder"
	"github.com/zhashkevych/todo-app/pcg/loging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db_chat struct {
	collection *mongo.Collection
	logger     loging.Logger
}


func (d *db_chat) FindAllChat(ctx context.Context) (u []chat_naw.Chat, err error) {
	panic("implement me")
}

func (d *db_chat) FindOneChat(ctx context.Context, chat_id string) (chat_naw.Chat, error) {
	panic("implement me")
}

func (d *db_chat) UpdateChat(ctx context.Context, chat chat_naw.Chat) error {
	panic("implement me")
}

func (d *db_chat) DeleteChat(ctx context.Context, chat_id string) error {
	panic("implement me")
}

func (d *db_chat) CreateChat(ctx context.Context,chats chat_naw.Chat) (string, error) {
	d.logger.Debug("create chat")
	result, err := d.collection.InsertOne(ctx, chats)
	if err != nil {
		return "", fmt.Errorf("failed to create user due to error: %v", err)
	}

	d.logger.Debug("convert InsertedID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(chats)
	return "", fmt.Errorf("failed to convert objectid to hex. probably oid: %s", oid)
}

func (d *db_chat) FindAll(ctx context.Context) (u []chat_naw.Chat, err error) {
	cursor, err := d.collection.Find(ctx, bson.M{})
	if cursor.Err() != nil {
		return u, fmt.Errorf("failed to find all users due to error: %v", err)
	}

	if err = cursor.All(ctx, &u); err != nil {
		return u, fmt.Errorf("failed to read all documents from cursor. error: %v", err)
	}

	return u, nil
}

func (d *db_chat) FindOne(ctx context.Context, id string) (u chat_naw.Chat, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to objectid. hex: %s", id)
	}

	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return u, nil
		}
		return u, fmt.Errorf("failed to find one user by id: %s due to error: %v", id, err)
	}
	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user (id:%s) from DB due to error: %v", id, err)
	}

	return u, nil
}

func (d *db_chat) Update(ctx context.Context,chat chat_naw.Chat) error {
	objectID, err := primitive.ObjectIDFromHex(chat.ChatID)
	if err != nil {
		return fmt.Errorf("failed to convert user ID to ObjectID. ID=%s", chat.ChatID)
	}

	filter := bson.M{"_id": objectID}

	chatBytes, err := bson.Marshal(chat)
	if err != nil {
		return fmt.Errorf("failed to marhsal user. error: %v", err)
	}

	var updateChatObj bson.M
	err = bson.Unmarshal(chatBytes, &updateChatObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal user bytes. error: %v", err)
	}

	delete(updateChatObj, "_id")

	update := bson.M{
		"$set": updateChatObj,
	}

	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute update user query. error: %v", err)
	}

	if result.MatchedCount == 0 {
		return nil
	}

	d.logger.Tracef("Matched %d documents and Modified %d documents", result.MatchedCount, result.ModifiedCount)

	return nil

}

func (d *db_chat) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert user ID to ObjectID. ID=%s", id)
	}

	filter := bson.M{"_id": objectID}

	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %v", err)
	}
	if result.DeletedCount == 0 {
		return nil
	}
	d.logger.Tracef("Deleted %d documents", result.DeletedCount)

	return nil
}

func NewStorage(database *mongo.Database, collection string, logger loging.Logger) future_chats_folder.StorageChat {
	return &db_chat{
		collection: database.Collection(collection),
		logger:     logger,
	}
}