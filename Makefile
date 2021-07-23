all: build run

build:
	docker build -t chatengine/server:latest .

run: build
	docker-compose up --build -d