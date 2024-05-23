## Golang Backend API with JWT, MySQL & Tests

[YTvid](https://www.youtube.com/watch?v=7VLmLOiQ3ck)
[GHsource](https://github.com/sikozonpc/ecom)

### create a HTTP server with Handlers

- in Go it is convention to store entry points in a CMD directory
  - two entry points: API itself, and migration

- main file within the CMD directory is the entry point function, that will create a new API instance (a server instance)

### Services Dir

- Handlers will be in the Services directory

### Config Dir

- this holds all of the server configurations