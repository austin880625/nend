# nend
Just got tired of setting up a whole nginx installation just to resolve CORS error during local development. DO NOT USE IT IN PRODUCTION.

## Usage

Building

```
go build .
```

Example usage

```
./nend serve -p 8888 -f 3000 -b 8080 -a /api/v1
```