# StorageService

## Deployment

We recommand usage of Docker.

### Docker

```bash
$ docker run -v db:/data/db -e GIN_MODE=release -e STORAGESERVICE_MONGO_IP=mongo -e STORAGESERVICE_MONGO_PORT=27017 -e STORAGESERVICE_MONGO_USER=StorageService -e STORAGESERVICE_MONGO_PASSWORD=StorageService -e STORAGESERVICE_MONGO_DB=StorageService -e STORAGESERVICE_PORT=80 uprefer/storageservice:v0.1.0
```

### Docker Compose

```docker-compose
version: "3.3"

services:
  mongo:
    image: mongo:3-stretch
    volumes:
    - db:/data/db
    environment:
      MONGO_INITDB_ROOT_USERNAME: StorageService
      MONGO_INITDB_ROOT_PASSWORD: StorageService
      MONGO_INITDB_DATABASE: StorageService

  storage_service:
    image: uprefer/storageservice:v0.1.0
    ports:
    - 8080:80
    depends_on:
    - mongo
    environment:
      GIN_MODE: release
      STORAGESERVICE_MONGO_IP: mongo
      STORAGESERVICE_MONGO_PORT: 27017
      STORAGESERVICE_MONGO_USER: StorageService
      STORAGESERVICE_MONGO_PASSWORD: StorageService
      STORAGESERVICE_MONGO_DB: StorageService
      STORAGESERVICE_PORT: 80
    links:
    - mongo:mongo

volumes:
db:
```

### From source code
`StorageService` requires `Go 1.11` or later.

```bash
$ go get github.com/UPrefer/StorageService
$ ${GO_PATH}/bin/StorageService
```

## Usage

[Insomnia](https://insomnia.rest/) configuration file [is available in repository](insomnia_conf.json).

### Create file

Creates a new resource on StorageService

#### Response

Header `Location` contains created resource URI.

### Upload file

Uploads file to a previously created resource.

You must replace `artifact_id` by id of the previously created artifact.

### Download file

You must replace `artifact_id` by id of the previously created artifact.

Downloads the previously uploaded file.

### Get Metainfos

Retrieves all meta informations about an artifact.
