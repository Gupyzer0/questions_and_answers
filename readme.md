# Questions and answers

#### Execution

> go run cmd/main.go [flags]

*It's better to migrate and seed first:*

> go run cmd/main.go -migrate -seed

#### flags

**migrate:** Creates database tables.

**seed:** Inserts seed data for testing.

#### TODO:
logging
implement gRPC
