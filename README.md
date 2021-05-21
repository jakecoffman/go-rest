# go-rest

This is a non-ORM solution to a RESTful HTTP server in Go. 

## dependencies

- gin
  - This is a solid and fast router that has been around for a while
- crud
  - Provides an easy way to get OpenAPI docs and validation middleware
- pq
  - Database driver for postgres, can be swapped out easily
- sqlx
  - Make mapping struct fields to sql queries much easier
  
TODO: 
- Need a sql builder to support passing URL query variables e.g. limit=10, name=Bob
