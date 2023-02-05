all: build run

build:
	docker build -t pluralityserver/server:latest .

run:
	docker-compose up -d
