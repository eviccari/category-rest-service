docker run -d --rm -p 27017:27017 --name mongo-local -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=123 mongo:latest