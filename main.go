package main

import (
	"context"
	"example/foo/echo"
	"log"
	"net"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
)

func main() {
	p1, p2 := net.Pipe()
	// Instantiate a local EchoServer.
	go func() {
		server := echo.EchoServer{}

		// Derive a client capability that points to the server.  Note the
		// return type of echo.ServerToClient.  It is of type echo.Echo,
		// which is the client capability.  This capability is bound to the
		// server instance above; calling client methods will result in RPC
		// against the corresponding server method.
		//
		// The client can be shared over the network.
		client := echo.Echo_ServerToClient(server) // TODO: Send down path

		// Expose the client over the network.  The 'rwc' parameter can be any
		// io.ReadWriteCloser.  In practice, it is almost always a net.Conn.
		//
		// Note the BootstrapClient option.  This tells the RPC connection to
		// immediately make the supplied client -- an echo.Echo, in our case
		// -- to the remote endpoint.  The capability that an rpc.Conn exports
		// by default is called the "bootstrap capability".

		// net.pipe
		conn := rpc.NewConn(rpc.NewStreamTransport(p1), &rpc.Options{
			// The BootstrapClient is the RPC interface that will be made available
			// to the remote endpoint by default.  In this case, Echo.
			BootstrapClient: capnp.Client(client),
		})
		defer conn.Close()
		<-conn.Done()
		// Block until the connection terminates.
	}()

	// As before, rwc can be any io.ReadWriteCloser, and will typically be
	// a net.Conn.  The rpc.Options can be nil, if you don't want to override
	// the defaults.
	//
	// Here, we expect to receive an echo.Echo from the remote side.  The
	// remote side is not expecting a capability in return, however, so we
	// don't need to define a bootstrap interface.
	//
	// This last point bears emphasis:  capnp RPC is fully bidirectional!  Both
	// sides of a connection MAY export a boostrap interface, and in such cases,
	// the bootstrap interfaces need not be the same!
	//
	// Again, for the avoidance of doubt:  only the remote side is exporting a
	// bootstrap interface in this example.
	conn := rpc.NewConn(rpc.NewStreamTransport(p2), nil)
	defer conn.Close()

	// Now we resolve the bootstrap interface from the remote EchoServer.
	// Thanks to Cap'n Proto's promise pipelining, this function call does
	// NOT block.  We can start making RPC calls with 'a' immediately, and
	// these will transparently resolve when bootstrapping completes.
	//
	// The context can be used to time-out or otherwise abort the bootstrap
	// call.   It is safe to cancel the context after the first method call
	// on 'a' completes.
	a := echo.Echo(conn.Bootstrap(context.Background()))

	// Okay! Let's make an RPC call!  Remember:  RPC is performed simply by
	// calling a's methods.
	//
	// There are couple of interesting things to note here:
	//  1. We pass a callback function to set parameters on the RPC call.  If the
	//     call takes no arguments, you MAY pass nil.
	//  2. We return a Future type, representing the in-flight RPC call.  As with
	//     the earlier call to Bootstrap, a's methods do not block.  They instead
	//     return a future that eventually resolves with the RPC results. We also
	//     return a release function, which MUST be called when you're done with
	//     the RPC call and its results.
	f, release := a.Send(context.Background(), func(ps echo.Echo_send_Params) error {
		ps.SetMsg("Heyo ma dude!!")
		return nil
	})
	defer release()

	// You can do other things while the RPC call is in-flight.  Everything
	// is asynchronous. For simplicity, we're going to block until the call
	// completes.
	res, err := f.Struct()
	if err != nil {
		panic(err)
	}

	// Lastly, let's print the result.  Recall that 'product' is the name of
	// the return value that we defined in the schema file.
	resp, _ := res.Response()
	log.Println(resp) // prints 84

}

// package main

// import (
// 	"context"
// 	"example/foo/arith"
// 	"log"
// 	"net"

// 	"capnproto.org/go/capnp/v3"
// 	"capnproto.org/go/capnp/v3/rpc"
// )

// func main() {
// 	p1, p2 := net.Pipe()
// 	// Instantiate a local ArithServer.
// 	go func() {
// 		server := arith.ArithServer{}

// 		// Derive a client capability that points to the server.  Note the
// 		// return type of arith.ServerToClient.  It is of type arith.Arith,
// 		// which is the client capability.  This capability is bound to the
// 		// server instance above; calling client methods will result in RPC
// 		// against the corresponding server method.
// 		//
// 		// The client can be shared over the network.
// 		client := arith.Arith_ServerToClient(server)

// 		// Expose the client over the network.  The 'rwc' parameter can be any
// 		// io.ReadWriteCloser.  In practice, it is almost always a net.Conn.
// 		//
// 		// Note the BootstrapClient option.  This tells the RPC connection to
// 		// immediately make the supplied client -- an arith.Arith, in our case
// 		// -- to the remote endpoint.  The capability that an rpc.Conn exports
// 		// by default is called the "bootstrap capability".

// 		// net.pipe
// 		conn := rpc.NewConn(rpc.NewStreamTransport(p1), &rpc.Options{
// 			// The BootstrapClient is the RPC interface that will be made available
// 			// to the remote endpoint by default.  In this case, Arith.
// 			BootstrapClient: capnp.Client(client),
// 		})
// 		defer conn.Close()
// 		<-conn.Done()
// 		// Block until the connection terminates.
// 	}()

// 	// As before, rwc can be any io.ReadWriteCloser, and will typically be
// 	// a net.Conn.  The rpc.Options can be nil, if you don't want to override
// 	// the defaults.
// 	//
// 	// Here, we expect to receive an arith.Arith from the remote side.  The
// 	// remote side is not expecting a capability in return, however, so we
// 	// don't need to define a bootstrap interface.
// 	//
// 	// This last point bears emphasis:  capnp RPC is fully bidirectional!  Both
// 	// sides of a connection MAY export a boostrap interface, and in such cases,
// 	// the bootstrap interfaces need not be the same!
// 	//
// 	// Again, for the avoidance of doubt:  only the remote side is exporting a
// 	// bootstrap interface in this example.
// 	conn := rpc.NewConn(rpc.NewStreamTransport(p2), nil)
// 	defer conn.Close()

// 	// Now we resolve the bootstrap interface from the remote ArithServer.
// 	// Thanks to Cap'n Proto's promise pipelining, this function call does
// 	// NOT block.  We can start making RPC calls with 'a' immediately, and
// 	// these will transparently resolve when bootstrapping completes.
// 	//
// 	// The context can be used to time-out or otherwise abort the bootstrap
// 	// call.   It is safe to cancel the context after the first method call
// 	// on 'a' completes.
// 	a := arith.Arith(conn.Bootstrap(context.Background()))

// 	// Okay! Let's make an RPC call!  Remember:  RPC is performed simply by
// 	// calling a's methods.
// 	//
// 	// There are couple of interesting things to note here:
// 	//  1. We pass a callback function to set parameters on the RPC call.  If the
// 	//     call takes no arguments, you MAY pass nil.
// 	//  2. We return a Future type, representing the in-flight RPC call.  As with
// 	//     the earlier call to Bootstrap, a's methods do not block.  They instead
// 	//     return a future that eventually resolves with the RPC results. We also
// 	//     return a release function, which MUST be called when you're done with
// 	//     the RPC call and its results.
// 	f, release := a.Multiply(context.Background(), func(ps arith.Arith_multiply_Params) error {
// 		ps.SetA(2)
// 		ps.SetB(42)
// 		return nil
// 	})
// 	defer release()

// 	// You can do other things while the RPC call is in-flight.  Everything
// 	// is asynchronous. For simplicity, we're going to block until the call
// 	// completes.
// 	res, err := f.Struct()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Lastly, let's print the result.  Recall that 'product' is the name of
// 	// the return value that we defined in the schema file.
// 	log.Println(res.Product()) // prints 84

// }
