package test

import (
	"context"
	"database/sql"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/googleapis/go-sql-spanner"
	"github.com/k1LoW/runn"
	"github.com/tom-uchida/go-api-test/internal"
	"github.com/tom-uchida/go-api-test/internal/db"
)

const (
	projectID  = "test-project"
	instanceID = "test-instance"
	dbName     = "test-db"
)

var parent string

func TestMain(m *testing.M) {
	ctx := context.Background()

	os.Setenv("PROJECT_ID", projectID)
	os.Setenv("INSTANCE_ID", instanceID)

	if container, err := setupSpannerEmulator(ctx); err != nil {
		log.Fatalf("failed to start spanner emulator: %v", err)
	} else {
		defer container.Terminate(ctx)
	}

	var err error
	parent, err = db.CreateSpannerInstance(ctx)
	if err != nil {
		log.Fatalf("failed to create spanner instance: %v", err)
	}

	os.Exit(m.Run())
}

func Test_CreateUser(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	os.Setenv("DB_NAME", dbName)

	dsn, err := db.CreateDatabase(ctx, parent)
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}
	client, err := db.NewClient(ctx, dsn)
	if err != nil {
		log.Fatalf("failed to create spanner client: %v", err)
	}
	defer client.Close()

	// APIテスト用のサーバー起動
	mux := internal.NewHandler(client)
	server := httptest.NewServer(mux)
	defer server.Close()

	// APIテスト自体はrunn側で実施する
	sqlDB, err := sql.Open("spanner", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer sqlDB.Close()
	o, err := runn.Load("runbook/go-test-ver.yaml", []runn.Option{
		runn.T(t),
		runn.Runner("req", server.URL),
		runn.DBRunner("dsn", sqlDB),
	}...)
	if err != nil {
		t.Fatal(err)
	}
	if err := o.RunN(ctx); err != nil {
		t.Fatal(err)
	}
}
