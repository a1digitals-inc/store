# Store
Store is an online e-commerce web application designed for small online stores.

## Project Structure
```
/api -- Handlers
/cmd -- Main file
/db -- Database
/models
```

## Endpoints
```
GET /api/products
GET /api/products/:name
PUT /api/product -- Auth required
POST /api/login
GET /api/admin/products -- Auth required
GET /api/admin/product/:name -- Auth required
POST /api/cart
```
### TODO
```
GET /api/admin/stock/:name -- Auth required
PUT /api/admin/stock/:name -- Auth required
```

## TODO
* General code cleanup
* Add product stock manager
* Add shopping cart
* Add purchasing
