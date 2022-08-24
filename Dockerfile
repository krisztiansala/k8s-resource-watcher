FROM golang:1.18-buster AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY ./ ./
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make build

FROM gcr.io/distroless/static-debian11
WORKDIR /
EXPOSE 8080
USER nonroot:nonroot
COPY --from=build --chown=nonroot:nonroot /app/bin/k8s-resource-watcher /resource-watcher
ENTRYPOINT ["/resource-watcher"]
