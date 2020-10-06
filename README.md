<p align="center">
    <img src="img/neon_logo.png?raw=true" width="600"/>
</p>

# Authorization service

[![CodeFactor](https://www.codefactor.io/repository/github/lavash256/neon-auth/badge)](https://www.codefactor.io/repository/github/lavash256/neon-auth)  [![Go Report Card](https://goreportcard.com/badge/github.com/Lavash95/neon-auth)](https://goreportcard.com/report/github.com/Lavash95/neon-auth)  [![codecov](https://codecov.io/gh/lavash256/neon-auth/branch/master/graph/badge.svg)](https://codecov.io/gh/lavash256/neon-auth) [![Build Status](https://travis-ci.org/lavash256/neon-auth.svg?branch=master)](https://travis-ci.org/lavash256/neon-auth)

Neon auth is an easily extensible and adaptive OAuth service based on the GRPC protocol, making it easy to integrate into a microservice architecture.
Using Email and Password for authorization

## Quick start

```
docker pull lavash256/neon-auth
docker run lavash256/neon-auth migrate
docker run lavash256/neon-auth run
```

Before running migrations and services, you need to change the config or add values ​​for environment

## Installation

#### local environment
```
git clone https://github.com/lavash256/neon-auth.git
cd neon-auth
make
```

After executing this command, the bin / directory will appear, where the program binary file is located. It remains to change the config file in the directory or config / or substitute environment variables.

```
./bin/neon-auth -config {path to config file}
```

#### docker environment
```
git clone https://github.com/lavash256/neon-auth.git
cd neon-auth
docker build -t neon/auth . 
```

The required environment variables are written in config / config.yaml

## Migrations
