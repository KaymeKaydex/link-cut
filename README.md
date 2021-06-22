# Link cut
Link shortening service
## About
:link: Through the grpc protocol, the service accepts a link and shortens it for a fictitious domain l.ru 
## Commands
#### Build :cd:
Bulding bin file into ./bin/
```shell
make build
```
#### Run :checkered_flag:
How to fast run application
```shell
cp .dist.env .env
docker-compose up -d
make migrate
make build
```
## Future 
1. New config type & config parse
2. Replace all constants with the config
3. Endpoint module 
4. Metrics
