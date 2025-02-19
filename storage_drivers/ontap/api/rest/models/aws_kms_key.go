// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AwsKmsKey aws kms key
//
// swagger:model aws_kms_key
type AwsKmsKey struct {

	// Key identifier of the AWS KMS key encryption key.
	// Example: key01
	KeyID *string `json:"key_id,omitempty"`
}

// Validate validates this aws kms key
func (m *AwsKmsKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this aws kms key based on context it is used
func (m *AwsKmsKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AwsKmsKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsKmsKey) UnmarshalBinary(b []byte) error {
	var res AwsKmsKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
