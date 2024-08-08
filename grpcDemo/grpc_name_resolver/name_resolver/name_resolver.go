package name_resolver

import "google.golang.org/grpc/resolver"

// Following is an example name resolver. It includes a
// ResolverBuilder(https://godoc.org/google.golang.org/grpc/resolver#Builder)
// and a Resolver(https://godoc.org/google.golang.org/grpc/resolver#Resolver).
//
// A ResolverBuilder is registered for a scheme (in this example, "example" is
// the scheme). When a ClientConn is created for this scheme, the
// ResolverBuilder will be picked to build a Resolver. Note that a new Resolver
// is built for each ClientConn. The Resolver will watch the updates for the
// target, and send updates to the ClientConn.

// exampleResolverBuilder is a
// ResolverBuilder(https://godoc.org/google.golang.org/grpc/resolver#Builder).
type CustomResolverBuilder struct{}
type customResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

const GrpcServiceScheme = "ns"
const GrpcServiceName = "resolver.scheme.hydracode.io"
const grpcServerAddress = "localhost:9000"

func (r *customResolver) ResolveNow(o resolver.ResolveNowOptions) {
	// 直接从map中取出对于的addrList
	addrStrs := r.addrsStore[r.target.Endpoint()]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}

func (*customResolver) Close() {}

func (*CustomResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &customResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			GrpcServiceName: {grpcServerAddress},
		},
	}
	r.ResolveNow(resolver.ResolveNowOptions{})
	return r, nil
}
func (*CustomResolverBuilder) Scheme() string { return GrpcServiceScheme }
