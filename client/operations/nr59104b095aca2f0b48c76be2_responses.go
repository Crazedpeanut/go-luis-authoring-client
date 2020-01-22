// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// Nr59104b095aca2f0b48c76be2Reader is a Reader for the Nr59104b095aca2f0b48c76be2 structure.
type Nr59104b095aca2f0b48c76be2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr59104b095aca2f0b48c76be2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNr59104b095aca2f0b48c76be2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr59104b095aca2f0b48c76be2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr59104b095aca2f0b48c76be2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNr59104b095aca2f0b48c76be2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNr59104b095aca2f0b48c76be2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr59104b095aca2f0b48c76be2OK creates a Nr59104b095aca2f0b48c76be2OK with default headers values
func NewNr59104b095aca2f0b48c76be2OK() *Nr59104b095aca2f0b48c76be2OK {
	return &Nr59104b095aca2f0b48c76be2OK{}
}

/*Nr59104b095aca2f0b48c76be2OK handles this case with default header values.

Nr59104b095aca2f0b48c76be2OK 59104b095aca2f0b48c76be2 o k
*/
type Nr59104b095aca2f0b48c76be2OK struct {
}

func (o *Nr59104b095aca2f0b48c76be2OK) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/customprebuiltdomains/{domainName}][%d] 59104b095aca2f0b48c76be2OK ", 200)
}

func (o *Nr59104b095aca2f0b48c76be2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr59104b095aca2f0b48c76be2BadRequest creates a Nr59104b095aca2f0b48c76be2BadRequest with default headers values
func NewNr59104b095aca2f0b48c76be2BadRequest() *Nr59104b095aca2f0b48c76be2BadRequest {
	return &Nr59104b095aca2f0b48c76be2BadRequest{}
}

/*Nr59104b095aca2f0b48c76be2BadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing or malformed.
*/
type Nr59104b095aca2f0b48c76be2BadRequest struct {
}

func (o *Nr59104b095aca2f0b48c76be2BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/customprebuiltdomains/{domainName}][%d] 59104b095aca2f0b48c76be2BadRequest ", 400)
}

func (o *Nr59104b095aca2f0b48c76be2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr59104b095aca2f0b48c76be2Unauthorized creates a Nr59104b095aca2f0b48c76be2Unauthorized with default headers values
func NewNr59104b095aca2f0b48c76be2Unauthorized() *Nr59104b095aca2f0b48c76be2Unauthorized {
	return &Nr59104b095aca2f0b48c76be2Unauthorized{}
}

/*Nr59104b095aca2f0b48c76be2Unauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type Nr59104b095aca2f0b48c76be2Unauthorized struct {
}

func (o *Nr59104b095aca2f0b48c76be2Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/customprebuiltdomains/{domainName}][%d] 59104b095aca2f0b48c76be2Unauthorized ", 401)
}

func (o *Nr59104b095aca2f0b48c76be2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr59104b095aca2f0b48c76be2Forbidden creates a Nr59104b095aca2f0b48c76be2Forbidden with default headers values
func NewNr59104b095aca2f0b48c76be2Forbidden() *Nr59104b095aca2f0b48c76be2Forbidden {
	return &Nr59104b095aca2f0b48c76be2Forbidden{}
}

/*Nr59104b095aca2f0b48c76be2Forbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type Nr59104b095aca2f0b48c76be2Forbidden struct {
}

func (o *Nr59104b095aca2f0b48c76be2Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/customprebuiltdomains/{domainName}][%d] 59104b095aca2f0b48c76be2Forbidden ", 403)
}

func (o *Nr59104b095aca2f0b48c76be2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr59104b095aca2f0b48c76be2TooManyRequests creates a Nr59104b095aca2f0b48c76be2TooManyRequests with default headers values
func NewNr59104b095aca2f0b48c76be2TooManyRequests() *Nr59104b095aca2f0b48c76be2TooManyRequests {
	return &Nr59104b095aca2f0b48c76be2TooManyRequests{}
}

/*Nr59104b095aca2f0b48c76be2TooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type Nr59104b095aca2f0b48c76be2TooManyRequests struct {
}

func (o *Nr59104b095aca2f0b48c76be2TooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/customprebuiltdomains/{domainName}][%d] 59104b095aca2f0b48c76be2TooManyRequests ", 429)
}

func (o *Nr59104b095aca2f0b48c76be2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}