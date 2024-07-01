package main

import (
	"restxgrpc/grpcserver"
	"restxgrpc/restserver"
)

func main() {
	grpcserver.Run()
	restserver.Run()
}
