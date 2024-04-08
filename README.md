# chrono-querist
A Backend Service written in Golang which takes advantage of Youtube API to store and search videos library

# Table of Contents
- [Chrono Querist](#chrono-querist)
- [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Linux/MacOS](#linuxmacos)
    - [Docker](#docker)
      - [Troubleshooting](#troubleshooting)
  - [API Documentation](#api-documentation)
  - [Explanation](#explanation)
      - [Database](#database)
    - [API](#api)
      - [Testing](#testing)
    - [Youtube Fetcher](#youtube-fetcher)
    - [DbApp](#dbapp)

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.
### Prerequisites
In case of running using docker, you need to have the following installed on your machine:
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Google Cloud Platform](https://cloud.google.com/) account with Youtube Data API enabled

In case of running the application directly, you need to have the following installed on your machine:
- [Go](https://golang.org/)
- [MySQL](https://www.mysql.com/)
- [Google Cloud Platform](https://cloud.google.com/) account with Youtube Data API enabled

### How to start the application - Docker
1. Clone the repository
```bash
git clone git@github.com:Ashish0077/chrono-querist.git 
cd chrono-querist
```
2. Fill the appropriate values in the environment variables in docker-compose
For `YOUTUBE_API_KEYS`, you can add multiple keys separated by a comma. The application will switch to the next key if the quota for the current key is exhausted.

3.. Run the following command to run the application
```bash
docker-compose up --build
```

The application should be running on `http://localhost:9000` by default.

Note: The setup is used for local or development purposes. For production, the environment variables should be set in a more secure way and additional measures should be taken.

## API Documentation
<!-- internal/adapters/primary/web/videos/openapi-spec.yaml -->
You can get the API documentation in the OpenAPI format from this [link](internal/adapters/primary/web/videos/openapi-spec.yaml).
You can directly import this file to Postman or Swagger to get the API documentation.

## Explanation
The application is following the Hexagonal Architecture

### Hexagonal Architecture
![alt text](https://miro.medium.com/v2/resize:fit:1170/1*aD3zDFzcF5Y2_27dvU213Q.png)

```
.
├── Dockerfile
├── LICENSE
├── README.md
├── cmd
│   └── web
│       └── main.go
├── docker-compose.yaml
├── go.mod
├── go.sum
├── internal
│   ├── adapters
│   │   ├── primary
│   │   │   └── web
│   │   │       ├── router.go
│   │   │       ├── server.go
│   │   │       └── videos
│   │   │           ├── handler.go
│   │   │           └── openapi-spec.yaml
│   │   └── secondary
│   │       └── repository
│   │           └── videos
│   │               ├── index.go
│   │               ├── upsert.go
│   │               └── video.go
│   ├── core
│   │   ├── jobs
│   │   │   └── videos
│   │   │       └── youtube_data_loader_job.go
│   │   ├── lib
│   │   │   └── external_apis
│   │   │       └── youtube
│   │   │           ├── client.go
│   │   │           └── search_by_query.go
│   │   ├── models
│   │   │   └── video.go
│   │   └── services
│   │       └── videos
│   │           ├── index.go
│   │           └── video.go
│   └── ports
│       ├── primary
│       │   └── API
│       │       └── videos.go
│       └── secondary
│           └── repository
│               └── videos.go
└── utils
    └── db
        ├── connection.go
        └── migrate.go
```

   
We use our own command line tool, [DbApp](#dbapp) to handle one-time database operations.

### Database
We are using MySQL as the database. The database is used to store the videos fetched from Youtube. The database is managed using the [GORM](https://gorm.io/) library. The database schema is defined in the `internal/core/models/video.go`. The database schema is defined using the GORM tags.

### API
The API is the main part of the application. It is responsible for handling the requests from the client and returning the appropriate response. The API is built using the [Go Fiber](https://docs.gofiber.io/api/fiber).