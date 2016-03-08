# serve
A simple Web Server to serve your static files on a URL

## Usage

```go
serve // serves the files in current directory on a default URL(localhost:4000)

or

serve ./foobar // serves the files in `foobar` directory
```

If you do not want to open browser by default, you can use `-s` flag.
You can serve on a specified local IP as `serve -a=127.0.0.1:8000`

## Build from source

* Clone this repository
* Make sure you have Go installed
* Do `go build`
* Copy the `serve` output file to `$PATH`(like `/usr/local/bin` for unix systems)
