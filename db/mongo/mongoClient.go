package mongo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	guuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/abdulmoeid7112/impact-analysis-api-svc/config"
	"github.com/abdulmoeid7112/impact-analysis-api-svc/db"
	"github.com/abdulmoeid7112/impact-analysis-api-svc/models"
)

const (
	collectionName = "users"
)

func init() {
	db.Register("mongo", NewMongoClient)
}

type client struct {
	conn *mongo.Client
}

// NewMongoClient initializes a mongo database connection.
func NewMongoClient(conf db.Option) (db.DataStore, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s", viper.GetString(config.DbUser), viper.GetString(config.DbPass), viper.GetString(config.DbHost))
	log().Infof("initializing mongodb: %s", uri)
	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect")
	}

	return &client{conn: cli}, nil
}

func (c client) AddUser(context context.Context, user *models.User) (*models.User, error) {
	if user.ID == "" {
		user.ID = guuid.NewV4().String()
	}

	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(collectionName)
	if _, err := collection.UpdateOne(context, bson.M{"_id": user.ID}, bson.M{"$set": user}, options.Update().SetUpsert(true)); err != nil {
		return nil, errors.Wrap(err, "failed to add user")
	}

	return user, nil
}

func (c client) GetUserByName(context context.Context, name string) (*models.User, error) {
	var user *models.User

	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(collectionName)

	if err := collection.FindOne(context, bson.M{"username": name}).Decode(&user); err != nil {
		if mongo.ErrNoDocuments != nil {
			return nil, errors.Wrap(err, "failed to fetch user....not found")
		}

		return nil, err
	}

	return user, nil
}

func (c client) ListUser(context context.Context) ([]*models.User, error) {
	var users []*models.User
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(collectionName)

	cursor, err := collection.Find(context, &options.FindOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve student")
	}

	defer cursor.Close(context)
	for cursor.Next(context) {
		var em *models.User
		if err = cursor.Decode(&em); err != nil {
			return nil, errors.Wrap(err, "failed to scan retrieved rows")
		}

		users = append(users, em)
	}

	return users, nil
}
