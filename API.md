# API documentation

## Authorization

Your API token should be provided in `Authorization` header. Please see environment variables in [README.md](./README.md)

## Create a user(or get a new token)

Request:
```
POST http://localhost:8080/mtgs
{
	"name": "username"
}
```

Response:
```
{
    "name": "testusdsgsdger",
    "secret": "dd851e19bb60a6d137a97b87b6193bb70e"
}
```

## Delete a user

Request:
```
DELETE http://localhost:8080/mtgs/username
```

Response:
```
{
    "status": "deleted"
}
```

## Get all users
Request:
```
GET http://localhost:8080/mtgs
```

Response:
```
[
    {
        "name": "username1",
        "secret": "dd851e19bb60a6d137a97b87b6193bb70e"
    },
    {
        "name": "username2",
        "secret": "dd7fcae07fe22f4a8fb332db5a0419f019"
    }
]
```
