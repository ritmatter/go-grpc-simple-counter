# Counter gRPC Server

A simple example of using Go with Protobuf and gRPC.

Useful for an example of the setup as well as a very simple demonstration
of concurrency in Go using `atomic` to ensure thread-safe updates to a shared
variable and `sync.WaitGroup` to wait for a collection of goroutines to finish.
## To use:

1. Make sure `$GOPATH` has a `src` dir inside.

2. Navigate to the `counter` directory and build the `counter` service:

    ```
    cd counter
    protoc counter.proto -I. --go_out=:$GOPATH/src --go-grpc_out=:$GOPATH/src
    ```

3. In the .mod files for `client` and `server`, update the `replace` line to point to
   your `$GOPATH` rather than the one that is set up for my machine.

4. Start the server:

    ```
    cd server
    go run main.go
    ```

5. Start the client:

    ```
    cd client
    go run main.go
    ```
