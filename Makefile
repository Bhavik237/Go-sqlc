postgres:
	docker run  --name postgres14 -p 5430:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=pass -d postgres:14-alpine 
createdb:
	docker exec -it postgres14 createdb --username=postgres --owner=postgres course

migrateup:
	migrate -path db/migration -database "postgresql://postgres:pass@localhost:5430/course?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:pass@localhost:5430/course?sslmode=disable" -verbose down

