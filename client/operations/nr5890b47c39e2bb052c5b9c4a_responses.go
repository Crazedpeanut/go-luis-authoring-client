// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// Nr5890b47c39e2bb052c5b9c4aReader is a Reader for the Nr5890b47c39e2bb052c5b9c4a structure.
type Nr5890b47c39e2bb052c5b9c4aReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr5890b47c39e2bb052c5b9c4aReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNr5890b47c39e2bb052c5b9c4aOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr5890b47c39e2bb052c5b9c4aBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr5890b47c39e2bb052c5b9c4aUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewNr5890b47c39e2bb052c5b9c4aGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNr5890b47c39e2bb052c5b9c4aTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr5890b47c39e2bb052c5b9c4aOK creates a Nr5890b47c39e2bb052c5b9c4aOK with default headers values
func NewNr5890b47c39e2bb052c5b9c4aOK() *Nr5890b47c39e2bb052c5b9c4aOK {
	return &Nr5890b47c39e2bb052c5b9c4aOK{}
}

/*Nr5890b47c39e2bb052c5b9c4aOK handles this case with default header values.

A list containing user's external keys info.
*/
type Nr5890b47c39e2bb052c5b9c4aOK struct {
}

func (o *Nr5890b47c39e2bb052c5b9c4aOK) Error() string {
	return fmt.Sprintf("[GET /externalKeys][%d] 5890b47c39e2bb052c5b9c4aOK ", 200)
}

func (o *Nr5890b47c39e2bb052c5b9c4aOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c4aBadRequest creates a Nr5890b47c39e2bb052c5b9c4aBadRequest with default headers values
func NewNr5890b47c39e2bb052c5b9c4aBadRequest() *Nr5890b47c39e2bb052c5b9c4aBadRequest {
	return &Nr5890b47c39e2bb052c5b9c4aBadRequest{}
}

/*Nr5890b47c39e2bb052c5b9c4aBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type Nr5890b47c39e2bb052c5b9c4aBadRequest struct {
}

func (o *Nr5890b47c39e2bb052c5b9c4aBadRequest) Error() string {
	return fmt.Sprintf("[GET /externalKeys][%d] 5890b47c39e2bb052c5b9c4aBadRequest ", 400)
}

func (o *Nr5890b47c39e2bb052c5b9c4aBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c4aUnauthorized creates a Nr5890b47c39e2bb052c5b9c4aUnauthorized with default headers values
func NewNr5890b47c39e2bb052c5b9c4aUnauthorized() *Nr5890b47c39e2bb052c5b9c4aUnauthorized {
	return &Nr5890b47c39e2bb052c5b9c4aUnauthorized{}
}

/*Nr5890b47c39e2bb052c5b9c4aUnauthorized handles this case with default header values.

Nr5890b47c39e2bb052c5b9c4aUnauthorized 5890b47c39e2bb052c5b9c4a unauthorized
*/
type Nr5890b47c39e2bb052c5b9c4aUnauthorized struct {
}

func (o *Nr5890b47c39e2bb052c5b9c4aUnauthorized) Error() string {
	return fmt.Sprintf("[GET /externalKeys][%d] 5890b47c39e2bb052c5b9c4aUnauthorized ", 401)
}

func (o *Nr5890b47c39e2bb052c5b9c4aUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c4aGone creates a Nr5890b47c39e2bb052c5b9c4aGone with default headers values
func NewNr5890b47c39e2bb052c5b9c4aGone() *Nr5890b47c39e2bb052c5b9c4aGone {
	return &Nr5890b47c39e2bb052c5b9c4aGone{}
}

/*Nr5890b47c39e2bb052c5b9c4aGone handles this case with default header values.

THIS API IS DEPRECATED.
*/
type Nr5890b47c39e2bb052c5b9c4aGone struct {
}

func (o *Nr5890b47c39e2bb052c5b9c4aGone) Error() string {
	return fmt.Sprintf("[GET /externalKeys][%d] 5890b47c39e2bb052c5b9c4aGone ", 410)
}

func (o *Nr5890b47c39e2bb052c5b9c4aGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c4aTooManyRequests creates a Nr5890b47c39e2bb052c5b9c4aTooManyRequests with default headers values
func NewNr5890b47c39e2bb052c5b9c4aTooManyRequests() *Nr5890b47c39e2bb052c5b9c4aTooManyRequests {
	return &Nr5890b47c39e2bb052c5b9c4aTooManyRequests{}
}

/*Nr5890b47c39e2bb052c5b9c4aTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type Nr5890b47c39e2bb052c5b9c4aTooManyRequests struct {
}

func (o *Nr5890b47c39e2bb052c5b9c4aTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /externalKeys][%d] 5890b47c39e2bb052c5b9c4aTooManyRequests ", 429)
}

func (o *Nr5890b47c39e2bb052c5b9c4aTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
