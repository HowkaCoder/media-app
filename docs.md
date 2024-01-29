

# API Documentation

## Base URL
```
http://your-domain.com/api
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

## Attributes

### Get All Attributes
```
GET /attributes
```
#### Response
```json
[
{
"id": 1,
"nameUZ": "Attribute 1",
"nameKK": "Attribute 1",
"nameRU": "Attribute 1",
"nameEN": "Attribute 1"
},
{
"id": 2,
"nameUZ": "Attribute 2",
"nameKK": "Attribute 2",
"nameRU": "Attribute 2",
"nameEN": "Attribute 2"
}
]
```

### Get Attribute by ID
```
GET /attributes/:id
```
#### Response
```json
{
"id": 1,
"nameUZ": "Attribute 1",
"nameKK": "Attribute 1",
"nameRU": "Attribute 1",
"nameEN": "Attribute 1"
}
```

### Create Attribute
```
POST /attributes
```
#### Request Body
```json
{
"nameUZ": "New Attribute",
"nameKK": "New Attribute",
"nameRU": "New Attribute",
"nameEN": "New Attribute"
}
```
#### Response
```json
{
"id": 3,
"nameUZ": "New Attribute",
"nameKK": "New Attribute",
"nameRU": "New Attribute",
"nameEN": "New Attribute"
}
```

### Update Attribute
```
PATCH /attributes/:id
```
#### Request Body
```json
{
"nameUZ": "Updated Attribute",
"nameKK": "Updated Attribute",
"nameRU": "Updated Attribute",
"nameEN": "Updated Attribute"
}
```
#### Response
```json
{
"id": 1,
"nameUZ": "Updated Attribute",
"nameKK": "Updated Attribute",
"nameRU": "Updated Attribute",
"nameEN": "Updated Attribute"
}
```

### Delete Attribute
```
DELETE /attributes/:id
```
#### Response
```json
{
"message": "Attribute successfully deleted"
}
```

## Products

### Get All Products
```
GET /products
```
#### Response
```json
[
{
"id": 1,
"name": "Product 1",
"category": {
"id": 1,
"name": "Category 1"
},
"images": [
{
"id": 1,
"path": "/path/to/image1.jpg"
},
{
"id": 2,
"path": "/path/to/image2.jpg"
}
],
"attributes": [
{
"id": 1,
"nameUZ": "Attribute 1",
"nameKK": "Attribute 1",
"nameRU": "Attribute 1",
"nameEN": "Attribute 1",
"valueUZ": "Value 1",
"valueKK": "Value 1",
"valueRU": "Value 1",
"valueEN": "Value 1"
},
{
"id": 2,
"nameUZ": "Attribute 2",
"nameKK": "Attribute 2",
"nameRU": "Attribute 2",
"nameEN": "Attribute 2",
"valueUZ": "Value 2",
"valueKK": "Value 2",
"valueRU": "Value 2",
"valueEN": "Value 2"
}
]
},
{
"id": 2,
"name": "Product 2",
...
}
]
```

### Create Product
```
POST /products
```
#### Request Body
```json
{
"product": {
"name": "New Product",
"category_id": 1
},
"images": [
{
"path": "/path/to/image1.jpg"
},
{
"path": "/path/to/image2.jpg"
}
],
"values": [
{
"attribute_id": 1,
"valueUZ": "Value 1",
"valueKK": "Value 1",
"valueRU": "Value 1",
"valueEN": "Value 1"
},
{
"attribute_id": 2,
"valueUZ": "Value 2",
"valueKK": "Value 2",
"valueRU": "Value 2",
"valueEN": "Value 2"
}
]
}
```
#### Response
```json
{
"id": 3,
"name": "New Product",
...
}
```

### Update Product
```
PATCH /products/:id
```
#### Request Body
```json
{
"product": {
"name": "Updated Product",
...
},
"images": [
{
"id": 1,
"path": "/path/to/image1.jpg"
},
{
"id": 2,
"path": "/path/to/image2.jpg"
}
],
"values": [
{
"id": 1,
"attribute_id": 1,
"valueUZ": "Updated Value 1",
...
},
{
"id": 2,
"attribute_id": 2,
"valueUZ": "Updated Value 2",
...
}
]
}
```
#### Response
```json
{
"id": 1,
"name": "Updated Product",
...
}
```

### Delete Product
```
DELETE /products/:id
```
#### Response
```json


{
"message": "Product successfully deleted"
}
```

