DBUSER:=root
DBPASSWORD:=password
DBPORT:=5432
DBNAME:=hogehoge
DOCKER_DNS:=db
SERVICE:=
COMMAND:=

.PHONY: dc/command
dc/command:
	docker-compose $(COMMAND)

.PHONY: dc/build
dc/build:
	docker-compose build

.PHONY: dc/up
dc/up:
	docker-compose up db backend adminer

.PHONY: dc/up-d
dc/up-d:
	docker-compose up -d db backend adminer

.PHONY: dc/down
dc/down:
	docker-compose down

.PHONY: dc/logs
dc/logs:
	docker-compose logs -f

.PHONY: dc/down-remove
__dc/down-remove:
	docker-compose down --volumes

MIGRATION_SERVICE:=migration
FLYWAY_CONF?=-url=jdbc:postgresql://$(DOCKER_DNS):$(DBPORT)/$(DBNAME) -user=$(DBUSER) -password=$(DBPASSWORD)
.PHONY: flyway/info
flyway/info:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) info

.PHONY: flyway/validate
flyway/validate:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) validate

.PHONY: flyway/migrate
flyway/migrate:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) migrate

.PHONY: flyway/repair
flyway/repair:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) repair

.PHONY: flyway/baseline
flyway/baseline:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) baseline

DB_SERVICE:=db
.PHONY: db/init
db/init:
	docker-compose exec $(DB_SERVICE) psql --username=$(DBUSER) --command="create database $(DBNAME)"

.PHONY: __db/drop
__db/drop:
	docker-compose exec $(DB_SERVICE) psql --username=$(DBUSER) --command="drop database $(DBNAME)"
