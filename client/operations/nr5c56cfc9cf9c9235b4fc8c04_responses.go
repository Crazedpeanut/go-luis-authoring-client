// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// Nr5c56cfc9cf9c9235b4fc8c04Reader is a Reader for the Nr5c56cfc9cf9c9235b4fc8c04 structure.
type Nr5c56cfc9cf9c9235b4fc8c04Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr5c56cfc9cf9c9235b4fc8c04Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNr5c56cfc9cf9c9235b4fc8c04OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr5c56cfc9cf9c9235b4fc8c04BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr5c56cfc9cf9c9235b4fc8c04Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNr5c56cfc9cf9c9235b4fc8c04Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNr5c56cfc9cf9c9235b4fc8c04TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr5c56cfc9cf9c9235b4fc8c04OK creates a Nr5c56cfc9cf9c9235b4fc8c04OK with default headers values
func NewNr5c56cfc9cf9c9235b4fc8c04OK() *Nr5c56cfc9cf9c9235b4fc8c04OK {
	return &Nr5c56cfc9cf9c9235b4fc8c04OK{}
}

/*Nr5c56cfc9cf9c9235b4fc8c04OK handles this case with default header values.

An object containing the culture info and the supported application tokenizer versions.
*/
type Nr5c56cfc9cf9c9235b4fc8c04OK struct {
}

func (o *Nr5c56cfc9cf9c9235b4fc8c04OK) Error() string {
	return fmt.Sprintf("[GET /apps/cultures/{cultureCode}/tokenizerversions][%d] 5c56cfc9cf9c9235b4fc8c04OK ", 200)
}

func (o *Nr5c56cfc9cf9c9235b4fc8c04OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5c56cfc9cf9c9235b4fc8c04BadRequest creates a Nr5c56cfc9cf9c9235b4fc8c04BadRequest with default headers values
func NewNr5c56cfc9cf9c9235b4fc8c04BadRequest() *Nr5c56cfc9cf9c9235b4fc8c04BadRequest {
	return &Nr5c56cfc9cf9c9235b4fc8c04BadRequest{}
}

/*Nr5c56cfc9cf9c9235b4fc8c04BadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type Nr5c56cfc9cf9c9235b4fc8c04BadRequest struct {
}

func (o *Nr5c56cfc9cf9c9235b4fc8c04BadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/cultures/{cultureCode}/tokenizerversions][%d] 5c56cfc9cf9c9235b4fc8c04BadRequest ", 400)
}

func (o *Nr5c56cfc9cf9c9235b4fc8c04BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5c56cfc9cf9c9235b4fc8c04Unauthorized creates a Nr5c56cfc9cf9c9235b4fc8c04Unauthorized with default headers values
func NewNr5c56cfc9cf9c9235b4fc8c04Unauthorized() *Nr5c56cfc9cf9c9235b4fc8c04Unauthorized {
	return &Nr5c56cfc9cf9c9235b4fc8c04Unauthorized{}
}

/*Nr5c56cfc9cf9c9235b4fc8c04Unauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type Nr5c56cfc9cf9c9235b4fc8c04Unauthorized struct {
}

func (o *Nr5c56cfc9cf9c9235b4fc8c04Unauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/cultures/{cultureCode}/tokenizerversions][%d] 5c56cfc9cf9c9235b4fc8c04Unauthorized ", 401)
}

func (o *Nr5c56cfc9cf9c9235b4fc8c04Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5c56cfc9cf9c9235b4fc8c04Forbidden creates a Nr5c56cfc9cf9c9235b4fc8c04Forbidden with default headers values
func NewNr5c56cfc9cf9c9235b4fc8c04Forbidden() *Nr5c56cfc9cf9c9235b4fc8c04Forbidden {
	return &Nr5c56cfc9cf9c9235b4fc8c04Forbidden{}
}

/*Nr5c56cfc9cf9c9235b4fc8c04Forbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type Nr5c56cfc9cf9c9235b4fc8c04Forbidden struct {
}

func (o *Nr5c56cfc9cf9c9235b4fc8c04Forbidden) Error() string {
	return fmt.Sprintf("[GET /apps/cultures/{cultureCode}/tokenizerversions][%d] 5c56cfc9cf9c9235b4fc8c04Forbidden ", 403)
}

func (o *Nr5c56cfc9cf9c9235b4fc8c04Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5c56cfc9cf9c9235b4fc8c04TooManyRequests creates a Nr5c56cfc9cf9c9235b4fc8c04TooManyRequests with default headers values
func NewNr5c56cfc9cf9c9235b4fc8c04TooManyRequests() *Nr5c56cfc9cf9c9235b4fc8c04TooManyRequests {
	return &Nr5c56cfc9cf9c9235b4fc8c04TooManyRequests{}
}

/*Nr5c56cfc9cf9c9235b4fc8c04TooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type Nr5c56cfc9cf9c9235b4fc8c04TooManyRequests struct {
}

func (o *Nr5c56cfc9cf9c9235b4fc8c04TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/cultures/{cultureCode}/tokenizerversions][%d] 5c56cfc9cf9c9235b4fc8c04TooManyRequests ", 429)
}

func (o *Nr5c56cfc9cf9c9235b4fc8c04TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}