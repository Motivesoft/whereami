# whereami
Call a geolocation service with an IP address to obtain location information

## Usage

`whereami [ip address]`

## Building

```shell
go build
GOOS=windows GOARCH=amd64 go build -o myapp.exe
```

## Requirements

The web service used by this software requires a user ID and key. These should be stored in a file alongside the executable.

Create a file called `.env` containing the following:

```yaml
User-ID: <id>
API-Key: <key>
```

The values for `<id>` and `<key>` can be obtained by registering with [Neutrino API](https://www.neutrinoapi.com).