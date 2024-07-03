Report Generator is an example application to export customers table
from Postgress database into CSV file.

### Usage 
```shell
    ./report-generator [config.json]
```

config is optional as all of the config settings can be provided through
environment variables.

Config format:
```json
{
  "host": "localhost",
  "port": 5432,
  "user": "postgres",
  "password": "my-password",
  "dbname": "database-name",
  "schedule": "@every 5s",
  "reports_path": "output/directory"
}

```
This can be also provided through environment:
`CSV_DBHOST`, `CSV_DBPORT`, `CSV_DBUSER`, `CSV_DBPASS`, `CSV_DBNAME`,
`CSV_SCHEDULE`, `CSV_REPORTS_PATH`.

Schedule format is similar to Unix [cron](https://en.wikipedia.org/wiki/Cron).

To build from sources do `make build`. To run in container through docker compose
do `compose-up` / `compose-down`. To run unit tests do `make test`

