# Demonstration of slow HTTP2 push in Golang

* Run `go run main.go`
* Open in browser: https://localhost:8080/small
* Observe time it took to handle the request (during which 100 pushes is made, with 10 KB large responses), e.g.: `main /small: 7.865`, which means 7 ms.
* Open in browser: https://localhost:8080/large
* Observe time it took to handle the request (during which 100 pushes is made, with 1 MB large responses), e.g.: `main /large: 42.865`, which means 42 ms.
