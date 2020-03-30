# Store
Store is an online e-commerce web application designed for small online stores.

## Project Structure
```
/api -- Handlers
/cmd -- Main file
/db -- Database
/models
/client -- Vue frontend
```

## Endpoints
```
GET /

GET /api/products
GET /api/products/:name
POST /api/cart
POST /api/checkout

POST /api/login

GET /api/admin/products -- Auth required
GET /api/admin/product/:name -- Auth required
PUT /api/admin/product -- Auth required
POST /api/admin/product -- Auth required

POST /api/admin/image -- Auth required

GET /api/admin/stocks/:name -- Auth required
PUT /api/admin/stocks/:name -- Auth required
POST /api/admin/stocks/:name -- Auth required
```

## TODO
* General code cleanup
* Endpoint address cleanup
* Add checkout confirm and cancel endpoint
* Standarize JSON responses
