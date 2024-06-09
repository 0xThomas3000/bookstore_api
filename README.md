## How To Run This Project

#### Run with Docker

- Install Docker.
- Run the command below(to download + start docker container):

  docker run -d --name bookstore --privileged=true -e MYSQL_ROOT_PASSWORD="admin" -e MYSQL_USER="user" -e MYSQL_PASSWORD="1234567" -e MYSQL_DATABASE="bookstore" -p 3306:3306 mysql:latest

- Start MYSQL(use MYSQL connection information which is stored in 'app.env' file)
- Create the 'books' table (run the file "bookstore.sql" in 'util' folder in MYSQL)
- At the root folder of the project, run the command: `go run .`
- Access APIs using `http://localhost:8080`

### API Request and Response

1. Upload file

- Request: (`/Users/thomas/Desktop/harry_potter.jpg` below is as the path to the .jpg file)

```
curl --location 'http://localhost:8080/g1/upload' \
--form 'file=@"/Users/thomas/Desktop/harry_potter.jpg"'
```

- Response: (the local .jpg will be uploaded and stored into './static' folder)

```json
{
  "status_code": 200,
  "message": "OK",
  "data": {
    "url": "http://localhost:8080/static/book1.jpg"
  }
}
```

2. Add a book

- Request:

```
curl --location 'http://localhost:8080/g1/books' \
--header 'Content-Type: application/json' \
--data '{
  "title": "Harry Potter",
  "author": "J. K. Rowling",
  "isbn": "3456753644",
  "published_date": "1997-06-26",
  "number_of_pages": 223,
  "cover_image": {
      "url": "http://localhost:8080/static/harry_potter.jpg"
  },
  "language": "English"
}'
```

- Response

```json
{
  "id": "DCZjYGxLJLZzE4k",
  "log": "book successfully added",
  "message": "created",
  "status_code": 201
}
```

- Note: the `url` for `cover_image` is pointed to the image file path in the `static` folder. Additionally, the `id` returned to display publicly in `json` will be encoded for security (except `id` stored into the database)

3. List all books

- Request:

```
curl --location 'http://localhost:8080/g1/books'
```

- Response

```json
{
  "status_code": 200,
  "message": "OK",
  "paging": {
    "page": 1,
    "limit": 50,
    "total": 30,
    "cursor": "",
    "next_cursor": "e532qos8jjM2"
  },
  "data": [
    {
      "id": "DCZjYGxLJLZzE4k",
      "created_at": "2024-06-10T00:05:02+07:00",
      "updated_at": "2024-06-10T00:05:02+07:00",
      "title": "Harry Potter",
      "author": "J. K. Rowling",
      "published_date": "1997-06-26T00:00:00+07:00",
      "isbn": "3456753644",
      "number_of_pages": 223,
      "cover_image": {
        "url": "http://localhost:8080/static/harry_potter.jpg"
      },
      "language": "English"
    },
    {
      "id": "DCWW2AWgRZf7Ltn",
      "created_at": "2024-06-09T18:06:15+07:00",
      "updated_at": "2024-06-09T18:06:15+07:00",
      "title": "Introduction to java",
      "author": "Y. Daniel Liang",
      "published_date": "2014-07-03T00:00:00+07:00",
      "isbn": "3456753644",
      "number_of_pages": 1000,
      "cover_image": {
        "url": "http://localhost:8080/static/intro_to_java.jpg"
      },
      "language": "efwef"
    }
    .......
  ]
}
```

- Note: we can add a query param `limit` for paging (curl --location 'http://localhost:8080/g1/books?limit=2')

4. Find a book by `id`

- Request:

```
curl --location 'http://localhost:8080/g1/books/DCZjYGxLJLZzE4k'
```

- Response:

```json
{
  "status_code": 200,
  "message": "OK",
  "data": {
    "id": "DCZjYGxLJLZzE4k",
    "created_at": "2024-06-10T00:05:02+07:00",
    "updated_at": "2024-06-10T00:05:02+07:00",
    "title": "Harry Potter",
    "author": "J. K. Rowling",
    "published_date": "1997-06-26T00:00:00+07:00",
    "isbn": "3456753644",
    "number_of_pages": 223,
    "cover_image": {
      "url": "http://localhost:8080/static/harry_potter.jpg"
    },
    "language": "English"
  }
}
```

- Note: the `id` input param must be encoded to display in json (tip: we can retrieve all the encoded `id` of each book by using the API listing the books' list above)

5. Update a book by `id`

- Request:

```
curl --location --request PUT 'http://localhost:8080/g1/books/DCZjYGxLJLZzE4k' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Truyen Harry Potter",
    "author": "J. K. Rowling",
    "isbn": "4567545678",
    "published_date": "2023-11-04",
    "number_of_pages": 300,
    "cover_image": {
        "url": "http://localhost:8080/static/vietnamese_harry_potter.jpg"
    },
    "language": "Vietnamese"
}'
```

- Response:

```json
{
  "status_code": 200,
  "message": "OK",
  "log": "book information updated"
}
```

6. Delete a book by `id`

- Request:

```
curl --location --request DELETE 'http://localhost:8080/g1/books/DCZjYGxLJLZzE4k'
```

### Unit-test

In this project, we have some mock code for entities, business layer:

- Therefore, we can then navigate to each of these folders: `entities`, `business` to run `go test .` command for unit testing
