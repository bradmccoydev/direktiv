// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateServiceRequest UpdateServiceRequest UpdateServiceRequest update service request
//
// swagger:model updateServiceRequest
type UpdateServiceRequest struct {

	// cmd
	// Required: true
	Cmd *string `json:"cmd"`

	// image
	// Required: true
	Image *string `json:"image"`

	// minScale
	// Required: true
	MinScale *int32 `json:"minScale"`

	// size
	// Required: true
	Size *int32 `json:"size"`

	// trafficPercent
	// Required: true
	TrafficPercent *int64 `json:"trafficPercent"`
}

// Validate validates this update service request
func (m *UpdateServiceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCmd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinScale(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrafficPercent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateServiceRequest) validateCmd(formats strfmt.Registry) error {

	if err := validate.Required("cmd", "body", m.Cmd); err != nil {
		return err
	}

	return nil
}

func (m *UpdateServiceRequest) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

func (m *UpdateServiceRequest) validateMinScale(formats strfmt.Registry) error {

	if err := validate.Required("minScale", "body", m.MinScale); err != nil {
		return err
	}

	return nil
}

func (m *UpdateServiceRequest) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *UpdateServiceRequest) validateTrafficPercent(formats strfmt.Registry) error {

	if err := validate.Required("trafficPercent", "body", m.TrafficPercent); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update service request based on context it is used
func (m *UpdateServiceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateServiceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateServiceRequest) UnmarshalBinary(b []byte) error {
	var res UpdateServiceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
