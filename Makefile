  # Load environment variables
  include .env

up:
	@echo "Starting mongodb containers..."
	docker-compose  up --build -d --remove-orphans

down:
		@echo "Stopping containers..."
		docker-compose down

build:
	go build -o ${BINARY} ./cmd/

start:
	MONGODB_URI=${MONGODB_URI} ./${BINARY}

restart: build start

#clean:
#	docker-compose down -v
#	rm -f ${BINARY}
#

