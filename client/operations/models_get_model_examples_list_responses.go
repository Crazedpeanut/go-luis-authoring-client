// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ModelsGetModelExamplesListReader is a Reader for the ModelsGetModelExamplesList structure.
type ModelsGetModelExamplesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModelsGetModelExamplesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModelsGetModelExamplesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewModelsGetModelExamplesListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewModelsGetModelExamplesListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewModelsGetModelExamplesListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewModelsGetModelExamplesListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModelsGetModelExamplesListOK creates a ModelsGetModelExamplesListOK with default headers values
func NewModelsGetModelExamplesListOK() *ModelsGetModelExamplesListOK {
	return &ModelsGetModelExamplesListOK{}
}

/*ModelsGetModelExamplesListOK handles this case with default header values.

A list of examples for a specific requested model.
*/
type ModelsGetModelExamplesListOK struct {
}

func (o *ModelsGetModelExamplesListOK) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/models/{modelId}/examples][%d] modelsGetModelExamplesListOK ", 200)
}

func (o *ModelsGetModelExamplesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModelsGetModelExamplesListBadRequest creates a ModelsGetModelExamplesListBadRequest with default headers values
func NewModelsGetModelExamplesListBadRequest() *ModelsGetModelExamplesListBadRequest {
	return &ModelsGetModelExamplesListBadRequest{}
}

/*ModelsGetModelExamplesListBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type ModelsGetModelExamplesListBadRequest struct {
}

func (o *ModelsGetModelExamplesListBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/models/{modelId}/examples][%d] modelsGetModelExamplesListBadRequest ", 400)
}

func (o *ModelsGetModelExamplesListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModelsGetModelExamplesListUnauthorized creates a ModelsGetModelExamplesListUnauthorized with default headers values
func NewModelsGetModelExamplesListUnauthorized() *ModelsGetModelExamplesListUnauthorized {
	return &ModelsGetModelExamplesListUnauthorized{}
}

/*ModelsGetModelExamplesListUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type ModelsGetModelExamplesListUnauthorized struct {
}

func (o *ModelsGetModelExamplesListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/models/{modelId}/examples][%d] modelsGetModelExamplesListUnauthorized ", 401)
}

func (o *ModelsGetModelExamplesListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModelsGetModelExamplesListForbidden creates a ModelsGetModelExamplesListForbidden with default headers values
func NewModelsGetModelExamplesListForbidden() *ModelsGetModelExamplesListForbidden {
	return &ModelsGetModelExamplesListForbidden{}
}

/*ModelsGetModelExamplesListForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type ModelsGetModelExamplesListForbidden struct {
}

func (o *ModelsGetModelExamplesListForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/models/{modelId}/examples][%d] modelsGetModelExamplesListForbidden ", 403)
}

func (o *ModelsGetModelExamplesListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModelsGetModelExamplesListTooManyRequests creates a ModelsGetModelExamplesListTooManyRequests with default headers values
func NewModelsGetModelExamplesListTooManyRequests() *ModelsGetModelExamplesListTooManyRequests {
	return &ModelsGetModelExamplesListTooManyRequests{}
}

/*ModelsGetModelExamplesListTooManyRequests handles this case with default header values.

Rate limit is exceeded.


*/
type ModelsGetModelExamplesListTooManyRequests struct {
}

func (o *ModelsGetModelExamplesListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/models/{modelId}/examples][%d] modelsGetModelExamplesListTooManyRequests ", 429)
}

func (o *ModelsGetModelExamplesListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}