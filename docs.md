

## Media App API Documentation for Frontend Developers

This document outlines the available endpoints and their usage for managing media categories in the Media App backend.

### Base URL

All endpoints are relative to the base URL of the Media App API.

```
http://your-api-url.com/api/categories
```

### Retrieve All Categories

```
GET /api/categories
```

#### Description

Retrieves all categories from the database.

#### Parameters

None

#### Response

- **200 OK**: Returns an array of category objects.
- **500 Internal Server Error**: Indicates a server error.

### Create a New Category

```
POST /api/categories
```

#### Description

Creates a new category in the database.

#### Parameters

- **Request Body**: Category object containing category details.

#### Response

- **201 Created**: Returns the created category object.
- **400 Bad Request**: Indicates invalid request parameters.
- **500 Internal Server Error**: Indicates a server error.

### Update an Existing Category

```
PATCH /api/categories/:id
```

#### Description

Updates an existing category in the database.

#### Parameters

- **:id**: ID of the category to be updated.

#### Request Body

Category object containing updated category details.

#### Response

- **200 OK**: Indicates successful update.
- **404 Not Found**: Indicates that the category with the specified ID was not found.
- **500 Internal Server Error**: Indicates a server error.

### Retrieve a Single Category

```
GET /api/categories/:id
```

#### Description

Retrieves a single category from the database by its ID.

#### Parameters

- **:id**: ID of the category to retrieve.

#### Response

- **200 OK**: Returns the category object.
- **404 Not Found**: Indicates that the category with the specified ID was not found.
- **500 Internal Server Error**: Indicates a server error.

### Delete a Category

```
DELETE /api/categories/:id
```

#### Description

Deletes a category from the database by its ID.

#### Parameters

- **:id**: ID of the category to delete.

#### Response

- **200 OK**: Indicates successful deletion.
- **404 Not Found**: Indicates that the category with the specified ID was not found.
- **500 Internal Server Error**: Indicates a server error.

