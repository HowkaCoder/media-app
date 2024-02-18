
# E-commerce API Documentation 

URL - https://ecommerce.icedev.uz/


### The documentation is not complete yet

Structure of documentation
- [Login](#login) 
- [CallOrders](#call-orders)
    - [Get all call orders](#get-all-call-orders)
    - [Get a single call order](#get-a-single-call-order)
    - [Add new call order](#add-new-call-order)
    - [Delete a call order](#delete-a-call-order)
- [Countries](#countries)
  - [Get all countries](#get-all-countries)
  - [Get a single country](#get-a-single-country)
  - [Add new country](#add-new-country)
  - [Update a country](#update-a-country)
  - [Delete a country](#delete-a-country)
- [Categories](#categories)
  - [Get all categories](#get-all-categories)
  - [Get a single category](#get-a-single-category)
  - [Add new category](#add-new-category)
  - [Update a category](#update-a-category)
  - [Delete a category](#delete-a-category)
- [Products](#products)
   - [Sort products by min and max price](#sort-products-by-min-and-max-price)
   - [Filter products by attribute values](#filter-products-by-attributes)
   - [Sort products](#sort-products)
   - [Add product attributes]()
   - [Delete product attributes]()
- [Users](#users)
  - [Add new user](#add-new-user)
  - [Add new admin](#add-new-admin)
- [Attributes](#attributes)
- [Orders](#orders)
- [Reviews](#reviews)
- [Comments](#comments)

## Login
```
fetch('https://ecommerce.icedev.uz/token',{
          method: 'POST',
          body: JSON.stringify({
              username: 'any username',
              password: 'simple_password'
          }),
          headers: {
              'Content-type': 'application/x-www-form-urlencoded;charset=UTF-8',
          }
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "access_token": ""
  }
</pre>
</details>

## Call Orders
#### Get all call orders
```
fetch('https://ecommerce.icedev.uz/call_orders/')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
[
  {
    "full_name": "some name",
    "phone_number": "+998991234567",
    "start_time": "12:00:00",
    "end_time": "14:00:00",
    "comment": "string",
    "id": 1
  },
  {
    "full_name": " name",
    "phone_number": "+998991234567",
    "start_time": "10:00:00",
    "end_time": "14:00:00",
    "comment": null,
    "id": 3
  }
]
</pre>
</details>

#### Get a single call order
```
fetch('https://ecommerce.icedev.uz/call_orders/1')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "full_name": "some name",
    "phone_number": "+998991234567",
    "start_time": "12:00:00",
    "end_time": "14:00:00",
    "comment": "string",
    "id": 1
  }
</pre>
</details>


#### Add new call order
```
fetch('https://ecommerce.icedev.uz/call_orders/',{
            method:"POST",
            body:JSON.stringify(
                {
                    "full_name": "some name",
                    "phone_number": "+998991234567",
                    "start_time": "12:00:00",
                    "end_time": "14:00:00",
                    "comment": "some comment"
                }
            )
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "full_name": "some name",
    "phone_number": "+998991234567",
    "start_time": "12:00:00",
    "end_time": "14:00:00",
    "comment": "some comment",
    "id": 1
  }
</pre>
</details>


#### Delete a call order
> **Note**
> *Only admin can delete*

```
fetch('https://ecommerce.icedev.uz/call_orders/7',{
            method:"DELETE"
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "ok": True
  }
</pre>
</details>

## Countries
#### Get all countries
```
fetch('https://ecommerce.icedev.uz/countries')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
[
  {
    "country_name": "string",
    "id": 1
  },
  /*...*/
  {
    "country_name": "string",
    "id": 10
  }
]
</pre>
</details>

#### Get a single country
```
fetch('https://ecommerce.icedev.uz/countries/1')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "country_name": "string",
    "id": 1
  }
</pre>
</details>


#### Limit results
```
fetch('https://ecommerce.icedev.uz/countries?limit=5&offset=5')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
[
  {
    "country_name": "string",
    "id": 5
  },
  /*...*/
  {
    "country_name": "string",
    "id": 10
  }
]
</pre>
</details>


#### Add new country
> **Note**
> *Only admin can add*
```
fetch('https://ecommerce.icedev.uz/countries',{
            method:"POST",
            body:JSON.stringify(
                {
                    "country_name": "new country name",
                }
            )
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "country_name": "string",
    "id": 11
  }
</pre>
</details>

#### Update a country
> **Note**
> *Only admin can update*
```
fetch('https://ecommerce.icedev.uz/countries/7',{
            method:"PUT",
            body:JSON.stringify(
                {
                    "country_name": "Updated country name",
                }
            )
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "country_name": "Updated country name",
    "id": 7
  }
</pre>
</details>

#### Delete a country
> **Note**
> *Only admin can delete*
```
fetch('https://ecommerce.icedev.uz/countries/7',{
            method:"DELETE"
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "ok": True
  }
</pre>
</details>

## Categories

#### Get all categories
```
fetch('https://ecommerce.icedev.uz/categories')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
[
  {
    "name": "Электроника",
    "id": 1,
    "children_category": [
      {
        "name": "Телефоны и гаджеты",
        "id": 6
      }
    ],
    "parent_category": null
  },
  /*...*/
  {
    "name": "Телефоны и гаджеты",
    "id": 6,
    "children_category": [
      {
        "name": "Смартфоны",
        "id": 7
      }
    ],
    "parent_category": {
      "name": "Электроника",
      "id": 3
    }
  },
  {
    "name": "Смартфоны",
    "id": 7,
    "children_category": [],
    "parent_category": {
      "name": "Телефоны и гаджеты",
      "id": 6
    }
  }
]
</pre>
</details>

#### Get a single category
```
fetch('https://ecommerce.icedev.uz/categories/1')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "name": "Электроника",
    "id": 1,
    "children_category": [
      {
        "name": "Телефоны и гаджеты",
        "id": 6
      }
    ],
    "parent_category": null
  }
</pre>
</details>

#### Limit results
```
fetch('https://ecommerce.icedev.uz/categories?limit=5&offset=5')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
[
  {
    "name": "Электроника",
    "id": 5,
    "children_category": [
      {
        "name": "Телефоны и гаджеты",
        "id": 6
      }
    ],
    "parent_category": null
  },
  /*...*/
  {
    "name": "Телефоны и гаджеты",
    "id": 9,
    "children_category": [
      {
        "name": "Смартфоны",
        "id": 7
      }
    ],
    "parent_category": {
      "name": "Электроника",
      "id": 3
    }
  },
  {
    "name": "Смартфоны",
    "id": 10,
    "children_category": [],
    "parent_category": {
      "name": "Телефоны и гаджеты",
      "id": 6
    }
  }
]
</pre>
</details>

#### Add new category
> **Note**
> *Only admin can add*
```
fetch('https://ecommerce.icedev.uz/categories',{
            method:"POST",
            body:JSON.stringify(
                {
                   "name": "new category",
                   "parent_category_id": null/int
                }
            )
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "name": "new category",
    "id": 1,
    "children_category": [],
    "parent_category": {
        "name": "parent category",
        "id": 6
      }
  }
</pre>
</details>

#### Update a category
> **Note**
> *Only admin can update*
```
fetch('https://ecommerce.icedev.uz/categories/7',{
            method:"PUT",
            body:JSON.stringify(
                {
                   "name": "updated category",
                   "parent_category_id": null/int
                }
            )
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "name": "updated category",
    "id": 1,
    "children_category": [],
    "parent_category": {
        "name": "parent category",
        "id": 6
      }
  }
</pre>
</details>

#### Delete a category
> **Note**
> *Only admin can delete*
```
fetch('https://ecommerce.icedev.uz/categories/7',{
            method:"DELETE"
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "ok": True
  }
</pre>
</details>


## Products

#### Get all products
```
fetch('https://ecommerce.icedev.uz/products/')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
[
     {
        "name": "iPhone",
        "price": 2000,
        "description": "Some description ....",
        "quantity": 50,
        "discount": 0,
        "id": 1,
        "images": [
            {
                "product_id": 1,
                "product_variants_id": null,
                "image_path": "www.website.com/product_image1.png",
                "id": 1
            },
            {
                "product_id": 1,
                "product_variants_id": null,
                "image_path": "www.website.com/product_image2.png",
                "id": 2
            }
        ],
        "category": {
            "name": "string",
            "id": 1,
            "children_category": [],
            "parent_category": null
        }
    },
    /*...*/
    {
        "name": "TV",
        "price": 500,
        "description": "Some description ....",
        "quantity": 100,
        "discount": 0,
        "id": 10,
        "images": [
            {
                "product_id": 1,
                "product_variants_id": null,
                "image_path": "www.website.com/product_image1.png",
                "id": 1
            },
            {
                "product_id": 1,
                "product_variants_id": null,
                "image_path": "www.website.com/product_image2.png",
                "id": 2
            }
        ],
        "category": {
            "name": "string",
            "id": 5,
            "children_category": [],
            "parent_category": null
        }
    }
]
</pre>
</details>

#### Limit products
```
fetch('https://ecommerce.icedev.uz/products?limit=5')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
[
     {
        "name": "iPhone",
        "price": 2000,
        "description": "Some description ....",
        "quantity": 50,
        "discount": 0,
        "id": 1,
        "images": [
            {
                "product_id": 1,
                "product_variants_id": null,
                "image_path": "www.website.com/product_image1.png",
                "id": 1
            },
            {
                "product_id": 1,
                "product_variants_id": null,
                "image_path": "www.website.com/product_image2.png",
                "id": 2
            }
        ],
        "category": {
            "name": "string",
            "id": 1,
            "children_category": [],
            "parent_category": null
        }
    },
    /*...*/
    {
        "name": "TV",
        "price": 500,
        "description": "Some description ....",
        "quantity": 100,
        "discount": 0,
        "id": 5,
        "images": [
            {
                "product_id": 1,
                "product_variants_id": null,
                "image_path": "www.website.com/product_image1.png",
                "id": 1
            },
            {
                "product_id": 1,
                "product_variants_id": null,
                "image_path": "www.website.com/product_image2.png",
                "id": 2
            }
        ],
        "category": {
            "name": "string",
            "id": 5,
            "children_category": [],
            "parent_category": null
        }
    }
]
</pre>
</details>

#### Get products by category
```
fetch('https://ecommerce.icedev.uz/categories/1/products/')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
[
     {
        "name": "iPhone",
        "price": 2000,
        "description": "Some description ....",
        "quantity": 50,
        "discount": 0,
        "id": 1,
        "images": [
            {
                "product_id": 1,
                "product_variants_id": null,
                "image_path": "www.website.com/product_image1.png",
                "id": 1
            },
            {
                "product_id": 1,
                "product_variants_id": null,
                "image_path": "www.website.com/product_image2.png",
                "id": 2
            }
        ],
        "category": {
            "name": "string",
            "id": 1,
            "children_category": [],
            "parent_category": null
        }
    },
    /*...*/
    {
        "name": "TV",
        "price": 500,
        "description": "Some description ....",
        "quantity": 100,
        "discount": 0,
        "id": 5,
        "images": [
            {
                "product_id": 1,
                "product_variants_id": null,
                "image_path": "www.website.com/product_image1.png",
                "id": 1
            },
            {
                "product_id": 1,
                "product_variants_id": null,
                "image_path": "www.website.com/product_image2.png",
                "id": 2
            }
        ],
        "category": {
            "name": "string",
            "id": 1,
            "children_category": [],
            "parent_category": null
        }
    }
]
</pre>
</details>

#### Sort products by min and max price
```
fetch('https://ecommerce.icedev.uz/categories/{category_id}/products?min_price=500&max_price=1000')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

#### Filter products by attributes
```
fetch('https://ecommerce.icedev.uz/categories/{category_id}/products?filters=variant_id&filters=variant_id&filters=variant_id&....')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

#### Sort products

> **sort=cheap**   *# sort products by ascending price*
>
> **sort=expensive**   *# sort products by descending price*

```
fetch('https://ecommerce.icedev.uz/categories/{category_id}/products?sort=cheap')
            .then(res=>res.json())
            .then(json=>console.log(json))
```


#### Get a single product
```
fetch('https://ecommerce.icedev.uz/products/1')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
     {
        "name": "iPhone",
        "price": 2000,
        "description": "Some description ....",
        "quantity": 50,
        "discount": 0,
        "id": 1,
        "images": [
            {
                "product_id": 1,
                "product_variants_id": null,
                "image_path": "www.website.com/product_image1.png",
                "id": 1
            },
            {
                "product_id": 1,
                "product_variants_id": null,
                "image_path": "www.website.com/product_image2.png",
                "id": 2
            }
        ],
        "category": {
            "name": "Электроника",
            "id": 1,
            "children_category": [],
            "parent_category": null
        }
    }
</pre>
</details>


#### Add new product
> **Note**
> *Only admin can add*
```
fetch('https://ecommerce.icedev.uz/products',{
            method:"POST",
            body:JSON.stringify(
                {
                    "product": {
                        "name": "iPhone",
                        "price": 2000,
                        "description": "Some description here",
                        "quantity": 50,
                        "discount": 0,
                        "category_id": 1
                    },
                    "product_images": [
                        {
                            "image_path": "www.test.com/product_image1"
                        },
                        {
                            "image_path": "www.test.com/product_image2"
                        }
                    ]
                }
            )
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
     {
        "name": "iPhone",
        "price": 2000,
        "description": "Some description ....",
        "quantity": 50,
        "discount": 0,
        "id": 1,
        "images": [
            {
                "product_id": 1,
                "product_variants_id": null,
                "image_path": "www.website.com/product_image1.png",
                "id": 1
            },
            {
                "product_id": 1,
                "product_variants_id": null,
                "image_path": "www.website.com/product_image2.png",
                "id": 2
            }
        ],
        "category": {
            "name": "Электроника",
            "id": 1,
            "children_category": [],
            "parent_category": null
        }
    }
</pre>
</details>

#### Update a product
> **Note**
> *Only admin can update*
```
fetch('https://ecommerce.icedev.uz/products/1',{
            method:"PUT",
            body:JSON.stringify(
                {
                    "name": "iPhone new",
                    "price": 1000,
                    "description": "Some description here updated",
                    "quantity": 100,
                    "discount": 10,
                    "category_id": 1
                }
            )
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
      "name": "iPhone new",
      "price": 1000,
      "description": "Some description here updated",
      "quantity": 100,
      "discount": 10,
      "id": 1,
      "images": [
          {
              "product_id": 1,
              "product_variants_id": null,
              "image_path": "www.website.com/product_image1.png",
              "id": 1
          },
          {
              "product_id": 1,
              "product_variants_id": null,
              "image_path": "www.website.com/product_image2.png",
              "id": 2
          }
      ],
      "category": {
          "name": "Электроника",
          "id": 1,
          "children_category": [],
          "parent_category": null
      }
  }
</pre>
</details>

#### Delete a product
> **Note**
> *Only admin can delete*
```
fetch('https://ecommerce.icedev.uz/products/7',{
            method:"DELETE"
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "ok": True
  }
</pre>
</details>


#### Add new product variant
> **Note**
> *Only admin can add*
```
fetch('https://ecommerce.icedev.uz/products/{product_id}/attributes',{
            method:"POST",
            body:JSON.stringify(
                {
                  "variant_id": 8,
                }
            )
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
    {
    "name": "iPhone 14",
    "price": 1400,
    "description": "string",
    "quantity": 10,
    "discount": 0,
    "id": 1,
    "images": [
      {
        "product_id": 1,
        "image_path": "product1.jpg",
        "id": 1
      }
    ],
    "category": {
      "name": "Smartphones",
      "id": 2,
      "children_category": [],
      "parent_category": null
    },
    "attributes": [
      {
        "value": "green",
        "id": 8,
        "attribute": {
          "id": 3,
          "attribute_name": "color",
          "category_id": 2
        }
      }
    ]
  }
</pre>
</details>



#### Get product attributes
```
fetch('https://ecommerce.icedev.uz/products/{product_id}/attributes')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
[
  {
    "value": "green",
    "id": 8,
    "attribute": {
      "id": 3,
      "attribute_name": "color",
      "category_id": 2
    }
  }
]
</pre>
</details>


#### Delete a product attribute
> **Note**
> *Only admin can delete*
```
fetch('https://ecommerce.icedev.uz/products/{product_id}/{variant_id}',{
            method:"DELETE"
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "ok": True
  }
</pre>
</details>

## Users
#### Get all users
> **Note**
> *Only admin can get*
```
fetch('https://ecommerce.icedev.uz/users/')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>

[
    {
        "username": "string",
        "is_admin": true,
        "id": 1,
        "user_detail": {
            "first_name": "Any name",
            "last_name": "Any surname",
            "user_image": "www.somewebsite.com/user1.jpg",
            "id": 1
        },
        "phone_numbers": [
            {
                "phone_number": "+998991234567",
                "type": "mobile",
                "id": 1
            }
        ],
        "addresses": [
            {
                "street_address": "street 50",
                "postal_code": "123100",
                "city": "Nukus",
                "id": 1,
                "country": {
                    "country_name": "Uzbekistan",
                    "id": 2
                }
            }
        ]
    }, 
    /***/
    {},
    {}
]
</pre>
</details>

#### Limit results
```
fetch('https://ecommerce.icedev.uz/users?limit=5')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>

[
    {
        "username": "string",
        "is_admin": true,
        "id": 1,
        "user_detail": {
            "first_name": "Any name",
            "last_name": "Any surname",
            "user_image": "www.somewebsite.com/user1.jpg",
            "id": 1
        },
        "phone_numbers": [
            {
                "phone_number": "+998991234567",
                "type": "mobile",
                "id": 1
            }
        ],
        "addresses": [
            {
                "street_address": "street 50",
                "postal_code": "123100",
                "city": "Nukus",
                "id": 1,
                "country": {
                    "country_name": "Uzbekistan",
                    "id": 2
                }
            }
        ]
    }, 
    /***/
    {"id":5},
]
</pre>
</details>

#### Get a single user
```
fetch('https://ecommerce.icedev.uz/users/1')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
{
    "username": "string",
    "is_admin": true,
    "id": 1,
    "user_detail": {
        "first_name": "Any name",
        "last_name": "Any surname",
        "user_image": "www.somewebsite.com/user1.jpg",
        "id": 1
    },
    "phone_numbers": [
        {
            "phone_number": "+998991234567",
            "type": "mobile",
            "id": 1
        }
    ],
    "addresses": [
        {
            "street_address": "street 50",
            "postal_code": "123100",
            "city": "Nukus",
            "id": 1,
            "country": {
                "country_name": "Uzbekistan",
                "id": 2
            }
        }
    ]
}
</pre>
</details>


#### Add new user
```
fetch('https://ecommerce.icedev.uz/users',{
            method:"POST",
            body:JSON.stringify(
                {
                    "user": {
                        "username": "string",
                        "password": "string"
                    },
                    "user_detail": {
                        "first_name": "Any name",
                        "last_name": "Any surname",
                        "user_image": "www.somewebsite.com/user1.jpg"
                    },
                    "user_phones": [
                        {
                            "phone_number": "+998991234567",
                            "type": "mobile"
                        }
                    ],
                    "user_address": {
                        "street_address": "street 50",
                        "postal_code": "123100",
                        "city": "Nukus",
                        "country_id": 2
                    }
                }
            )
        })
        .then(res=>res.json())
        .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
{
    "username": "string",
    "id": 1,
    "user_detail": {
        "first_name": "Any name",
        "last_name": "Any surname",
        "user_image": "www.somewebsite.com/user1.jpg",
        "id": 1
    },
    "phone_numbers": [
        {
            "phone_number": "+998991234567",
            "type": "mobile",
            "id": 1
        }
    ],
    "addresses": [
        {
            "street_address": "street 50",
            "postal_code": "123100",
            "city": "Nukus",
            "id": 1,
            "country": {
                "country_name": "Uzbekistan",
                "id": 2
            }
        }
    ]
}
</pre>
</details>

#### Add new admin
> **Note**
> *Only admin can add*
```
fetch('https://ecommerce.icedev.uz/users/admin',{
            method:"POST",
            body:JSON.stringify(
                {
                    "user": {
                        "username": "string",
                        "password": "string"
                    },
                    "user_detail": {
                        "first_name": "Any name",
                        "last_name": "Any surname",
                        "user_image": "www.somewebsite.com/user1.jpg"
                    },
                    "user_phones": [
                        {
                            "phone_number": "+998991234567",
                            "type": "mobile"
                        }
                    ],
                    "user_address": {
                        "street_address": "street 50",
                        "postal_code": "123100",
                        "city": "Nukus",
                        "country_id": 2
                    }
                }
            )
        })
        .then(res=>res.json())
        .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
{
    "username": "string",
    "id": 1,
    "user_detail": {
        "first_name": "Any name",
        "last_name": "Any surname",
        "user_image": "www.somewebsite.com/user1.jpg",
        "id": 1
    },
    "phone_numbers": [
        {
            "phone_number": "+998991234567",
            "type": "mobile",
            "id": 1
        }
    ],
    "addresses": [
        {
            "street_address": "street 50",
            "postal_code": "123100",
            "city": "Nukus",
            "id": 1,
            "country": {
                "country_name": "Uzbekistan",
                "id": 2
            }
        }
    ]
}
</pre>
</details>

#### Update a user
> **Note**
> *Only authenticated user can update*
```
fetch('https://ecommerce.icedev.uz/users/7',{
            method:"PUT",
            body:JSON.stringify(
                {
                    "user": {
                        "username": "new username",
                    },
                    "user_detail": {
                        "first_name": "new name",
                        "last_name": "Mambnew updated surname",
                        "user_image": "www.somewebsite.com/user1.jpg"
                    }
                }
            )
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
    {
        "username": "new username",
        "id": 1,
        "user_detail": {
            "first_name": "new name",
            "last_name": "new updated surname",
            "user_image": "www.somewebsite.com/user1.jpg",
            "id": 1
        },
        "phone_numbers": [
            {
                "phone_number": "+998991234567",
                "type": "mobile",
                "id": 1
            }
        ],
        "addresses": [
            {
                "street_address": "street 50",
                "postal_code": "123100",
                "city": "Nukus",
                "id": 1,
                "country": {
                    "country_name": "Uzbekistan",
                    "id": 2
                }
            }
        ]
    }
</pre>
</details>

#### Delete a user
> **Note**
> *Only admin can delete*
```
fetch('https://ecommerce.icedev.uz/users/7',{
            method:"DELETE"
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "ok": True
  }
</pre>
</details>


## Attributes

#### Get category attributes
```
fetch('https://ecommerce.icedev.uz/categories/{category_id}/attributes')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
[
  {
    "attribute_name": "color",
    "category_id": 1,
    "id": 1,
    "variants": {
      "value": "green",
      "id": 1
    },
    {
      "value": "red",
      "id": 2
    }
  },
  /*...*/
  {
    "attribute_name": "size",
    "category_id": 1,
    "id": 10,
    "variants": {
      "value": "300",
      "id": 5
    },
    {
      "value": "500",
      "id": 6
    }
  },
]
</pre>
</details>

#### Get a single attribute
```
fetch('https://ecommerce.icedev.uz/attributes/1')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "attribute_name": "color",
    "category_id": 1,
    "id": 1,
    "variants": {
      "value": "green",
      "id": 1
    },
    {
      "value": "red",
      "id": 2
    }
  },
</pre>
</details>


#### Limit results
```
fetch('https://ecommerce.icedev.uz/categories/{category_id}/attributes?limit=5&offset=0')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
[
  {
    "attribute_name": "color",
    "category_id": 1,
    "id": 1,
    "variants": {
      "value": "green",
      "id": 1
    },
    {
      "value": "red",
      "id": 2
    }
  },
  /*...*/
  {
    "attribute_name": "size",
    "category_id": 1,
    "id": 5,
    "variants": {
      "value": "300",
      "id": 5
    },
    {
      "value": "500",
      "id": 6
    }
  },
]
</pre>
</details>


#### Add new attribute
> **Note**
> *Only admin can add*
```
fetch('https://ecommerce.icedev.uz/attributes',{
            method:"POST",
            body:JSON.stringify(
                {
                  "attribute": {
                    "attribute_name": "color",
                    "category_id": 1
                  },
                  "variants": [
                    {
                      "value": "red"
                    },
                    {
                      "value": "green"
                    }
                  ]
                }
            )
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "attribute_name": "color",
    "category_id": 1,
    "id": 5
    "variants": [
        {
            "id": 7,
            "value": "red"
        },
        {
            "id": 8,
            "value": "green"
        }
    ]
  }
</pre>
</details>


#### Add new attribute variant
> **Note**
> *Only admin can add*
```
fetch('https://ecommerce.icedev.uz/attributes/{attribute_id}/variants',{
            method:"POST",
            body:JSON.stringify(
                {
                  "value": "blue",
                }
            )
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
      "value": "blue",
      "id": 2,
      "attribute": {
        "id": 1,
        "attribute_name": "color",
        "category_id": 2
      }
  }
</pre>
</details>

#### Update an attribute
> **Note**
> *Only admin can update*
```
fetch('https://ecommerce.icedev.uz/attributes/7',{
            method:"PUT",
            body:JSON.stringify(
                {
                    "attribute_name": 'Updated attribute name,
                    "category_id": 12
                }
            )
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
{
  "attribute_name": "updated attribute name",
  "category_id": 12,
  "id": 2,
  "variants": [
    {
      "value": "green",
      "id": 3
    },
    {
      "value": "red",
      "id": 4
    }
  ]
}
</pre>
</details>

#### Delete an attribute
> **Note**
> *Only admin can delete*
```
fetch('https://ecommerce.icedev.uz/attributes/7',{
            method:"DELETE"
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "ok": True
  }
</pre>
</details>

#### Delete an attribute variant
> **Note**
> *Only admin can delete*
```
fetch('https://ecommerce.icedev.uz/attributes/{attribute_id}/variants/{variant_id}',{
            method:"DELETE"
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "ok": True
  }
</pre>
</details>


### Orders
#### Get all order status
```
fetch('https://ecommerce.icedev.uz/orders/status/')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
[
  {
    "status": "string",
    "id": 1
  },
  /*...*/
  {
    "status": "string",
    "id": 10
  }
]
</pre>
</details>

#### Get a single order status
```
fetch('https://ecommerce.icedev.uz/orders/status/1/')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "status": "string",
    "id": 1
  }
</pre>
</details>

#### Add new order status
> **Note**
> *Only admin can add*
```
fetch('https://ecommerce.icedev.uz/orders/status',{
            method:"POST",
            body:JSON.stringify(
                {
                    "status": "string"
                }
            )
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "status": "string",
    "id": 11
  }
</pre>
</details>

#### Update an order status
> **Note**
> *Only admin can update*
```
fetch('https://ecommerce.icedev.uz/orders/status/7',{
            method:"PUT",
            body:JSON.stringify(
                {
                    "status": "string",
                }
            )
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "status": "Updated status",
    "id": 7
  }
</pre>
</details>

#### Delete an order status
> **Note**
> *Only admin can delete*
```
fetch('https://ecommerce.icedev.uz/orders/status/7',{
            method:"DELETE"
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "ok": True
  }
</pre>
</details>



#### Get all orders
> **Note**
> *Only admin can get*
```
fetch('https://ecommerce.icedev.uz/orders/')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
[
  {
    "user_id": 1,
    "order_date": "2023-01-31",
    "address_id": 1,
    "id": 1,
    "order_details": [],
    "order_status": {
      "status": "string",
      "id": 0
    }
  }
  /*...*/
  {
    "user_id": 5,
    "order_date": "2022-11-29",
    "address_id": 7,
    "id": 10,
    "order_details": [],
    "order_status": {
      "status": "string",
      "id": 0
    }
  }
]
</pre>
</details>


#### Get all user orders
```
fetch('https://ecommerce.icedev.uz/orders/{user_id}')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
[
  {
    "user_id": 1,
    "order_date": "2023-01-31",
    "address_id": 1,
    "id": 1,
    "order_details": [],
    "order_status": {
      "status": "string",
      "id": 0
    }
  }
  /*...*/
  {
    "user_id": 1,
    "order_date": "2022-11-29",
    "address_id": 1,
    "id": 10,
    "order_details": [],
    "order_status": {
      "status": "string",
      "id": 0
    }
  }
]
</pre>
</details>

#### Get a single user order
```
fetch('https://ecommerce.icedev.uz/orders/{user_id}/{order_id}')
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "user_id": 1,
    "order_date": "2022-11-29",
    "address_id": 1,
    "id": 10,
    "order_details": [],
    "order_status": {
      "status": "string",
      "id": 0
    }
  }
</pre>
</details>

#### Add new order
```
fetch('https://ecommerce.icedev.uz/orders/',{
            method:"POST",
            body:JSON.stringify(
                {
                  "order": {
                    "user_id": 1,
                    "order_date": "2023-01-26",
                    "address_id": 1,
                    "order_status_id": 1
                  },
                  "order_details": [
                    {
                      "product_id": 1,
                      "quantity": 2,
                      "price": 1000
                    },
                    {
                      "product_id": 3,
                      "quantity": 1,
                      "price": 500
                    }
                  ]
                }
            )
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
{
  "user_id": 1,
  "order_date": "2023-01-26",
  "address_id": 1,
  "id": 1,
  "order_details": [
    {
        "order_id": 1,
        "product_id": 1,
        "quantity": 2,
        "price": 1000
    },
    {
        "order_id": 1,
        "product_id": 3,
        "quantity": 1,
        "price": 500
    }
  ],
  "order_status": {
    "status": "string",
    "id": 0
  }
}
</pre>
</details>

#### Update an order
```
fetch('https://ecommerce.icedev.uz/orders/{order_id}',{
            method:"PUT",
            body:JSON.stringify(
                {
                  "user_id": 1,
                  "order_date": "2023-01-16",
                  "address_id": 1,
                  "order_status_id": 1
                }
            )
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
  "user_id": 1,
  "order_date": "2023-01-26",
  "address_id": 1,
  "order_status_id": 1,
  "id": 1,
  "order_details": [
    {
      "product_id": 1,
      "quantity": 1,
      "price": 1000
    }
  ],
  "order_status": {
    "status": " status",
    "id": 1
  }
}
</pre>
</details>

#### Delete an order
> **Note**
> *Only admin can delete*
```
fetch('https://ecommerce.icedev.uz/orders/{user_id}/{order_id}',{
            method:"DELETE"
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {
    "ok": True
  }
</pre>
</details>

## Reviews

#### Add new review
> **Note**
> *Only authenticated user can leave review*
```
fetch('https://ecommerce.icedev.uz/products/{product_id}/reviews',{
            method:"POST",
            body:JSON.stringify(
                {
                  "review": {
                    "stars": 5
                  },
                  "comment": {
                    "comment": "some comment",
                    "created_date": "2023-02-16 20:10:20"
                  }
                }
            )
        })
            .then(res=>res.json())
            .then(json=>console.log(json))
```

<details><summary>Output</summary>
<pre>
  {

  }
</pre>
</details>


## Comments
Coming soon...
