# Mock DA

Mock DA implements the [go-da][go-da] interface over a mock Data Availability service.

It is intended to be used for testing DA layers without having to set up the actual services.

## Usage

```sh
make build
./mock-da
```

should output

```sh
2024/04/11 12:23:34 Listening on: localhost:7980
```

Which exposes the [go-da] interface over JSONRPC and can be accessed with an HTTP client like [xh][xh]:

You can also run mock-da with a `-listen-all` flag which will make the process listen on 0.0.0.0 so that it can be accessed from other machines.

```sh
$ ./mock-da -listen-all
2024/04/11 12:23:34 Listening on: 0.0.0.0:7980
```

### MaxBlobSize

```sh
xh http://127.0.0.1:7980 id=1 method=da.MaxBlobSize | jq
```

output:

```sh
{
  "jsonrpc": "2.0",
  "result": 1974272,
  "id": "1"
}
```

#### Submit

```sh
xh http://127.0.0.1:7980 id=1 method=da.Submit 'params[][]=SGVsbG8gd28ybGQh' 'params[]:=-2'  'params[]=AAAAAAAAAAAAAAAAAAAAAAAAAAECAwQFBgcICRA=' | jq
```

output:

```sh
{
    "jsonrpc": "2.0",
    "result": [
        "AQAAAAAAAADfgz5/IP20UeF81iRRzDBu/qC8eXr9DUyplrfXod3VOA=="
    ],
    "id": "1"
}
```

#### Get

```sh
xh http://127.0.0.1:7980 id=1 method=da.Get 'params[][]=AQAAAAAAAADfgz5/IP20UeF81iRRzDBu/qC8eXr9DUyplrfXod3VOA==' 'params[]=AAAAAAAAAAAAAAAAAAAAAAAAAAECAwQFBgcICRA='
```

output:

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

[2] [xh][xh]

[go-da]: https://github.com/rollkit/go-da
[xh]: https://github.com/ducaale/xh
