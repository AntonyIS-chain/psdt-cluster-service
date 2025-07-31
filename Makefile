build:
	go build -o bin/backend/user-service
	
serve: build
	ENV=development ./bin/backend/user-service

serve-dev-test: build
	ENV=development_test go test -v ./...

docker-push:
	docker build -t github.com/AntonyIS-chain/psdt-cluster-service:latest --build-arg ENV=docker .
	docker push github.com/AntonyIS-chain/psdt-cluster-service:latest

docker-run:
	docker run -p 8081:8081 ENV=docker github.com/AntonyIS-chain/psdt-cluster-service:latest

docker-test:
	ENV=docker_test go test -v ./...