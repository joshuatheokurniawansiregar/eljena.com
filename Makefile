export POSTGRESQL_URL = "postgres://postgres:admin@localhost:5432/sdf?sslmode=disable"

migrate-create:
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-up:
	@ migrate -database ${POSTGRESQL_URL} -path scripts/migrations up

migrate-down:
	@ migrate -database ${POSTGRESQL_URL} -path scripts/migrations down
air:
	air