// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNr58fcccdd5aca2f08a4104342Params creates a new Nr58fcccdd5aca2f08a4104342Params object
// with the default values initialized.
func NewNr58fcccdd5aca2f08a4104342Params() *Nr58fcccdd5aca2f08a4104342Params {
	var ()
	return &Nr58fcccdd5aca2f08a4104342Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewNr58fcccdd5aca2f08a4104342ParamsWithTimeout creates a new Nr58fcccdd5aca2f08a4104342Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr58fcccdd5aca2f08a4104342ParamsWithTimeout(timeout time.Duration) *Nr58fcccdd5aca2f08a4104342Params {
	var ()
	return &Nr58fcccdd5aca2f08a4104342Params{

		timeout: timeout,
	}
}

// NewNr58fcccdd5aca2f08a4104342ParamsWithContext creates a new Nr58fcccdd5aca2f08a4104342Params object
// with the default values initialized, and the ability to set a context for a request
func NewNr58fcccdd5aca2f08a4104342ParamsWithContext(ctx context.Context) *Nr58fcccdd5aca2f08a4104342Params {
	var ()
	return &Nr58fcccdd5aca2f08a4104342Params{

		Context: ctx,
	}
}

// NewNr58fcccdd5aca2f08a4104342ParamsWithHTTPClient creates a new Nr58fcccdd5aca2f08a4104342Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr58fcccdd5aca2f08a4104342ParamsWithHTTPClient(client *http.Client) *Nr58fcccdd5aca2f08a4104342Params {
	var ()
	return &Nr58fcccdd5aca2f08a4104342Params{
		HTTPClient: client,
	}
}

/*Nr58fcccdd5aca2f08a4104342Params contains all the parameters to send to the API endpoint
for the 58fcccdd5aca2f08a4104342 operation typically these are written to a http.Request
*/
type Nr58fcccdd5aca2f08a4104342Params struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*Body
	  A JSON object containing the user's email address.


	*/
	Body interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the 58fcccdd5aca2f08a4104342 params
func (o *Nr58fcccdd5aca2f08a4104342Params) WithTimeout(timeout time.Duration) *Nr58fcccdd5aca2f08a4104342Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 58fcccdd5aca2f08a4104342 params
func (o *Nr58fcccdd5aca2f08a4104342Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 58fcccdd5aca2f08a4104342 params
func (o *Nr58fcccdd5aca2f08a4104342Params) WithContext(ctx context.Context) *Nr58fcccdd5aca2f08a4104342Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 58fcccdd5aca2f08a4104342 params
func (o *Nr58fcccdd5aca2f08a4104342Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 58fcccdd5aca2f08a4104342 params
func (o *Nr58fcccdd5aca2f08a4104342Params) WithHTTPClient(client *http.Client) *Nr58fcccdd5aca2f08a4104342Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 58fcccdd5aca2f08a4104342 params
func (o *Nr58fcccdd5aca2f08a4104342Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the 58fcccdd5aca2f08a4104342 params
func (o *Nr58fcccdd5aca2f08a4104342Params) WithAppID(appID string) *Nr58fcccdd5aca2f08a4104342Params {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the 58fcccdd5aca2f08a4104342 params
func (o *Nr58fcccdd5aca2f08a4104342Params) SetAppID(appID string) {
	o.AppID = appID
}

// WithBody adds the body to the 58fcccdd5aca2f08a4104342 params
func (o *Nr58fcccdd5aca2f08a4104342Params) WithBody(body interface{}) *Nr58fcccdd5aca2f08a4104342Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the 58fcccdd5aca2f08a4104342 params
func (o *Nr58fcccdd5aca2f08a4104342Params) SetBody(body interface{}) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *Nr58fcccdd5aca2f08a4104342Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
