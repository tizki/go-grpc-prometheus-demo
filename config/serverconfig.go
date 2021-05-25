package config

import "os"

const (
	defaultHost = "localhost"
	defaultPort = "9090"
)


func GetServerHost() string {
	host := os.Getenv("GRPC_SERVER_HOST")
	if host != "" {
		return host
	}
	return defaultHost

}

func GetServerPort() string {
	port := os.Getenv("GRPC_SERVER_PORT")
	if port != "" {
		return port
	}
	return defaultPort

}