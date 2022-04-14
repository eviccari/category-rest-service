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

The GET method search categories using given parameters that have been sent over URI query string. Actually, API supports **ID** and **Name** as query parameters (optionals):

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
