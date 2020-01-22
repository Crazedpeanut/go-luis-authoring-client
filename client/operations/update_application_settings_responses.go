// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateApplicationSettingsReader is a Reader for the UpdateApplicationSettings structure.
type UpdateApplicationSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateApplicationSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateApplicationSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateApplicationSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateApplicationSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateApplicationSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateApplicationSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateApplicationSettingsOK creates a UpdateApplicationSettingsOK with default headers values
func NewUpdateApplicationSettingsOK() *UpdateApplicationSettingsOK {
	return &UpdateApplicationSettingsOK{}
}

/*UpdateApplicationSettingsOK handles this case with default header values.

UpdateApplicationSettingsOK update application settings o k
*/
type UpdateApplicationSettingsOK struct {
}

func (o *UpdateApplicationSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/settings][%d] updateApplicationSettingsOK ", 200)
}

func (o *UpdateApplicationSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationSettingsBadRequest creates a UpdateApplicationSettingsBadRequest with default headers values
func NewUpdateApplicationSettingsBadRequest() *UpdateApplicationSettingsBadRequest {
	return &UpdateApplicationSettingsBadRequest{}
}

/*UpdateApplicationSettingsBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.

This error can be returned if the request's body is incorrect meaning the JSON is missing, malformed, or too large.
*/
type UpdateApplicationSettingsBadRequest struct {
}

func (o *UpdateApplicationSettingsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/settings][%d] updateApplicationSettingsBadRequest ", 400)
}

func (o *UpdateApplicationSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationSettingsUnauthorized creates a UpdateApplicationSettingsUnauthorized with default headers values
func NewUpdateApplicationSettingsUnauthorized() *UpdateApplicationSettingsUnauthorized {
	return &UpdateApplicationSettingsUnauthorized{}
}

/*UpdateApplicationSettingsUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls


*/
type UpdateApplicationSettingsUnauthorized struct {
}

func (o *UpdateApplicationSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/settings][%d] updateApplicationSettingsUnauthorized ", 401)
}

func (o *UpdateApplicationSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationSettingsForbidden creates a UpdateApplicationSettingsForbidden with default headers values
func NewUpdateApplicationSettingsForbidden() *UpdateApplicationSettingsForbidden {
	return &UpdateApplicationSettingsForbidden{}
}

/*UpdateApplicationSettingsForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type UpdateApplicationSettingsForbidden struct {
}

func (o *UpdateApplicationSettingsForbidden) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/settings][%d] updateApplicationSettingsForbidden ", 403)
}

func (o *UpdateApplicationSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationSettingsTooManyRequests creates a UpdateApplicationSettingsTooManyRequests with default headers values
func NewUpdateApplicationSettingsTooManyRequests() *UpdateApplicationSettingsTooManyRequests {
	return &UpdateApplicationSettingsTooManyRequests{}
}

/*UpdateApplicationSettingsTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type UpdateApplicationSettingsTooManyRequests struct {
}

func (o *UpdateApplicationSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/settings][%d] updateApplicationSettingsTooManyRequests ", 429)
}

func (o *UpdateApplicationSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}