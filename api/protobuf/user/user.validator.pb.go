// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/protobuf/user/user.proto

package user

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/mwitkow/go-proto-validators"

import time "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

func (this *User) Validate() error {
	if this.CreatedAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if this.UpdatedAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.UpdatedAt); err != nil {
			return go_proto_validators.FieldError("UpdatedAt", err)
		}
	}
	return nil
}
func (this *ID) Validate() error {
	return nil
}
func (this *Phone) Validate() error {
	return nil
}
func (this *Email) Validate() error {
	return nil
}
func (this *Nickname) Validate() error {
	return nil
}
func (this *Account) Validate() error {
	return nil
}
func (this *IDPassword) Validate() error {
	return nil
}
func (this *IDPhone) Validate() error {
	return nil
}
func (this *IDEmail) Validate() error {
	return nil
}
func (this *IDNickname) Validate() error {
	return nil
}
func (this *IDAvatar) Validate() error {
	return nil
}
func (this *IDStatus) Validate() error {
	return nil
}