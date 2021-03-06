// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePhraseListFeatureReader is a Reader for the DeletePhraseListFeature structure.
type DeletePhraseListFeatureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePhraseListFeatureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePhraseListFeatureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeletePhraseListFeatureBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeletePhraseListFeatureUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePhraseListFeatureForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeletePhraseListFeatureTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePhraseListFeatureOK creates a DeletePhraseListFeatureOK with default headers values
func NewDeletePhraseListFeatureOK() *DeletePhraseListFeatureOK {
	return &DeletePhraseListFeatureOK{}
}

/*DeletePhraseListFeatureOK handles this case with default header values.

DeletePhraseListFeatureOK delete phrase list feature o k
*/
type DeletePhraseListFeatureOK struct {
}

func (o *DeletePhraseListFeatureOK) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/phraselists/{phraselistId}][%d] deletePhraseListFeatureOK ", 200)
}

func (o *DeletePhraseListFeatureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePhraseListFeatureBadRequest creates a DeletePhraseListFeatureBadRequest with default headers values
func NewDeletePhraseListFeatureBadRequest() *DeletePhraseListFeatureBadRequest {
	return &DeletePhraseListFeatureBadRequest{}
}

/*DeletePhraseListFeatureBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing or malformed.
*/
type DeletePhraseListFeatureBadRequest struct {
}

func (o *DeletePhraseListFeatureBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/phraselists/{phraselistId}][%d] deletePhraseListFeatureBadRequest ", 400)
}

func (o *DeletePhraseListFeatureBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePhraseListFeatureUnauthorized creates a DeletePhraseListFeatureUnauthorized with default headers values
func NewDeletePhraseListFeatureUnauthorized() *DeletePhraseListFeatureUnauthorized {
	return &DeletePhraseListFeatureUnauthorized{}
}

/*DeletePhraseListFeatureUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls

*/
type DeletePhraseListFeatureUnauthorized struct {
}

func (o *DeletePhraseListFeatureUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/phraselists/{phraselistId}][%d] deletePhraseListFeatureUnauthorized ", 401)
}

func (o *DeletePhraseListFeatureUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePhraseListFeatureForbidden creates a DeletePhraseListFeatureForbidden with default headers values
func NewDeletePhraseListFeatureForbidden() *DeletePhraseListFeatureForbidden {
	return &DeletePhraseListFeatureForbidden{}
}

/*DeletePhraseListFeatureForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type DeletePhraseListFeatureForbidden struct {
}

func (o *DeletePhraseListFeatureForbidden) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/phraselists/{phraselistId}][%d] deletePhraseListFeatureForbidden ", 403)
}

func (o *DeletePhraseListFeatureForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePhraseListFeatureTooManyRequests creates a DeletePhraseListFeatureTooManyRequests with default headers values
func NewDeletePhraseListFeatureTooManyRequests() *DeletePhraseListFeatureTooManyRequests {
	return &DeletePhraseListFeatureTooManyRequests{}
}

/*DeletePhraseListFeatureTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type DeletePhraseListFeatureTooManyRequests struct {
}

func (o *DeletePhraseListFeatureTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/phraselists/{phraselistId}][%d] deletePhraseListFeatureTooManyRequests ", 429)
}

func (o *DeletePhraseListFeatureTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
