run: run-nginx run-client run-ticket run-payment

create-network:
	docker network create simtix

run-client:
	cd simtix-client && docker-compose up -d --build

run-ticket:
	cd simtix-ticketing && docker-compose up -d --build

run-payment:
	cd simtix-payment && docker-compose up -d --build

run-nginx:
	docker compose up -d --build

stop-all: stop-client stop-ticket stop-payment stop-nginx

remove-network:
	docker network remove simtix

stop-client:
	cd simtix-client && docker-compose down

stop-ticket:
	cd simtix-ticketing && docker-compose down

stop-payment:
	cd simtix-payment && docker-compose down

stop-nginx:
	docker-compose down

seed:
	py seeder.py