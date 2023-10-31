// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cosmos/staking/v1beta1/tx.proto

package stakingv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_CreateValidator_FullMethodName             = "/cosmos.staking.v1beta1.Msg/CreateValidator"
	Msg_EditValidator_FullMethodName               = "/cosmos.staking.v1beta1.Msg/EditValidator"
	Msg_Delegate_FullMethodName                    = "/cosmos.staking.v1beta1.Msg/Delegate"
	Msg_BeginRedelegate_FullMethodName             = "/cosmos.staking.v1beta1.Msg/BeginRedelegate"
	Msg_Undelegate_FullMethodName                  = "/cosmos.staking.v1beta1.Msg/Undelegate"
	Msg_CancelUnbondingDelegation_FullMethodName   = "/cosmos.staking.v1beta1.Msg/CancelUnbondingDelegation"
	Msg_UpdateParams_FullMethodName                = "/cosmos.staking.v1beta1.Msg/UpdateParams"
	Msg_UnbondValidator_FullMethodName             = "/cosmos.staking.v1beta1.Msg/UnbondValidator"
	Msg_TokenizeShares_FullMethodName              = "/cosmos.staking.v1beta1.Msg/TokenizeShares"
	Msg_RedeemTokensForShares_FullMethodName       = "/cosmos.staking.v1beta1.Msg/RedeemTokensForShares"
	Msg_TransferTokenizeShareRecord_FullMethodName = "/cosmos.staking.v1beta1.Msg/TransferTokenizeShareRecord"
	Msg_DisableTokenizeShares_FullMethodName       = "/cosmos.staking.v1beta1.Msg/DisableTokenizeShares"
	Msg_EnableTokenizeShares_FullMethodName        = "/cosmos.staking.v1beta1.Msg/EnableTokenizeShares"
	Msg_ValidatorBond_FullMethodName               = "/cosmos.staking.v1beta1.Msg/ValidatorBond"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// CreateValidator defines a method for creating a new validator.
	CreateValidator(ctx context.Context, in *MsgCreateValidator, opts ...grpc.CallOption) (*MsgCreateValidatorResponse, error)
	// EditValidator defines a method for editing an existing validator.
	EditValidator(ctx context.Context, in *MsgEditValidator, opts ...grpc.CallOption) (*MsgEditValidatorResponse, error)
	// Delegate defines a method for performing a delegation of coins
	// from a delegator to a validator.
	Delegate(ctx context.Context, in *MsgDelegate, opts ...grpc.CallOption) (*MsgDelegateResponse, error)
	// BeginRedelegate defines a method for performing a redelegation
	// of coins from a delegator and source validator to a destination validator.
	BeginRedelegate(ctx context.Context, in *MsgBeginRedelegate, opts ...grpc.CallOption) (*MsgBeginRedelegateResponse, error)
	// Undelegate defines a method for performing an undelegation from a
	// delegate and a validator.
	Undelegate(ctx context.Context, in *MsgUndelegate, opts ...grpc.CallOption) (*MsgUndelegateResponse, error)
	// CancelUnbondingDelegation defines a method for performing canceling the unbonding delegation
	// and delegate back to previous validator.
	//
	// Since: cosmos-sdk 0.46
	CancelUnbondingDelegation(ctx context.Context, in *MsgCancelUnbondingDelegation, opts ...grpc.CallOption) (*MsgCancelUnbondingDelegationResponse, error)
	// UpdateParams defines an operation for updating the x/staking module
	// parameters.
	// Since: cosmos-sdk 0.47
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	// UnbondValidator defines a method for performing the status transition for a validator
	// from bonded to unbonding
	UnbondValidator(ctx context.Context, in *MsgUnbondValidator, opts ...grpc.CallOption) (*MsgUnbondValidatorResponse, error)
	// TokenizeShares defines a method for tokenizing shares from a validator.
	TokenizeShares(ctx context.Context, in *MsgTokenizeShares, opts ...grpc.CallOption) (*MsgTokenizeSharesResponse, error)
	// RedeemTokensForShares defines a method for redeeming tokens from a validator for
	// shares.
	RedeemTokensForShares(ctx context.Context, in *MsgRedeemTokensForShares, opts ...grpc.CallOption) (*MsgRedeemTokensForSharesResponse, error)
	// TransferTokenizeShareRecord defines a method to transfer ownership of
	// TokenizeShareRecord
	TransferTokenizeShareRecord(ctx context.Context, in *MsgTransferTokenizeShareRecord, opts ...grpc.CallOption) (*MsgTransferTokenizeShareRecordResponse, error)
	// DisableTokenizeShares defines a method to prevent the tokenization of an addresses stake
	DisableTokenizeShares(ctx context.Context, in *MsgDisableTokenizeShares, opts ...grpc.CallOption) (*MsgDisableTokenizeSharesResponse, error)
	// EnableTokenizeShares defines a method to re-enable the tokenization of an addresseses stake
	// after it has been disabled
	EnableTokenizeShares(ctx context.Context, in *MsgEnableTokenizeShares, opts ...grpc.CallOption) (*MsgEnableTokenizeSharesResponse, error)
	// ValidatorBond defines a method for performing a validator self-bond
	ValidatorBond(ctx context.Context, in *MsgValidatorBond, opts ...grpc.CallOption) (*MsgValidatorBondResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateValidator(ctx context.Context, in *MsgCreateValidator, opts ...grpc.CallOption) (*MsgCreateValidatorResponse, error) {
	out := new(MsgCreateValidatorResponse)
	err := c.cc.Invoke(ctx, Msg_CreateValidator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) EditValidator(ctx context.Context, in *MsgEditValidator, opts ...grpc.CallOption) (*MsgEditValidatorResponse, error) {
	out := new(MsgEditValidatorResponse)
	err := c.cc.Invoke(ctx, Msg_EditValidator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Delegate(ctx context.Context, in *MsgDelegate, opts ...grpc.CallOption) (*MsgDelegateResponse, error) {
	out := new(MsgDelegateResponse)
	err := c.cc.Invoke(ctx, Msg_Delegate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BeginRedelegate(ctx context.Context, in *MsgBeginRedelegate, opts ...grpc.CallOption) (*MsgBeginRedelegateResponse, error) {
	out := new(MsgBeginRedelegateResponse)
	err := c.cc.Invoke(ctx, Msg_BeginRedelegate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Undelegate(ctx context.Context, in *MsgUndelegate, opts ...grpc.CallOption) (*MsgUndelegateResponse, error) {
	out := new(MsgUndelegateResponse)
	err := c.cc.Invoke(ctx, Msg_Undelegate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelUnbondingDelegation(ctx context.Context, in *MsgCancelUnbondingDelegation, opts ...grpc.CallOption) (*MsgCancelUnbondingDelegationResponse, error) {
	out := new(MsgCancelUnbondingDelegationResponse)
	err := c.cc.Invoke(ctx, Msg_CancelUnbondingDelegation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UnbondValidator(ctx context.Context, in *MsgUnbondValidator, opts ...grpc.CallOption) (*MsgUnbondValidatorResponse, error) {
	out := new(MsgUnbondValidatorResponse)
	err := c.cc.Invoke(ctx, Msg_UnbondValidator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TokenizeShares(ctx context.Context, in *MsgTokenizeShares, opts ...grpc.CallOption) (*MsgTokenizeSharesResponse, error) {
	out := new(MsgTokenizeSharesResponse)
	err := c.cc.Invoke(ctx, Msg_TokenizeShares_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RedeemTokensForShares(ctx context.Context, in *MsgRedeemTokensForShares, opts ...grpc.CallOption) (*MsgRedeemTokensForSharesResponse, error) {
	out := new(MsgRedeemTokensForSharesResponse)
	err := c.cc.Invoke(ctx, Msg_RedeemTokensForShares_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransferTokenizeShareRecord(ctx context.Context, in *MsgTransferTokenizeShareRecord, opts ...grpc.CallOption) (*MsgTransferTokenizeShareRecordResponse, error) {
	out := new(MsgTransferTokenizeShareRecordResponse)
	err := c.cc.Invoke(ctx, Msg_TransferTokenizeShareRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DisableTokenizeShares(ctx context.Context, in *MsgDisableTokenizeShares, opts ...grpc.CallOption) (*MsgDisableTokenizeSharesResponse, error) {
	out := new(MsgDisableTokenizeSharesResponse)
	err := c.cc.Invoke(ctx, Msg_DisableTokenizeShares_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) EnableTokenizeShares(ctx context.Context, in *MsgEnableTokenizeShares, opts ...grpc.CallOption) (*MsgEnableTokenizeSharesResponse, error) {
	out := new(MsgEnableTokenizeSharesResponse)
	err := c.cc.Invoke(ctx, Msg_EnableTokenizeShares_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ValidatorBond(ctx context.Context, in *MsgValidatorBond, opts ...grpc.CallOption) (*MsgValidatorBondResponse, error) {
	out := new(MsgValidatorBondResponse)
	err := c.cc.Invoke(ctx, Msg_ValidatorBond_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// CreateValidator defines a method for creating a new validator.
	CreateValidator(context.Context, *MsgCreateValidator) (*MsgCreateValidatorResponse, error)
	// EditValidator defines a method for editing an existing validator.
	EditValidator(context.Context, *MsgEditValidator) (*MsgEditValidatorResponse, error)
	// Delegate defines a method for performing a delegation of coins
	// from a delegator to a validator.
	Delegate(context.Context, *MsgDelegate) (*MsgDelegateResponse, error)
	// BeginRedelegate defines a method for performing a redelegation
	// of coins from a delegator and source validator to a destination validator.
	BeginRedelegate(context.Context, *MsgBeginRedelegate) (*MsgBeginRedelegateResponse, error)
	// Undelegate defines a method for performing an undelegation from a
	// delegate and a validator.
	Undelegate(context.Context, *MsgUndelegate) (*MsgUndelegateResponse, error)
	// CancelUnbondingDelegation defines a method for performing canceling the unbonding delegation
	// and delegate back to previous validator.
	//
	// Since: cosmos-sdk 0.46
	CancelUnbondingDelegation(context.Context, *MsgCancelUnbondingDelegation) (*MsgCancelUnbondingDelegationResponse, error)
	// UpdateParams defines an operation for updating the x/staking module
	// parameters.
	// Since: cosmos-sdk 0.47
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	// UnbondValidator defines a method for performing the status transition for a validator
	// from bonded to unbonding
	UnbondValidator(context.Context, *MsgUnbondValidator) (*MsgUnbondValidatorResponse, error)
	// TokenizeShares defines a method for tokenizing shares from a validator.
	TokenizeShares(context.Context, *MsgTokenizeShares) (*MsgTokenizeSharesResponse, error)
	// RedeemTokensForShares defines a method for redeeming tokens from a validator for
	// shares.
	RedeemTokensForShares(context.Context, *MsgRedeemTokensForShares) (*MsgRedeemTokensForSharesResponse, error)
	// TransferTokenizeShareRecord defines a method to transfer ownership of
	// TokenizeShareRecord
	TransferTokenizeShareRecord(context.Context, *MsgTransferTokenizeShareRecord) (*MsgTransferTokenizeShareRecordResponse, error)
	// DisableTokenizeShares defines a method to prevent the tokenization of an addresses stake
	DisableTokenizeShares(context.Context, *MsgDisableTokenizeShares) (*MsgDisableTokenizeSharesResponse, error)
	// EnableTokenizeShares defines a method to re-enable the tokenization of an addresseses stake
	// after it has been disabled
	EnableTokenizeShares(context.Context, *MsgEnableTokenizeShares) (*MsgEnableTokenizeSharesResponse, error)
	// ValidatorBond defines a method for performing a validator self-bond
	ValidatorBond(context.Context, *MsgValidatorBond) (*MsgValidatorBondResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreateValidator(context.Context, *MsgCreateValidator) (*MsgCreateValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateValidator not implemented")
}
func (UnimplementedMsgServer) EditValidator(context.Context, *MsgEditValidator) (*MsgEditValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditValidator not implemented")
}
func (UnimplementedMsgServer) Delegate(context.Context, *MsgDelegate) (*MsgDelegateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delegate not implemented")
}
func (UnimplementedMsgServer) BeginRedelegate(context.Context, *MsgBeginRedelegate) (*MsgBeginRedelegateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginRedelegate not implemented")
}
func (UnimplementedMsgServer) Undelegate(context.Context, *MsgUndelegate) (*MsgUndelegateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Undelegate not implemented")
}
func (UnimplementedMsgServer) CancelUnbondingDelegation(context.Context, *MsgCancelUnbondingDelegation) (*MsgCancelUnbondingDelegationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelUnbondingDelegation not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) UnbondValidator(context.Context, *MsgUnbondValidator) (*MsgUnbondValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnbondValidator not implemented")
}
func (UnimplementedMsgServer) TokenizeShares(context.Context, *MsgTokenizeShares) (*MsgTokenizeSharesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenizeShares not implemented")
}
func (UnimplementedMsgServer) RedeemTokensForShares(context.Context, *MsgRedeemTokensForShares) (*MsgRedeemTokensForSharesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedeemTokensForShares not implemented")
}
func (UnimplementedMsgServer) TransferTokenizeShareRecord(context.Context, *MsgTransferTokenizeShareRecord) (*MsgTransferTokenizeShareRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferTokenizeShareRecord not implemented")
}
func (UnimplementedMsgServer) DisableTokenizeShares(context.Context, *MsgDisableTokenizeShares) (*MsgDisableTokenizeSharesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableTokenizeShares not implemented")
}
func (UnimplementedMsgServer) EnableTokenizeShares(context.Context, *MsgEnableTokenizeShares) (*MsgEnableTokenizeSharesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableTokenizeShares not implemented")
}
func (UnimplementedMsgServer) ValidatorBond(context.Context, *MsgValidatorBond) (*MsgValidatorBondResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorBond not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_CreateValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateValidator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateValidator(ctx, req.(*MsgCreateValidator))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_EditValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEditValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EditValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_EditValidator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EditValidator(ctx, req.(*MsgEditValidator))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Delegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDelegate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Delegate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Delegate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Delegate(ctx, req.(*MsgDelegate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BeginRedelegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBeginRedelegate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BeginRedelegate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_BeginRedelegate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BeginRedelegate(ctx, req.(*MsgBeginRedelegate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Undelegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUndelegate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Undelegate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Undelegate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Undelegate(ctx, req.(*MsgUndelegate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelUnbondingDelegation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelUnbondingDelegation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelUnbondingDelegation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CancelUnbondingDelegation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelUnbondingDelegation(ctx, req.(*MsgCancelUnbondingDelegation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UnbondValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnbondValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UnbondValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UnbondValidator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UnbondValidator(ctx, req.(*MsgUnbondValidator))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TokenizeShares_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTokenizeShares)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TokenizeShares(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_TokenizeShares_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TokenizeShares(ctx, req.(*MsgTokenizeShares))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RedeemTokensForShares_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRedeemTokensForShares)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RedeemTokensForShares(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RedeemTokensForShares_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RedeemTokensForShares(ctx, req.(*MsgRedeemTokensForShares))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TransferTokenizeShareRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferTokenizeShareRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransferTokenizeShareRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_TransferTokenizeShareRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransferTokenizeShareRecord(ctx, req.(*MsgTransferTokenizeShareRecord))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DisableTokenizeShares_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDisableTokenizeShares)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DisableTokenizeShares(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DisableTokenizeShares_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DisableTokenizeShares(ctx, req.(*MsgDisableTokenizeShares))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_EnableTokenizeShares_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEnableTokenizeShares)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EnableTokenizeShares(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_EnableTokenizeShares_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EnableTokenizeShares(ctx, req.(*MsgEnableTokenizeShares))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ValidatorBond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgValidatorBond)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ValidatorBond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ValidatorBond_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ValidatorBond(ctx, req.(*MsgValidatorBond))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.staking.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateValidator",
			Handler:    _Msg_CreateValidator_Handler,
		},
		{
			MethodName: "EditValidator",
			Handler:    _Msg_EditValidator_Handler,
		},
		{
			MethodName: "Delegate",
			Handler:    _Msg_Delegate_Handler,
		},
		{
			MethodName: "BeginRedelegate",
			Handler:    _Msg_BeginRedelegate_Handler,
		},
		{
			MethodName: "Undelegate",
			Handler:    _Msg_Undelegate_Handler,
		},
		{
			MethodName: "CancelUnbondingDelegation",
			Handler:    _Msg_CancelUnbondingDelegation_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "UnbondValidator",
			Handler:    _Msg_UnbondValidator_Handler,
		},
		{
			MethodName: "TokenizeShares",
			Handler:    _Msg_TokenizeShares_Handler,
		},
		{
			MethodName: "RedeemTokensForShares",
			Handler:    _Msg_RedeemTokensForShares_Handler,
		},
		{
			MethodName: "TransferTokenizeShareRecord",
			Handler:    _Msg_TransferTokenizeShareRecord_Handler,
		},
		{
			MethodName: "DisableTokenizeShares",
			Handler:    _Msg_DisableTokenizeShares_Handler,
		},
		{
			MethodName: "EnableTokenizeShares",
			Handler:    _Msg_EnableTokenizeShares_Handler,
		},
		{
			MethodName: "ValidatorBond",
			Handler:    _Msg_ValidatorBond_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/staking/v1beta1/tx.proto",
}
