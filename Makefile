.DEFAULT_GOAL = all
all: test vet fmt build

test:
	go test ./...
vet:
	go vet ./...
fmt:
	go fmt ./...
build:
	go build -o bin/ ./...
dev:
	cd cmd/k8s-resource-watcher && nodemon --exec go run . --port=4000 --signal SIGTERM -e go
run:
	cd cmd/k8s-resource-watcher && go run .
