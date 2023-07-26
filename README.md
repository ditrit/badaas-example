# Example of BaDaaS

- [Example of BaDaaS](#example-of-badaas)
  - [Set up](#set-up)
  - [Authentication](#authentication)
  - [Custom route](#custom-route)
  - [Explanation](#explanation)

## Set up

This project uses `badctl` to generate the files that allow us to run this example. For installing it, use:

```bash
go install github.com/ditrit/badaas/tools/badctl@249d3c0
```

Then generate files to make this project work with `cockroach` as database:

```bash
badctl gen
```

For more information about `badctl` refer to [badctl Docs](https://github.com/ditrit/badaas/tools/badctl/README.md).

Finally, you can run the api with:

```bash
make badaas_run
```

The api will be available at <http://localhost:8000>.

httpie util will be used in the examples below, but it works with curl or any similar tools.

## Authentication

Currently we only support a basic authentication using an email and a password.
The default credentials for the user are ̀`admin-no-reply@badaas.com` and `admin`.

```bash
http POST localhost:8000/login email=admin-no-reply@badaas.com password=admin
```

## Custom route

Let's check the route this example adds:

```bash
http localhost:8000/hello
```

```json
HTTP/1.1 200 OK
Content-Length: 13
Content-Type: application/json
Date: Thu, 04 May 2023 09:32:29 GMT

"hello world"
```

## Explanation

To understand why this example was made in this way refer to [BadAas Docs](https://github.com/ditrit/badaas/README.md#step-by-step-instructions).
