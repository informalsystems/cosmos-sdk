package feegrantv1beta1

import (
	fmt "fmt"
	io "io"
	reflect "reflect"
	sync "sync"

	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/query/v1beta1"
)

var (
	md_QueryAllowanceRequest         protoreflect.MessageDescriptor
	fd_QueryAllowanceRequest_granter protoreflect.FieldDescriptor
	fd_QueryAllowanceRequest_grantee protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_feegrant_v1beta1_query_proto_init()
	md_QueryAllowanceRequest = File_cosmos_feegrant_v1beta1_query_proto.Messages().ByName("QueryAllowanceRequest")
	fd_QueryAllowanceRequest_granter = md_QueryAllowanceRequest.Fields().ByName("granter")
	fd_QueryAllowanceRequest_grantee = md_QueryAllowanceRequest.Fields().ByName("grantee")
}

var _ protoreflect.Message = (*fastReflection_QueryAllowanceRequest)(nil)

type fastReflection_QueryAllowanceRequest QueryAllowanceRequest

func (x *QueryAllowanceRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryAllowanceRequest)(x)
}

func (x *QueryAllowanceRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_feegrant_v1beta1_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryAllowanceRequest_messageType fastReflection_QueryAllowanceRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryAllowanceRequest_messageType{}

type fastReflection_QueryAllowanceRequest_messageType struct{}

func (x fastReflection_QueryAllowanceRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryAllowanceRequest)(nil)
}
func (x fastReflection_QueryAllowanceRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryAllowanceRequest)
}
func (x fastReflection_QueryAllowanceRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryAllowanceRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryAllowanceRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryAllowanceRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryAllowanceRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryAllowanceRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryAllowanceRequest) New() protoreflect.Message {
	return new(fastReflection_QueryAllowanceRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryAllowanceRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryAllowanceRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryAllowanceRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Granter != "" {
		value := protoreflect.ValueOfString(x.Granter)
		if !f(fd_QueryAllowanceRequest_granter, value) {
			return
		}
	}
	if x.Grantee != "" {
		value := protoreflect.ValueOfString(x.Grantee)
		if !f(fd_QueryAllowanceRequest_grantee, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryAllowanceRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowanceRequest.granter":
		return x.Granter != ""
	case "cosmos.feegrant.v1beta1.QueryAllowanceRequest.grantee":
		return x.Grantee != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowanceRequest"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowanceRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowanceRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowanceRequest.granter":
		x.Granter = ""
	case "cosmos.feegrant.v1beta1.QueryAllowanceRequest.grantee":
		x.Grantee = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowanceRequest"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowanceRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryAllowanceRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowanceRequest.granter":
		value := x.Granter
		return protoreflect.ValueOfString(value)
	case "cosmos.feegrant.v1beta1.QueryAllowanceRequest.grantee":
		value := x.Grantee
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowanceRequest"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowanceRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowanceRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowanceRequest.granter":
		x.Granter = value.Interface().(string)
	case "cosmos.feegrant.v1beta1.QueryAllowanceRequest.grantee":
		x.Grantee = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowanceRequest"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowanceRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowanceRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowanceRequest.granter":
		panic(fmt.Errorf("field granter of message cosmos.feegrant.v1beta1.QueryAllowanceRequest is not mutable"))
	case "cosmos.feegrant.v1beta1.QueryAllowanceRequest.grantee":
		panic(fmt.Errorf("field grantee of message cosmos.feegrant.v1beta1.QueryAllowanceRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowanceRequest"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowanceRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryAllowanceRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowanceRequest.granter":
		return protoreflect.ValueOfString("")
	case "cosmos.feegrant.v1beta1.QueryAllowanceRequest.grantee":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowanceRequest"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowanceRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryAllowanceRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.feegrant.v1beta1.QueryAllowanceRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryAllowanceRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowanceRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryAllowanceRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryAllowanceRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryAllowanceRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Granter)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Grantee)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryAllowanceRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Grantee) > 0 {
			i -= len(x.Grantee)
			copy(dAtA[i:], x.Grantee)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Grantee)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Granter) > 0 {
			i -= len(x.Granter)
			copy(dAtA[i:], x.Granter)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Granter)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryAllowanceRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryAllowanceRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryAllowanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Granter = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Grantee = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_QueryAllowanceResponse           protoreflect.MessageDescriptor
	fd_QueryAllowanceResponse_allowance protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_feegrant_v1beta1_query_proto_init()
	md_QueryAllowanceResponse = File_cosmos_feegrant_v1beta1_query_proto.Messages().ByName("QueryAllowanceResponse")
	fd_QueryAllowanceResponse_allowance = md_QueryAllowanceResponse.Fields().ByName("allowance")
}

var _ protoreflect.Message = (*fastReflection_QueryAllowanceResponse)(nil)

type fastReflection_QueryAllowanceResponse QueryAllowanceResponse

func (x *QueryAllowanceResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryAllowanceResponse)(x)
}

func (x *QueryAllowanceResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_feegrant_v1beta1_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryAllowanceResponse_messageType fastReflection_QueryAllowanceResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryAllowanceResponse_messageType{}

type fastReflection_QueryAllowanceResponse_messageType struct{}

func (x fastReflection_QueryAllowanceResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryAllowanceResponse)(nil)
}
func (x fastReflection_QueryAllowanceResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryAllowanceResponse)
}
func (x fastReflection_QueryAllowanceResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryAllowanceResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryAllowanceResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryAllowanceResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryAllowanceResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryAllowanceResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryAllowanceResponse) New() protoreflect.Message {
	return new(fastReflection_QueryAllowanceResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryAllowanceResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryAllowanceResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryAllowanceResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Allowance != nil {
		value := protoreflect.ValueOfMessage(x.Allowance.ProtoReflect())
		if !f(fd_QueryAllowanceResponse_allowance, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryAllowanceResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowanceResponse.allowance":
		return x.Allowance != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowanceResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowanceResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowanceResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowanceResponse.allowance":
		x.Allowance = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowanceResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowanceResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryAllowanceResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowanceResponse.allowance":
		value := x.Allowance
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowanceResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowanceResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowanceResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowanceResponse.allowance":
		x.Allowance = value.Message().Interface().(*Grant)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowanceResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowanceResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowanceResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowanceResponse.allowance":
		if x.Allowance == nil {
			x.Allowance = new(Grant)
		}
		return protoreflect.ValueOfMessage(x.Allowance.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowanceResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowanceResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryAllowanceResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowanceResponse.allowance":
		m := new(Grant)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowanceResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowanceResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryAllowanceResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.feegrant.v1beta1.QueryAllowanceResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryAllowanceResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowanceResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryAllowanceResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryAllowanceResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryAllowanceResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Allowance != nil {
			l = options.Size(x.Allowance)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryAllowanceResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Allowance != nil {
			encoded, err := options.Marshal(x.Allowance)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryAllowanceResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryAllowanceResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryAllowanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Allowance == nil {
					x.Allowance = &Grant{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Allowance); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_QueryAllowancesRequest            protoreflect.MessageDescriptor
	fd_QueryAllowancesRequest_grantee    protoreflect.FieldDescriptor
	fd_QueryAllowancesRequest_pagination protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_feegrant_v1beta1_query_proto_init()
	md_QueryAllowancesRequest = File_cosmos_feegrant_v1beta1_query_proto.Messages().ByName("QueryAllowancesRequest")
	fd_QueryAllowancesRequest_grantee = md_QueryAllowancesRequest.Fields().ByName("grantee")
	fd_QueryAllowancesRequest_pagination = md_QueryAllowancesRequest.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryAllowancesRequest)(nil)

type fastReflection_QueryAllowancesRequest QueryAllowancesRequest

func (x *QueryAllowancesRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryAllowancesRequest)(x)
}

func (x *QueryAllowancesRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_feegrant_v1beta1_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryAllowancesRequest_messageType fastReflection_QueryAllowancesRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryAllowancesRequest_messageType{}

type fastReflection_QueryAllowancesRequest_messageType struct{}

func (x fastReflection_QueryAllowancesRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryAllowancesRequest)(nil)
}
func (x fastReflection_QueryAllowancesRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryAllowancesRequest)
}
func (x fastReflection_QueryAllowancesRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryAllowancesRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryAllowancesRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryAllowancesRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryAllowancesRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryAllowancesRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryAllowancesRequest) New() protoreflect.Message {
	return new(fastReflection_QueryAllowancesRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryAllowancesRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryAllowancesRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryAllowancesRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Grantee != "" {
		value := protoreflect.ValueOfString(x.Grantee)
		if !f(fd_QueryAllowancesRequest_grantee, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryAllowancesRequest_pagination, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryAllowancesRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesRequest.grantee":
		return x.Grantee != ""
	case "cosmos.feegrant.v1beta1.QueryAllowancesRequest.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesRequest"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowancesRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesRequest.grantee":
		x.Grantee = ""
	case "cosmos.feegrant.v1beta1.QueryAllowancesRequest.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesRequest"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryAllowancesRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesRequest.grantee":
		value := x.Grantee
		return protoreflect.ValueOfString(value)
	case "cosmos.feegrant.v1beta1.QueryAllowancesRequest.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesRequest"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowancesRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesRequest.grantee":
		x.Grantee = value.Interface().(string)
	case "cosmos.feegrant.v1beta1.QueryAllowancesRequest.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageRequest)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesRequest"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowancesRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesRequest.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageRequest)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	case "cosmos.feegrant.v1beta1.QueryAllowancesRequest.grantee":
		panic(fmt.Errorf("field grantee of message cosmos.feegrant.v1beta1.QueryAllowancesRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesRequest"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryAllowancesRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesRequest.grantee":
		return protoreflect.ValueOfString("")
	case "cosmos.feegrant.v1beta1.QueryAllowancesRequest.pagination":
		m := new(v1beta1.PageRequest)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesRequest"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryAllowancesRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.feegrant.v1beta1.QueryAllowancesRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryAllowancesRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowancesRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryAllowancesRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryAllowancesRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryAllowancesRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Grantee)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryAllowancesRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Grantee) > 0 {
			i -= len(x.Grantee)
			copy(dAtA[i:], x.Grantee)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Grantee)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryAllowancesRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryAllowancesRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryAllowancesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Grantee = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageRequest{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_QueryAllowancesResponse_1_list)(nil)

type _QueryAllowancesResponse_1_list struct {
	list *[]*Grant
}

func (x *_QueryAllowancesResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryAllowancesResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryAllowancesResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Grant)
	(*x.list)[i] = concreteValue
}

func (x *_QueryAllowancesResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Grant)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryAllowancesResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(Grant)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryAllowancesResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryAllowancesResponse_1_list) NewElement() protoreflect.Value {
	v := new(Grant)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryAllowancesResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryAllowancesResponse            protoreflect.MessageDescriptor
	fd_QueryAllowancesResponse_allowances protoreflect.FieldDescriptor
	fd_QueryAllowancesResponse_pagination protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_feegrant_v1beta1_query_proto_init()
	md_QueryAllowancesResponse = File_cosmos_feegrant_v1beta1_query_proto.Messages().ByName("QueryAllowancesResponse")
	fd_QueryAllowancesResponse_allowances = md_QueryAllowancesResponse.Fields().ByName("allowances")
	fd_QueryAllowancesResponse_pagination = md_QueryAllowancesResponse.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryAllowancesResponse)(nil)

type fastReflection_QueryAllowancesResponse QueryAllowancesResponse

func (x *QueryAllowancesResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryAllowancesResponse)(x)
}

func (x *QueryAllowancesResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_feegrant_v1beta1_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryAllowancesResponse_messageType fastReflection_QueryAllowancesResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryAllowancesResponse_messageType{}

type fastReflection_QueryAllowancesResponse_messageType struct{}

func (x fastReflection_QueryAllowancesResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryAllowancesResponse)(nil)
}
func (x fastReflection_QueryAllowancesResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryAllowancesResponse)
}
func (x fastReflection_QueryAllowancesResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryAllowancesResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryAllowancesResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryAllowancesResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryAllowancesResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryAllowancesResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryAllowancesResponse) New() protoreflect.Message {
	return new(fastReflection_QueryAllowancesResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryAllowancesResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryAllowancesResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryAllowancesResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Allowances) != 0 {
		value := protoreflect.ValueOfList(&_QueryAllowancesResponse_1_list{list: &x.Allowances})
		if !f(fd_QueryAllowancesResponse_allowances, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryAllowancesResponse_pagination, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryAllowancesResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesResponse.allowances":
		return len(x.Allowances) != 0
	case "cosmos.feegrant.v1beta1.QueryAllowancesResponse.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowancesResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesResponse.allowances":
		x.Allowances = nil
	case "cosmos.feegrant.v1beta1.QueryAllowancesResponse.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryAllowancesResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesResponse.allowances":
		if len(x.Allowances) == 0 {
			return protoreflect.ValueOfList(&_QueryAllowancesResponse_1_list{})
		}
		listValue := &_QueryAllowancesResponse_1_list{list: &x.Allowances}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.feegrant.v1beta1.QueryAllowancesResponse.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowancesResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesResponse.allowances":
		lv := value.List()
		clv := lv.(*_QueryAllowancesResponse_1_list)
		x.Allowances = *clv.list
	case "cosmos.feegrant.v1beta1.QueryAllowancesResponse.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowancesResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesResponse.allowances":
		if x.Allowances == nil {
			x.Allowances = []*Grant{}
		}
		value := &_QueryAllowancesResponse_1_list{list: &x.Allowances}
		return protoreflect.ValueOfList(value)
	case "cosmos.feegrant.v1beta1.QueryAllowancesResponse.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageResponse)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryAllowancesResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesResponse.allowances":
		list := []*Grant{}
		return protoreflect.ValueOfList(&_QueryAllowancesResponse_1_list{list: &list})
	case "cosmos.feegrant.v1beta1.QueryAllowancesResponse.pagination":
		m := new(v1beta1.PageResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryAllowancesResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.feegrant.v1beta1.QueryAllowancesResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryAllowancesResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowancesResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryAllowancesResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryAllowancesResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryAllowancesResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Allowances) > 0 {
			for _, e := range x.Allowances {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryAllowancesResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Allowances) > 0 {
			for iNdEx := len(x.Allowances) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Allowances[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryAllowancesResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryAllowancesResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryAllowancesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Allowances", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Allowances = append(x.Allowances, &Grant{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Allowances[len(x.Allowances)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageResponse{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_QueryAllowancesByGranterRequest            protoreflect.MessageDescriptor
	fd_QueryAllowancesByGranterRequest_granter    protoreflect.FieldDescriptor
	fd_QueryAllowancesByGranterRequest_pagination protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_feegrant_v1beta1_query_proto_init()
	md_QueryAllowancesByGranterRequest = File_cosmos_feegrant_v1beta1_query_proto.Messages().ByName("QueryAllowancesByGranterRequest")
	fd_QueryAllowancesByGranterRequest_granter = md_QueryAllowancesByGranterRequest.Fields().ByName("granter")
	fd_QueryAllowancesByGranterRequest_pagination = md_QueryAllowancesByGranterRequest.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryAllowancesByGranterRequest)(nil)

type fastReflection_QueryAllowancesByGranterRequest QueryAllowancesByGranterRequest

func (x *QueryAllowancesByGranterRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryAllowancesByGranterRequest)(x)
}

func (x *QueryAllowancesByGranterRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_feegrant_v1beta1_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryAllowancesByGranterRequest_messageType fastReflection_QueryAllowancesByGranterRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryAllowancesByGranterRequest_messageType{}

type fastReflection_QueryAllowancesByGranterRequest_messageType struct{}

func (x fastReflection_QueryAllowancesByGranterRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryAllowancesByGranterRequest)(nil)
}
func (x fastReflection_QueryAllowancesByGranterRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryAllowancesByGranterRequest)
}
func (x fastReflection_QueryAllowancesByGranterRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryAllowancesByGranterRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryAllowancesByGranterRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryAllowancesByGranterRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryAllowancesByGranterRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryAllowancesByGranterRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryAllowancesByGranterRequest) New() protoreflect.Message {
	return new(fastReflection_QueryAllowancesByGranterRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryAllowancesByGranterRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryAllowancesByGranterRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryAllowancesByGranterRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Granter != "" {
		value := protoreflect.ValueOfString(x.Granter)
		if !f(fd_QueryAllowancesByGranterRequest_granter, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryAllowancesByGranterRequest_pagination, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryAllowancesByGranterRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest.granter":
		return x.Granter != ""
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowancesByGranterRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest.granter":
		x.Granter = ""
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryAllowancesByGranterRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest.granter":
		value := x.Granter
		return protoreflect.ValueOfString(value)
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowancesByGranterRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest.granter":
		x.Granter = value.Interface().(string)
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageRequest)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowancesByGranterRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageRequest)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest.granter":
		panic(fmt.Errorf("field granter of message cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryAllowancesByGranterRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest.granter":
		return protoreflect.ValueOfString("")
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest.pagination":
		m := new(v1beta1.PageRequest)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryAllowancesByGranterRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryAllowancesByGranterRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowancesByGranterRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryAllowancesByGranterRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryAllowancesByGranterRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryAllowancesByGranterRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Granter)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryAllowancesByGranterRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Granter) > 0 {
			i -= len(x.Granter)
			copy(dAtA[i:], x.Granter)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Granter)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryAllowancesByGranterRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryAllowancesByGranterRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryAllowancesByGranterRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Granter = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageRequest{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_QueryAllowancesByGranterResponse_1_list)(nil)

type _QueryAllowancesByGranterResponse_1_list struct {
	list *[]*Grant
}

func (x *_QueryAllowancesByGranterResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryAllowancesByGranterResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryAllowancesByGranterResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Grant)
	(*x.list)[i] = concreteValue
}

func (x *_QueryAllowancesByGranterResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Grant)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryAllowancesByGranterResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(Grant)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryAllowancesByGranterResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryAllowancesByGranterResponse_1_list) NewElement() protoreflect.Value {
	v := new(Grant)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryAllowancesByGranterResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryAllowancesByGranterResponse            protoreflect.MessageDescriptor
	fd_QueryAllowancesByGranterResponse_allowances protoreflect.FieldDescriptor
	fd_QueryAllowancesByGranterResponse_pagination protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_feegrant_v1beta1_query_proto_init()
	md_QueryAllowancesByGranterResponse = File_cosmos_feegrant_v1beta1_query_proto.Messages().ByName("QueryAllowancesByGranterResponse")
	fd_QueryAllowancesByGranterResponse_allowances = md_QueryAllowancesByGranterResponse.Fields().ByName("allowances")
	fd_QueryAllowancesByGranterResponse_pagination = md_QueryAllowancesByGranterResponse.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryAllowancesByGranterResponse)(nil)

type fastReflection_QueryAllowancesByGranterResponse QueryAllowancesByGranterResponse

func (x *QueryAllowancesByGranterResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryAllowancesByGranterResponse)(x)
}

func (x *QueryAllowancesByGranterResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_feegrant_v1beta1_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryAllowancesByGranterResponse_messageType fastReflection_QueryAllowancesByGranterResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryAllowancesByGranterResponse_messageType{}

type fastReflection_QueryAllowancesByGranterResponse_messageType struct{}

func (x fastReflection_QueryAllowancesByGranterResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryAllowancesByGranterResponse)(nil)
}
func (x fastReflection_QueryAllowancesByGranterResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryAllowancesByGranterResponse)
}
func (x fastReflection_QueryAllowancesByGranterResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryAllowancesByGranterResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryAllowancesByGranterResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryAllowancesByGranterResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryAllowancesByGranterResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryAllowancesByGranterResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryAllowancesByGranterResponse) New() protoreflect.Message {
	return new(fastReflection_QueryAllowancesByGranterResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryAllowancesByGranterResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryAllowancesByGranterResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryAllowancesByGranterResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Allowances) != 0 {
		value := protoreflect.ValueOfList(&_QueryAllowancesByGranterResponse_1_list{list: &x.Allowances})
		if !f(fd_QueryAllowancesByGranterResponse_allowances, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryAllowancesByGranterResponse_pagination, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_QueryAllowancesByGranterResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse.allowances":
		return len(x.Allowances) != 0
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowancesByGranterResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse.allowances":
		x.Allowances = nil
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryAllowancesByGranterResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse.allowances":
		if len(x.Allowances) == 0 {
			return protoreflect.ValueOfList(&_QueryAllowancesByGranterResponse_1_list{})
		}
		listValue := &_QueryAllowancesByGranterResponse_1_list{list: &x.Allowances}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowancesByGranterResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse.allowances":
		lv := value.List()
		clv := lv.(*_QueryAllowancesByGranterResponse_1_list)
		x.Allowances = *clv.list
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowancesByGranterResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse.allowances":
		if x.Allowances == nil {
			x.Allowances = []*Grant{}
		}
		value := &_QueryAllowancesByGranterResponse_1_list{list: &x.Allowances}
		return protoreflect.ValueOfList(value)
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageResponse)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryAllowancesByGranterResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse.allowances":
		list := []*Grant{}
		return protoreflect.ValueOfList(&_QueryAllowancesByGranterResponse_1_list{list: &list})
	case "cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse.pagination":
		m := new(v1beta1.PageResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryAllowancesByGranterResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryAllowancesByGranterResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryAllowancesByGranterResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_QueryAllowancesByGranterResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryAllowancesByGranterResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryAllowancesByGranterResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Allowances) > 0 {
			for _, e := range x.Allowances {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryAllowancesByGranterResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Allowances) > 0 {
			for iNdEx := len(x.Allowances) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Allowances[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*QueryAllowancesByGranterResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryAllowancesByGranterResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryAllowancesByGranterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Allowances", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Allowances = append(x.Allowances, &Grant{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Allowances[len(x.Allowances)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageResponse{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Since: cosmos-sdk 0.43

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/feegrant/v1beta1/query.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// QueryAllowanceRequest is the request type for the Query/Allowance RPC method.
type QueryAllowanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// granter is the address of the user granting an allowance of their funds.
	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	// grantee is the address of the user being granted an allowance of another user's funds.
	Grantee string `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
}

func (x *QueryAllowanceRequest) Reset() {
	*x = QueryAllowanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_feegrant_v1beta1_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAllowanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAllowanceRequest) ProtoMessage() {}

// Deprecated: Use QueryAllowanceRequest.ProtoReflect.Descriptor instead.
func (*QueryAllowanceRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_feegrant_v1beta1_query_proto_rawDescGZIP(), []int{0}
}

func (x *QueryAllowanceRequest) GetGranter() string {
	if x != nil {
		return x.Granter
	}
	return ""
}

func (x *QueryAllowanceRequest) GetGrantee() string {
	if x != nil {
		return x.Grantee
	}
	return ""
}

// QueryAllowanceResponse is the response type for the Query/Allowance RPC method.
type QueryAllowanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// allowance is a allowance granted for grantee by granter.
	Allowance *Grant `protobuf:"bytes,1,opt,name=allowance,proto3" json:"allowance,omitempty"`
}

func (x *QueryAllowanceResponse) Reset() {
	*x = QueryAllowanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_feegrant_v1beta1_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAllowanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAllowanceResponse) ProtoMessage() {}

// Deprecated: Use QueryAllowanceResponse.ProtoReflect.Descriptor instead.
func (*QueryAllowanceResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_feegrant_v1beta1_query_proto_rawDescGZIP(), []int{1}
}

func (x *QueryAllowanceResponse) GetAllowance() *Grant {
	if x != nil {
		return x.Allowance
	}
	return nil
}

// QueryAllowancesRequest is the request type for the Query/Allowances RPC method.
type QueryAllowancesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grantee string `protobuf:"bytes,1,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// pagination defines an pagination for the request.
	Pagination *v1beta1.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryAllowancesRequest) Reset() {
	*x = QueryAllowancesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_feegrant_v1beta1_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAllowancesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAllowancesRequest) ProtoMessage() {}

// Deprecated: Use QueryAllowancesRequest.ProtoReflect.Descriptor instead.
func (*QueryAllowancesRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_feegrant_v1beta1_query_proto_rawDescGZIP(), []int{2}
}

func (x *QueryAllowancesRequest) GetGrantee() string {
	if x != nil {
		return x.Grantee
	}
	return ""
}

func (x *QueryAllowancesRequest) GetPagination() *v1beta1.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryAllowancesResponse is the response type for the Query/Allowances RPC method.
type QueryAllowancesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// allowances are allowance's granted for grantee by granter.
	Allowances []*Grant `protobuf:"bytes,1,rep,name=allowances,proto3" json:"allowances,omitempty"`
	// pagination defines an pagination for the response.
	Pagination *v1beta1.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryAllowancesResponse) Reset() {
	*x = QueryAllowancesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_feegrant_v1beta1_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAllowancesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAllowancesResponse) ProtoMessage() {}

// Deprecated: Use QueryAllowancesResponse.ProtoReflect.Descriptor instead.
func (*QueryAllowancesResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_feegrant_v1beta1_query_proto_rawDescGZIP(), []int{3}
}

func (x *QueryAllowancesResponse) GetAllowances() []*Grant {
	if x != nil {
		return x.Allowances
	}
	return nil
}

func (x *QueryAllowancesResponse) GetPagination() *v1beta1.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryAllowancesByGranterRequest is the request type for the Query/AllowancesByGranter RPC method.
type QueryAllowancesByGranterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	// pagination defines an pagination for the request.
	Pagination *v1beta1.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryAllowancesByGranterRequest) Reset() {
	*x = QueryAllowancesByGranterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_feegrant_v1beta1_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAllowancesByGranterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAllowancesByGranterRequest) ProtoMessage() {}

// Deprecated: Use QueryAllowancesByGranterRequest.ProtoReflect.Descriptor instead.
func (*QueryAllowancesByGranterRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_feegrant_v1beta1_query_proto_rawDescGZIP(), []int{4}
}

func (x *QueryAllowancesByGranterRequest) GetGranter() string {
	if x != nil {
		return x.Granter
	}
	return ""
}

func (x *QueryAllowancesByGranterRequest) GetPagination() *v1beta1.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryAllowancesByGranterResponse is the response type for the Query/AllowancesByGranter RPC method.
type QueryAllowancesByGranterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// allowances that have been issued by the granter.
	Allowances []*Grant `protobuf:"bytes,1,rep,name=allowances,proto3" json:"allowances,omitempty"`
	// pagination defines an pagination for the response.
	Pagination *v1beta1.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryAllowancesByGranterResponse) Reset() {
	*x = QueryAllowancesByGranterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_feegrant_v1beta1_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAllowancesByGranterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAllowancesByGranterResponse) ProtoMessage() {}

// Deprecated: Use QueryAllowancesByGranterResponse.ProtoReflect.Descriptor instead.
func (*QueryAllowancesByGranterResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_feegrant_v1beta1_query_proto_rawDescGZIP(), []int{5}
}

func (x *QueryAllowancesByGranterResponse) GetAllowances() []*Grant {
	if x != nil {
		return x.Allowances
	}
	return nil
}

func (x *QueryAllowancesByGranterResponse) GetPagination() *v1beta1.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_cosmos_feegrant_v1beta1_query_proto protoreflect.FileDescriptor

var file_cosmos_feegrant_v1beta1_query_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x66, 0x65,
	0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x26,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x15, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x22, 0x56, 0x0a, 0x16,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x61, 0x6e, 0x63, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x32, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x65, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa2, 0x01, 0x0a, 0x17,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x9d, 0x01, 0x0a, 0x1f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x42, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xab, 0x01, 0x0a, 0x20, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x42, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x9f,
	0x04, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0xac, 0x01, 0x0a, 0x09, 0x41, 0x6c, 0x6c,
	0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x38, 0x12,
	0x36, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x61,
	0x6e, 0x63, 0x65, 0x2f, 0x7b, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x7d, 0x2f, 0x7b, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x7d, 0x12, 0xa6, 0x01, 0x0a, 0x0a, 0x41, 0x6c, 0x6c, 0x6f,
	0x77, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x2f, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2f, 0x12, 0x2d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x67, 0x72,
	0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x7d,
	0x12, 0xbd, 0x01, 0x0a, 0x13, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x42,
	0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x38, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x42, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x39, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x66, 0x65, 0x65, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x42, 0x79, 0x47, 0x72,
	0x61, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x12, 0x29, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x66,
	0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x2f, 0x7b, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x7d,
	0x42, 0xf1, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x42, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x46, 0x58, 0xaa, 0x02,
	0x17, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x46, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74,
	0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x17, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x5c, 0x46, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xe2, 0x02, 0x23, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x46, 0x65, 0x65, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x3a, 0x3a, 0x46, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_feegrant_v1beta1_query_proto_rawDescOnce sync.Once
	file_cosmos_feegrant_v1beta1_query_proto_rawDescData = file_cosmos_feegrant_v1beta1_query_proto_rawDesc
)

func file_cosmos_feegrant_v1beta1_query_proto_rawDescGZIP() []byte {
	file_cosmos_feegrant_v1beta1_query_proto_rawDescOnce.Do(func() {
		file_cosmos_feegrant_v1beta1_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_feegrant_v1beta1_query_proto_rawDescData)
	})
	return file_cosmos_feegrant_v1beta1_query_proto_rawDescData
}

var file_cosmos_feegrant_v1beta1_query_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cosmos_feegrant_v1beta1_query_proto_goTypes = []interface{}{
	(*QueryAllowanceRequest)(nil),            // 0: cosmos.feegrant.v1beta1.QueryAllowanceRequest
	(*QueryAllowanceResponse)(nil),           // 1: cosmos.feegrant.v1beta1.QueryAllowanceResponse
	(*QueryAllowancesRequest)(nil),           // 2: cosmos.feegrant.v1beta1.QueryAllowancesRequest
	(*QueryAllowancesResponse)(nil),          // 3: cosmos.feegrant.v1beta1.QueryAllowancesResponse
	(*QueryAllowancesByGranterRequest)(nil),  // 4: cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest
	(*QueryAllowancesByGranterResponse)(nil), // 5: cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse
	(*Grant)(nil),                            // 6: cosmos.feegrant.v1beta1.Grant
	(*v1beta1.PageRequest)(nil),              // 7: cosmos.base.query.v1beta1.PageRequest
	(*v1beta1.PageResponse)(nil),             // 8: cosmos.base.query.v1beta1.PageResponse
}
var file_cosmos_feegrant_v1beta1_query_proto_depIdxs = []int32{
	6,  // 0: cosmos.feegrant.v1beta1.QueryAllowanceResponse.allowance:type_name -> cosmos.feegrant.v1beta1.Grant
	7,  // 1: cosmos.feegrant.v1beta1.QueryAllowancesRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	6,  // 2: cosmos.feegrant.v1beta1.QueryAllowancesResponse.allowances:type_name -> cosmos.feegrant.v1beta1.Grant
	8,  // 3: cosmos.feegrant.v1beta1.QueryAllowancesResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	7,  // 4: cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	6,  // 5: cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse.allowances:type_name -> cosmos.feegrant.v1beta1.Grant
	8,  // 6: cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	0,  // 7: cosmos.feegrant.v1beta1.Query.Allowance:input_type -> cosmos.feegrant.v1beta1.QueryAllowanceRequest
	2,  // 8: cosmos.feegrant.v1beta1.Query.Allowances:input_type -> cosmos.feegrant.v1beta1.QueryAllowancesRequest
	4,  // 9: cosmos.feegrant.v1beta1.Query.AllowancesByGranter:input_type -> cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest
	1,  // 10: cosmos.feegrant.v1beta1.Query.Allowance:output_type -> cosmos.feegrant.v1beta1.QueryAllowanceResponse
	3,  // 11: cosmos.feegrant.v1beta1.Query.Allowances:output_type -> cosmos.feegrant.v1beta1.QueryAllowancesResponse
	5,  // 12: cosmos.feegrant.v1beta1.Query.AllowancesByGranter:output_type -> cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_cosmos_feegrant_v1beta1_query_proto_init() }
func file_cosmos_feegrant_v1beta1_query_proto_init() {
	if File_cosmos_feegrant_v1beta1_query_proto != nil {
		return
	}
	file_cosmos_feegrant_v1beta1_feegrant_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cosmos_feegrant_v1beta1_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAllowanceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_feegrant_v1beta1_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAllowanceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_feegrant_v1beta1_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAllowancesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_feegrant_v1beta1_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAllowancesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_feegrant_v1beta1_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAllowancesByGranterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_feegrant_v1beta1_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAllowancesByGranterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_feegrant_v1beta1_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_feegrant_v1beta1_query_proto_goTypes,
		DependencyIndexes: file_cosmos_feegrant_v1beta1_query_proto_depIdxs,
		MessageInfos:      file_cosmos_feegrant_v1beta1_query_proto_msgTypes,
	}.Build()
	File_cosmos_feegrant_v1beta1_query_proto = out.File
	file_cosmos_feegrant_v1beta1_query_proto_rawDesc = nil
	file_cosmos_feegrant_v1beta1_query_proto_goTypes = nil
	file_cosmos_feegrant_v1beta1_query_proto_depIdxs = nil
}
