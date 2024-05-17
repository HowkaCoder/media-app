# API Documentation

## Base URL
```
https://mediaapp-uksut1q4.b4a.run/api/
```

## Categories

### Get All Categories
```
GET /categories
```
#### Response
```json
[
{
"id": 1,
"parent_category_id":0,
"level":1,
"name": "Category 1"
},
{
"id": 2,
"name": "Category 2"
}
]
```

### Get Category by ID
```
GET /categories/:id
```
#### Response
```json
{
"id": 1,
"name": "Category 1"
}
```

### Create Category
```
POST /categories
```
#### Request Body
```json
{
"name": "New Category"
}
```
#### Response
```json
{
"id": 3,
"name": "New Category"
}
```

### Update Category
```
PATCH /categories/:id
```
#### Request Body
```json
{
"name": "Updated Category"
}
```
#### Response
```json
{
"id": 1,
"name": "Updated Category"
}
```

### Delete Category
```
DELETE /categories/:id
```
#### Response
```json
{
"message": "Category successfully deleted"
}
```



### Languages

#### Get All Languages
```
GET /languages
```
##### Response
```json
[
{
"id": 1,
"name": "Language 1"
},
{
"id": 2,
"name": "Language 2"
}
]
```

#### Get Language by ID
```
GET /languages/:id
```
##### Response
```json
{
"id": 1,
"name": "Language 1"
}
```

#### Create Language
```
POST /languages
```
##### Request Body
```json
{
"name": "New Language"
}
```
##### Response
```json
{
"id": 3,
"name": "New Language"
}
```

#### Update Language
```
PATCH /languages/:id
```
##### Request Body
```json
{
"name": "Updated Language"
}
```
##### Response
```json
{
"id": 1,
"name": "Updated Language"
}
```

#### Delete Language
```
DELETE /languages/:id
```
##### Response
```json
{
"message": "Language successfully deleted"
}
```





### Products

#### Product Structure
```json
{
"id": 1,
"name": "Product Name",
"price": 100,
"discount": 0,
"quantity": 10,
"category_id": 1,
"category": {
"id": 1,
"name": "Category Name"
},
"images": [
{
"id": 1,
"path": "/uploads/photo/product1.jpg"
}
],
"characteristics": [
{
"id": 1,
"value": "Value 1",
"description": "Description 1"
}
]
}
```

#### Get All Products
```
GET /products
```
Query Parameters

limit (optional): Limit the number of products returned.
min (optional): Filter products by minimum price.
max (optional): Filter products by maximum price.
Value and Description (optional)


**Response**
```json
[
{
"id": 1,
"name": "Product Name",
"price": 100,
"discount": 0,
"quantity": 10,
"category_id": 1,
"category": {
"id": 1,
"name": "Category Name"
},
"images": [
{
"id": 1,
"path": "/uploads/photo/product1.jpg"
}
],
"characteristics": [
{
"id": 1,
"value": "Value 1",
"description": "Description 1"
}
]
}
]
```

#### Get Product by ID
```
GET /products/:id
```
**Response**
```json
{
"id": 1,
"name": "Product Name",
"price": 100,
"discount": 0,
"quantity": 10,
"category_id": 1,
"category": {
"id": 1,
"name": "Category Name"
},
"images": [
{
"id": 1,
"path": "/uploads/photo/product1.jpg"
}
],
"characteristics": [
{
"id": 1,
"value": "Value 1",
"description": "Description 1"
}
]
}
```

#### Create Product
```
POST /products
```
**Request Body**
```json
    {
    "product": {
    "name": "New Product",
    "price": 150,
    "discount": 0,
    "quantity": 20,
    "category_id": 1
    },
    "images": [
    {
    "path": "/uploads/photo/new_product.jpg"
    }
    ],
    "characteristics": [
    {
    "value": "New Value",
    "description": "New Description"
    }
    ]
    }
```
**Response**
```json
{
"id": 2,
"name": "New Product",
"price": 150,
"discount": 0,
"quantity": 20,
"category_id": 1,
"category": {
"id": 1,
"name": "Category Name"
},
"images": [
{
"id": 2,
"path": "/uploads/photo/new_product.jpg"
}
],
"characteristics": [
{
"id": 2,
"value": "New Value",
"description": "New Description"
}
]
}
```

#### Update Product
```
PATCH /products/:id
```
**Request Body**
```json
{
"product": {
"name": "Updated Product",
"price": 200,
"quantity": 15,
"category_id": 2
},
"images": [
{
"id": 1,
"path": "/uploads/photo/updated_product.jpg"
}
],
"characteristics": [
{
"id": 1,
"value": "Updated Value",
"description": "Updated Description"
}
]
}
```
**Response**
```json
{
"message": "successfully updated"
}
```

#### Delete Product
```
DELETE /products/:id
```
**Response**
```json
{
"message": "successfully deleted"
}
```

#### Get Products by Category
```
GET /categories/:id/products
```

Query Parameters

sort ( optional ) - "cheap" or "expensive"

**Response**
```json
[
{
"id": 1,
"name": "Product Name",
"price": 100,
"discount": 0,
"quantity": 10,
"category_id": 1,
"category": {
"id": 1,
"name": "Category Name"
},
"images": [
{
"id": 1,
"path": "/uploads/photo/product1.jpg"
}
],
"characteristics": [
{
"id": 1,
"value": "Value 1",
"description": "Description 1"
}
]
}
]
```
