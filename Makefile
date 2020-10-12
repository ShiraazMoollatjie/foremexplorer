run: start-testdb
	go run cmd/foremexplorer/main.go --dbSchema=foremexplorer --dbUser=foremexplorer --dbPassword=password

run-app:
	go run cmd/foremexplorer/main.go --dbSchema=foremexplorer --dbUser=foremexplorer --dbPassword=password

seed: 
	go run cmd/seeder/main.go --dbSchema=foremexplorer --dbUser=foremexplorer --dbPassword=password

start-testdb:
	docker-compose -f deploy/docker-compose.yml up -d

stop-testdb:
	docker-compose -f deploy/docker-compose.yml down 

ci: vet test

vet:
	go vet ./...

test: 
	go test ./... -coverprofile=c.out
