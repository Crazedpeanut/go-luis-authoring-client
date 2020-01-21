// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// Nr5890b47c39e2bb052c5b9c39Reader is a Reader for the Nr5890b47c39e2bb052c5b9c39 structure.
type Nr5890b47c39e2bb052c5b9c39Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr5890b47c39e2bb052c5b9c39Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNr5890b47c39e2bb052c5b9c39OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr5890b47c39e2bb052c5b9c39BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr5890b47c39e2bb052c5b9c39Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNr5890b47c39e2bb052c5b9c39Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNr5890b47c39e2bb052c5b9c39TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr5890b47c39e2bb052c5b9c39OK creates a Nr5890b47c39e2bb052c5b9c39OK with default headers values
func NewNr5890b47c39e2bb052c5b9c39OK() *Nr5890b47c39e2bb052c5b9c39OK {
	return &Nr5890b47c39e2bb052c5b9c39OK{}
}

/*Nr5890b47c39e2bb052c5b9c39OK handles this case with default header values.

Nr5890b47c39e2bb052c5b9c39OK 5890b47c39e2bb052c5b9c39 o k
*/
type Nr5890b47c39e2bb052c5b9c39OK struct {
}

func (o *Nr5890b47c39e2bb052c5b9c39OK) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}][%d] 5890b47c39e2bb052c5b9c39OK ", 200)
}

func (o *Nr5890b47c39e2bb052c5b9c39OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c39BadRequest creates a Nr5890b47c39e2bb052c5b9c39BadRequest with default headers values
func NewNr5890b47c39e2bb052c5b9c39BadRequest() *Nr5890b47c39e2bb052c5b9c39BadRequest {
	return &Nr5890b47c39e2bb052c5b9c39BadRequest{}
}

/*Nr5890b47c39e2bb052c5b9c39BadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing or malformed.
*/
type Nr5890b47c39e2bb052c5b9c39BadRequest struct {
}

func (o *Nr5890b47c39e2bb052c5b9c39BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}][%d] 5890b47c39e2bb052c5b9c39BadRequest ", 400)
}

func (o *Nr5890b47c39e2bb052c5b9c39BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c39Unauthorized creates a Nr5890b47c39e2bb052c5b9c39Unauthorized with default headers values
func NewNr5890b47c39e2bb052c5b9c39Unauthorized() *Nr5890b47c39e2bb052c5b9c39Unauthorized {
	return &Nr5890b47c39e2bb052c5b9c39Unauthorized{}
}

/*Nr5890b47c39e2bb052c5b9c39Unauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type Nr5890b47c39e2bb052c5b9c39Unauthorized struct {
}

func (o *Nr5890b47c39e2bb052c5b9c39Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}][%d] 5890b47c39e2bb052c5b9c39Unauthorized ", 401)
}

func (o *Nr5890b47c39e2bb052c5b9c39Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c39Forbidden creates a Nr5890b47c39e2bb052c5b9c39Forbidden with default headers values
func NewNr5890b47c39e2bb052c5b9c39Forbidden() *Nr5890b47c39e2bb052c5b9c39Forbidden {
	return &Nr5890b47c39e2bb052c5b9c39Forbidden{}
}

/*Nr5890b47c39e2bb052c5b9c39Forbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type Nr5890b47c39e2bb052c5b9c39Forbidden struct {
}

func (o *Nr5890b47c39e2bb052c5b9c39Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}][%d] 5890b47c39e2bb052c5b9c39Forbidden ", 403)
}

func (o *Nr5890b47c39e2bb052c5b9c39Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c39TooManyRequests creates a Nr5890b47c39e2bb052c5b9c39TooManyRequests with default headers values
func NewNr5890b47c39e2bb052c5b9c39TooManyRequests() *Nr5890b47c39e2bb052c5b9c39TooManyRequests {
	return &Nr5890b47c39e2bb052c5b9c39TooManyRequests{}
}

/*Nr5890b47c39e2bb052c5b9c39TooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type Nr5890b47c39e2bb052c5b9c39TooManyRequests struct {
}

func (o *Nr5890b47c39e2bb052c5b9c39TooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}][%d] 5890b47c39e2bb052c5b9c39TooManyRequests ", 429)
}

func (o *Nr5890b47c39e2bb052c5b9c39TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
