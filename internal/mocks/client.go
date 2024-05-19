// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	context "context"

	client "github.com/MetalBlockchain/metal-network-runner/client"

	mock "github.com/stretchr/testify/mock"

	rpcpb "github.com/MetalBlockchain/metal-network-runner/rpcpb"

	time "time"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// AddNode provides a mock function with given fields: ctx, name, execPath, opts
func (_m *Client) AddNode(ctx context.Context, name string, execPath string, opts ...client.OpOption) (*rpcpb.AddNodeResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, execPath)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *rpcpb.AddNodeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...client.OpOption) (*rpcpb.AddNodeResponse, error)); ok {
		return rf(ctx, name, execPath, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...client.OpOption) *rpcpb.AddNodeResponse); ok {
		r0 = rf(ctx, name, execPath, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.AddNodeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, ...client.OpOption) error); ok {
		r1 = rf(ctx, name, execPath, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddPermissionlessDelegator provides a mock function with given fields: ctx, validatorSpec
func (_m *Client) AddPermissionlessDelegator(ctx context.Context, validatorSpec []*rpcpb.PermissionlessStakerSpec) (*rpcpb.AddPermissionlessDelegatorResponse, error) {
	ret := _m.Called(ctx, validatorSpec)

	var r0 *rpcpb.AddPermissionlessDelegatorResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*rpcpb.PermissionlessStakerSpec) (*rpcpb.AddPermissionlessDelegatorResponse, error)); ok {
		return rf(ctx, validatorSpec)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*rpcpb.PermissionlessStakerSpec) *rpcpb.AddPermissionlessDelegatorResponse); ok {
		r0 = rf(ctx, validatorSpec)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.AddPermissionlessDelegatorResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*rpcpb.PermissionlessStakerSpec) error); ok {
		r1 = rf(ctx, validatorSpec)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddPermissionlessValidator provides a mock function with given fields: ctx, validatorSpec
func (_m *Client) AddPermissionlessValidator(ctx context.Context, validatorSpec []*rpcpb.PermissionlessStakerSpec) (*rpcpb.AddPermissionlessValidatorResponse, error) {
	ret := _m.Called(ctx, validatorSpec)

	var r0 *rpcpb.AddPermissionlessValidatorResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*rpcpb.PermissionlessStakerSpec) (*rpcpb.AddPermissionlessValidatorResponse, error)); ok {
		return rf(ctx, validatorSpec)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*rpcpb.PermissionlessStakerSpec) *rpcpb.AddPermissionlessValidatorResponse); ok {
		r0 = rf(ctx, validatorSpec)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.AddPermissionlessValidatorResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*rpcpb.PermissionlessStakerSpec) error); ok {
		r1 = rf(ctx, validatorSpec)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddSubnetValidators provides a mock function with given fields: ctx, validatorSpec
func (_m *Client) AddSubnetValidators(ctx context.Context, validatorSpec []*rpcpb.SubnetValidatorsSpec) (*rpcpb.AddSubnetValidatorsResponse, error) {
	ret := _m.Called(ctx, validatorSpec)

	var r0 *rpcpb.AddSubnetValidatorsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*rpcpb.SubnetValidatorsSpec) (*rpcpb.AddSubnetValidatorsResponse, error)); ok {
		return rf(ctx, validatorSpec)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*rpcpb.SubnetValidatorsSpec) *rpcpb.AddSubnetValidatorsResponse); ok {
		r0 = rf(ctx, validatorSpec)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.AddSubnetValidatorsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*rpcpb.SubnetValidatorsSpec) error); ok {
		r1 = rf(ctx, validatorSpec)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AttachPeer provides a mock function with given fields: ctx, nodeName
func (_m *Client) AttachPeer(ctx context.Context, nodeName string) (*rpcpb.AttachPeerResponse, error) {
	ret := _m.Called(ctx, nodeName)

	var r0 *rpcpb.AttachPeerResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*rpcpb.AttachPeerResponse, error)); ok {
		return rf(ctx, nodeName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *rpcpb.AttachPeerResponse); ok {
		r0 = rf(ctx, nodeName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.AttachPeerResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nodeName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *Client) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateBlockchains provides a mock function with given fields: ctx, blockchainSpecs
func (_m *Client) CreateBlockchains(ctx context.Context, blockchainSpecs []*rpcpb.BlockchainSpec) (*rpcpb.CreateBlockchainsResponse, error) {
	ret := _m.Called(ctx, blockchainSpecs)

	var r0 *rpcpb.CreateBlockchainsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*rpcpb.BlockchainSpec) (*rpcpb.CreateBlockchainsResponse, error)); ok {
		return rf(ctx, blockchainSpecs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*rpcpb.BlockchainSpec) *rpcpb.CreateBlockchainsResponse); ok {
		r0 = rf(ctx, blockchainSpecs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.CreateBlockchainsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*rpcpb.BlockchainSpec) error); ok {
		r1 = rf(ctx, blockchainSpecs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSubnets provides a mock function with given fields: ctx, subnetSpecs
func (_m *Client) CreateSubnets(ctx context.Context, subnetSpecs []*rpcpb.SubnetSpec) (*rpcpb.CreateSubnetsResponse, error) {
	ret := _m.Called(ctx, subnetSpecs)

	var r0 *rpcpb.CreateSubnetsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*rpcpb.SubnetSpec) (*rpcpb.CreateSubnetsResponse, error)); ok {
		return rf(ctx, subnetSpecs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*rpcpb.SubnetSpec) *rpcpb.CreateSubnetsResponse); ok {
		r0 = rf(ctx, subnetSpecs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.CreateSubnetsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*rpcpb.SubnetSpec) error); ok {
		r1 = rf(ctx, subnetSpecs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSnapshotNames provides a mock function with given fields: ctx
func (_m *Client) GetSnapshotNames(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Health provides a mock function with given fields: ctx
func (_m *Client) Health(ctx context.Context) (*rpcpb.HealthResponse, error) {
	ret := _m.Called(ctx)

	var r0 *rpcpb.HealthResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*rpcpb.HealthResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *rpcpb.HealthResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.HealthResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBlockchains provides a mock function with given fields: ctx
func (_m *Client) ListBlockchains(ctx context.Context) ([]*rpcpb.CustomChainInfo, error) {
	ret := _m.Called(ctx)

	var r0 []*rpcpb.CustomChainInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*rpcpb.CustomChainInfo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*rpcpb.CustomChainInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*rpcpb.CustomChainInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRpcs provides a mock function with given fields: ctx
func (_m *Client) ListRpcs(ctx context.Context) ([]*rpcpb.BlockchainRpcs, error) {
	ret := _m.Called(ctx)

	var r0 []*rpcpb.BlockchainRpcs
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*rpcpb.BlockchainRpcs, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*rpcpb.BlockchainRpcs); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*rpcpb.BlockchainRpcs)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSubnets provides a mock function with given fields: ctx
func (_m *Client) ListSubnets(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadSnapshot provides a mock function with given fields: ctx, snapshotName, opts
func (_m *Client) LoadSnapshot(ctx context.Context, snapshotName string, opts ...client.OpOption) (*rpcpb.LoadSnapshotResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, snapshotName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *rpcpb.LoadSnapshotResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...client.OpOption) (*rpcpb.LoadSnapshotResponse, error)); ok {
		return rf(ctx, snapshotName, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...client.OpOption) *rpcpb.LoadSnapshotResponse); ok {
		r0 = rf(ctx, snapshotName, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.LoadSnapshotResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...client.OpOption) error); ok {
		r1 = rf(ctx, snapshotName, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PauseNode provides a mock function with given fields: ctx, name
func (_m *Client) PauseNode(ctx context.Context, name string) (*rpcpb.PauseNodeResponse, error) {
	ret := _m.Called(ctx, name)

	var r0 *rpcpb.PauseNodeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*rpcpb.PauseNodeResponse, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *rpcpb.PauseNodeResponse); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.PauseNodeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ping provides a mock function with given fields: ctx
func (_m *Client) Ping(ctx context.Context) (*rpcpb.PingResponse, error) {
	ret := _m.Called(ctx)

	var r0 *rpcpb.PingResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*rpcpb.PingResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *rpcpb.PingResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.PingResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RPCVersion provides a mock function with given fields: ctx
func (_m *Client) RPCVersion(ctx context.Context) (*rpcpb.RPCVersionResponse, error) {
	ret := _m.Called(ctx)

	var r0 *rpcpb.RPCVersionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*rpcpb.RPCVersionResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *rpcpb.RPCVersionResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.RPCVersionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveNode provides a mock function with given fields: ctx, name
func (_m *Client) RemoveNode(ctx context.Context, name string) (*rpcpb.RemoveNodeResponse, error) {
	ret := _m.Called(ctx, name)

	var r0 *rpcpb.RemoveNodeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*rpcpb.RemoveNodeResponse, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *rpcpb.RemoveNodeResponse); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.RemoveNodeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveSnapshot provides a mock function with given fields: ctx, snapshotName
func (_m *Client) RemoveSnapshot(ctx context.Context, snapshotName string) (*rpcpb.RemoveSnapshotResponse, error) {
	ret := _m.Called(ctx, snapshotName)

	var r0 *rpcpb.RemoveSnapshotResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*rpcpb.RemoveSnapshotResponse, error)); ok {
		return rf(ctx, snapshotName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *rpcpb.RemoveSnapshotResponse); ok {
		r0 = rf(ctx, snapshotName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.RemoveSnapshotResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, snapshotName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveSubnetValidator provides a mock function with given fields: ctx, validatorSpec
func (_m *Client) RemoveSubnetValidator(ctx context.Context, validatorSpec []*rpcpb.RemoveSubnetValidatorSpec) (*rpcpb.RemoveSubnetValidatorResponse, error) {
	ret := _m.Called(ctx, validatorSpec)

	var r0 *rpcpb.RemoveSubnetValidatorResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*rpcpb.RemoveSubnetValidatorSpec) (*rpcpb.RemoveSubnetValidatorResponse, error)); ok {
		return rf(ctx, validatorSpec)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*rpcpb.RemoveSubnetValidatorSpec) *rpcpb.RemoveSubnetValidatorResponse); ok {
		r0 = rf(ctx, validatorSpec)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.RemoveSubnetValidatorResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*rpcpb.RemoveSubnetValidatorSpec) error); ok {
		r1 = rf(ctx, validatorSpec)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestartNode provides a mock function with given fields: ctx, name, opts
func (_m *Client) RestartNode(ctx context.Context, name string, opts ...client.OpOption) (*rpcpb.RestartNodeResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *rpcpb.RestartNodeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...client.OpOption) (*rpcpb.RestartNodeResponse, error)); ok {
		return rf(ctx, name, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...client.OpOption) *rpcpb.RestartNodeResponse); ok {
		r0 = rf(ctx, name, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.RestartNodeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...client.OpOption) error); ok {
		r1 = rf(ctx, name, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResumeNode provides a mock function with given fields: ctx, name
func (_m *Client) ResumeNode(ctx context.Context, name string) (*rpcpb.ResumeNodeResponse, error) {
	ret := _m.Called(ctx, name)

	var r0 *rpcpb.ResumeNodeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*rpcpb.ResumeNodeResponse, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *rpcpb.ResumeNodeResponse); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.ResumeNodeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveSnapshot provides a mock function with given fields: ctx, snapshotName
func (_m *Client) SaveSnapshot(ctx context.Context, snapshotName string) (*rpcpb.SaveSnapshotResponse, error) {
	ret := _m.Called(ctx, snapshotName)

	var r0 *rpcpb.SaveSnapshotResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*rpcpb.SaveSnapshotResponse, error)); ok {
		return rf(ctx, snapshotName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *rpcpb.SaveSnapshotResponse); ok {
		r0 = rf(ctx, snapshotName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.SaveSnapshotResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, snapshotName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendOutboundMessage provides a mock function with given fields: ctx, nodeName, peerID, op, msgBody
func (_m *Client) SendOutboundMessage(ctx context.Context, nodeName string, peerID string, op uint32, msgBody []byte) (*rpcpb.SendOutboundMessageResponse, error) {
	ret := _m.Called(ctx, nodeName, peerID, op, msgBody)

	var r0 *rpcpb.SendOutboundMessageResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, uint32, []byte) (*rpcpb.SendOutboundMessageResponse, error)); ok {
		return rf(ctx, nodeName, peerID, op, msgBody)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, uint32, []byte) *rpcpb.SendOutboundMessageResponse); ok {
		r0 = rf(ctx, nodeName, peerID, op, msgBody)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.SendOutboundMessageResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, uint32, []byte) error); ok {
		r1 = rf(ctx, nodeName, peerID, op, msgBody)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields: ctx, execPath, opts
func (_m *Client) Start(ctx context.Context, execPath string, opts ...client.OpOption) (*rpcpb.StartResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, execPath)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *rpcpb.StartResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...client.OpOption) (*rpcpb.StartResponse, error)); ok {
		return rf(ctx, execPath, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...client.OpOption) *rpcpb.StartResponse); ok {
		r0 = rf(ctx, execPath, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.StartResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...client.OpOption) error); ok {
		r1 = rf(ctx, execPath, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Status provides a mock function with given fields: ctx
func (_m *Client) Status(ctx context.Context) (*rpcpb.StatusResponse, error) {
	ret := _m.Called(ctx)

	var r0 *rpcpb.StatusResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*rpcpb.StatusResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *rpcpb.StatusResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.StatusResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stop provides a mock function with given fields: ctx
func (_m *Client) Stop(ctx context.Context) (*rpcpb.StopResponse, error) {
	ret := _m.Called(ctx)

	var r0 *rpcpb.StopResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*rpcpb.StopResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *rpcpb.StopResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.StopResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StreamStatus provides a mock function with given fields: ctx, pushInterval
func (_m *Client) StreamStatus(ctx context.Context, pushInterval time.Duration) (<-chan *rpcpb.ClusterInfo, error) {
	ret := _m.Called(ctx, pushInterval)

	var r0 <-chan *rpcpb.ClusterInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Duration) (<-chan *rpcpb.ClusterInfo, error)); ok {
		return rf(ctx, pushInterval)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Duration) <-chan *rpcpb.ClusterInfo); ok {
		r0 = rf(ctx, pushInterval)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *rpcpb.ClusterInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Duration) error); ok {
		r1 = rf(ctx, pushInterval)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransformElasticSubnets provides a mock function with given fields: ctx, elasticSubnetSpecs
func (_m *Client) TransformElasticSubnets(ctx context.Context, elasticSubnetSpecs []*rpcpb.ElasticSubnetSpec) (*rpcpb.TransformElasticSubnetsResponse, error) {
	ret := _m.Called(ctx, elasticSubnetSpecs)

	var r0 *rpcpb.TransformElasticSubnetsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*rpcpb.ElasticSubnetSpec) (*rpcpb.TransformElasticSubnetsResponse, error)); ok {
		return rf(ctx, elasticSubnetSpecs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*rpcpb.ElasticSubnetSpec) *rpcpb.TransformElasticSubnetsResponse); ok {
		r0 = rf(ctx, elasticSubnetSpecs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.TransformElasticSubnetsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*rpcpb.ElasticSubnetSpec) error); ok {
		r1 = rf(ctx, elasticSubnetSpecs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// URIs provides a mock function with given fields: ctx
func (_m *Client) URIs(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VMID provides a mock function with given fields: ctx, vmName
func (_m *Client) VMID(ctx context.Context, vmName string) (string, error) {
	ret := _m.Called(ctx, vmName)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, vmName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, vmName)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, vmName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitForHealthy provides a mock function with given fields: ctx
func (_m *Client) WaitForHealthy(ctx context.Context) (*rpcpb.WaitForHealthyResponse, error) {
	ret := _m.Called(ctx)

	var r0 *rpcpb.WaitForHealthyResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*rpcpb.WaitForHealthyResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *rpcpb.WaitForHealthyResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.WaitForHealthyResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
