package storage

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	// Options contains db connection options
	Options struct {
		DBHost string
		DBPort string
		DBName string
		DBUser string
		DBPass string
	}

	// Store provides access to the database and holds Store interface implementations
	driver struct {
		SurveyFormRepository
		SurveyRepository
	}
)

// New sets up a new database repository.
func New(ctx context.Context, opts *Options) (repo Repository, err error) {
	db, err := connect(ctx, opts)
	if err != nil {
		return
	}

	err = seedDB(ctx, db)
	if err != nil {
		return
	}

	repo = &driver{
		SurveyFormRepository: NewSurveyFormRepository(db),
		SurveyRepository:     NewSurveyRepository(db),
	}
	return
}

func connect(ctx context.Context, opts *Options) (db *mongo.Database, err error) {
	log.Println("===== setting up db =====")

	uri := fmt.Sprintf("mongodb://%s:%s", opts.DBHost, opts.DBPort)
	auth := options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    "admin",
		Username:      opts.DBUser,
		Password:      opts.DBPass,
	}

	clientOpts := options.Client().ApplyURI(uri).SetAuth(auth)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return
	}

	err = testConnection(ctx, client)
	if err != nil {
		return
	}

	db = client.Database(opts.DBName)
	return
}

// testConnection tests that the driver can properly connect to the Mongo Server.
func testConnection(ctx context.Context, client *mongo.Client) error {
	log.Println("===== testing connection =====")
	return client.Ping(ctx, nil)
}
