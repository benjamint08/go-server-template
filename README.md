#
<img src="gopher-idea.png">

# Go Server Template

By [benjamint08](https://github.com/benjamint08)

## Description

This is a template for a Go API.
## Installation

1. Clone the repository
2. Run `go run main.go` to run the application.

## Adding API endpoints

To add an API endpoint, create a new file in the `handlers` directory. Check the `handlers/hello.go` file for an example.

You can then add the handler to the router in the `main.go` file.

## Production (Docker)

1. Run `docker build -t <your_image_name> .`
2. Run `docker run -d -p 8080:8080 <your_image_name>`
3. Visit `http://localhost:8080` in your browser
4. Success