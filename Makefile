all: build run

build:
	docker build -t fedigram/server:latest .

run: build
	docker-compose up --build -d
