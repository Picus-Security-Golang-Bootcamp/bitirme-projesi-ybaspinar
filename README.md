# ybaspinar-bitirme-projesi
## Description

Patika Picus Security Go bootcamp bitirme projesi.
Gin, Postgres, Gorm ve JWT kullanılarak geliştirilmiş bir API.

## Installation

```
git clone https://github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar.git
cd bitirme-projesi-ybaspinar

```

If you want to check out project, you can use the following command:

```
go run .
```

If you want to build project, you can use the following command:

```
go build .
```

## Configuration

Change configuration file to your needs.

Its in pkg/config/config-local.yaml.

# API Examples

For api without any changes to config file, you can use the following endpoints:

## Root

```
http://localhost:8080/api/v1/
```

## Cart

### List user's cart

#### Request

```
GET :8080/api/v1/cart
    HEADER: Authorization: Bearer <token>
```

#### Response

```
{
    "page": 1,
    "pageSize": 100,
    "pageCount": 1,
    "totalCount": 1,
    "items": [
        {
            "id": "6a19487d-f927-43c6-afe5-247f467ff946",
            "userid": "3d68eb31-130b-4cf1-9951-fe0835ad10cf",
            "CreatedAt": "2022-04-18T07:16:33.887799+03:00",
            "UpdatedAt": "2022-04-18T07:16:33.887799+03:00",
            "DeletedAt": null,
            "products": null
        }
    ]
}
```

### Create new cart with given UserID and Products

#### Request

```
    POST :8080/api/v1/cart/create
    HEADER: Authorization: Bearer <token>
    {
    "userid":"3d68eb31-130b-4cf1-9951-fe0835ad10cf",
    "products":[{
    "id":"d6504704-b64f-46fc-aff0-3eda18205328",
    "name": "ayakkabı",
    "sku": "1234567",
    "description": "ayakkabi",
    "price": 12,
    "stock": 6,
    "categoryid":" 1"
    },{
    "id":"6fb1d4c8-978b-4713-a5e0-7e3747088877",
    "name": "ayakkabı",
    "sku": "1234567",
    "description": "ayakkabi",
    "price": 12,
    "stock": 6,
    "categoryid":" 1"}]
    }
```

#### Response

```
    BODY: {
    "id": "6a19487d-f927-43c6-afe5-247f467ff946",
    "userid": "3d68eb31-130b-4cf1-9951-fe0835ad10cf",
    "CreatedAt": "2022-04-18T07:16:33.887799613+03:00",
    "UpdatedAt": "2022-04-18T07:16:33.887799613+03:00",
    "DeletedAt": null,
    "products": [
        {
            "id": "2fc43178-3130-4d67-abfd-29f7123c4509",
            "CreatedAt": "2022-04-18T07:16:33.888915281+03:00",
            "UpdatedAt": "2022-04-18T07:16:33.888915281+03:00",
            "DeletedAt": null,
            "name": "ayakkabı",
            "sku": "1234567",
            "description": "ayakkabi",
            "price": 12,
            "stock": 6,
            "categoryid": " 1"
        },
        {
            "id": "432a2bb7-520e-4796-8c87-a3e3dcaf92ca",
            "CreatedAt": "2022-04-18T07:16:33.888915281+03:00",
            "UpdatedAt": "2022-04-18T07:16:33.888915281+03:00",
            "DeletedAt": null,
            "name": "ayakkabı",
            "sku": "1234567",
            "description": "ayakkabi",
            "price": 12,
            "stock": 6,
            "categoryid": " 1"
        }
    ]
}
```

### Update Cart with given products

#### Request

```
POST :8080/api/v1/cart/update
    HEADER: Authorization: Bearer <token>
    BODY: {
        {
    "userid":"3d68eb31-130b-4cf1-9951-fe0835ad10cf",
    "products":[{
    "id":"d6504704-b64f-46fc-aff0-3eda18205328",
    "name": "ayakkabı",
    "sku": "1234567",
    "description": "ayakkabi",
    "price": 12,
    "stock": 6,
    "categoryid":" 1"
    },{
    "id":"d6504704-b64f-46fc-aff0-3eda18205328",
    "name": "ayakkabı",
    "sku": "1234567",
    "description": "ayakkabi",
    "price": 12,
    "stock": 6,
    "categoryid":" 1"
    },{
    "id":"6fb1d4c8-978b-4713-a5e0-7e3747088877",
    "name": "ayakkabı",
    "sku": "1234567",
    "description": "ayakkabi",
    "price": 12,
    "stock": 6,
    "categoryid":" 1"}]
}
    }
```

#### Response

```
{
    "id": "d302d890-93e4-4682-8a94-07d9b241bbdb",
    "userid": "3d68eb31-130b-4cf1-9951-fe0835ad10cf",
    "CreatedAt": "2022-04-18T07:31:00.725764429+03:00",
    "UpdatedAt": "2022-04-18T07:31:00.725764429+03:00",
    "DeletedAt": null,
    "products": [
        {
            "id": "561e2304-71fb-496d-a723-53610b20b96b",
            "CreatedAt": "2022-04-18T07:31:00.72729133+03:00",
            "UpdatedAt": "2022-04-18T07:31:00.72729133+03:00",
            "DeletedAt": null,
            "name": "ayakkabı",
            "sku": "1234567",
            "description": "ayakkabi",
            "price": 12,
            "stock": 6,
            "categoryid": " 1"
        },
        {
            "id": "b5a53099-e517-4410-aa06-3f239ffe5691",
            "CreatedAt": "2022-04-18T07:31:00.72729133+03:00",
            "UpdatedAt": "2022-04-18T07:31:00.72729133+03:00",
            "DeletedAt": null,
            "name": "ayakkabı",
            "sku": "1234567",
            "description": "ayakkabi",
            "price": 12,
            "stock": 6,
            "categoryid": " 1"
        },
        {
            "id": "df954380-1fa6-470d-a201-d9469ec1257c",
            "CreatedAt": "2022-04-18T07:31:00.72729133+03:00",
            "UpdatedAt": "2022-04-18T07:31:00.72729133+03:00",
            "DeletedAt": null,
            "name": "ayakkabı",
            "sku": "1234567",
            "description": "ayakkabi",
            "price": 12,
            "stock": 6,
            "categoryid": " 1"
        }
    ]
}
```

### Delete users cart

#### Request

```
DELETE :8080/api/v1/cart/
    HEADER: Authorization: Bearer <token>
    BODY :{
    "userid":"3d68eb31-130b-4cf1-9951-fe0835ad10cf"
}
```

#### Response

```
{
    "id": "00000000-0000-0000-0000-000000000000",
    "userid": "3d68eb31-130b-4cf1-9951-fe0835ad10cf",
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": "2022-04-18T07:35:56.845606558+03:00",
    "products": null
}
```

## Category

### Get all categories

#### Request

```
GET :8080/api/v1/category/

```

#### Response

```
        RESPONSE
        {
        "message": "200",
        "data": {
    "page": 1,
    "pageSize": 100,
    "pageCount": 1,
    "totalCount": 16,
    "items": [
        {
            "id": "1",
            "CreatedAt": "2022-04-18T05:56:10.61603+03:00",
            "UpdatedAt": "2022-04-18T05:56:10.61603+03:00",
            "DeletedAt": null,
            "name": "q"
        },
        {
            "id": "2",
            "CreatedAt": "2022-04-18T05:56:10.63311+03:00",
            "UpdatedAt": "2022-04-18T05:56:10.63311+03:00",
            "DeletedAt": null,
            "name": "w"
        }}
```

### Create category from CSV file

#### Request

```
POST :8080/api/v1/category/create
REQUİRED:
HEADER: Authorization: Bearer <token>
BODY: {file : "dumy_data.csv"}
```

#### Response

```
    RESPONSE
        {
        "message": "200",
        "data": {{
        "id": "1",
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "name": "q"
    },
    {
        "id": "2",
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "name": "w"
    }
        }
        {
        "message": "400",
        "ERROR": "File is not valid"
        }
        {
        "message": "401",
        "ERROR": "Status unauthorized"
        }
```

## Orders

### Get all orders from user

#### Request

```
GET :8080/api/v1/orders/
        REQUİRED:
            HEADER: Authorization: Bearer <token>
            BODY: {userId : "e2ae3c48-b660-4247-bc3d-c6ac4fa2254f"}
```

#### Response

```
        {
        "message": "200",
        "data": {
    "page": 1,
    "pageSize": 100,
    "pageCount": 1,
    "totalCount": 1,
    "items": [
        {
            "id": "b9bf3fa0-7a28-4899-a16f-a25498503466",
            "CreatedAt": "2022-04-18T08:12:06.603315+03:00",
            "UpdatedAt": "2022-04-18T08:22:21.253484+03:00",
            "DeletedAt": null,
            "userid": "3d68eb31-130b-4cf1-9951-fe0835ad10cf",
            "cartid": "6a19487d-f927-43c6-afe5-247f467ff946",
            "status": "delivered"
        }
    ]
}
        {"message": "400",
        "ERROR": "File is not valid"}
        {"message": "401",
        "ERROR": "Status unauthorized"}
```

### Create order from user request

#### Request

```
POST :8080/api/v1/orders/create
        REQUİRED:
            HEADER: Authorization: Bearer <token>
            BODY: {userId : "e2ae3c48-b660-4247-bc3d-c6ac4fa2254f"
                   cartId : "f6ba83b0-1494-4bb4-90bf-15eb8f3ab369"}
```

#### Response

```
        {
        "message": "200",
        {
    "id": "0d008c64-17a2-4286-aef1-adaff7e05545",
    "CreatedAt": "2022-04-18T08:06:26.18371696+03:00",
    "UpdatedAt": "2022-04-18T08:06:26.18371696+03:00",
    "DeletedAt": null,
    "userid": "3d68eb31-130b-4cf1-9951-fe0835ad10cf",
    "cartid": "6a19487d-f927-43c6-afe5-247f467ff946",
    "status": ""
}
        "message": "400",
        "ERROR": "File is not valid"
        }
        {
        "message": "401",
        "ERROR": "Status unauthorized"
        }
```

### Cancel order from user request if 14 days not passed

#### Request

```
POST :8080/api/v1/orders/cancel
        REQUİRED:
            HEADER: Authorization: Bearer <token>
            BODY: {  "id":"b9bf3fa0-7a28-4899-a16f-a25498503466",
                    userId : "e2ae3c48-b660-4247-bc3d-c6ac4fa2254f"
                   cartId : "f6ba83b0-1494-4bb4-90bf-15eb8f3ab369"}
```

#### Response

```
        {
        "message": "200",
        {
    "id": "b9bf3fa0-7a28-4899-a16f-a25498503466",
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "userid": "3d68eb31-130b-4cf1-9951-fe0835ad10cf",
    "cartid": "6a19487d-f927-43c6-afe5-247f467ff946",
    "status": ""
}
        {
        "message": "400",
        "ERROR": "File is not valid"
        }
        {
        "message": "401",
        "ERROR": "Status unauthorized"
        }
```

### Get confirmation from user request

#### Request

```
        POST :8080/api/v1/orders/confirm
        REQUEST
        REQUİRED:
            HEADER: Authorization: Bearer <token>
            BODY: {  "id":"b9bf3fa0-7a28-4899-a16f-a25498503466",
                    userId : "e2ae3c48-b660-4247-bc3d-c6ac4fa2254f"
                   cartId : "f6ba83b0-1494-4bb4-90bf-15eb8f3ab369"}
```

#### Response

```
        {
        "message": "200",
        "data": {
    "id": "b9bf3fa0-7a28-4899-a16f-a25498503466",
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "userid": "3d68eb31-130b-4cf1-9951-fe0835ad10cf",
    "cartid": "6a19487d-f927-43c6-afe5-247f467ff946",
    "status": ""
}
        {
        "message": "400",
        "ERROR": "File is not valid"
        }
        {
        "message": "401",
        "ERROR": "Status unauthorized"
        }
```

## User

### Login with given email and password

#### Request

```
POST :8080/api/v1/user/login
        REQUEST
        REQUİRED:
            BODY: {"email" : "y.baspinar@yandex.com"
                   "password" : "123456"}
```

#### Response

```

        RESPONSE
         {
        "message": "200",
        "JWT Toke": eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InkuYmFzcGluYXJAeWFuZGV4LmNvbSIsImV4cCI6MTY1MDMzNTcwNSwiaWF0IjoxNjUwMjQ5MzA1LCJpc0FkbWluIjpmYWxzZSwiaXNzIjoiIiwidXNlcmlkIjoiM2Q2OGViMzEtMTMwYi00Y2YxLTk5NTEtZmUwODM1YWQxMGNmIn0.z9nkJzBcYQ_cKtbh1CvCFj3fnqQy6QQBIRlDBstsyaA
        {
        {"message": "400",
         "error": {}}
```

### Login with given data

#### Request

```
POST :8080/api/v1/user/signup
        REQUEST
        REQUİRED:
            BODY: JSON
                   {"firstname" : "Yusuf"
                   "lastname" : "Başpınar"
                   "email" : "y.baspinar@yandex.com"   //unique,required
                   "password" : "123456"}               //required
```

#### Response

```
        RESPONSE
        {
        "message": "200",
        "JWT Toke": eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InkuYmFzcGluYXJAeWFuZGV4LmNvbSIsImV4cCI6MTY1MDMzNTcwNSwiaWF0IjoxNjUwMjQ5MzA1LCJpc0FkbWluIjpmYWxzZSwiaXNzIjoiIiwidXNlcmlkIjoiM2Q2OGViMzEtMTMwYi00Y2YxLTk5NTEtZmUwODM1YWQxMGNmIn0.z9nkJzBcYQ_cKtbh1CvCFj3fnqQy6QQBIRlDBstsyaA
        {
        "message": "400",
        "error": {
        "Severity": "ERROR",
        "Code": "23505",
        "Message": "duplicate key value violates unique constraint \"users_email_key\"",
        "Detail": "Key (email)=(y.baspinar@yandex.com) already exists.",
        "Hint": "",
        "Position": 0,
        "InternalPosition": 0,
        "InternalQuery": "",
        "Where": "",
        "SchemaName": "public",
        "TableName": "users",
        "ColumnName": "",
        "DataTypeName": "",
        "ConstraintName": "users_email_key",
        "File": "nbtinsert.c",
        "Line": 649,
        "Routine": "_bt_check_unique"
```

## Products

### Get all the products

#### Request

```
get :8080/api/v1/product/
```

#### Response

```
        {
                "message": "200",

           }
{
    "page": 1,
    "pageSize": 100,
    "pageCount": 1,
    "totalCount": 2,
    "items": [
        {
            "id": "27653112-bf2b-48d5-88e2-076fb478c4a2",
            "CreatedAt": "2022-04-18T06:13:07.343131+03:00",
            "UpdatedAt": "2022-04-18T06:13:07.343131+03:00",
            "DeletedAt": null,
            "name": "ayakkabı",
            "sku": "1234567",
            "description": "ayakkabi",
            "price": 12,
            "stock": 6,
            "categoryid": " 1"
        },
        {
            "id": "d6504704-b64f-46fc-aff0-3eda18205328",
            "CreatedAt": "2022-04-18T06:14:04.15513+03:00",
            "UpdatedAt": "2022-04-18T06:14:04.15513+03:00",
            "DeletedAt": null,
            "name": "ayakkabı",
            "sku": "1234567",
            "description": "ayakkabi",
            "price": 12,
            "stock": 6,
            "categoryid": " 1"
        }
    ]
}
```

### Create product from given json

#### Request

```
POST :8080/api/v1/product/create
        REQUİRED:
            AUTHERIZATION: JWT token
            BODY: {
    "name": "ayakkabı",
    "sku": "1234567",
    "description": "ayakkabi",
    "price": 12,
    "stock": 6,
    "categoryid":" 1"
}
```

#### Response

```
        {
        "message": "200",
        {
        "id": "27653112-bf2b-48d5-88e2-076fb478c4a2",
        "CreatedAt": "2022-04-18T06:13:07.343131912+03:00",
        "UpdatedAt": "2022-04-18T06:13:07.343131912+03:00",
        "DeletedAt": null,
        "name": "ayakkabı",
        "sku": "1234567",
        "description": "ayakkabi",
        "price": 12,
        "stock": 6,
         "categoryid": " 1"
        }
        {
        "message": "200",
                }
        "message": "400",
        "ERROR": "File is not valid"
        }
        {
        "message": "401",
        "ERROR": "Status unauthorized"
        }


```

### Update product from given json

#### Request

```
PUT :8080/api/v1/product/
        REQUİRED:
            AUTHERIZATION: JWT token
            BODY: {
    "name": "ayakkabı",
    "sku": "1234567",
    "description": "ayakkabi",
    "price": 12,
    "stock": 6,
    "categoryid":" 1"
}
```

#### Response

```
        {
        "message": "200",
        {
        "id": "27653112-bf2b-48d5-88e2-076fb478c4a2",
        "CreatedAt": "2022-04-18T06:13:07.343131912+03:00",
        "UpdatedAt": "2022-04-18T06:13:07.343131912+03:00",
        "DeletedAt": null,
        "name": "ayakkabı",
        "sku": "1234567",
        "description": "ayakkabi",
        "price": 12,
        "stock": 6,
         "categoryid": " 1"
        }
        {
        "message": "200",
                }
        "message": "400",
        "ERROR": "File is not valid"
        }
        {
        "message": "401",
        "ERROR": "Status unauthorized"
        }


```

### Search product by with given query

#### Request

```
GET :8080/api/v1/product/search?q=REQUESTEDQERYHERE
```

#### Response

````
        {
        "message": "200",
        [
    {
        "id": "27653112-bf2b-48d5-88e2-076fb478c4a2",
        "CreatedAt": "2022-04-18T06:13:07.343131+03:00",
        "UpdatedAt": "2022-04-18T06:13:07.343131+03:00",
        "DeletedAt": null,
        "name": "ayakkabı",
        "sku": "1234567",
        "description": "ayakkabi",
        "price": 12,
        "stock": 6,
        "categoryid": " 1"
    },
    {
        "id": "d6504704-b64f-46fc-aff0-3eda18205328",
        "CreatedAt": "2022-04-18T06:14:04.15513+03:00",
        "UpdatedAt": "2022-04-18T06:14:04.15513+03:00",
        "DeletedAt": null,
        "name": "ayakkabı",
        "sku": "1234567",
        "description": "ayakkabi",
        "price": 12,
        "stock": 6,
        "categoryid": " 1"
    },
    {
        "id": "6fb1d4c8-978b-4713-a5e0-7e3747088877",
        "CreatedAt": "2022-04-18T06:21:01.579221+03:00",
        "UpdatedAt": "2022-04-18T06:21:01.579221+03:00",
        "DeletedAt": null,
        "name": "ayakkabı",
        "sku": "12345662",
        "description": "ayakkabi",
        "price": 12,
        "stock": 6,
        "categoryid": " 1"
    }
        "message": "400",
        "ERROR": "File is not valid"
        }
        {
        "message": "401",
        "ERROR": "Status unauthorized"
        }


 ```~~

### Delete product with given id
#### Request
````

### Delete product with given id

#### Request

```
DELETE :8080/api/v1/product/
REQUİRED:
AUTHERIZATION: JWT token
BODY: {"id": "27653112-bf2b-48d5-88e2-076fb478c4a2"}
```

#### Response

```
        {
        "message": "200",
        {
         "message": "Product deleted"
        }
        "message": "400",
        "ERROR": "File is not valid"
        }
        {
        "message": "401",
        "ERROR": "Status unauthorized"
        }


```
