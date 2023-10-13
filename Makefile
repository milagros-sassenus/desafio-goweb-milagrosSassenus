build:
	docker build . -t go-containerized:latest

run:
  docker run -e PORT=8080 -p 8080:8080 go-containerized:latest

run:
  docker compose up

