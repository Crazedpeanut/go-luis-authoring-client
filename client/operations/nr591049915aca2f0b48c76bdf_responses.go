// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// Nr591049915aca2f0b48c76bdfReader is a Reader for the Nr591049915aca2f0b48c76bdf structure.
type Nr591049915aca2f0b48c76bdfReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr591049915aca2f0b48c76bdfReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNr591049915aca2f0b48c76bdfOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr591049915aca2f0b48c76bdfBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr591049915aca2f0b48c76bdfUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNr591049915aca2f0b48c76bdfForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNr591049915aca2f0b48c76bdfTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr591049915aca2f0b48c76bdfOK creates a Nr591049915aca2f0b48c76bdfOK with default headers values
func NewNr591049915aca2f0b48c76bdfOK() *Nr591049915aca2f0b48c76bdfOK {
	return &Nr591049915aca2f0b48c76bdfOK{}
}

/*Nr591049915aca2f0b48c76bdfOK handles this case with default header values.

Returns a list of all custom prebuilt models and their representations.
*/
type Nr591049915aca2f0b48c76bdfOK struct {
}

func (o *Nr591049915aca2f0b48c76bdfOK) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/customprebuiltmodels][%d] 591049915aca2f0b48c76bdfOK ", 200)
}

func (o *Nr591049915aca2f0b48c76bdfOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr591049915aca2f0b48c76bdfBadRequest creates a Nr591049915aca2f0b48c76bdfBadRequest with default headers values
func NewNr591049915aca2f0b48c76bdfBadRequest() *Nr591049915aca2f0b48c76bdfBadRequest {
	return &Nr591049915aca2f0b48c76bdfBadRequest{}
}

/*Nr591049915aca2f0b48c76bdfBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type Nr591049915aca2f0b48c76bdfBadRequest struct {
}

func (o *Nr591049915aca2f0b48c76bdfBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/customprebuiltmodels][%d] 591049915aca2f0b48c76bdfBadRequest ", 400)
}

func (o *Nr591049915aca2f0b48c76bdfBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr591049915aca2f0b48c76bdfUnauthorized creates a Nr591049915aca2f0b48c76bdfUnauthorized with default headers values
func NewNr591049915aca2f0b48c76bdfUnauthorized() *Nr591049915aca2f0b48c76bdfUnauthorized {
	return &Nr591049915aca2f0b48c76bdfUnauthorized{}
}

/*Nr591049915aca2f0b48c76bdfUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type Nr591049915aca2f0b48c76bdfUnauthorized struct {
}

func (o *Nr591049915aca2f0b48c76bdfUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/customprebuiltmodels][%d] 591049915aca2f0b48c76bdfUnauthorized ", 401)
}

func (o *Nr591049915aca2f0b48c76bdfUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr591049915aca2f0b48c76bdfForbidden creates a Nr591049915aca2f0b48c76bdfForbidden with default headers values
func NewNr591049915aca2f0b48c76bdfForbidden() *Nr591049915aca2f0b48c76bdfForbidden {
	return &Nr591049915aca2f0b48c76bdfForbidden{}
}

/*Nr591049915aca2f0b48c76bdfForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type Nr591049915aca2f0b48c76bdfForbidden struct {
}

func (o *Nr591049915aca2f0b48c76bdfForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/customprebuiltmodels][%d] 591049915aca2f0b48c76bdfForbidden ", 403)
}

func (o *Nr591049915aca2f0b48c76bdfForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr591049915aca2f0b48c76bdfTooManyRequests creates a Nr591049915aca2f0b48c76bdfTooManyRequests with default headers values
func NewNr591049915aca2f0b48c76bdfTooManyRequests() *Nr591049915aca2f0b48c76bdfTooManyRequests {
	return &Nr591049915aca2f0b48c76bdfTooManyRequests{}
}

/*Nr591049915aca2f0b48c76bdfTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type Nr591049915aca2f0b48c76bdfTooManyRequests struct {
}

func (o *Nr591049915aca2f0b48c76bdfTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/customprebuiltmodels][%d] 591049915aca2f0b48c76bdfTooManyRequests ", 429)
}

func (o *Nr591049915aca2f0b48c76bdfTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
