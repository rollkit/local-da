# Mock DA

Mock DA implements the [go-da][go-da] interface over a mock Data Availability service.

It is intended to be used for testing DA layers without having to setup the actual services.

### Usage

```sh
    make build
    ./mock-da
```

should output


```sh
    2024/04/11 12:23:34 Listening on: localhost:7980
```

Which exposes the [go-da] interface over JSONRPC and can be accessed with a http client like [xh][xh]:

```sh
    xh http://127.0.0.1:7980 id=1 method=da.MaxBlobSize | jq
```

should output

```sh
{
  "jsonrpc": "2.0",
  "result": 1974272,
  "id": "1"
}
```

```sh
    xh http://127.0.0.1:7980 id=1 method=da.Submit 'params[][]=SGVsbG8gd28ybGQh' 'params[]:=-2'  'params[]=AAAAAAAAAAAAAAAAAAAAAAAAAAECAwQFBgcICRA=' | jq
```

should output

```sh
{
    "jsonrpc": "2.0",
    "result": [
        "AQAAAAAAAADfgz5/IP20UeF81iRRzDBu/qC8eXr9DUyplrfXod3VOA=="
    ],
    "id": "1"
}
```

```sh
    xh http://127.0.0.1:7980 id=1 method=da.Get 'params[][]=AQAAAAAAAADfgz5/IP20UeF81iRRzDBu/qC8eXr9DUyplrfXod3VOA==' 'params[]=AAAAAAAAAAAAAAAAAAAAAAAAAAECAwQFBgcICRA='
```

should output

```sh
{
    "jsonrpc": "2.0",
    "result": [
        "SGVsbG8gd28ybGQh"
    ],
    "id": "1"
}
```

## References

[1] [go-da][go-da]
[1] [xh][xh]

[go-da]: https://github.com/rollkit/go-da
[xh]: https://github.com/ducaale/xh
