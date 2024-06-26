include .env

ifeq ($(PRODUCTION), true)
	FILE=./conf/docker/prod/docker-compose.yml
	PROJECT=em_test_task_prod
else
	FILE=./conf/docker/dev/docker-compose.yml
	PROJECT=em_test_task_dev
endif

build:
	docker compose -f $(FILE) --env-file .env --project-directory . -p $(PROJECT) build

start:
	docker compose -f $(FILE) --env-file .env --project-directory . -p $(PROJECT) up

logs:
	docker compose -p $(PROJECT) logs app -f

stop:
	docker compose -p $(PROJECT) stop

clean:
	docker compose -p $(PROJECT) down --rmi "local" -v