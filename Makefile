all : run test docker
.PHONY : all

run:
	@air

test:
	@go test ./... -v -cover

docker:
	@docker compose up --build --force-recreate -V
