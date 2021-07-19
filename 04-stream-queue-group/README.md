## Usage

In terminal 1:

```console
nats-server -js
```

In terminal 2:

Create a stream.

```console
> nats stream create ticket-new --subjects ticket.new --storage memory --retention work --discard old --max-msgs=-1 --max-bytes=-1 --max-age=-1 --max-msg-size=-1 --dupe-window=2m
Stream ticket-new was created

Information for Stream ticket-new

Configuration:

             Subjects: ticket.new
     Acknowledgements: true
            Retention: Memory - WorkQueue
             Replicas: 1
       Discard Policy: Old
     Duplicate Window: 2m0s
     Maximum Messages: unlimited
        Maximum Bytes: unlimited
          Maximum Age: 0s
 Maximum Message Size: unlimited
    Maximum Consumers: unlimited

State:

            Messages: 0
               Bytes: 0 B
            FirstSeq: 0
             LastSeq: 0
    Active Consumers: 0
```

Run the workers:

```console
go run cmd/worker/main.go
```

In terminal 3:

```console
go run cmd/publisher/main.go
```

The messages will be received by all 3 pull subscribers.

Stop and restart a few times the worker. Messages lost should be
received successfully at next launch. You can also start the publisher
first and not lose any messages.
