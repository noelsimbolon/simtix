run: run-client run-ticket run-payment

create-network:
	docker network create simtix

run-client:
	cd simtix-client && docker-compose up -d --build

run-ticket:
	cd simtix-ticketing

run-payment:
	cd simtix-payment && docker-compose up -d --build

stop-all: remove-network stop-client stop-ticket stop-payment

remove-network:
	docker network remove simtix

stop-client:
	cd simtix-client && docker-compose down

stop-ticket:
	cd simtix-ticketing

stop-payment:
	cd simtix-payment && docker-compose down
