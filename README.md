# Store
Store is an online e-commerce web application designed for small online stores.

## Project Structure
```
/api -- Handlers
/cmd -- Main file
/db -- Database
/views -- HTML templates
/src -- Webpack entrypoints
/src/components -- React Components
/static -- .js bundles
```

## Endpoints
```
GET /api/products
GET /api/products/:name
PUT /api/product -- Auth required
POST /api/login
GET /api/admin/products -- Auth required
GET /api/admin/product/:name -- Auth required
```
### TODO
```
GET /api/admin/stock/:name -- Auth required
PUT /api/admin/stock/:name -- Auth required
```

## TODO
* Separate models
* General code cleanup
* Separate backend and frontend
* Add product stock manager
* Add shopping cart
* Add purchasing
