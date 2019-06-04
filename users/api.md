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
{
    "status": "deleted"
}
```