tests:
	go test -race -covermode=atomic ./... 

start:
ifdef env
ifeq ($(env), dev)
	$(shell docker-compose down)
	$(shell chmod +x .docker/entrypoint.sh)
	$(shell cp config.dev.toml config.toml)\
	$(shell cp docker-compose.dev.yml docker-compose.yml)
	$(shell docker-compose up -d)
else 
	$(info invalid parameter, we accept only dev as env and if it is not set, then we will use env = production)
endif
else
	cp config.production.toml config.toml
	cp docker-compose.production.yml docker-compose.yml
endif


	