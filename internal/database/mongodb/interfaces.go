package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"os"
)

var dbName string

func init() {
	dbName = os.Getenv("MONGO_DB")
}

func getDbName() string {
	if dbName == "" {
		dbName = os.Getenv("MONGO_DB")
	}
	return dbName
}

// QueryOptions represents optional parameters for database operations
type QueryOptions struct {
	Limit *int64
	Sort  interface{}
	Skip  *int64
}

// DatabaseOperations interface defines common database operations
type DatabaseOperations interface {
	// InsertOne inserts a single document into the collection
	InsertOne(ctx context.Context, collection string, document interface{}) (*mongo.InsertOneResult, error)

	// InsertMany inserts multiple documents into the collection
	InsertMany(ctx context.Context, collection string, documents []interface{}) (*mongo.InsertManyResult, error)

	// FindOne finds a single document matching the filter
	FindOne(ctx context.Context, collection string, filter interface{}, result interface{}, opts *QueryOptions) error

	// Find finds all documents matching the filter
	Find(ctx context.Context, collection string, filter interface{}, results interface{}, opts *QueryOptions) error

	// UpdateOne updates a single document matching the filter
	UpdateOne(ctx context.Context, collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error)

	// UpdateMany updates all documents matching the filter
	UpdateMany(ctx context.Context, collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error)

	// DeleteOne deletes a single document matching the filter
	DeleteOne(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error)

	// DeleteMany deletes all documents matching the filter
	DeleteMany(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error)
}

// MongoOperations implements DatabaseOperations interface
type MongoOperations struct{}

// NewMongoOperations creates a new MongoOperations instance
func NewMongoOperations() DatabaseOperations {
	return &MongoOperations{}
}

func (m *MongoOperations) InsertOne(ctx context.Context, collection string, document interface{}) (*mongo.InsertOneResult, error) {
	db := GetClient()
	coll := db.Database(getDbName()).Collection(collection)
	return coll.InsertOne(ctx, document)
}

func (m *MongoOperations) InsertMany(ctx context.Context, collection string, documents []interface{}) (*mongo.InsertManyResult, error) {
	db := GetClient()
	coll := db.Database(getDbName()).Collection(collection)
	return coll.InsertMany(ctx, documents)
}

func (m *MongoOperations) FindOne(ctx context.Context, collection string, filter interface{}, result interface{}, opts *QueryOptions) error {
	db := GetClient()
	coll := db.Database(getDbName()).Collection(collection)

	findOptions := options.FindOne()
	if opts != nil {
		if opts.Sort != nil {
			findOptions.SetSort(opts.Sort)
		}
	}

	return coll.FindOne(ctx, filter, findOptions).Decode(result)
}

func (m *MongoOperations) Find(ctx context.Context, collection string, filter interface{}, results interface{}, opts *QueryOptions) error {
	db := GetClient()
	coll := db.Database(getDbName()).Collection(collection)

	findOptions := options.Find()
	if opts != nil {
		if opts.Limit != nil {
			findOptions.SetLimit(*opts.Limit)
		}
		if opts.Skip != nil {
			findOptions.SetSkip(*opts.Skip)
		}
		if opts.Sort != nil {
			findOptions.SetSort(opts.Sort)
		}
	}

	cursor, err := coll.Find(ctx, filter, findOptions)
	if err != nil {
		return err
	}
	return cursor.All(ctx, results)
}

func (m *MongoOperations) UpdateOne(ctx context.Context, collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	db := GetClient()
	coll := db.Database(getDbName()).Collection(collection)
	return coll.UpdateOne(ctx, filter, update)
}

func (m *MongoOperations) UpdateMany(ctx context.Context, collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	db := GetClient()
	coll := db.Database(getDbName()).Collection(collection)
	return coll.UpdateMany(ctx, filter, update)
}

func (m *MongoOperations) DeleteOne(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error) {
	db := GetClient()
	coll := db.Database(getDbName()).Collection(collection)
	return coll.DeleteOne(ctx, filter)
}

func (m *MongoOperations) DeleteMany(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error) {
	db := GetClient()
	coll := db.Database(getDbName()).Collection(collection)
	return coll.DeleteMany(ctx, filter)
}
