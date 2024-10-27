# Favicon Fetcher

A simple service for retrieving favicons (website icons) given a domain name. It uses the `go.deanishe.net/favicon` library to discover favicon URLs and `go-resty/resty` to fetch the icon data.  The service is built with the Fiber framework for performance and uses a caching middleware to improve response times for requests for the same domain.

## Features

* **Favicon Discovery:**  Automatically identifies and retrieves the most appropriate favicon for a given domain.
* **Error Handling:**  Gracefully handles cases where a favicon cannot be found, returning a default icon.
* **Caching:** Caches successful responses to reduce load times and bandwidth consumption.
* **Content Type Handling:** Sets the correct `Content-Type` header for the favicon, ensuring proper display in browsers.

## Installation

1. **Prerequisites:** Make sure you have Go installed on your system.

2. **Clone the repository:**

   ```bash
   git clone https://github.com/radeeyate/favicon.git 
   cd favicon
   ```

3. **Download dependencies:**

   ```bash
   go get
   ```

4. **Build the server:**

   ```bash
   go build .
   ```

5. **(Optional) Provide a default favicon:** Create a file named `default.ico` in the project root. This icon will be served if a favicon cannot be found for a given domain.  You can use the provided `default.ico` in this repository (stolen from DuckDuckGo's favicon service).

## Usage

1. **Start the server:**

   ```bash
   ./favicon
   ```

2. **Access the favicon via HTTP:**

   ```http
   GET http://localhost:7000/get/example.com
   ```

   Replace `example.com` with the domain you want.

## Configuration

The server listens on port 7000 by default. You can change this by modifying the `app.Listen(":7000")` line in `main.go`.

## Future Improvements

* [] Icon Resizing
* [] Configuration Options


## Contributing

Contributions are welcome!  Please feel free to submit pull requests or open issues for bug reports and feature requests.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
