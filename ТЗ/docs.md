## Category API Documentation

This documentation provides details about the endpoints available in the Category API.

### Base URL

The base URL for all endpoints is `/api`.

### Endpoints

#### Get All Categories

- **URL:** `/categories`
- **Method:** `GET`
- **Description:** Retrieves all categories.
- **Query Parameters:**
- `limit` (optional): Limits the number of categories returned. Default is all categories.
- `offset` (optional): Specifies the offset for paginated results. Default is 0.
- **Response:**
- **Status Code:** 200 OK
- **Body:**
```json
[
{
"ID": 1,
"ParentCategoryID": 0,
"Level": 1,
"NameUZ": "Category Name UZ",
"NameKK": "Category Name KK",
"NameRU": "Category Name RU",
"NameEN": "Category Name EN",
"ParentCategory": null,
"ChildrenCategories": []
},
...
]
```

#### Create Category

- **URL:** `/categories`
- **Method:** `POST`
- **Description:** Creates a new category.
- **Request Body:**
```json
{
"ParentCategoryID": 0,
"NameUZ": "Category Name UZ",
"NameKK": "Category Name KK",
"NameRU": "Category Name RU",
"NameEN": "Category Name EN"
}
```
- **Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
"ID": 1,
"ParentCategoryID": 0,
"Level": 1,
"NameUZ": "Category Name UZ",
"NameKK": "Category Name KK",
"NameRU": "Category Name RU",
"NameEN": "Category Name EN",
"ParentCategory": null,
"ChildrenCategories": []
}
```

#### Update Category

- **URL:** `/categories/:id`
- **Method:** `PATCH`
- **Description:** Updates an existing category.
- **URL Parameters:**
- `id`: The ID of the category to be updated.
- **Request Body:** (Fields are optional, only include those to be updated)
```json
{
"NameUZ": "Updated Category Name UZ",
"NameKK": "Updated Category Name KK",
"NameRU": "Updated Category Name RU",
"NameEN": "Updated Category Name EN",
"ParentCategoryID": 0
}
```
- **Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
"message": "updated successfully"
}
```

#### Get Category by ID

- **URL:** `/categories/:id`
- **Method:** `GET`
- **Description:** Retrieves a category by its ID.
- **URL Parameters:**
- `id`: The ID of the category to retrieve.
- **Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
"ID": 1,
"ParentCategoryID": 0,
"Level": 1,
"NameUZ": "Category Name UZ",
"NameKK": "Category Name KK",
"NameRU": "Category Name RU",
"NameEN": "Category Name EN",
"ParentCategory": null,
"ChildrenCategories": []
}
```

#### Delete Category

- **URL:** `/categories/:id`
- **Method:** `DELETE`
- **Description:** Deletes a category by its ID.
- **URL Parameters:**
- `id`: The ID of the category to delete.
- **Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
"message": "Successfully deleted"
}
```

## Language API Documentation

This documentation provides details about the endpoints available in the Language API.

### Base URL

The base URL for all endpoints is `/api`.

### Endpoints

#### Get All Languages

- **URL:** `/languages`
- **Method:** `GET`
- **Description:** Retrieves all languages.
- **Response:**
- **Status Code:** 200 OK
- **Body:**
```json
[
{
"ID": 1,
"Name": "Language Name 1"
},
...
]
```

#### Get Language by ID

- **URL:** `/languages/:id`
- **Method:** `GET`
- **Description:** Retrieves a language by its ID.
- **URL Parameters:**
- `id`: The ID of the language to retrieve.
- **Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
"ID": 1,
"Name": "Language Name 1"
}
```

#### Create Language

- **URL:** `/languages`
- **Method:** `POST`
- **Description:** Creates a new language.
- **Request Body:**
```json
{
"Name": "New Language Name"
}
```
- **Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
"ID": 1,
"Name": "New Language Name"
}
```

#### Update Language

- **URL:** `/languages/:id`
- **Method:** `PATCH`
- **Description:** Updates an existing language.
- **URL Parameters:**
- `id`: The ID of the language to be updated.
- **Request Body:**
```json
{
"Name": "Updated Language Name"
}
```
- **Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
"message": "Language updated successfully"
}
```

#### Delete Language

- **URL:** `/languages/:id`
- **Method:** `DELETE`
- **Description:** Deletes a language by its ID.
- **URL Parameters:**
- `id`: The ID of the language to delete.
- **Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
"message": "Language deleted successfully"
}
```

## Translation API Documentation

This documentation provides details about the endpoints available in the Translation API.

### Base URL

The base URL for all endpoints is `/api`.

### Product Translations

#### Get Product Translations by Product ID

- **URL:** `/products/:product_id/translations`
- **Method:** `GET`
- **Description:** Retrieves all translations of a product by its ID.
- **URL Parameters:**
- `product_id`: The ID of the product to retrieve translations for.
- **Response:**
- **Status Code:** 200 OK
- **Body:**
```json
[
{
"ID": 1,
"ProductID": 1,
"LanguageID": 1,
"Name": "Product Name Translation 1"
},
...
]
```

#### Create Product Translation

- **URL:** `/translations/product`
- **Method:** `POST`
- **Description:** Creates a new translation for a product.
- **Request Body:**
```json
{
"product_id": 1,
"language_id": 1,
"name": "New Product Name Translation"
}
```
- **Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
"message": "Successfully Created"
}
```

#### Update Product Translation

- **URL:** `/translations/product/:id`
- **Method:** `PATCH`
- **Description:** Updates an existing translation of a product.
- **URL Parameters:**
- `id`: The ID of the translation to be updated.
- **Request Body:**
```json
{
"name": "Updated Product Name Translation"
}
```
- **Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
"message": "Successfully Updated"
}
```

#### Delete Product Translation

- **URL:** `/translations/product/:id`
- **Method:** `DELETE`
- **Description:** Deletes a translation of a product by its ID.
- **URL Parameters:**
- `id`: The ID of the translation to delete.
- **Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
"message": "Successfully Deleted"
}
```

### Characteristic Translations

#### Get Characteristic Translations by Characteristic ID

- **URL:** `/characteristics/:characteristic_id/translations`
- **Method:** `GET`
- **Description:** Retrieves all translations of a characteristic by its ID.
- **URL Parameters:**
- `characteristic_id`: The ID of the characteristic to retrieve translations for.
- **Response:**
- **Status Code:** 200 OK
- **Body:**
```json
[
{
"ID": 1,
"CharacteristicID": 1,
"LanguageID": 1,
"Value": "Characteristic Value Translation 1",
"Description": "Characteristic Description Translation 1"
},
...
]
```

#### Create Characteristic Translation

- **URL:** `/translations/characteristic`
- **Method:** `POST`
- **Description:** Creates a new translation for a characteristic.
- **Request Body:**
```json
{
"characteristic_id": 1,
"language_id": 1,
"value": "New Characteristic Value Translation",
"description": "New Characteristic Description Translation"
}
```
- **Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
"message": "Successfully Created"
}
```

#### Update Characteristic Translation

- **URL:** `/translations/characteristic/:id`
- **Method:** `PATCH`
- **Description:** Updates an existing translation of a characteristic.
- **URL Parameters:**
- `id`: The ID of the translation to be updated.
- **Request Body:**
```json
{
"value": "Updated Characteristic Value Translation",
"description": "Updated Characteristic Description Translation"
}
```
- **Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
"message": "Successfully Updated"
}
```

#### Delete Characteristic Translation

- **URL:** `/translations/characteristic/:id`
- **Method:** `DELETE`
- **Description:** Deletes a translation of a characteristic by its ID.
- **URL Parameters:**
- `id`: The ID of the translation to delete.
- **Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
"message": "Successfully Deleted"
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
