#How To Use The Demo Server And Client

##Server
The server is located in the server package. It starts a server 
that exposes grpc and http endpoints on the same port. The default port is 9090
and can be change by passing GRPC_SERVER_PORT as environment variable.
When it's running, you should be able to open locahost:9090 in your browser
and see prometheus metrics

##Client
The client for the demo is the client located in democlient package. The 
client will try calling SayHello grpc method on the server. The default server
host is localhost and the port is 9090. Both can be changed by passing the environment
variable GRPC_SERVER_HOST and GRPC_SERVER_PORT.
A successful run of the client will print the response of SayHello method to 
the console.

##Configuration
The demo uses the config package and routeguide (proto) package.