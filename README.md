# Store
Store is an online e-commerce web application designed for small online stores.

## Project Structure
```
/client -- Vue client

/server -- Go API server
/server/api -- Handlers
/server/cmd -- Main file
/server/db -- Database
/server/models
/server/client -- Vue frontend
```

## API server endpoints
```
GET /

GET /api/products
GET /api/products/:name
POST /api/cart
POST /api/checkout

POST /api/login
POST /api/refresh -- Auth required

GET /api/admin/products -- Auth required
GET /api/admin/product/:name -- Auth required
PUT /api/admin/product -- Auth required
POST /api/admin/product -- Auth required

POST /api/admin/image -- Auth required

GET /api/admin/stocks/:name -- Auth required
PUT /api/admin/stocks/:name -- Auth required
POST /api/admin/stocks/:name -- Auth required
```

## Running with Docker
```
$ git clone https://github.com/sergiosegrera/store
$ docker-compose up
```

## TODO
* General code cleanup
* Endpoint address cleanup
* Add checkout confirm and cancel endpoint
* Standarize JSON responses
