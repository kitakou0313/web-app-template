.PHONY: swagger/version
swagger/version:
	docker-compose run --rm go-swagger version

.PHONY: run
run:
	reflex -r '\.go$$' -s -- sh -c 'go run cmd/api/main.go'