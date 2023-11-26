DOCKER_LINT_IMAGE = golangci/golangci-lint:latest
DOCKER_LINT_FLAGS = --rm -v $(CURDIR):/app -w/app -e GOGC=30
DOCKER_LINT_PARAMS = golangci-lint run -v -c .golangci.yml ./...

docker-lint:
	docker run $(DOCKER_LINT_FLAGS) $(DOCKER_LINT_IMAGE) $(DOCKER_LINT_PARAMS)

lint:
	golangci-lint run ./...

test:
	go test ./... -count=1 -race