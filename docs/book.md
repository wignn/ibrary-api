# Book API Spec

## Create Book

Endpoint: POST /api/v1/books

Request Header :

- authorization : TOKEN

Request Body :

```json
{
  "cover": "file://exemple.com",
  "title": "no longer human",
  "author": "dazai osamu",
  "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,",
  "published_date":"22-12-2021"
}
```

Response Body (Success) :

```json
{
  "message": "Book created successfully"
}
```

Response Body (Failed) :

```json
"errors": "Internal server error"
```

## Get By Id

GET /api/v1/book/:id

Request Header :

- authorization : TOKEN

Response Body (Success) :

```json
{
    "id": "cm2qyq4pp0000d3g9bzmdwykt",
    "title": "no longer human",
    "author": "dazai osamu",
    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,",
    "published_date":"22-12-2021"
}
```

Response Body (Failed) :

```json
{
    "errors": "Internal server error"
}
```

## Get List

GET /api/v1/books

Request Header :

- authorization : TOKEN

Response Body (Success) :

```json
{
  "data": [
    {
      "id": "cm2qyq4pp0000d3g9bzmdwykt",
      "title": "no longer human",
      "author": "dazai osamu",
      "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,",
      "published_date":"22-12-2021"
    },
    {
      "id": "cm2qyq4pp0000d3g9bzmdwykt",
      "title": "no longer human2",
      "author": "dazai osamu",
      "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,",
      "published_date":"22-12-2021"
    }
  ]
}
```

Response Body (Failed) :

```json
{
    "errors": "Internal server error"
}
```

## Update

PUT /api/v1/books/:id

Request Header :

- authorization : TOKEN

Request Body :

```json
{
  "cover": "file://exemple.com",/optional
  "title": "no longer human",/optional
  "author": "dazai osamu",/optional
  "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,"/optional,
  "published_date":"22-12-2021"/optional

}
```

Response Body (Success) :

```json
{
    "message": "Book updated successfully"
}
```

Response Body (Failed) :

```json
{
    "errors": "Internal server error"
}
```

## Delete

DELETE /api/books/:id

- authorization : TOKEN

Response Body (Success) :

```json
{
    "message": "Book deleted successfully"
}
```

Response Body (Failed) :

```json
{
    "errors": "Internal server error"
}
```
