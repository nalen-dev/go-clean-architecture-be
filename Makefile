migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@localhost/be_test" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:password@localhost/be_test" -verbose down