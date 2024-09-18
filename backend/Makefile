run:
	@templ generate
	@go run cmd/api/main.go

build:
	# @templ generate
	@go build -o bin/main cmd/api/main.go

docker:
	@docker build -t web .

drun:
	@docker run --rm -it -p 8000:8000 web

stress:
	go-wrk -c 80 -d 5 http://localhost:8000/projects
