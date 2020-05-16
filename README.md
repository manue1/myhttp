# myhttp - CLI tool

This tool makes http requests and prints the address of the request along with the
MD5 hash of the response.

## Usage

In order to run this tool it first needs to be built. Please ensure to have [Go 1.14](https://golang.org/dl/) installed.

```bash
make build
```

The usage of the tool is limited to the following arguments:

```
./myhttp [options...] <url> [<url>...]

Options:
  -parallel <int>   Number of requests in parallel (default 10)

Examples:
  ./myhttp http://www.adjust.com ​ http://google.com

  ./myhttp -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com

```

Please note that currently only the HTTP protocol is supported.

## Tests

For running the unit-tests a make target is set up.

```bash
make unit-test
```
