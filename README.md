# Introduction

This is a very simple example of how to use nginx to act as a load balancer within docker-compose.

## How it works

This project is really simple, it's built up from two parts:

1. An nginx server
2. A go server which just serves a simple message on the "/hello" endpoint with a message driven from an environment variable

There is then a [docker-compose](./docker-compose.yml) file which spins up one nginx server and two instances of the go server. Each of the go servers is running with a different message to server. The ports on these go servers aren't exposed outside of the docker-compose network, however, the nginx server is, that then balances requests between these two go servers automatically.

## Running the project

To run the project you just need to run the following command:

```bash
docker-compose up -d
```

And then you can hit the end points using curl with:

```bash
curl localhost:8080/hello
```

You should see the message "Hello from service 1" or "Hello from service 2" depending on which server you hit.

