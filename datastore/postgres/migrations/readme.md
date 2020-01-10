Create new migration from project root

`migrate create -ext sql -dir datastore/postgres/migrations -seq create_users_table`

Note: ensure database has been created before running migration

Run migration from root folder
`migrate -database "postgres://postgres:postgres-admin-secret@localhost:5432/weather_monster?sslmode=disable" -path datastore/postgres/migrations up`

https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md
