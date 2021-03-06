// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// TrainVersionReader is a Reader for the TrainVersion structure.
type TrainVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TrainVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewTrainVersionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTrainVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTrainVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTrainVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewTrainVersionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTrainVersionAccepted creates a TrainVersionAccepted with default headers values
func NewTrainVersionAccepted() *TrainVersionAccepted {
	return &TrainVersionAccepted{}
}

/*TrainVersionAccepted handles this case with default header values.

This response indicates the initial training status. If the response is <b>Queued</b>, keep polling at the get training status API to check for training completion.
*/
type TrainVersionAccepted struct {
}

func (o *TrainVersionAccepted) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/train][%d] trainVersionAccepted ", 202)
}

func (o *TrainVersionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTrainVersionBadRequest creates a TrainVersionBadRequest with default headers values
func NewTrainVersionBadRequest() *TrainVersionBadRequest {
	return &TrainVersionBadRequest{}
}

/*TrainVersionBadRequest handles this case with default header values.

This response can occur if either of the <b>appId</b> or <b>versionId</b> parameters are incorrect.
*/
type TrainVersionBadRequest struct {
}

func (o *TrainVersionBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/train][%d] trainVersionBadRequest ", 400)
}

func (o *TrainVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTrainVersionUnauthorized creates a TrainVersionUnauthorized with default headers values
func NewTrainVersionUnauthorized() *TrainVersionUnauthorized {
	return &TrainVersionUnauthorized{}
}

/*TrainVersionUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls

*/
type TrainVersionUnauthorized struct {
}

func (o *TrainVersionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/train][%d] trainVersionUnauthorized ", 401)
}

func (o *TrainVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTrainVersionForbidden creates a TrainVersionForbidden with default headers values
func NewTrainVersionForbidden() *TrainVersionForbidden {
	return &TrainVersionForbidden{}
}

/*TrainVersionForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type TrainVersionForbidden struct {
}

func (o *TrainVersionForbidden) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/train][%d] trainVersionForbidden ", 403)
}

func (o *TrainVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTrainVersionTooManyRequests creates a TrainVersionTooManyRequests with default headers values
func NewTrainVersionTooManyRequests() *TrainVersionTooManyRequests {
	return &TrainVersionTooManyRequests{}
}

/*TrainVersionTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type TrainVersionTooManyRequests struct {
}

func (o *TrainVersionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/train][%d] trainVersionTooManyRequests ", 429)
}

func (o *TrainVersionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
