<h1 align="center">
  <img src="https://i.ibb.co/LhGDRPq/Solution-5-File-Metadata.png" alt="Logo" width="600">
</h1>

## Features

- Filename, size, and type data from uploaded file

## Tech Stack

**Server:** Golang (Gin)

## Run Locally

**Your server is accessible in port 3000**
Here is the steps to run it with `golang`

```bash
# Move to directory
$ cd workspace

# Clone this repository
$ git clone https://github.com/madjiebimaa/fcc-file-metadata-ms.git

# Move to project
$ cd fcc-file-metadata-ms

# Set gin to release mode
$ export GIN_MODE=release

# Run the application
$ go run main.go
```

Here is the steps to run it with `docker-compose`

```bash
# Move to directory
$ cd workspace

# Clone this repository
$ git clone https://github.com/madjiebimaa/fcc-file-metadata-ms.git

# Move to project
$ cd fcc-file-metadata-ms

# Set gin to release mode
$ export GIN_MODE=release

# Download, setup, and run the image
$ docker-compose up -d

# Stops containers and removes containers, networks, volumes, and images created by up
$ docker-compose down
```

## Running Tests

To run tests and get the percentage of code coverage, run the following command

```bash
  go test ./... -cover
```

## Lessons Learned

- How to create API with Golang (Gin)
- How to create middleware
- How to create gin handler
- How to create unit test for handler
- How to process file

## API Reference

[Test API with REST Client Extension](https://github.com/madjiebimaa/fcc-file-metadata-ms/tree/main/docs/apis)
