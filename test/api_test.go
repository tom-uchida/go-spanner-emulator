package main

import (
	"context"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/tom-uchida/go-api-test/internal"
	"github.com/tom-uchida/go-api-test/internal/db"
)

func TestSomething(t *testing.T) {
	ctx := context.Background()

	parent, err := db.CreateSpannerInstance(ctx)
	if err != nil {
		log.Fatalf("failed to create spanner instance: %v", err)
	}
	dsn, err := db.CreateDatabase(ctx, parent)
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}
	client, err := db.NewClient(ctx, dsn)
	if err != nil {
		log.Fatalf("failed to create spanner client: %v", err)
	}
	defer client.Close()

	mux := internal.NewHandler(client)
	server := httptest.NewServer(mux)
	defer server.Close()
}
