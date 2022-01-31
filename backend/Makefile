.PHONY: swagger/version
swagger/version:
	docker-compose run --rm go-swagger version

.PHONY: go/run_dev
go/run_dev:
	reflex -r '\.go$$' -s -- sh -c 'go run cmd/api/main.go'