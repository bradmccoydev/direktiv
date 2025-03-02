// Code generated by go-swagger; DO NOT EDIT.

package variables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetInstanceVariableReader is a Reader for the GetInstanceVariable structure.
type GetInstanceVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstanceVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstanceVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInstanceVariableOK creates a GetInstanceVariableOK with default headers values
func NewGetInstanceVariableOK() *GetInstanceVariableOK {
	return &GetInstanceVariableOK{}
}

/* GetInstanceVariableOK describes a response with status code 200, with default header values.

successfully got instance variable
*/
type GetInstanceVariableOK struct {
}

func (o *GetInstanceVariableOK) Error() string {
	return fmt.Sprintf("[GET /api/namespaces/{namespace}/instances/{instance}/vars/{variable}][%d] getInstanceVariableOK ", 200)
}

func (o *GetInstanceVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
