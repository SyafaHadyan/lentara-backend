# Lentara Backend

## Usage

`go run main.go [args]`

## Supported Request

|Request|Route Handler|Function|
|:---|:---|:---|
|GET|/products|Get All Products|
|GET|/products/:id|Get Product by ID|
|GET|/produts/category/:category|Get Products by Category|
|GET|/search/:title|Search Product by Name|
|GET|/productspec/:id|Get Product Specification by ID|
|POST|/products|Create New Product|
|POST|/productspec/:id|Create New Product Specifications|
|PATCH|/products/:id|Edit Product by ID|
|DELETE|/products/:id|Delete Product by ID|
|DELETE|/productspec/:id|Delete Product Specifications|

## Args

`--migrate`: Migrate database

## Sample Response

### Get All Products

#### Response Body

```json
{
    "message": [
        {
            "id": "437afdc1-32d4-44d0-af5b-7cee2c15b55e",
            "title": "Go Programming Language",
            "description": "Go book",
            "category": "book",
            "price": 256144,
            "stock": 128,
            "rent_count": 16,
            "rating": 4.8,
            "photo_url": "https://image.example.com/go-book.png",
            "created_at": "2025-02-26T02:17:17Z",
            "updated_at": "2025-02-26T14:24:25Z"
        },
        {
            "id": "4de2993e-811b-462c-a2a1-7788934426db",
            "title": "Pemanas Air",
            "description": "Berfungsi untuk memanaskan air",
            "category": "electronics",
            "price": 12800,
            "stock": 16,
            "rent_count": 0,
            "rating": 0,
            "photo_url": "https://image.example.com/water-heater.png",
            "created_at": "2025-02-28T15:22:14Z",
            "updated_at": "2025-02-28T15:22:14Z"
        },
        {
            "id": "5895dad4-c317-43ec-8a16-ff07b87395e9",
            "title": "Playing Card",
            "description": "Playing card set",
            "category": "hobby",
            "price": 3200,
            "stock": 128,
            "rent_count": 0,
            "rating": 0,
            "photo_url": "https://image.example.com/card-set.png",
            "created_at": "2025-02-28T13:47:37Z",
            "updated_at": "2025-02-28T13:47:37Z"
        },
        {
            "id": "71e0affc-f42f-4a48-839e-3b731b4e3aa6",
            "title": "Java Programming Language",
            "description": "Java book",
            "category": "book",
            "price": 128256,
            "stock": 32,
            "rent_count": 256,
            "rating": 4.8,
            "photo_url": "https://image.example.com/java-book.png",
            "created_at": "2025-02-26T10:16:17Z",
            "updated_at": "2025-02-26T14:30:38Z"
        },
        {
            "id": "74ff6e6f-7fa1-4a04-82c7-e0c3d5808ac6",
            "title": "Java Programming Language",
            "description": "Java OOP Book",
            "category": "book",
            "price": 128256,
            "stock": 32,
            "rent_count": 256,
            "rating": 4.8,
            "photo_url": "https://image.example.com/java-book.png",
            "created_at": "2025-02-26T12:06:54Z",
            "updated_at": "2025-02-26T14:23:34Z"
        },
        {
            "id": "b23e1d35-0098-4480-a7f9-7ce713ffdfc2",
            "title": "Operating System Concepts",
            "description": "Operating system book",
            "category": "book",
            "price": 144000,
            "stock": 64,
            "rent_count": 32,
            "rating": 4.9,
            "photo_url": "https://image.example.com/os-book.png",
            "created_at": "2025-02-26T00:31:54Z",
            "updated_at": "2025-02-26T14:19:29Z"
        },
        {
            "id": "b2d71c54-b53a-4291-babf-533ea61aaee5",
            "title": "Rust Programming Language",
            "description": "Rust book",
            "category": "book",
            "price": 144000,
            "stock": 64,
            "rent_count": 16,
            "rating": 4.9,
            "photo_url": "https://image.example.com/rust-book.png",
            "created_at": "2025-02-26T00:38:01Z",
            "updated_at": "2025-02-26T14:22:28Z"
        },
        {
            "id": "e48ccdfd-359a-4771-a033-58b98851d8c0",
            "title": "Java Car",
            "description": "Java car",
            "category": "car",
            "price": 512256,
            "stock": 2,
            "rent_count": 2,
            "rating": 5,
            "photo_url": "https://image.example.com/java-car.png",
            "created_at": "2025-02-27T16:48:58Z",
            "updated_at": "2025-02-27T16:49:22Z"
        },
        {
            "id": "f0313184-2398-4035-af46-3f8fd3314101",
            "title": "Toaster",
            "description": "Toaster with electricity support",
            "category": "electronics",
            "price": 12800,
            "stock": 32,
            "rent_count": 0,
            "rating": 0,
            "photo_url": "https://image.example.com/bread-toaster.png",
            "created_at": "2025-02-28T13:45:35Z",
            "updated_at": "2025-02-28T13:45:35Z"
        }
    ]
}
```

### Get Product by ID

#### Response Body

```json
{
    "id": "4de2993e-811b-462c-a2a1-7788934426db",
    "title": "Pemanas Air",
    "description": "Berfungsi untuk memanaskan air",
    "category": "electronics",
    "price": 12800,
    "stock": 16,
    "rent_count": 0,
    "rating": 0,
    "photo_url": "https://image.example.com/water-heater.png",
    "created_at": "2025-02-28T15:22:14Z",
    "updated_at": "2025-02-28T15:22:14Z"
}
```

### Get Product by Category

#### Response Body

```json
{
    "payload": [
        {
            "id": "4de2993e-811b-462c-a2a1-7788934426db",
            "title": "Pemanas Air",
            "description": "Berfungsi untuk memanaskan air",
            "category": "electronics",
            "price": 12800,
            "stock": 16,
            "rent_count": 0,
            "rating": 0,
            "photo_url": "https://image.example.com/water-heater.png",
            "created_at": "2025-02-28T15:22:14Z",
            "updated_at": "2025-02-28T15:22:14Z"
        },
        {
            "id": "f0313184-2398-4035-af46-3f8fd3314101",
            "title": "Toaster",
            "description": "Toaster with electricity support",
            "category": "electronics",
            "price": 12800,
            "stock": 32,
            "rent_count": 0,
            "rating": 0,
            "photo_url": "https://image.example.com/bread-toaster.png",
            "created_at": "2025-02-28T13:45:35Z",
            "updated_at": "2025-02-28T13:45:35Z"
        }
    ]
}
```

### Search Product by Name

#### Response Body

```json
{
    "payload": [
        {
            "id": "71e0affc-f42f-4a48-839e-3b731b4e3aa6",
            "title": "Java Programming Language",
            "description": "Java book",
            "category": "book",
            "price": 128256,
            "stock": 32,
            "rent_count": 256,
            "rating": 4.8,
            "photo_url": "https://image.example.com/java-book.png",
            "created_at": "2025-02-26T10:16:17Z",
            "updated_at": "2025-02-26T14:30:38Z"
        },
        {
            "id": "74ff6e6f-7fa1-4a04-82c7-e0c3d5808ac6",
            "title": "Java Programming Language",
            "description": "Java OOP Book",
            "category": "book",
            "price": 128256,
            "stock": 32,
            "rent_count": 256,
            "rating": 4.8,
            "photo_url": "https://image.example.com/java-book.png",
            "created_at": "2025-02-26T12:06:54Z",
            "updated_at": "2025-02-26T14:23:34Z"
        },
        {
            "id": "e48ccdfd-359a-4771-a033-58b98851d8c0",
            "title": "Java Car",
            "description": "Java car",
            "category": "car",
            "price": 512256,
            "stock": 2,
            "rent_count": 2,
            "rating": 5,
            "photo_url": "https://image.example.com/java-car.png",
            "created_at": "2025-02-27T16:48:58Z",
            "updated_at": "2025-02-27T16:49:22Z"
        }
    ]
}
```

### Create New Product

#### Request Body

```json
{
    "title": "Kompor Listrik",
    "description": "Memasak dengan listrik",
    "category": "electronics",
    "price": 6400,
    "stock": 8,
    "photo_url": "https://image.example.com/electric-stove.png"
}
```

#### Response

```json
{
    "message": "succesfully created product",
    "payload": {
        "id": "eea949da-eba8-4e7b-8b02-ead902ea7644",
        "title": "Kompor Listrik",
        "description": "Memasak dengan listrik",
        "category": "electronics",
        "price": 6400,
        "stock": 8,
        "rent_count": 0,
        "rating": 0,
        "photo_url": "https://image.example.com/electric-stove.png",
        "created_at": "2025-02-28T19:45:46.655Z",
        "updated_at": "2025-02-28T19:45:46.655Z"
    }
}
```

### Edit Product by ID

#### Request Body

```json
{
    "rent_count" : 2,
    "rating": 4.8
}
```

#### Response

```json
{
    "id": "eea949da-eba8-4e7b-8b02-ead902ea7644",
    "title": "Kompor Listrik",
    "description": "Memasak dengan listrik",
    "category": "electronics",
    "price": 6400,
    "stock": 8,
    "rent_count": 2,
    "rating": 4.8,
    "photo_url": "https://image.example.com/electric-stove.png",
    "created_at": "2025-02-28T19:45:46Z",
    "updated_at": "2025-02-28T19:47:22Z"
}
```

> Note: data fields are optional

Request below won't change any value except `updated_at`

```json
{

}
```

```json
{
    "id": "eea949da-eba8-4e7b-8b02-ead902ea7644",
    "title": "Kompor Listrik",
    "description": "Memasak dengan listrik",
    "category": "electronics",
    "price": 6400,
    "stock": 8,
    "rent_count": 2,
    "rating": 4.8,
    "photo_url": "https://image.example.com/electric-stove.png",
    "created_at": "2025-02-28T19:45:46Z",
    "updated_at": "2025-02-28T19:49:28Z"
}
```

### Create Product Specification

#### Request Body

```json
{
    "product_id": "18991ddc-4852-480c-a263-7b974b267506",
    "specification_1": "Hardened wood",
    "specification_2": "Rounded Edge",
    "specification_3": "",
    "specification_4": "",
    "specification_5": ""
}
```

#### Response Body

```json
{
    "payload": {
        "product_id": "18991ddc-4852-480c-a263-7b974b267506",
        "specification_1": "Hardened wood",
        "specification_2": "Rounded edge",
        "specification_3": "",
        "specification_4": "",
        "specification_5": ""
    }
}
```

### Edit Product Specification

#### Request Body

```json
{
    "product_id": "18991ddc-4852-480c-a263-7b974b267506",
    "specification_1": "Hardened wood",
    "specification_2": "Rounded Edge",
    "specification_3": "",
    "specification_4": "",
    "specification_5": ""
}
```

#### Response Body

```json
{
    "payload": {
        "product_id": "18991ddc-4852-480c-a263-7b974b267506",
        "specification_1": "Hardened wood",
        "specification_2": "Rounded edge",
        "specification_3": "",
        "specification_4": "",
        "specification_5": ""
    }
}
```

### Delete Product Specification

#### Response Body

```json
{
    "payload": {
        "product_id": "18991ddc-4852-480c-a263-7b974b267506"
    }
}
```

### Delete Product by ID

#### Response Body

```json
{
    "payload": {
        "id": "fd0e46ef-2ac5-47a9-b171-e605c439f9b4"
    }
}
```
