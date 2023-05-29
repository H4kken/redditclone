.PHONY: postgres adminer migrate

postgres:
	docker run --rm -ti -e POSTGRES_PASSWORD=secret postgres

adminer:
	docker run --rm -ti -p 8080:8080 adminer

migrate:
	migrate -source file://migration \
			-database postgres://postgres:+Serdnol2011@localhost/postgres?sslmode=disable up

migrate-down:
	migrate -source file://migration \
			-database postgres://postgres:+Serdnol2011@localhost/postgres?sslmode=disable down