# Game-Service Service

This is the Game-Service service

Generated with

```
micro new github.com/hellodudu/Ultimate/game-service --namespace=ultimate.service --type=game
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: ultimate.service.game
- Type: srv
- Alias: game-service

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./ultimate-service-game
```

Build a docker image
```
make docker
```
