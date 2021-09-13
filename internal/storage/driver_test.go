package storage

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	opts = &Options{
		DBHost: "localhost",
		DBPort: "27017",
		DBName: "survey_db",
		DBUser: "admin",
		DBPass: "secret",
	}
)

func TestDriverNew_integration(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	d, err := New(ctx, opts)

	log.Println(err)
	assert.Nil(t, err)
	assert.NotNil(t, d)
}
