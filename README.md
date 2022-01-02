# tcp-proxy

A proxy that forwords from a host to another.

## Building

```shell
go build -ldflags="-X 'main.Version=$(git describe --tags $(git rev-list --tags --max-count=1) || echo "dev")' -X 'main.BuildDate=$(date -u --rfc-2822)' -X 'main.CommitHash=$(git rev-parse HEAD)'"
```

## Usage

```
Usage: tcp-proxy [options]
Options:
  -h, --help                 print help
  -l, --listen-port string   Listen port (default ":8000")
  -t, --target string        Target host
  -p, --target-port int      Target port
  -v, --version              print version
```

## License

Copyright (c) 2022, [@mrinjamul](https://github.com/mrinjamul)
