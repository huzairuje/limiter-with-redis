local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build

test:
	go test -cover ./...

clean:
	docker system prune -f

run:
	go run main.go