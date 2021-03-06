// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetUsageScenariosReader is a Reader for the GetUsageScenarios structure.
type GetUsageScenariosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsageScenariosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsageScenariosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsageScenariosBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUsageScenariosUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsageScenariosForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUsageScenariosTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsageScenariosOK creates a GetUsageScenariosOK with default headers values
func NewGetUsageScenariosOK() *GetUsageScenariosOK {
	return &GetUsageScenariosOK{}
}

/*GetUsageScenariosOK handles this case with default header values.

A list object containing the available application usage scenarios.
*/
type GetUsageScenariosOK struct {
}

func (o *GetUsageScenariosOK) Error() string {
	return fmt.Sprintf("[GET /apps/usagescenarios][%d] getUsageScenariosOK ", 200)
}

func (o *GetUsageScenariosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsageScenariosBadRequest creates a GetUsageScenariosBadRequest with default headers values
func NewGetUsageScenariosBadRequest() *GetUsageScenariosBadRequest {
	return &GetUsageScenariosBadRequest{}
}

/*GetUsageScenariosBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type GetUsageScenariosBadRequest struct {
}

func (o *GetUsageScenariosBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/usagescenarios][%d] getUsageScenariosBadRequest ", 400)
}

func (o *GetUsageScenariosBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsageScenariosUnauthorized creates a GetUsageScenariosUnauthorized with default headers values
func NewGetUsageScenariosUnauthorized() *GetUsageScenariosUnauthorized {
	return &GetUsageScenariosUnauthorized{}
}

/*GetUsageScenariosUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type GetUsageScenariosUnauthorized struct {
}

func (o *GetUsageScenariosUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/usagescenarios][%d] getUsageScenariosUnauthorized ", 401)
}

func (o *GetUsageScenariosUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsageScenariosForbidden creates a GetUsageScenariosForbidden with default headers values
func NewGetUsageScenariosForbidden() *GetUsageScenariosForbidden {
	return &GetUsageScenariosForbidden{}
}

/*GetUsageScenariosForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type GetUsageScenariosForbidden struct {
}

func (o *GetUsageScenariosForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/usagescenarios][%d] getUsageScenariosForbidden ", 403)
}

func (o *GetUsageScenariosForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsageScenariosTooManyRequests creates a GetUsageScenariosTooManyRequests with default headers values
func NewGetUsageScenariosTooManyRequests() *GetUsageScenariosTooManyRequests {
	return &GetUsageScenariosTooManyRequests{}
}

/*GetUsageScenariosTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type GetUsageScenariosTooManyRequests struct {
}

func (o *GetUsageScenariosTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/usagescenarios][%d] getUsageScenariosTooManyRequests ", 429)
}

func (o *GetUsageScenariosTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
