# api-template-juanguaje

api template juanguaje description

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make
$ ./bin/api-template-juanguaje
```

### Configs

#### Configuration file config, loaded from `./config/config.go` file.

```yaml
server:
  port: 8080

api:
  pharmacy:
    url: https://farmanet.minsal.cl/index.php/ws/getLocales
```

### Tools

- [Go](https://go.dev/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Viper](https://github.com/spf13/viper)


### Endpoints

#### host: `localhost:8080`


### Request

### All Pharmacies

curl --location 'http://localhost:8080'

### Filter Commune Response Json

curl --location --request GET 'http://localhost:8080' \
--header 'Content-Type: application/json' \
--data '{
    "filtroValor":"LA CALERA",
    "tipoRespuesta":"JSON"
}'

### Filter Commune Response XML

curl --location --request GET 'http://localhost:8080' \
--header 'Content-Type: application/json' \
--data '{
    "filtroValor":"LA CALERA",
    "tipoRespuesta":"XML"
}'