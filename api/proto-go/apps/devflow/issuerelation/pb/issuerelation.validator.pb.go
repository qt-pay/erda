// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issuerelation.proto

package pb

import (
	fmt "fmt"
	math "math"

	_ "github.com/erda-project/erda-proto-go/common/pb"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *IssueRelation) Validate() error {
	if this.TimeCreated != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TimeCreated); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TimeCreated", err)
		}
	}
	if this.TimeUpdated != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TimeUpdated); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TimeUpdated", err)
		}
	}
	if this.SoftDeletedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SoftDeletedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SoftDeletedAt", err)
		}
	}
	return nil
}
func (this *CreateIssueRelationRequest) Validate() error {
	return nil
}
func (this *CreateIssueRelationResponse) Validate() error {
	if this.IssueRelation != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.IssueRelation); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("IssueRelation", err)
		}
	}
	return nil
}
func (this *DeleteIssueRelationRequest) Validate() error {
	return nil
}
func (this *DeleteIssueRelationResponse) Validate() error {
	return nil
}
func (this *ListIssueRelationRequest) Validate() error {
	return nil
}
func (this *ListIssueRelationResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}