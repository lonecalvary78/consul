// Code generated by protoc-gen-grpc-inmem. DO NOT EDIT.

package pbdataplane

import (
	"context"

	grpc "google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// compile-time check to ensure that the generator is implementing all
// of the grpc client interfaces methods.
var _ DataplaneServiceClient = CloningDataplaneServiceClient{}

// IsCloningDataplaneServiceClient is an interface that can be used to detect
// that a DataplaneServiceClient is using the in-memory transport and has already
// been wrapped with a with a CloningDataplaneServiceClient.
type IsCloningDataplaneServiceClient interface {
	IsCloningDataplaneServiceClient() bool
}

// CloningDataplaneServiceClient implements the DataplaneServiceClient interface by wrapping
// another implementation and copying all protobuf messages that pass through the client.
// This is mainly useful to wrap the an in-process client to insulate users of that
// client from having to care about potential immutability of data they receive or having
// the server implementation mutate their internal memory.
type CloningDataplaneServiceClient struct {
	DataplaneServiceClient
}

func NewCloningDataplaneServiceClient(client DataplaneServiceClient) DataplaneServiceClient {
	if cloner, ok := client.(IsCloningDataplaneServiceClient); ok && cloner.IsCloningDataplaneServiceClient() {
		// prevent a double clone if the underlying client is already the cloning client.
		return client
	}

	return CloningDataplaneServiceClient{
		DataplaneServiceClient: client,
	}
}

// IsCloningDataplaneServiceClient implements the IsCloningDataplaneServiceClient interface. This
// is only used to detect wrapped clients that would be double cloning data and prevent that.
func (c CloningDataplaneServiceClient) IsCloningDataplaneServiceClient() bool {
	return true
}

func (c CloningDataplaneServiceClient) GetSupportedDataplaneFeatures(ctx context.Context, in *GetSupportedDataplaneFeaturesRequest, opts ...grpc.CallOption) (*GetSupportedDataplaneFeaturesResponse, error) {
	in = proto.Clone(in).(*GetSupportedDataplaneFeaturesRequest)

	out, err := c.DataplaneServiceClient.GetSupportedDataplaneFeatures(ctx, in)
	if err != nil {
		return nil, err
	}

	return proto.Clone(out).(*GetSupportedDataplaneFeaturesResponse), nil
}

func (c CloningDataplaneServiceClient) GetEnvoyBootstrapParams(ctx context.Context, in *GetEnvoyBootstrapParamsRequest, opts ...grpc.CallOption) (*GetEnvoyBootstrapParamsResponse, error) {
	in = proto.Clone(in).(*GetEnvoyBootstrapParamsRequest)

	out, err := c.DataplaneServiceClient.GetEnvoyBootstrapParams(ctx, in)
	if err != nil {
		return nil, err
	}

	return proto.Clone(out).(*GetEnvoyBootstrapParamsResponse), nil
}