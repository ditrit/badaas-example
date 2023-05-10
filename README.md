# Example of BadAas authentication and object storage

- [Example of BadAas authentication and object storage](#example-of-badaas-authentication-and-object-storage)
  - [Set up](#set-up)
  - [Model definition](#model-definition)
  - [Authentication](#authentication)
  - [Test it](#test-it)
    - [Custom routes](#custom-routes)
    - [EAV routes](#eav-routes)
    - [CRUD routes](#crud-routes)
  - [Explanation](#explanation)

## Set up

This project uses `badctl` to generate the files that allow us to run this example. For installing it, use:

<!-- TODO remove commit when badctl has a first tagged version -->
```bash
go install github.com/ditrit/badaas/tools/badctl@cbd4c9e035709de25df59ec17e4b302b3a7b9931
```

Then generate files to make this project work with `cockroach` as database:

```bash
badctl gen --db_provider cockroachdb
```

For more information about `badctl` refer to [badctl Docs](https://github.com/ditrit/badaas/tools/badctl/README.md).

Finally, you can run the api with:

```bash
badctl run
```

The api will be available at <http://localhost:8000>.

## Model definition

This example defines the following model [here](example.go):

```mermaid
classDiagram
  class Post {
    title String
    body String
    ownerID String
  }
  class Profile {
    displayName String
    urlPic String
    userID String
  }
  class User {
  }
  Profile "1" <--> "*" Post
  Profile "1" <--> "1" User
```

The Users are stored on a classic sql table outside of the object storage engine (EAV), so the userID will be a value. We will use the Profile object to store info about the users.

For now we need to setup the schema manually using a invoke function that will run at the initialization of the programme. Make sure that the database you are gonna use is empty.

## Authentication

Currently we only support a basic authentication using an email and a password.
The default credentials for the user are ̀`admin-no-reply@badaas.com` and `admin`.

## Test it

httpie util will be used in the examples below, but it works with curl or any similar tools.

### Custom routes

Let's first start by checking the route this example adds:

```bash
http localhost:8000/hello
```

```json
"hello world"
```

### EAV routes

Then, we can test the routes provided by BadAss

Get all the profiles:

```bash
http localhost:8000/eav/objects/profile
```

```json
[
    {
        "attrs": {
            "urlPic": "The Super Admin",
            "userId": "wowASuperCoolUserID"
        },
        "createdAt": "2023-01-05T12:47:41.166796+01:00",
        "id": "67c5c6dc-b82d-4c2a-9ab8-f7847a732086",
        "type": "profile",
        "updatedAt": "2023-01-05T12:47:41.166796+01:00"
    }
]
```

Get all the posts posted by this user:

```bash
http GET localhost:8000/eav/objects/post ownerID=wowASuperCoolUserID
```

```json
[
    {
        "attrs": {
            "body": "Lorem ipsum dolor sit amet, consectetur adipiscing elit.\n\n\t\tIn consectetur, ex at hendrerit lobortis, tellus lorem blandit eros, vel ornare odio lorem eget nisi.\n\n\t\tIn erat mi, pharetra ut lacinia at, facilisis vitae nunc.\n\t",
            "ownerID": "wowASuperCoolUserID",
            "title": "Why cats like mice ?"
        },
        "createdAt": "2023-01-06T09:18:53.313565+01:00",
        "id": "68facfa6-fe77-479b-a64e-f26344393a31",
        "type": "post",
        "updatedAt": "2023-01-06T09:18:53.313565+01:00"
    }
]
```

### CRUD routes

Get all the sales:

```bash
http localhost:8000/objects/sale
```

```json
[
    {
        "CreatedAt": "2023-05-10T08:32:11.754637Z",
        "DeletedAt": null,
        "ID": "a9ca9271-8e5e-4774-ab45-7f8ee6328d87",
        "Product": null,
        "ProductID": "64f3331e-77df-403c-a548-5c66df6f0e81",
        "Seller": null,
        "SellerID": "60f87294-6d78-4da8-b1a9-ec5418900ce5",
        "UpdatedAt": "2023-05-10T08:32:11.754637Z"
    },
    {
        "CreatedAt": "2023-05-10T08:32:11.769282Z",
        "DeletedAt": null,
        "ID": "deabdeda-3730-4399-b99f-3268fabdd591",
        "Product": null,
        "ProductID": "19708413-f245-41a0-b9ec-6154c35e2e0a",
        "Seller": null,
        "SellerID": "28086169-269d-493a-9121-69b78b777a27",
        "UpdatedAt": "2023-05-10T08:32:11.769282Z"
    }
]
```

Get all the sales done by a seller (adapt the id according to the response you obtained in last step):

```bash
http GET localhost:8000/objects/sale seller_id=29b027c0-184a-42a7-950e-a5c9b9d6b6e2
```

```json
[
    {
        "CreatedAt": "2023-05-10T08:32:11.754637Z",
        "DeletedAt": null,
        "ID": "a9ca9271-8e5e-4774-ab45-7f8ee6328d87",
        "Product": null,
        "ProductID": "64f3331e-77df-403c-a548-5c66df6f0e81",
        "Seller": null,
        "SellerID": "60f87294-6d78-4da8-b1a9-ec5418900ce5",
        "UpdatedAt": "2023-05-10T08:32:11.754637Z"
    }
]
```json

This is equivalent to:

```bash
http GET localhost:8000/objects/sale seller:='{"id":"29b027c0-184a-42a7-950e-a5c9b9d6b6e2"}'
```

We can also query the attributes of the related objects:

```bash
http GET localhost:8000/objects/sale seller:='{"name":"franco"}'
```

And so on:

```bash
http GET localhost:8000/objects/sale seller:='{"company":{"name":"ditrit"}}'
```

## Explanation

To understand why this example was made in this way refer to [BadAas Docs](https://github.com/ditrit/badaas/README.md#step-by-step-instructions).
