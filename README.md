# go-rest

This is a non-ORM solution to a RESTful HTTP server in Go. 

## dependencies

- [crud](https://github.com/jakecoffman/crud)
  - Provides an easy way to get OpenAPI docs and validation middleware
- gin
  - This is a solid and fast router that has been around for a while
- pgx
  - Database driver for postgres, can be swapped out easily
- sqlboiler
  - Generate models from db schema
