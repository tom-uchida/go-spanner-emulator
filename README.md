# go-api-test

## API test using Go test

1. Run Go test

```shell
go test ./test/... -v
2025/10/12 02:05:14 github.com/testcontainers/testcontainers-go - Connected to docker: 
  Server Version: 25.0.2
  API Version: 1.44
  Operating System: Docker Desktop
  Total Memory: 7941 MB
  Testcontainers for Go Version: v0.39.0
  Resolved Docker Host: unix:///Users/uchidatomomasa/.docker/run/docker.sock
  Resolved Docker Socket Path: /var/run/docker.sock
  Test SessionID: 43879c1cb16a86c8764d8f6f74ce2a00954159357831a7efffc1a8239d947e0d
  Test ProcessID: 3378935c-06a8-43fc-ab0a-4319cc6bfa82
2025/10/12 02:05:14 🐳 Creating container for image gcr.io/cloud-spanner-emulator/emulator:latest
2025/10/12 02:05:14 🐳 Creating container for image testcontainers/ryuk:0.13.0
2025/10/12 02:05:14 ✅ Container created: e69acb697005
2025/10/12 02:05:14 🐳 Starting container: e69acb697005
2025/10/12 02:05:14 ✅ Container started: e69acb697005
2025/10/12 02:05:14 ⏳ Waiting for container id e69acb697005 image: testcontainers/ryuk:0.13.0. Waiting for: &{Port:8080/tcp timeout:<nil> PollInterval:100ms skipInternalCheck:false skipExternalCheck:false}
2025/10/12 02:05:15 Shell not executable in container, only external port validated
2025/10/12 02:05:15 🔔 Container is ready: e69acb697005
2025/10/12 02:05:15 ✅ Container created: 69356a870ae1
2025/10/12 02:05:15 🐳 Starting container: 69356a870ae1
2025/10/12 02:05:15 ✅ Container started: 69356a870ae1
2025/10/12 02:05:15 ⏳ Waiting for container id 69356a870ae1 image: gcr.io/cloud-spanner-emulator/emulator:latest. Waiting for: &{timeout:<nil> Log:Cloud Spanner emulator running IsRegexp:false Occurrence:1 PollInterval:100ms check:<nil> submatchCallback:<nil> re:<nil> log:[]}
2025/10/12 02:05:16 🔔 Container is ready: 69356a870ae1
Spanner emulator running at: localhost:50941
SPANNER_EMULATOR_HOST: localhost:50941
Instance created: projects/test-project/instances/test-instance
=== RUN   TestSomething
=== PAUSE TestSomething
=== CONT  TestSomething
Database created: test-db

=== RUN   TestSomething/runbook/go-test-ver.yaml(01559caa60f32e80a95eb254be15cecbbcd8f405)
{
  "user_id": "87b1036f-dbb7-408d-80ca-b993ea97cca9"
}
{
  "user_id": "a46792c9-740e-4573-ac3b-485bf788a26e"
}
[
  {
    "Name": "test-name-2",
    "UserID": "a46792c9-740e-4573-ac3b-485bf788a26e"
  },
  {
    "Name": "test-name-1",
    "UserID": "87b1036f-dbb7-408d-80ca-b993ea97cca9"
  }
]
--- PASS: TestSomething (0.22s)
    --- PASS: TestSomething/runbook/go-test-ver.yaml(01559caa60f32e80a95eb254be15cecbbcd8f405) (0.12s)
PASS
ok  	github.com/tom-uchida/go-api-test/test	4.209s
```

## API test using Docker containers

1. Start all containers

```shell
> docker compose up --build

[+] Building 12.5s (15/15) FINISHED                                               docker:desktop-linux
 => [api-server internal] load build definition from Dockerfile                                   0.0s
 => => transferring dockerfile: 321B                                                              0.0s
 => [api-server internal] load metadata for docker.io/library/alpine:latest                       1.3s
 => [api-server internal] load metadata for docker.io/library/golang:1.24-alpine                  1.3s
 => [api-server internal] load .dockerignore                                                      0.0s
 => => transferring context: 2B                                                                   0.0s
 => [api-server builder 1/6] FROM docker.io/library/golang:1.24-alpine@sha256:fc2cff6625f3c1c92e  0.0s
 => [api-server stage-1 1/3] FROM docker.io/library/alpine:latest@sha256:4bcff63911fcb4448bd4fda  0.0s
 => [api-server internal] load build context                                                      0.0s
 => => transferring context: 12.19kB                                                              0.0s
 => CACHED [api-server builder 2/6] WORKDIR /app                                                  0.0s
 => CACHED [api-server builder 3/6] COPY go.mod go.sum ./                                         0.0s
 => CACHED [api-server builder 4/6] RUN go mod download                                           0.0s
 => [api-server builder 5/6] COPY . .                                                             0.0s
 => [api-server builder 6/6] RUN go build -o /app/server ./cmd/main.go                           11.0s
 => CACHED [api-server stage-1 2/3] WORKDIR /app                                                  0.0s
 => [api-server stage-1 3/3] COPY --from=builder /app/server .                                    0.0s
 => [api-server] exporting to image                                                               0.1s
 => => exporting layers                                                                           0.1s
 => => writing image sha256:2026b7b66666eb24d7a50485bd577d63a76087c70c28756a90ea93defa28bbe1      0.0s
 => => naming to docker.io/library/build-api-server                                               0.0s
[+] Running 2/2
 ✔ Container build-spanner-emulator-1  Created                                                    0.0s 
 ✔ Container build-api-server-1        Recreated                                                  0.0s 
Attaching to api-server-1, spanner-emulator-1
spanner-emulator-1  | WARNING: proto: file "google/rpc/status.proto" is already registered
spanner-emulator-1  | 	previously from: "google.golang.org/genproto/googleapis/rpc/status"
spanner-emulator-1  | 	currently from:  "unknown"
spanner-emulator-1  | See https://protobuf.dev/reference/go/faq#namespace-conflict
spanner-emulator-1  | 
spanner-emulator-1  | WARNING: proto: file "google/rpc/status.proto" has a name conflict over google.rpc.Status
spanner-emulator-1  | 	previously from: "google.golang.org/genproto/googleapis/rpc/status"
spanner-emulator-1  | 	currently from:  "unknown"
spanner-emulator-1  | See https://protobuf.dev/reference/go/faq#namespace-conflict
spanner-emulator-1  | 
spanner-emulator-1  | WARNING: proto: message google.rpc.Status is already registered
spanner-emulator-1  | 	previously from: "google.golang.org/genproto/googleapis/rpc/status"
spanner-emulator-1  | 	currently from:  "unknown"
spanner-emulator-1  | See https://protobuf.dev/reference/go/faq#namespace-conflict
spanner-emulator-1  | 
spanner-emulator-1  | WARNING: All log messages before absl::InitializeLog() is called are written to STDERR
spanner-emulator-1  | I0000 00:00:1759591537.337902      11 emulator_main.cc:39] Cloud Spanner Emulator running.
spanner-emulator-1  | I0000 00:00:1759591537.337921      11 emulator_main.cc:40] Server address: 0.0.0.0:9010
api-server-1        | SPANNER_EMULATOR_HOST: spanner-emulator:9010
api-server-1        | Instance created: projects/test-project/instances/test-instance
api-server-1        | Database created: test-db
api-server-1        | 
api-server-1        | 2025/10/04 15:25:37 
api-server-1        | Server running at: localhost:8080
spanner-emulator-1  | 2025/10/04 15:25:38 gateway.go:151: Cloud Spanner emulator running.
spanner-emulator-1  | 2025/10/04 15:25:38 gateway.go:152: REST server listening at 0.0.0.0:9020
spanner-emulator-1  | 2025/10/04 15:25:38 gateway.go:153: gRPC server listening at 0.0.0.0:9010
```

2. Run runn

```shell
> SPANNER_EMULATOR_HOST=0.0.0.0:9010 runn run runbook/docker-containers-ver.yaml
{
  "user_id": "7b87e886-a73c-4820-a079-f44d12099d09"
}
{
  "user_id": "0813020f-65f4-47fc-adc1-4c244f80f8fb"
}
[
  {
    "Name": "test-name-1",
    "UserID": "7b87e886-a73c-4820-a079-f44d12099d09"
  },
  {
    "Name": "test-name-2",
    "UserID": "0813020f-65f4-47fc-adc1-4c244f80f8fb"
  }
]
.

1 scenario, 0 skipped, 0 failures
```
