// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GethierarchicalEntitiesReader is a Reader for the GethierarchicalEntities structure.
type GethierarchicalEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GethierarchicalEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGethierarchicalEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGethierarchicalEntitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGethierarchicalEntitiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGethierarchicalEntitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGethierarchicalEntitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGethierarchicalEntitiesOK creates a GethierarchicalEntitiesOK with default headers values
func NewGethierarchicalEntitiesOK() *GethierarchicalEntitiesOK {
	return &GethierarchicalEntitiesOK{}
}

/*GethierarchicalEntitiesOK handles this case with default header values.

A list of hierarchical entity model infos.
*/
type GethierarchicalEntitiesOK struct {
}

func (o *GethierarchicalEntitiesOK) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/hierarchicalentities][%d] gethierarchicalEntitiesOK ", 200)
}

func (o *GethierarchicalEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGethierarchicalEntitiesBadRequest creates a GethierarchicalEntitiesBadRequest with default headers values
func NewGethierarchicalEntitiesBadRequest() *GethierarchicalEntitiesBadRequest {
	return &GethierarchicalEntitiesBadRequest{}
}

/*GethierarchicalEntitiesBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type GethierarchicalEntitiesBadRequest struct {
}

func (o *GethierarchicalEntitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/hierarchicalentities][%d] gethierarchicalEntitiesBadRequest ", 400)
}

func (o *GethierarchicalEntitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGethierarchicalEntitiesUnauthorized creates a GethierarchicalEntitiesUnauthorized with default headers values
func NewGethierarchicalEntitiesUnauthorized() *GethierarchicalEntitiesUnauthorized {
	return &GethierarchicalEntitiesUnauthorized{}
}

/*GethierarchicalEntitiesUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type GethierarchicalEntitiesUnauthorized struct {
}

func (o *GethierarchicalEntitiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/hierarchicalentities][%d] gethierarchicalEntitiesUnauthorized ", 401)
}

func (o *GethierarchicalEntitiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGethierarchicalEntitiesForbidden creates a GethierarchicalEntitiesForbidden with default headers values
func NewGethierarchicalEntitiesForbidden() *GethierarchicalEntitiesForbidden {
	return &GethierarchicalEntitiesForbidden{}
}

/*GethierarchicalEntitiesForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type GethierarchicalEntitiesForbidden struct {
}

func (o *GethierarchicalEntitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/hierarchicalentities][%d] gethierarchicalEntitiesForbidden ", 403)
}

func (o *GethierarchicalEntitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGethierarchicalEntitiesTooManyRequests creates a GethierarchicalEntitiesTooManyRequests with default headers values
func NewGethierarchicalEntitiesTooManyRequests() *GethierarchicalEntitiesTooManyRequests {
	return &GethierarchicalEntitiesTooManyRequests{}
}

/*GethierarchicalEntitiesTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type GethierarchicalEntitiesTooManyRequests struct {
}

func (o *GethierarchicalEntitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/hierarchicalentities][%d] gethierarchicalEntitiesTooManyRequests ", 429)
}

func (o *GethierarchicalEntitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
