

# How to test

1. Start RabbitMQ with default username & password **guest:guest**
2. Start one ore more terminals, go to ./worker and run `go run .`
   - Running the consumers (worker) first will ensure the creation of the queue
3. Start one terminal, go to ./producer and run `go run .`

You will see that the workers will share the messages, like a queue.

# Example output

## Terminal consumer 1

```bash
go run .                  
2023/01/20 14:37:08 gorabbit INFO: Processing messages on 1 goroutines                                                                
2023/01/20 14:37:10 consumed: hello, world 0
2023/01/20 14:37:10 consumed: hello, world 2
2023/01/20 14:37:10 consumed: hello, world 4
2023/01/20 14:37:11 consumed: hello, world 6
2023/01/20 14:37:11 consumed: hello, world 9
2023/01/20 14:37:11 consumed: hello, world 10
2023/01/20 14:37:11 consumed: hello, world 12
2023/01/20 14:37:11 consumed: hello, world 14
2023/01/20 14:37:12 consumed: hello, world 16
2023/01/20 14:37:12 consumed: hello, world 18
2023/01/20 14:37:12 consumed: hello, world 21
2023/01/20 14:37:12 consumed: hello, world 23
2023/01/20 14:37:12 consumed: hello, world 25
2023/01/20 14:37:13 consumed: hello, world 27
2023/01/20 14:37:13 consumed: hello, world 29
2023/01/20 14:37:13 consumed: hello, world 30
2023/01/20 14:37:13 consumed: hello, world 32
2023/01/20 14:37:13 consumed: hello, world 34
2023/01/20 14:37:14 consumed: hello, world 36
```

## Terminal consumer 2

```bash                                                    
go run .                                                                                                                                                                    1 ↵ ──(Fri,Jan20)─┘
2023/01/20 14:37:08 gorabbit INFO: Processing messages on 1 goroutines
2023/01/20 14:37:10 consumed: hello, world 1
2023/01/20 14:37:10 consumed: hello, world 3
2023/01/20 14:37:10 consumed: hello, world 5
2023/01/20 14:37:11 consumed: hello, world 7
2023/01/20 14:37:11 consumed: hello, world 8
2023/01/20 14:37:11 consumed: hello, world 11
2023/01/20 14:37:11 consumed: hello, world 13
2023/01/20 14:37:11 consumed: hello, world 15
2023/01/20 14:37:12 consumed: hello, world 17
2023/01/20 14:37:12 consumed: hello, world 19
2023/01/20 14:37:12 consumed: hello, world 20
2023/01/20 14:37:12 consumed: hello, world 22
2023/01/20 14:37:12 consumed: hello, world 24
2023/01/20 14:37:13 consumed: hello, world 26
2023/01/20 14:37:13 consumed: hello, world 28
2023/01/20 14:37:13 consumed: hello, world 31
2023/01/20 14:37:13 consumed: hello, world 33
2023/01/20 14:37:13 consumed: hello, world 35
2023/01/20 14:37:14 consumed: hello, world 37
```

## Terminal producer

```bash
go run .                                                                                  
2023/01/20 14:32:48 gorabbit INFO: closing publisher...
2023/01/20 14:32:48 gorabbit INFO: closing connection manager...
```