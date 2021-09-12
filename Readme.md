# rabbitMQ-go > Example

> This is the simplest implementation of https://github.com/rabbitmq/amqp091-go .

# how to install `rabbitMQ-go`
```bash
    go get github.com/rabbitmq/amqp091-go
```

## This is how to you make connection

```go

	var Host string = "amqp://guest:guest@localhost:5672/"

	con, err := amqp.Dial(Host)

	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

```

## This is how you close a connection

```go
     con.Close()
```
> it's recommended to use Connection with `defer` so that it'll always run at the end of func.

## Read the code to get to know more ;-)

---

### Note :- 

i'm using Docker to spin-up RabbitMQ

### Help :-

if you want to know more about RabbitMQ, check out this video. 

- https://www.youtube.com/watch?v=e03c3CIGtYU