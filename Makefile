include .env

up:
		@echo "starting containers"
		docker-compose up --build -d --remove-orphans

down:
		@echo "stopong conatiners"
		docker-compose down

build: 	
		go build -o ${BINARY}

run: 
		./${BINARY}	

restart: build run
