# Remote Box API

## Connecting
The API uses :9074 as the standard port if not otherwise defined (See [.env Settings](#env-settings)).

```shell
curl -X {Method} http://localhost:9074/{Endpoint}
```

## Endpoints
Currently available endpoints are

| Method | Endpoint                    |              Body               | Description                        |
|:-------|:----------------------------|:-------------------------------:|:-----------------------------------|
| POST   | /tracks/main/               |         {"power": bool}         | Set power on main track            |
| POST   | /tracks/prog/               |         {"power": bool}         | Set power on prog track            |
| POST   | /vehicles/:dccAddress/speed | {"speed": int, "forward": bool} | Set speed and direction of Vehicle |
| POST   | /vehicles/:dccAddress/func  | {"function": int, "state": int} | Set function of Vehicle            |

## .env Settings {#env-settings}
You can change following settings

| Name        | Type   |   Default    | Description                               |
|:------------|:-------|:------------:|:------------------------------------------|
| GIN_PORT    | String |    :9074     | Port for Gin API (This API)               |
| SERIAL_PORT | String | /dev/ttyACM0 | Port for Serial Communication with DCC-EX |
| SERIAL_BAUD | Int    |    115200    | Baud Rate of Serial Port                  |