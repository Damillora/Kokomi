# Kokomi

A (WIP) customizable start / new tab page.

## Installation

The easiest way to get started is to use Docker:
```bash
docker pull damillora/kokomi
```

## Requirements

* PostgreSQL database

## Configuration

Kokomi is configured using environment variables:

* `POSTGRES_DATABASE`: DSN string of Postgres Database, see [Gorm documentation](https://gorm.io/docs/connecting_to_the_database.html)
* `AUTH_SECRET`: Secret used to sign JWTs
* `BASE_URL`: Accesible URL of the instance
* `DISABLE_REGISTRATION`: Optional, disable registration on the instance

## Contributing
Kokomi is still in an early stage, but contributions are welcome!

## License
[MIT](https://choosealicense.com/licenses/mit/)
