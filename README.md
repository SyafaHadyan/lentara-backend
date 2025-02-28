# Lentara Backend

## Usage

`go run main.go [args]`

## Supported Request

|Request|Route Handler|Function|
|:---|:---|:---|
|GET|/products|Get All Products|
|GET|/products/:id|Get Product by ID|
|GET|/produts/category:category|Get Products by Category|
|GET|/search/:title|Search Product by Name|
|POST|/products|Create New Product|
|PATCH|/products/:id|Edit Product by ID|
|DELETEL|/products/:id|Delete Product by ID|

## Args

`--migrate`

Migrate database
