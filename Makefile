all: build run

build:
	docker build -t PluralityServer/server:latest .

run:
	docker-compose up -d