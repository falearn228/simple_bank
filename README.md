## Simple bank service

The service that we’re going to build is a simple bank. It will provide APIs for the frontend to do following things:

1. Create and manage bank accounts, which are composed of owner’s name, balance, and currency (RUB, EUR, USD).
2. Record all balance changes to each of the account. So every time some money is added to or subtracted from the account, an account entry record will be created.
3. Perform a money transfer between 2 accounts. This should happen within a transaction, so that either both accounts’ balance are updated successfully or none of them are.

## Setup local development

### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Postman](https://www.postman.com/)
- [Sqlc](https://docs.sqlc.dev/en/latest/overview/install.html)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

    **For Linux dev**
    ```bash
    $ curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$os-$arch.tar.gz | tar xvz
    ```

- [DBML CLI](https://www.dbml.org/cli/#installation)

    ```bash
    npm install -g @dbml/cli
    dbml2sql --version
    ```

- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

    **For Ubuntu**
  
    ```bash
    sudo snap install sqlc
    ```

### Setup infrastructure

- Start postgres container:

    ```bash
    make start
    ```

- Create simple_bank database:

    ```bash
    make createdb
    ```

- Run db migration up all versions:

    ```bash
    make migrateup
    ```

- Run db migration down all versions:

    ```bash
    make migratedown
    ```

### How to generate code

- Generate SQL CRUD with sqlc:

    ```bash
    make sqlc
    ```

- Generate DB mock with gomock:

    ```bash
    make mock
    ```

### How to run

- Run server:

    ```bash
    make server
    ```

- Run test:

    ```bash
    make test
    ```
