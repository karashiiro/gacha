# gacha
Gacha microservice written in Go.

## Environment variables
```
MYSQL_CONNECTION_STRING: Connection string for local MySQL server, of the form `user:pass@tcp(localhost:3306)/`
REDIS_LOCATION: Location of local Redis server
```

## Architecture
![Architectural diagram](./architecture.png)

## Commands

### Roll
```json
{
    "command": "roll",
    "parameters": ["drop_table_name"]
}
```
Performs a single roll.

Example output:
#### Server
```
{"level":"info","ts":1617728443.9848087,"caller":"gacha/main.go:138","msg":"rolled Drop(id=5, object_id=1, rate=0.2, series_id=1)","correlation_id":"87410385-45a5-4f59-9ac9-9314d5d093b8"}
```

### Client
```
2021/04/06 10:00:43 Rolled object with ID: 1
```
