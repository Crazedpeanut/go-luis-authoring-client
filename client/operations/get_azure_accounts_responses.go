// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/crazedpeanut/go-luis-authoring-client/models"
)

// GetAzureAccountsReader is a Reader for the GetAzureAccounts structure.
type GetAzureAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAzureAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAzureAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAzureAccountsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAzureAccountsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAzureAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAzureAccountsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAzureAccountsOK creates a GetAzureAccountsOK with default headers values
func NewGetAzureAccountsOK() *GetAzureAccountsOK {
	return &GetAzureAccountsOK{}
}

/*GetAzureAccountsOK handles this case with default header values.

A List of azure account information objects.
*/
type GetAzureAccountsOK struct {
	Payload []*models.AzureAccountInfoObject
}

func (o *GetAzureAccountsOK) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/azureaccounts][%d] getAzureAccountsOK  %+v", 200, o.Payload)
}

func (o *GetAzureAccountsOK) GetPayload() []*models.AzureAccountInfoObject {
	return o.Payload
}

func (o *GetAzureAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAzureAccountsBadRequest creates a GetAzureAccountsBadRequest with default headers values
func NewGetAzureAccountsBadRequest() *GetAzureAccountsBadRequest {
	return &GetAzureAccountsBadRequest{}
}

/*GetAzureAccountsBadRequest handles this case with default header values.

The app ID is missing or invalid.
The includeResponse is invalid.
*/
type GetAzureAccountsBadRequest struct {
}

func (o *GetAzureAccountsBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/azureaccounts][%d] getAzureAccountsBadRequest ", 400)
}

func (o *GetAzureAccountsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAzureAccountsUnauthorized creates a GetAzureAccountsUnauthorized with default headers values
func NewGetAzureAccountsUnauthorized() *GetAzureAccountsUnauthorized {
	return &GetAzureAccountsUnauthorized{}
}

/*GetAzureAccountsUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls

*/
type GetAzureAccountsUnauthorized struct {
}

func (o *GetAzureAccountsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/azureaccounts][%d] getAzureAccountsUnauthorized ", 401)
}

func (o *GetAzureAccountsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAzureAccountsForbidden creates a GetAzureAccountsForbidden with default headers values
func NewGetAzureAccountsForbidden() *GetAzureAccountsForbidden {
	return &GetAzureAccountsForbidden{}
}

/*GetAzureAccountsForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type GetAzureAccountsForbidden struct {
}

func (o *GetAzureAccountsForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/azureaccounts][%d] getAzureAccountsForbidden ", 403)
}

func (o *GetAzureAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAzureAccountsTooManyRequests creates a GetAzureAccountsTooManyRequests with default headers values
func NewGetAzureAccountsTooManyRequests() *GetAzureAccountsTooManyRequests {
	return &GetAzureAccountsTooManyRequests{}
}

/*GetAzureAccountsTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type GetAzureAccountsTooManyRequests struct {
}

func (o *GetAzureAccountsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/azureaccounts][%d] getAzureAccountsTooManyRequests ", 429)
}

func (o *GetAzureAccountsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
