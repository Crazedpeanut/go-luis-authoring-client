// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AppsPackagepublishedapplicationasgzipReader is a Reader for the AppsPackagepublishedapplicationasgzip structure.
type AppsPackagepublishedapplicationasgzipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppsPackagepublishedapplicationasgzipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppsPackagepublishedapplicationasgzipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAppsPackagepublishedapplicationasgzipBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAppsPackagepublishedapplicationasgzipUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAppsPackagepublishedapplicationasgzipForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAppsPackagepublishedapplicationasgzipTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAppsPackagepublishedapplicationasgzipOK creates a AppsPackagepublishedapplicationasgzipOK with default headers values
func NewAppsPackagepublishedapplicationasgzipOK() *AppsPackagepublishedapplicationasgzipOK {
	return &AppsPackagepublishedapplicationasgzipOK{}
}

/*AppsPackagepublishedapplicationasgzipOK handles this case with default header values.

The GZip byte array of the exported published application.
*/
type AppsPackagepublishedapplicationasgzipOK struct {
}

func (o *AppsPackagepublishedapplicationasgzipOK) Error() string {
	return fmt.Sprintf("[GET /package/{appId}/slot/{slotName}/gzip][%d] appsPackagepublishedapplicationasgzipOK ", 200)
}

func (o *AppsPackagepublishedapplicationasgzipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAppsPackagepublishedapplicationasgzipBadRequest creates a AppsPackagepublishedapplicationasgzipBadRequest with default headers values
func NewAppsPackagepublishedapplicationasgzipBadRequest() *AppsPackagepublishedapplicationasgzipBadRequest {
	return &AppsPackagepublishedapplicationasgzipBadRequest{}
}

/*AppsPackagepublishedapplicationasgzipBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type AppsPackagepublishedapplicationasgzipBadRequest struct {
}

func (o *AppsPackagepublishedapplicationasgzipBadRequest) Error() string {
	return fmt.Sprintf("[GET /package/{appId}/slot/{slotName}/gzip][%d] appsPackagepublishedapplicationasgzipBadRequest ", 400)
}

func (o *AppsPackagepublishedapplicationasgzipBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAppsPackagepublishedapplicationasgzipUnauthorized creates a AppsPackagepublishedapplicationasgzipUnauthorized with default headers values
func NewAppsPackagepublishedapplicationasgzipUnauthorized() *AppsPackagepublishedapplicationasgzipUnauthorized {
	return &AppsPackagepublishedapplicationasgzipUnauthorized{}
}

/*AppsPackagepublishedapplicationasgzipUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls

*/
type AppsPackagepublishedapplicationasgzipUnauthorized struct {
}

func (o *AppsPackagepublishedapplicationasgzipUnauthorized) Error() string {
	return fmt.Sprintf("[GET /package/{appId}/slot/{slotName}/gzip][%d] appsPackagepublishedapplicationasgzipUnauthorized ", 401)
}

func (o *AppsPackagepublishedapplicationasgzipUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAppsPackagepublishedapplicationasgzipForbidden creates a AppsPackagepublishedapplicationasgzipForbidden with default headers values
func NewAppsPackagepublishedapplicationasgzipForbidden() *AppsPackagepublishedapplicationasgzipForbidden {
	return &AppsPackagepublishedapplicationasgzipForbidden{}
}

/*AppsPackagepublishedapplicationasgzipForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type AppsPackagepublishedapplicationasgzipForbidden struct {
}

func (o *AppsPackagepublishedapplicationasgzipForbidden) Error() string {
	return fmt.Sprintf("[GET /package/{appId}/slot/{slotName}/gzip][%d] appsPackagepublishedapplicationasgzipForbidden ", 403)
}

func (o *AppsPackagepublishedapplicationasgzipForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAppsPackagepublishedapplicationasgzipTooManyRequests creates a AppsPackagepublishedapplicationasgzipTooManyRequests with default headers values
func NewAppsPackagepublishedapplicationasgzipTooManyRequests() *AppsPackagepublishedapplicationasgzipTooManyRequests {
	return &AppsPackagepublishedapplicationasgzipTooManyRequests{}
}

/*AppsPackagepublishedapplicationasgzipTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type AppsPackagepublishedapplicationasgzipTooManyRequests struct {
}

func (o *AppsPackagepublishedapplicationasgzipTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /package/{appId}/slot/{slotName}/gzip][%d] appsPackagepublishedapplicationasgzipTooManyRequests ", 429)
}

func (o *AppsPackagepublishedapplicationasgzipTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
