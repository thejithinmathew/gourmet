package clients

import (
	"context"

	"github.com/thejithinmathew/gourmet/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Clients struct {
	DBClient *mongo.Client
}

func Init(c *config.Config) (*Clients, error) {
	opts := options.Client().
		ApplyURI(c.URI).
		SetAuth(options.Credential{
			Username: c.User,
			Password: c.Pass,
		})
	mClient, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	return &Clients{
		DBClient: mClient,
	}, nil
}
