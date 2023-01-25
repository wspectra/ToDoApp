all:
	docker-compose   up -d --build

logs:
	docker compose logs
build:
	go build -o todoApp cmd/todoApp/main.go

status:
	docker ps -a

test:
	go test internal/handler -v

clean:
	- docker-compose down

fclean:
	- docker-compose down
	- docker rm -vf $$(docker ps -aq)
	- docker rmi -f $$(docker images -aq)

.PHONY: all build inmemory test clean fclean status