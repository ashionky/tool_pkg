/**
 * @Author pibing
 * @create 2022/1/24 4:24 PM
 */

package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// conf
type Conf struct {
	Address       string
	MaxPoolSize   uint64
	SocketTimeout int
}

type mongoCollection struct {
	Collection *mongo.Collection
}

var mongoClient *mongo.Client

type B bson.M
type D bson.D
type Pipeline mongo.Pipeline

//string to objectId
func ObjectIDFromString(id string) (primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	return objectId, err
}

func NewClient(cfg *Conf) {
	clientOptions := options.Client().ApplyURI(cfg.Address)

	clientOptions.SetMaxPoolSize(cfg.MaxPoolSize)

	clientOptions.SetSocketTimeout(time.Duration(cfg.SocketTimeout) * time.Second)

	// Connect to MongoDB
	db, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	err = db.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	mongoClient = db
	fmt.Println("mongo is connected!")
}

//Operation entrance
func MongoDBCurd(database string, collection string, opts ...*options.CollectionOptions) *mongoCollection {
	if database == "" {
		panic("database is required")
	}
	dbCollection := getCollection(database, collection, opts...)
	return &mongoCollection{Collection: dbCollection}
}

func getCollection(database string, collection string, opts ...*options.CollectionOptions) *mongo.Collection {
	return mongoClient.Database(database).Collection(collection, opts...)
}

func GetClient(database string) *mongo.Client {
	return mongoClient.Database(database).Client()
}

func (m *mongoCollection) FindOne(ctx context.Context, filter interface{}, result interface{}, opts ...*options.FindOneOptions) (err error) {
	err = m.Collection.FindOne(ctx, filter, opts...).Decode(result)
	return
}

func (m *mongoCollection) Find(ctx context.Context, filter interface{}, result interface{}, opts ...*options.FindOptions) (err error) {
	find, err := m.Collection.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	return find.All(ctx, result)
}

func (m *mongoCollection) Insert(ctx context.Context, data []interface{}, opts ...*options.InsertManyOptions) (err error) {
	_, err = m.Collection.InsertMany(ctx, data, opts...)
	return
}

func (m *mongoCollection) InsertOne(ctx context.Context, data interface{}, opts ...*options.InsertOneOptions) (err error) {
	_, err = m.Collection.InsertOne(ctx, data, opts...)
	return
}

func (m *mongoCollection) Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (count int64, err error) {
	return m.Collection.CountDocuments(ctx, filter, opts...)
}

func (m *mongoCollection) UpdateOne(ctx context.Context, filter, update interface{}, opts ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	return m.Collection.UpdateOne(ctx, filter, update, opts...)
}

func (m *mongoCollection) UpSertOne(ctx context.Context, filter, update interface{}) (result *mongo.UpdateResult, err error) {
	opts := &options.UpdateOptions{}
	opts.SetUpsert(true)
	return m.Collection.UpdateOne(ctx, filter, update, opts)
}

func (m *mongoCollection) UpdateMany(ctx context.Context, filter, update interface{}, opts ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	return m.Collection.UpdateMany(ctx, filter, update, opts...)
}

func (m *mongoCollection) UpdateByID(ctx context.Context, id, update interface{}, opts ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	return m.Collection.UpdateByID(ctx, id, update, opts...)
}

func (m *mongoCollection) FindOneAndUpdate(ctx context.Context, filter, update interface{}, opts ...*options.FindOneAndUpdateOptions) (result *mongo.SingleResult) {
	return m.Collection.FindOneAndUpdate(ctx, filter, update, opts...)
}

func (m *mongoCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (result *mongo.DeleteResult, err error) {
	return m.Collection.DeleteOne(ctx, filter, opts...)
}

func (m *mongoCollection) FindOneAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) (result *mongo.SingleResult) {
	return m.Collection.FindOneAndDelete(ctx, filter, opts...)
}

func (m *mongoCollection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (result *mongo.DeleteResult, err error) {
	return m.Collection.DeleteMany(ctx, filter, opts...)
}

//bulk，pairs Length must be even
func (m *mongoCollection) BulkUpdate(ctx context.Context, pairs []interface{}) (result *mongo.BulkWriteResult, err error) {
	if len(pairs) == 0 {
		return nil, nil
	}
	writeModel := make([]mongo.WriteModel, 0)
	for i := 0; i < len(pairs); i += 2 {
		model := mongo.NewUpdateOneModel().SetFilter(pairs[i]).SetUpdate(pairs[i+1]).SetUpsert(true)
		writeModel = append(writeModel, model)
	}
	opts := &options.BulkWriteOptions{}
	opts.SetOrdered(false)

	return m.Collection.BulkWrite(ctx, writeModel, opts)
}

func (m *mongoCollection) Aggregate(ctx context.Context, pipeline interface{}, result interface{}, opts ...*options.AggregateOptions) error {
	cursor, err := m.Collection.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		return err
	}
	return cursor.All(ctx, result)
}
