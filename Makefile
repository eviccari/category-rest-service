clean:
	rm -rf ./dist
	rm -rf category-rest-service-runner

build:
	go mod download
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init
	go build -o category-rest-service-runner . 
	mkdir ./dist
	mv category-rest-service-runner ./dist
	cp .env ./dist
