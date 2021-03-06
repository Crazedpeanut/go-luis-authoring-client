// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetSubscriptionsReader is a Reader for the GetSubscriptions structure.
type GetSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewGetSubscriptionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSubscriptionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSubscriptionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSubscriptionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetSubscriptionsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSubscriptionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSubscriptionsCreated creates a GetSubscriptionsCreated with default headers values
func NewGetSubscriptionsCreated() *GetSubscriptionsCreated {
	return &GetSubscriptionsCreated{}
}

/*GetSubscriptionsCreated handles this case with default header values.

Returns the added subscription key information
*/
type GetSubscriptionsCreated struct {
}

func (o *GetSubscriptionsCreated) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] getSubscriptionsCreated ", 201)
}

func (o *GetSubscriptionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSubscriptionsBadRequest creates a GetSubscriptionsBadRequest with default headers values
func NewGetSubscriptionsBadRequest() *GetSubscriptionsBadRequest {
	return &GetSubscriptionsBadRequest{}
}

/*GetSubscriptionsBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.

This error can be returned if the request's body is incorrect meaning the JSON is missing, malformed, or too large.
*/
type GetSubscriptionsBadRequest struct {
}

func (o *GetSubscriptionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] getSubscriptionsBadRequest ", 400)
}

func (o *GetSubscriptionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSubscriptionsUnauthorized creates a GetSubscriptionsUnauthorized with default headers values
func NewGetSubscriptionsUnauthorized() *GetSubscriptionsUnauthorized {
	return &GetSubscriptionsUnauthorized{}
}

/*GetSubscriptionsUnauthorized handles this case with default header values.

GetSubscriptionsUnauthorized get subscriptions unauthorized
*/
type GetSubscriptionsUnauthorized struct {
}

func (o *GetSubscriptionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] getSubscriptionsUnauthorized ", 401)
}

func (o *GetSubscriptionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSubscriptionsForbidden creates a GetSubscriptionsForbidden with default headers values
func NewGetSubscriptionsForbidden() *GetSubscriptionsForbidden {
	return &GetSubscriptionsForbidden{}
}

/*GetSubscriptionsForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type GetSubscriptionsForbidden struct {
}

func (o *GetSubscriptionsForbidden) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] getSubscriptionsForbidden ", 403)
}

func (o *GetSubscriptionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSubscriptionsGone creates a GetSubscriptionsGone with default headers values
func NewGetSubscriptionsGone() *GetSubscriptionsGone {
	return &GetSubscriptionsGone{}
}

/*GetSubscriptionsGone handles this case with default header values.

THIS API IS DEPRECATED.
*/
type GetSubscriptionsGone struct {
}

func (o *GetSubscriptionsGone) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] getSubscriptionsGone ", 410)
}

func (o *GetSubscriptionsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSubscriptionsTooManyRequests creates a GetSubscriptionsTooManyRequests with default headers values
func NewGetSubscriptionsTooManyRequests() *GetSubscriptionsTooManyRequests {
	return &GetSubscriptionsTooManyRequests{}
}

/*GetSubscriptionsTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type GetSubscriptionsTooManyRequests struct {
}

func (o *GetSubscriptionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] getSubscriptionsTooManyRequests ", 429)
}

func (o *GetSubscriptionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
