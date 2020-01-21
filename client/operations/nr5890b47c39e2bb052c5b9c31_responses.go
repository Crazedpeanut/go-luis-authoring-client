// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// Nr5890b47c39e2bb052c5b9c31Reader is a Reader for the Nr5890b47c39e2bb052c5b9c31 structure.
type Nr5890b47c39e2bb052c5b9c31Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr5890b47c39e2bb052c5b9c31Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNr5890b47c39e2bb052c5b9c31Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr5890b47c39e2bb052c5b9c31BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr5890b47c39e2bb052c5b9c31Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNr5890b47c39e2bb052c5b9c31Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNr5890b47c39e2bb052c5b9c31TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr5890b47c39e2bb052c5b9c31Created creates a Nr5890b47c39e2bb052c5b9c31Created with default headers values
func NewNr5890b47c39e2bb052c5b9c31Created() *Nr5890b47c39e2bb052c5b9c31Created {
	return &Nr5890b47c39e2bb052c5b9c31Created{}
}

/*Nr5890b47c39e2bb052c5b9c31Created handles this case with default header values.

The ID of the imported application.
*/
type Nr5890b47c39e2bb052c5b9c31Created struct {
}

func (o *Nr5890b47c39e2bb052c5b9c31Created) Error() string {
	return fmt.Sprintf("[POST /apps/import][%d] 5890b47c39e2bb052c5b9c31Created ", 201)
}

func (o *Nr5890b47c39e2bb052c5b9c31Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c31BadRequest creates a Nr5890b47c39e2bb052c5b9c31BadRequest with default headers values
func NewNr5890b47c39e2bb052c5b9c31BadRequest() *Nr5890b47c39e2bb052c5b9c31BadRequest {
	return &Nr5890b47c39e2bb052c5b9c31BadRequest{}
}

/*Nr5890b47c39e2bb052c5b9c31BadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.

This error can be returned if the request's body is incorrect meaning the JSON is missing, malformed, or too large.

This error is returned is your import JSON is incorrect. This is usual due to using values in the JSON that are no longer used or missing properties that are required. The error message will reference the exact property that is causing the problem.
*/
type Nr5890b47c39e2bb052c5b9c31BadRequest struct {
}

func (o *Nr5890b47c39e2bb052c5b9c31BadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/import][%d] 5890b47c39e2bb052c5b9c31BadRequest ", 400)
}

func (o *Nr5890b47c39e2bb052c5b9c31BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c31Unauthorized creates a Nr5890b47c39e2bb052c5b9c31Unauthorized with default headers values
func NewNr5890b47c39e2bb052c5b9c31Unauthorized() *Nr5890b47c39e2bb052c5b9c31Unauthorized {
	return &Nr5890b47c39e2bb052c5b9c31Unauthorized{}
}

/*Nr5890b47c39e2bb052c5b9c31Unauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls

*/
type Nr5890b47c39e2bb052c5b9c31Unauthorized struct {
}

func (o *Nr5890b47c39e2bb052c5b9c31Unauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/import][%d] 5890b47c39e2bb052c5b9c31Unauthorized ", 401)
}

func (o *Nr5890b47c39e2bb052c5b9c31Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c31Forbidden creates a Nr5890b47c39e2bb052c5b9c31Forbidden with default headers values
func NewNr5890b47c39e2bb052c5b9c31Forbidden() *Nr5890b47c39e2bb052c5b9c31Forbidden {
	return &Nr5890b47c39e2bb052c5b9c31Forbidden{}
}

/*Nr5890b47c39e2bb052c5b9c31Forbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type Nr5890b47c39e2bb052c5b9c31Forbidden struct {
}

func (o *Nr5890b47c39e2bb052c5b9c31Forbidden) Error() string {
	return fmt.Sprintf("[POST /apps/import][%d] 5890b47c39e2bb052c5b9c31Forbidden ", 403)
}

func (o *Nr5890b47c39e2bb052c5b9c31Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c31TooManyRequests creates a Nr5890b47c39e2bb052c5b9c31TooManyRequests with default headers values
func NewNr5890b47c39e2bb052c5b9c31TooManyRequests() *Nr5890b47c39e2bb052c5b9c31TooManyRequests {
	return &Nr5890b47c39e2bb052c5b9c31TooManyRequests{}
}

/*Nr5890b47c39e2bb052c5b9c31TooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type Nr5890b47c39e2bb052c5b9c31TooManyRequests struct {
}

func (o *Nr5890b47c39e2bb052c5b9c31TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /apps/import][%d] 5890b47c39e2bb052c5b9c31TooManyRequests ", 429)
}

func (o *Nr5890b47c39e2bb052c5b9c31TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
