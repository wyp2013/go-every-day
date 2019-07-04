package reslover

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/dns"
	"google.golang.org/grpc/resolver/manual"
)

type resolverType int

const (
	resolverDNS resolverType = iota
	resolverManual

	resolverDNSStr     = "dns"
	resolverManualStr  = "manual"
	resolverUnknownStr = "unknown"
)

func ParseResolverType(rt string) (resolverType, error) {
	switch strings.ToLower(rt) {
	case resolverDNSStr:
		return resolverDNS, nil
	case resolverManualStr:
		return resolverManual, nil
	}

	return -1, fmt.Errorf("Unsupported resolver type: %s", rt)
}

func (r resolverType) String() string {
	switch r {
	case resolverDNS:
		return resolverDNSStr
	case resolverManual:
		return resolverManualStr
	default:
		return resolverUnknownStr
	}
}

func RegisterResolver(rt resolverType, serverIPs string) error {
	var builder resolver.Builder

	switch rt {
	case resolverDNS:
		builder = dns.NewBuilder()
	case resolverManual:
		b, _ := manual.GenerateAndRegisterManualResolver()
		addresses := []resolver.Address{}
		for _, addr := range strings.Split(serverIPs, ",") {
			addresses = append(addresses, resolver.Address{Addr: addr, Type: resolver.Backend})
		}
		b.InitialState(resolver.State{Addresses:addresses})
		builder = b
	default:
		return fmt.Errorf("Unsupported resolver type: %s", rt)
	}

	// Each resolver scheme package calls init() to register
	// itself. So we need to call Register() again here to
	// register the selected builder.
	resolver.Register(builder)
	resolver.SetDefaultScheme(builder.Scheme())
	return nil
}
