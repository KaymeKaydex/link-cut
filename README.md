# Link cut
Link shortening service
##Commands
####Build
Bulding bin file into ./bin/
```shell
make build
```
####Run
How to fast run application
```shell
cp .dist.env .env
docker-compose up -d
make migrate
make build
```
##Future
1. New config type & config parse