## Zoo builder API

To run, make sure you have Golang version 1.24.4:

```go run .```

The Endpoint is then at `http://localhost:8080/api/v1/zoos/{enclosures}`, as defined in the project specification.

### TODO improvements
To discuss:

1. Improved error handling - using custom error types to catch, handle and clearly log different types of errors throughout the system, e.g. loading the animals JSON data
2. Dependency injection - Loader for the Animals JSON data and the algorithm functions can be put in classes and injected. This would help with mocking and testing.
3. Logging - more verbose logging for monitoring and debugging, e.g. Loading the animals JSON data and within the algorithm process
4. General code tidy up - there are some functions that are not single-purpose, too long and messy. A small refactor would help clear things up a bit. Making function smaller and single-responsibility would help build unit tests to run and target specific areas of the code.
5. Follow more Go-specific conventions

6. *Zoo building algorithm* - unclear how to build an optimized and efficient algorithm for building the enclosures. I have used a simple greedy-ish algorithm, lots of room for improvement.
