# Simple Task API

## Overview

Simple RESTful CRUD API designed for front-end developers who are learning how
to build todo-list applications and work with an API. Offers a
straightforward task management system where users can create, retrieve, update,
and delete tasks. Each task is timestamped with `createdAt` and `updatedAt`
fields to track when tasks are added or modified.

This basic version does not utilize any actual database, instead just storing
the tasks in memory. There is also no authentication. I will be creating alternate
versions that are more advanced with authentication and a database.

> Please don't try to use this for a production application! This is meant
> purely for learning purposes and to be used locally.

## Getting Started

### Prerequisites

- Go version 1.22
- Dependencies (automatically installed with `go mod tidy`)
  - Chi
  - Zerolog
- Docker (optional)
- Basic knowledge of RESTful principles and HTTP requests

### Installation

1. Clone this repository or download the source code.

```bash
git clone github.com/nronzel/simpleTaskApi
```

2. Navigate to the project directory.

```bash
cd simpleTaskApi
```

3. Install the dependencies:

```bash
go mod tidy
```

4. Run the server from your terminal:

```bash
go run .
```

### Docker

If you have [Docker](https://docs.docker.com/get-docker/) installed, you can
run the API in a container.

> The latest versions of Docker Desktop come with `docker compose` pre-installed.

The easiest way to run the API in a container is to use `docker compose`. This
will build the image and run the container in the background.

```bash
docker compose up --build -d
```

Or you can build the image and run the container manually.

```bash
docker build -t taskapi .
docker run -p 9090:9090 taskapi
```

You can navigate to `http://localhost:9090/` in your browser to make sure it is working.

## Accessing the API Documentation

This API is documented using Swagger. Once the server is running, you can access
the Swagger UI to test the API endpoints directly through your browser:

```bash
http://localhost:9090/swagger/
```

## Endpoints

Below is a brief overview of the API endpoints. For detailed request and
response formats, refer to the Swagger documentation.

- List All Tasks

  - GET /tasks
  - Retrieves all tasks.

- Get a Task by ID

  - GET /tasks/{taskID}
  - Fetches a single task by its ID.

- Create a New Task

  - POST /tasks
  - Adds a new task. The request body should include the task field.

- Update a Task

  - PUT /tasks/{taskID}
  - Modifies an existing task. The request body should contain the updated task field.

- Delete a Task
  - DELETE /tasks/{taskID}
  - Removes a task by its ID.

## License

This project is open source and available under the [MIT License](./License.md).

## Features/Issues

If you have any feature requests or find any issues, please open an issue.

## Contributing

Contributions are welcome! Please open an issue before submitting a pull request.
