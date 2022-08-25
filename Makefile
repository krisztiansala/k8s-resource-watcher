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
	cd cmd/k8s-resource-watcher && nodemon --exec go run . --signal SIGTERM -e go
run:
	cd cmd/k8s-resource-watcher && go run .
portforward:
	kubectl port-forward service/k8s-resource-watcher-service 8000:8000
tf_local_apply:
	terraform -chdir=terraform/local apply -auto-approve
tf_local_destroy:
	terraform -chdir=terraform/local destroy -auto-approve
	