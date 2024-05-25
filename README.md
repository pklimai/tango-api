# Tango Api

Microservice for retrieving hardware values ​​that are taken from instruments during experiments at NICA.

https://stackoverflow.com/questions/24319662/from-inside-of-a-docker-container-how-do-i-connect-to-the-localhost-of-the-mach

## Deploying

### Build
```
docker build -t tango-api .
```
It creates image with name ```tango-api:latest```.

### Run
```
docker run -d --rm --network="host" --name tango-api tango-api:latest
```
There is dirty hack in flag --network="host". It is here because we run Go app in container with HOST (in config, default 127.0.0.1). Need to consider this question ([stackoverflow](https://stackoverflow.com/questions/24319662/from-inside-of-a-docker-container-how-do-i-connect-to-the-localhost-of-the-mach)) due to security risks.

### Monitor
```
docker logs tango-api -f
```
Collect logs from the app container. In future will be created a common solution for logs (Graylog).

### Clear
```
docker container stop 'container_name'
docker imager rmi 'image_name'
```

## Config
Image contains config/.env file that is copied during the building stage. 

You can find example of config in the config/local_example.env file. 
```
# APP
LOG_LEVEL = INFO # DEBUG, INFO, WARN, ERROR
PORT_HTTP = 7000
PORT_SWAGGER = 7001
PORT_GRPC = 7002
HOST = "127.0.0.1"

PG_DSN = "host=127.0.0.1 port=1331 dbname=tango-api user=tango-api-user password=1234 sslmode=disable"
```

For a PROD environment you need to specify PG_DSN with the right credentials. If you want to specify  ports:
* HTTP - need to change in config/.env, Dockerfile and api/tango_api_service/service.proto host with right port. Then execute ```make generate```.
* Swagger - change in config/.env and Dockerfile.
* gRPC - change in config/.env and Dockerfile.
