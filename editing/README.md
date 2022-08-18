# Editing OpenAPI Specification Documents

## swagger-editor

The swagger-editor docker container is a minimum frills editing environment that
can be used to modify an OpenAPI specification and generate go client and server
code. All file access is managed by the browser, so there is no need to mount a
volume to the container

```
docker run --rm -d -p 8080:8080 swaggerapi/swagger-editor

docker run -d -p 80:8080 -v $(pwd):/tmp -e SWAGGER_FILE=/tmp/swagger.json swaggerapi/swagger-editor

```


## apicurito

```
docker run --rm -d -p 8080:8080 apicurio/apicurito-ui
```
