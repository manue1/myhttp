# response-hasher

This tool makes http requests and prints the address of the request along with the
MD5 hash of the response.

## Usage

In order to run this tool it first needs to be built.

```bash
go build -o myhttp

myhttp -parallel 2 adjust.com google.com facebook.com
```

The usage of the tool is limited to the following arguments:

```
myhttp [options...] <url> [<url>...]
  -parallel <int>   Number of requests in parallel (default 10)
```

Please note that currently only the HTTP protocol is supported.
