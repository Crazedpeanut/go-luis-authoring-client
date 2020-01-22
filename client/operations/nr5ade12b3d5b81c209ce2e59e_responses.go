// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// Nr5ade12b3d5b81c209ce2e59eReader is a Reader for the Nr5ade12b3d5b81c209ce2e59e structure.
type Nr5ade12b3d5b81c209ce2e59eReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr5ade12b3d5b81c209ce2e59eReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNr5ade12b3d5b81c209ce2e59eOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr5ade12b3d5b81c209ce2e59eBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr5ade12b3d5b81c209ce2e59eUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNr5ade12b3d5b81c209ce2e59eForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNr5ade12b3d5b81c209ce2e59eTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr5ade12b3d5b81c209ce2e59eOK creates a Nr5ade12b3d5b81c209ce2e59eOK with default headers values
func NewNr5ade12b3d5b81c209ce2e59eOK() *Nr5ade12b3d5b81c209ce2e59eOK {
	return &Nr5ade12b3d5b81c209ce2e59eOK{}
}

/*Nr5ade12b3d5b81c209ce2e59eOK handles this case with default header values.

A list of the roles.
*/
type Nr5ade12b3d5b81c209ce2e59eOK struct {
}

func (o *Nr5ade12b3d5b81c209ce2e59eOK) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/roles][%d] 5ade12b3d5b81c209ce2e59eOK ", 200)
}

func (o *Nr5ade12b3d5b81c209ce2e59eOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade12b3d5b81c209ce2e59eBadRequest creates a Nr5ade12b3d5b81c209ce2e59eBadRequest with default headers values
func NewNr5ade12b3d5b81c209ce2e59eBadRequest() *Nr5ade12b3d5b81c209ce2e59eBadRequest {
	return &Nr5ade12b3d5b81c209ce2e59eBadRequest{}
}

/*Nr5ade12b3d5b81c209ce2e59eBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type Nr5ade12b3d5b81c209ce2e59eBadRequest struct {
}

func (o *Nr5ade12b3d5b81c209ce2e59eBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/roles][%d] 5ade12b3d5b81c209ce2e59eBadRequest ", 400)
}

func (o *Nr5ade12b3d5b81c209ce2e59eBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade12b3d5b81c209ce2e59eUnauthorized creates a Nr5ade12b3d5b81c209ce2e59eUnauthorized with default headers values
func NewNr5ade12b3d5b81c209ce2e59eUnauthorized() *Nr5ade12b3d5b81c209ce2e59eUnauthorized {
	return &Nr5ade12b3d5b81c209ce2e59eUnauthorized{}
}

/*Nr5ade12b3d5b81c209ce2e59eUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type Nr5ade12b3d5b81c209ce2e59eUnauthorized struct {
}

func (o *Nr5ade12b3d5b81c209ce2e59eUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/roles][%d] 5ade12b3d5b81c209ce2e59eUnauthorized ", 401)
}

func (o *Nr5ade12b3d5b81c209ce2e59eUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade12b3d5b81c209ce2e59eForbidden creates a Nr5ade12b3d5b81c209ce2e59eForbidden with default headers values
func NewNr5ade12b3d5b81c209ce2e59eForbidden() *Nr5ade12b3d5b81c209ce2e59eForbidden {
	return &Nr5ade12b3d5b81c209ce2e59eForbidden{}
}

/*Nr5ade12b3d5b81c209ce2e59eForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type Nr5ade12b3d5b81c209ce2e59eForbidden struct {
}

func (o *Nr5ade12b3d5b81c209ce2e59eForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/roles][%d] 5ade12b3d5b81c209ce2e59eForbidden ", 403)
}

func (o *Nr5ade12b3d5b81c209ce2e59eForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade12b3d5b81c209ce2e59eTooManyRequests creates a Nr5ade12b3d5b81c209ce2e59eTooManyRequests with default headers values
func NewNr5ade12b3d5b81c209ce2e59eTooManyRequests() *Nr5ade12b3d5b81c209ce2e59eTooManyRequests {
	return &Nr5ade12b3d5b81c209ce2e59eTooManyRequests{}
}

/*Nr5ade12b3d5b81c209ce2e59eTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type Nr5ade12b3d5b81c209ce2e59eTooManyRequests struct {
}

func (o *Nr5ade12b3d5b81c209ce2e59eTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/roles][%d] 5ade12b3d5b81c209ce2e59eTooManyRequests ", 429)
}

func (o *Nr5ade12b3d5b81c209ce2e59eTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}