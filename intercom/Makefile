dep:
	dep ensure

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -ldflags '-w -extldflags "-static"' -o bin/server ./cmd/server
	docker build -t intercom -f Dockerfile.server .

run:
	docker run intercom
