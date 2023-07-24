# Example of BadAas authentication

- [Example of BadAas authentication](#example-of-badaas-authentication)
  - [Set up](#set-up)
  - [Authentication](#authentication)
  - [Test it](#test-it)
    - [Custom routes](#custom-routes)
  - [Explanation](#explanation)

## Set up

This project uses `badctl` to generate the files that allow us to run this example. For installing it, use:

```bash
go install github.com/ditrit/badaas/tools/badctl
```

Then generate files to make this project work with `cockroach` as database:

```bash
badctl gen docker --db_provider cockroachdb
```

For more information about `badctl` refer to [badctl Docs](https://github.com/ditrit/badaas/tools/badctl/README.md).

Finally, you can run the api with:

```bash
make badaas_run
```

The api will be available at <http://localhost:8000>.

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

## Explanation

<!-- TODO add link to new docs -->
To understand why this example was made in this way refer to BaDaaS Docs.
