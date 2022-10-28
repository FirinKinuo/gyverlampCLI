# GyverLampCLI

Console interface for managing [GyverLamp](https://alexgyver.ru/gyverlamp/) at your network

## Install

Fetch the [latest release](https://github.com/FirinKinuo/gyverlamp-cli/releases) for your platform

**or** install by Makefile instruction

```bash
sudo make install
```

## Building

For your platform

```
go mod download
CGO_ENABLED=0 go build -v -tags release -o gyverLampCLI-$(cat VERSION) cmd/gyverlamp/main.go
```

**or** by Makefile

```
make build
```

**or** build for all tested platforms

```
make build-all
```