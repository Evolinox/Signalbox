# Remote Box API

## Connecting
The API uses :9074 as the standard port if not otherwise defined (See [.env Settings](#env-settings)).

```shell
curl -X {Method} http://localhost:9074/{Endpoint}
```

## Endpoints
Currently available endpoints are

| Method | Endpoint         | Body  | Description                   |
|:-------|:-----------------|:-----:|:------------------------------|
| POST   | /tracks/main/on  |  {}   | Turns on power on main track  |
| POST   | /tracks/main/off |  {}   | Turns off power on main track |

## .env Settings {#env-settings}
You can change following settings

| Name        | Type   |   Default    | Description                               |
|:------------|:-------|:------------:|:------------------------------------------|
| GIN_PORT    | String |    :9074     | Port for Gin API (This API)               |
| SERIAL_PORT | String | /dev/ttyACM0 | Port for Serial Communication with DCC-EX |
| SERIAL_BAUD | Int    |    115200    | Baud Rate of Serial Port                  |