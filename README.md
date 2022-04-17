# Category REST Service

### Demo API to learn about gin framework and golang

#

## Marjory Dependencies

- Golang 1.18
- MongoDB 5.0.6 Community Edition

#

## Description

Basic REST API to handle C.R.U.D. operations (**C**reate, **R**ead, **U**pdate and **D**elete) over Category Entity.

The category data structure looks like the follow:

```json
{
  "ID": "string",
  "Name": "string",
  "CreatedAt": "string",
  "UpdatedAt": "string"
}
```

To provide flexibility to change repository provider, the API use custom **ID** instead MongoDB.ObjectID. So, from client perspective, the entity identification should be ID as described before. Internally, the collection must provide a unique index with ID field to guarantee that behavior.

#

## Building Service

To build the executable service, make sure that the follow commands will be running into Unix terminal on the root of the project folder:

```bash
make clean && make build
```

These commands will create a **dist** folder with file **category-rest-service-runner**.

To run executable, create new file into dist folder with structure based on .env.example to configure MongoDB access properties:

```properties
MONGODB_HOST=YOUR_MONGODB_HOST
MONGODB_PORT=YOUR_MONGODB_PORT
MONGODB_PASS=YOUR_MONGODB_PASS
MONGODB_USER=YOUR_MONGODB_USER
MONGODB_COLL=categories
MONGODB_REPO=repositories
```

Run service:

```bash
cd ./dist
./category-rest-service-runner
```

#

## Prepare MongoDB Collection

Personally, I prefer MongoDB Compass (*https://www.mongodb.com/try/download/compass*) to manage collections, but feel free to use your favorite tool.

- **Running MongoDB with docker container**:
  On terminal:

```bash
docker run -d --rm -p 27017:27017 --name mongo-local -e MONGO_INITDB_ROOT_USERNAME=YOUR_MONGODB_USER -e MONGO_INITDB_ROOT_PASSWORD=YOUR_MONGODB_PASS mongo:latest
```

- **Create new connection:**

![Alt text](./images/create_connection.png?raw=true "Create DB Connection")

- **Create new repository called _repositories_ and collection _categories_ :**
  ![Alt text](./images/create_repository_and_collection.png?raw=true "Create repository and collection")

- **Create unique index category_index_by_ID:**
  Necessary because the category identifier must be provided and controlled by API instead _MongoDB ObjectID_ (avoid vendor lock in):

![Alt text](./images/create_index_by_id.png?raw=true "Create unique index by custom ID")

#

## Common Methods

- ### **/categories/v1/ -> POST:** Create new Category

To Create new category, clients must send the name and API should manage ID and create date time.

**Request:**

```bash
curl --request POST "http://localhost:8080/categories/v1/" \
  --header 'Content-Type: application/json' \
  --header 'Accept: application/json' \
  --data '{"name": "Sample Toys"}'
```

**Response**

```json
{
  "ID": "bd9def8f-71af-4f12-bfd3-cd648e31ab52",
  "Name": "Sample Toys",
  "CreatedAt": "2022-04-14T15:20:00.943459-03:00",
  "UpdatedAt": "2022-04-14T15:20:00.943459-03:00"
}
```

- ### **/categories/v1/ -> GET:** Get category(ies)

The GET method search categories using given parameters that have been sent over URI query string. Actually, API supports **ID** and **Name** as query parameters (optional):

**Request**

```bash
curl -X GET "http://localhost:8080/categories/v1/?ID=bd9def8f-71af-4f12-bfd3-cd648e31ab52" \
 --header 'Content-Type: application/json' \
 --header 'Accept: application/json'
```

**Response**

```json
[
  {
    "ID": "bd9def8f-71af-4f12-bfd3-cd648e31ab52",
    "Name": "Sample Toys",
    "CreatedAt": "2022-04-14T18:20:00.943Z",
    "UpdatedAt": "2022-04-14T18:20:00.943Z"
  }
]
```

- ### **/categories/v1/?ID={category id} -> PATCH:** Update (Override) category

Update single category using given **ID**. The request body must contain JSON payload with new **Name**:

**Request**

```bash
curl --request PATCH "http://localhost:8080/categories/v1/?ID=bd9def8f-71af-4f12-bfd3-cd648e31ab52" \
  --header 'Content-Type: application/json' \
  --header 'Accept: application/json' \
  --data '{"name": "NEW Sample Toys"}'
```

**Response**

```json
{
  "ID": "bd9def8f-71af-4f12-bfd3-cd648e31ab52",
  "Name": "NEW Sample Toys",
  "CreatedAt": "2022-04-14T18:20:00.943Z",
  "UpdatedAt": "2022-04-14T16:19:22.225603-03:00"
}
```

- ### **/categories/v1/?ID={category id} -> DELETE:** Delete category

Delete category using given **ID**:

**Request**

```bash
curl --request DELETE "http://localhost:8080/categories/v1/?ID=bd9def8f-71af-4f12-bfd3-cd648e31ab52"
```

**Response**

```json
{
  "Message": "category with id bd9def8f-71af-4f12-bfd3-cd648e31ab52 deleted"
}
```
