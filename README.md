# gacha
Gacha microservice written in Go.

## Environment variables
```
GACHA_RMQ_CONNECTION_STRING: Connection string for RabbitMQ server
GACHA_RMQ_CHANNEL: Binding channel for RabbitMQ
GACHA_MYSQL_CONNECTION_STRING: Connection string for local MySQL server, of the form `user:pass@tcp(localhost:3306)/`
GACHA_REDIS_LOCATION: Location of local Redis server
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

#### Client
```
2021/04/06 10:00:43 Rolled object with ID: 1
```

### Set drop table
```json
{
    "command": "set_drop_table",
    "parameters": ["new_drop_table_name", "[]DropInsert"]
}
```
Inserts the provided `DropInsert`s into the database under the provided drop series. The rates of the `DropInsert`s must sum to 1, or the command will fail.

Example output:
#### Server
```
{"level":"info","ts":1617760538.7823374,"caller":"gacha/main.go:213","msg":"set drop table: test","correlation_id":"ac013fa4-e5f3-4b4f-a999-1613138c67b5"}
```

#### Client
```
2021/04/06 18:55:38 Succeeded
```

### Delete drop table
```json
{
    "command": "delete_drop_table",
    "parameters": ["drop_table_name"]
}
```
Deletes a drop table and all drop rates associated with it.

Example output:
#### Server
```
{"level":"info","ts":1617760487.7681694,"caller":"gacha/main.go:254","msg":"deleted drop table: test","correlation_id":"f7d5e6a4-46a3-47f6-b343-bdab6b4221b5"}
```

#### Client
```
2021/04/06 18:54:47 Succeeded
```
