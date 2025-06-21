# API Reference

## Content
1. [**Auth**](#auth)
    - [Sign-Up](#sign-up)
    - [Sign-In](#sign-in)
2. [**API Keys**](#api-keys)
    - [Create API Key](#create-api-key)
    - [Get API Keys](#get-api-key)
    - [Delete API Key](#delete-api-key)
3. [**Collection**](#collection)
    - [Get Collection Name](#get-collection-name)
4. [**Storage**](#storage)
    - [Load Documents](#load-documents)
    - [Get Document](#get-document)
    - [Get All Documents](#get-all-documents)
    - [Delete Document](#delete-document)
5. [**Generation**](#generation)
    - [Generate](#generate)

## Auth

### Sign-Up
#### POST /api/v1/user/auth/sign-up
```
curl -X POST 'localhost:8080/api/v1/user/auth/sign-up' \
--header 'Content-Type: application/json' \
--data '{
    "login": "test",
    "password": "123"
}'
```

#### Request data
- *login* - user login
- *password* - user password

#### Response 
200 OK
```
{
    "status": 200,
    "data": {
        "collection": collection-uuid
    }
}
```

### Sign-In
#### POST /api/v1/user/auth/sign-in
```
curl -X POST 'localhost:8080/api/v1/user/auth/sign-in' \
--header 'Content-Type: application/json' \
--data '{
    "login": "test",
    "password": "123"
}'
```

#### Request data
- *login* - user login
- *password* - user password

#### Response 
200 OK
```
{
    "status": 200,
    "data": {
        "token": Bearer *****
    }
}
```

## API Keys

### Create API Key
#### POST /api/v1/user/api-keys
```
curl -X POST 'localhost:8080/api/v1/user/api-keys' \
--header 'Authorization: Bearer *****'
```

#### Headers
- *Authorization* - Bearer token auth

#### Response
200 OK
```
{
    "status": 200,
    "data": {
        "key": "some-hash"
    }
}
```

### Get API Keys
#### GET /api/v1/user/api-keys
```
curl -X GET 'localhost:8080/api/v1/user/api-keys' \
--header 'Authorization: Bearer *****'
```

#### Headers
- *Authorization* - Bearer token auth

#### Response
200 OK
```
{
    "status": 200,
    "data": {
        [
            {   "id": "some-uuid",
                "key": "some-hash"
            },
        ]
    }
}
```

### Delete API Key
#### DELETE /api/v1/user/api-keys
```
curl -X DELETE 'localhost:8080/api/v1/user/api-keys?id=some-uuid' \
--header 'Authorization: Bearer *****'
```

#### Headers
- *Authorization* - Bearer token auth

#### Query parameters
- *id* - key id

#### Response
200 OK
```
{
    "status": 200,
    "data": null
}
```

## Collections

### Get Collection Name
#### GET /api/v1/user/collection
```
curl -X GET 'localhost:8080/api/v1/user/collection' \
--header 'Authorization: Bearer *****' \
--data ''
```

#### Headers
- *Authorization* - Bearer token auth

#### Response
200 OK
```
{
    "status": 200,
    "data": {
        "collection": "some-uuid"
    }
}
```

## Storage

### Load Documents
#### POST /api/v1/storage/:collection
```
curl -X POST 'localhost:8080/api/v1/storage/:collection?api-key=some-hash' \
--header 'Authorization: Bearer *****' \
--data '{"documents": [
    {
        "content": "some-content",
        "metadata": {
            "url": "some-url"
        }
    },
    {
        "content": "some-content": {
            "url": "some-url"
        }
    },
]}'
```

#### Path parameters
- *collection* - collection id

#### Query parameters
- *api-key* - API Key

#### Request data
- *documents* - documents array
- *content* - some text
- *metadata* - metadata (only url for now)

#### Response
201 Accepted
```
{
    "status": 201,
    "data": null
}
```

### Get Document
#### GET /api/v1/storage/:collection/:id
```
curl -X GET 'localhost:8080/api/v1/storage/:collection/:id?api-key=some-hash'
```

#### Path parameters
- *collection* - collection id
- *id* - document id

#### Query parameters
- *api-key* - API Key

#### Response
200 OK
```
{
    "status": 200,
    "data": {
        "id": "some-uuid",
        "content": "some-text",
        "metadata": {
            "url": "some-url"
        }
    }
}
```

### Get All Documents
#### GET /api/v1/storage/:collection
```
curl -X GET 'localhost:8080/api/v1/storage/:collection?api-key=some-hash'
```

#### Path parameters
- *collection* - collection id

#### Query parameters
- *api-key* - API Key

#### Response
200 OK
```
{
    "status": 200,
    "data": {
        "documents": [
            {
                "id": "some-uuid",
                "content": "some-text",
                "metadata": {
                    "url": "some-url"
                }
            }
        ]
        "collection": "some-uuid"
    }
}
```

### Delete Document
#### DELETE /api/v1/storage/:collection/:id
```
curl -X DELETE 'localhost:8080/api/v1/storage/:collection/:id?api-key=some-hash'
```

#### Path parameters
- *collection* - collection id
- *id* - document id

#### Query parameters
- *api-key* - API Key

#### Response
200 OK
```
{
    "status": 200,
    "data": null
}
```

## Generation

### Generate
```
curl -X GET 'localhost:8080/api/v1/generation/:collection?query=%22some-query%22' 
```

#### Path parameters
- *collection* - collection id

#### Query parameters
- *query* - your question on context 

#### Response
*Server-Sent Event Stream*