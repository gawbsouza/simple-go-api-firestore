## A simple library API
This is a simple CRUD in Go Lang and Firebase

Start application:
```bash
docker-compose up
```

Endpoints:


```
GET		/books/
GET		/books/<id>
POST	/books
PUT		/books/<id>
DELETE	/books/<id>
```

Book JSON example to POST / PUT
```json
{
    "title":"My book",
    "authors": [
        "gabriel"
    ],
    "editions": 1,
    "year": 2022,
    "category": "tech"
}
```