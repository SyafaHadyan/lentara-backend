# Lentara Backend

## Usage

`go run main.go [args]`

### Args

`--migrate`: Migrate database

## Supported Request

> Append after `/api/v1/`

|Request|Route Handler|Function|
|:---|:---|:---|
|GET|/products|Get All Products|
|GET|/search/:title|Search Product by Name|
|GET|/products/:id|Get Product by Product ID|
|GET|/produts/category/:category|Get Products by Category|
|GET|/productspec/:id|Get Product Specifications by Product ID|
|GET|/productmedia/:id|Get Product Media by Product ID|
|GET|/cart/cartid/:id|Get Cart by Cart ID|
|POST|/products|Create New Product|
|POST|/productspec/:id|Create New Product Specifications From Product ID|
|POST|/productmedia/:id|Create New Product Media From Product ID|
|POST|/cart/:id|Create New Cart from User ID|
|POST|/users/register|Register New User|
|POST|/users/login|Login|
|PATCH|/products/:id|Edit Product by Product ID|
|PATCH|/productspec/:id|Edit Product Specifications by Product ID|
|PATCH|/productmedia/:id|Edit Product Media by Product ID|
|DELETE|/products/:id|Delete Product by Product ID|
|DELETE|/productspec/:id|Delete Product Specifications by Product ID|
|DELETE|/cart/cartid/:id|Delete Cart by Cart ID|
|DELETE|/cart/cartuser/:id|Delete All Carts from User ID|

## Sample Response

### Get All Products `/products`

#### Response Body

```json
{
    "payload": [
        {
            "id": "18991ddc-4852-480c-a263-7b974b267506",
            "title": "Hyper Tough #2 x 6 inch Screwdriver",
            "description": "Hyper Tough #2 x 6 inch Screwdriver",
            "category": "tools",
            "price": 7800,
            "stock": 1,
            "rent_count": 0,
            "rating": 0,
            "photo_url": "https://m.media-amazon.com/images/G/01/apparel/rcxgs/tile._CB483369110_.gif",
            "created_at": "2025-03-02T07:19:35Z",
            "updated_at": "2025-03-06T16:02:13Z"
        },
        {
            "id": "462b492c-991e-4393-9303-d56651558cdf",
            "title": "onn. 50” Class 4K UHD",
            "description": "onn. 50” Class 4K UHD (2160P) LED Roku Smart Television HDR",
            "category": "electronics",
            "price": 6400,
            "stock": 24500,
            "rent_count": 0,
            "rating": 0,
            "photo_url": "https://i5.walmartimages.com/seo/onn-50-Class-4K-UHD-2160P-LED-Roku-Smart-Television-HDR-100012585_5a6dd417-3795-4dc0-a964-f078638716a8.a3aa3bd9a2ef2a749dc3de23c504748e.jpeg?odnHeight=2000&odnWidth=2000&odnBg=FFFFFF",
            "created_at": "2025-03-05T14:00:44Z",
            "updated_at": "2025-03-05T14:00:44Z"
        },
        {
            "id": "ac6079b2-1bd7-439c-9d54-0e8c1696b46b",
            "title": "javaSport Car",
            "description": "javaSport Car",
            "category": "holiday",
            "price": 128512,
            "stock": 1,
            "rent_count": 0,
            "rating": 0,
            "photo_url": "https://images.pexels.com/photos/63764/pexels-photo-63764.jpeg?cs=srgb&dl=car-cars-lamborghini-aventador-63764.jpg&fm=jpg",
            "created_at": "2025-03-05T14:04:31Z",
            "updated_at": "2025-03-05T14:04:31Z"
        },
        {
            "id": "bc0aa396-83b0-408d-ae71-c5f187396ef8",
            "title": "HORUSDY 10 Pcs Magnetic Screwdriver Set",
            "description": "4 Phillips 5 Flat Head Tips",
            "category": "tools",
            "price": 6400,
            "stock": 12,
            "rent_count": 0,
            "rating": 0,
            "photo_url": "https://i5.walmartimages.com/asr/908e56dd-684c-479f-8d1c-51d494f52486.e48522af7416442373276b09a21aee3f.jpeg?odnHeight=640&odnWidth=640&odnBg=FFFFFF",
            "created_at": "2025-03-05T13:51:57Z",
            "updated_at": "2025-03-05T13:51:57Z"
        },
        {
            "id": "d18891ea-f7f3-4411-a1df-0060f381cc9c",
            "title": "Vim Complete Guide",
            "description": "Beginner's book for learning Vim",
            "category": "hobby",
            "price": 4800,
            "stock": 12,
            "rent_count": 0,
            "rating": 0,
            "photo_url": "https://i.redd.it/6lu41ha1o2461.jpg",
            "created_at": "2025-03-05T14:09:10Z",
            "updated_at": "2025-03-05T14:09:10Z"
        }
    ]
}
```

### Search Product by Name `/search/:title`

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

### Get Product by Product ID `/products/:id`

#### Response Body

```json
{
    "id": "940885ca-def9-467e-b5cd-e4551241bfbc",
    "title": "Outdoor Patio Dining Set Large",
    "description": "Mainstays Albany Lane Steel 6-Piece Outdoor Patio Dining Set with Umbrella: Large Version",
    "category": "holiday",
    "price": 12500,
    "stock": 3,
    "rent_count": 0,
    "rating": 0,
    "photo_url": "https://i5.walmartimages.com/seo/Mainstays-Albany-Lane-Steel-6-Piece-Outdoor-Patio-Dining-Set-with-Umbrella-Red_507de98f-62a3-4960-808c-46e0d3c4a822_1.c3440198a7932410a2e8a1ebd91177e5.jpeg?odnHeight=2000&odnWidth=2000&odnBg=FFFFFF",
    "created_at": "2025-03-07T17:21:07+07:00",
    "updated_at": "2025-03-07T17:21:07+07:00"
}
```

### Get Products by Category `products/category/:category`

#### Response Body

```json
{
    "payload": [
        {
            "id": "462b492c-991e-4393-9303-d56651558cdf",
            "title": "onn. 50” Class 4K UHD",
            "description": "onn. 50” Class 4K UHD (2160P) LED Roku Smart Television HDR",
            "category": "electronics",
            "price": 6400,
            "stock": 24500,
            "rent_count": 0,
            "rating": 0,
            "photo_url": "https://i5.walmartimages.com/seo/onn-50-Class-4K-UHD-2160P-LED-Roku-Smart-Television-HDR-100012585_5a6dd417-3795-4dc0-a964-f078638716a8.a3aa3bd9a2ef2a749dc3de23c504748e.jpeg?odnHeight=2000&odnWidth=2000&odnBg=FFFFFF",
            "created_at": "2025-03-05T14:00:44Z",
            "updated_at": "2025-03-05T14:00:44Z"
        },
        {
            "id": "87640b50-0fec-4fcf-aa72-cb42fdfdceab",
            "title": "PlayStation 5 Digital Console Slim",
            "description": "PlayStation 5 Digital Console Slim",
            "category": "electronics",
            "price": 7850,
            "stock": 11,
            "rent_count": 0,
            "rating": 0,
            "photo_url": "https://i5.walmartimages.com/seo/PlayStation-5-Digital-Console-Slim_330f0b1b-c9b6-4d17-8875-f35fea51bdfd.587fde46f23ab38eb3197552e46f5305.jpeg?odnHeight=2000&odnWidth=2000&odnBg=FFFFFF",
            "created_at": "2025-03-05T14:02:52Z",
            "updated_at": "2025-03-05T14:02:52Z"
        }
    ]
}
```

### Get Product Specifications by Product ID `/productspec/:id`

#### Response Body

```json
{
    "payload": [
        {
            "product_id": "18991ddc-4852-480c-a263-7b974b267506",
            "specification_1": "Hardened wood",
            "specification_2": "Rounded edge",
            "specification_3": "",
            "specification_4": "",
            "specification_5": ""
        }
    ]
}
```

### Get Product Media by Product ID `/productmedia/:id`

#### Response Body

```json
{
    "payload": [
        {
            "id": "18991ddc-4852-480c-a263-7b974b267506",
            "media_1": "https://i5.walmartimages.com/seo/Hyper-Tough-55-Piece-Screwdriver-Set_51e48a96-8faa-41b9-b7a2-cd88e2abf965.73a18f2dab5c7f856fcd5010376cedd5.jpeg?odnHeight=2000&odnWidth=2000&odnBg=FFFFFF",
            "media_2": "https://i5.walmartimages.com/asr/d0c5a616-91ac-4013-ae2d-ba2b8af7b860.9adb851dea5b587608b626aa0ed3ea24.jpeg?odnHeight=2000&odnWidth=2000&odnBg=FFFFFF",
            "media_3": "",
            "media_4": "",
            "media_5": "",
            "media_6": "",
            "media_7": "",
            "media_8": "",
            "media_9": "",
            "media_10": "test"
        }
    ]
}
```
