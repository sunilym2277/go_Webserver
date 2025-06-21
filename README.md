# Simple Go Web Server
This project is a basic web server implemented in Go. It's a beginner-friendly example that demonstrates how to handle different routes, serve static files, and process form submissions using Go's standard net/http package.

## Description
The primary goal of this project is to showcase the fundamentals of building a web application in Go. It sets up a server that listens for HTTP requests and responds by serving HTML pages or processing incoming data.

#### This server has three main routes:

/: Serves a static welcome page (index.html).

/hello: A simple handler that prints a "hello" message to the browser.

/form: Handles both GET and POST requests. It serves an HTML form (form.html) and processes the data submitted through it.

## Features
Static File Serving: Demonstrates how to serve static HTML files.

Routing: Implements basic routing to handle different URL paths.

Form Handling: Shows how to process data from an HTML form submission (POST request).

Built with Standard Library: Uses only Go's built-in packages, requiring no external frameworks.

### File Structure
.
├── go.mod        # Manages the project's dependencies
├── main.go       # Contains the core web server logic and route handlers
├── index.html    # The static home page for the root route (/)
└── form.html     # The HTML form served at the /form route

## Getting Started
Follow these instructions to get a copy of the project up and running on your local machine.

### Prerequisites
You need to have Go installed on your system. You can download it from the official website: golang.org/dl.

### Installation
Clone the repository:

git clone https://github.com/sunilym2277/go_Webserver.git

Navigate to the project directory:

cd go_Webserver

Usage
Run the web server:
Execute the main.go file from your terminal.

go run main.go

Access the application:
Once the server is running (it will likely listen on port 8080), open your web browser and navigate to the following URLs:

Home Page: http://localhost:8080/

Hello Page: http://localhost:8080/hello

Form Page: http://localhost:8080/form

You can fill out the form on the /form page and submit it to see the server process the POST request.

## Contributing
Contributions make the open-source community an amazing place to learn, create, and share. Any contributions you make are greatly appreciated.

If you have a suggestion that would make this project better, please fork the repository and create a pull request. You can also simply open an issue with the tag "enhancement".

### Fork the Project

Create your Feature Branch (git checkout -b feature/AmazingFeature)

Commit your Changes (git commit -m 'Add some AmazingFeature')

Push to the Branch (git push origin feature/AmazingFeature)

Open a Pull Request

License
This project is not currently licensed. It's recommended to add an open-source license like the MIT License to let others know how they can use your code.
