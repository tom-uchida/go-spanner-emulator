package test

import (
	"context"
	"fmt"
	"os"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func setupSpannerEmulator(ctx context.Context) (testcontainers.Container, error) {
	// Spanner Emulator の起動
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "gcr.io/cloud-spanner-emulator/emulator:latest",
			ExposedPorts: []string{"9010/tcp", "9020/tcp"},
			WaitingFor:   wait.ForLog("Cloud Spanner emulator running"),
		},
		Started: true,
	})
	if err != nil {
		return nil, err
	}

	// Spanner Emulator の接続先の設定
	host, _ := container.Host(ctx)
	port, _ := container.MappedPort(ctx, "9010")
	emulatorHost := fmt.Sprintf("%s:%s", host, port.Port())
	os.Setenv("SPANNER_EMULATOR_HOST", emulatorHost)
	fmt.Println("Spanner emulator running at:", emulatorHost)

	return container, nil
}
