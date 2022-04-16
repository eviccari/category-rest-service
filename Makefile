clean:
	rm -rf dist
	rm -rf category-rest-service-runner

build:
	go mod download
	go install github.com/swaggo/swag/cmd/swag@latest
	go mod verify
	go mod tidy
	swag init
	go build -o category-rest-service-runner . 
	mkdir dist
	mv category-rest-service-runner ./dist
	cp .env dist/

image:
	rm -rf dist/
	docker build -t category-rest-service . 
