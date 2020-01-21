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

// NewNr5adf73bbd5b81c09bc0db026Params creates a new Nr5adf73bbd5b81c09bc0db026Params object
// with the default values initialized.
func NewNr5adf73bbd5b81c09bc0db026Params() *Nr5adf73bbd5b81c09bc0db026Params {
	var ()
	return &Nr5adf73bbd5b81c09bc0db026Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewNr5adf73bbd5b81c09bc0db026ParamsWithTimeout creates a new Nr5adf73bbd5b81c09bc0db026Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr5adf73bbd5b81c09bc0db026ParamsWithTimeout(timeout time.Duration) *Nr5adf73bbd5b81c09bc0db026Params {
	var ()
	return &Nr5adf73bbd5b81c09bc0db026Params{

		timeout: timeout,
	}
}

// NewNr5adf73bbd5b81c09bc0db026ParamsWithContext creates a new Nr5adf73bbd5b81c09bc0db026Params object
// with the default values initialized, and the ability to set a context for a request
func NewNr5adf73bbd5b81c09bc0db026ParamsWithContext(ctx context.Context) *Nr5adf73bbd5b81c09bc0db026Params {
	var ()
	return &Nr5adf73bbd5b81c09bc0db026Params{

		Context: ctx,
	}
}

// NewNr5adf73bbd5b81c09bc0db026ParamsWithHTTPClient creates a new Nr5adf73bbd5b81c09bc0db026Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr5adf73bbd5b81c09bc0db026ParamsWithHTTPClient(client *http.Client) *Nr5adf73bbd5b81c09bc0db026Params {
	var ()
	return &Nr5adf73bbd5b81c09bc0db026Params{
		HTTPClient: client,
	}
}

/*Nr5adf73bbd5b81c09bc0db026Params contains all the parameters to send to the API endpoint
for the 5adf73bbd5b81c09bc0db026 operation typically these are written to a http.Request
*/
type Nr5adf73bbd5b81c09bc0db026Params struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*Body
	  A JSON object representing the new pattern.

	*/
	Body interface{}
	/*PatternID
	  Format - guid. The pattern ID.

	*/
	PatternID string
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the 5adf73bbd5b81c09bc0db026 params
func (o *Nr5adf73bbd5b81c09bc0db026Params) WithTimeout(timeout time.Duration) *Nr5adf73bbd5b81c09bc0db026Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 5adf73bbd5b81c09bc0db026 params
func (o *Nr5adf73bbd5b81c09bc0db026Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 5adf73bbd5b81c09bc0db026 params
func (o *Nr5adf73bbd5b81c09bc0db026Params) WithContext(ctx context.Context) *Nr5adf73bbd5b81c09bc0db026Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 5adf73bbd5b81c09bc0db026 params
func (o *Nr5adf73bbd5b81c09bc0db026Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 5adf73bbd5b81c09bc0db026 params
func (o *Nr5adf73bbd5b81c09bc0db026Params) WithHTTPClient(client *http.Client) *Nr5adf73bbd5b81c09bc0db026Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 5adf73bbd5b81c09bc0db026 params
func (o *Nr5adf73bbd5b81c09bc0db026Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the 5adf73bbd5b81c09bc0db026 params
func (o *Nr5adf73bbd5b81c09bc0db026Params) WithAppID(appID string) *Nr5adf73bbd5b81c09bc0db026Params {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the 5adf73bbd5b81c09bc0db026 params
func (o *Nr5adf73bbd5b81c09bc0db026Params) SetAppID(appID string) {
	o.AppID = appID
}

// WithBody adds the body to the 5adf73bbd5b81c09bc0db026 params
func (o *Nr5adf73bbd5b81c09bc0db026Params) WithBody(body interface{}) *Nr5adf73bbd5b81c09bc0db026Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the 5adf73bbd5b81c09bc0db026 params
func (o *Nr5adf73bbd5b81c09bc0db026Params) SetBody(body interface{}) {
	o.Body = body
}

// WithPatternID adds the patternID to the 5adf73bbd5b81c09bc0db026 params
func (o *Nr5adf73bbd5b81c09bc0db026Params) WithPatternID(patternID string) *Nr5adf73bbd5b81c09bc0db026Params {
	o.SetPatternID(patternID)
	return o
}

// SetPatternID adds the patternId to the 5adf73bbd5b81c09bc0db026 params
func (o *Nr5adf73bbd5b81c09bc0db026Params) SetPatternID(patternID string) {
	o.PatternID = patternID
}

// WithVersionID adds the versionID to the 5adf73bbd5b81c09bc0db026 params
func (o *Nr5adf73bbd5b81c09bc0db026Params) WithVersionID(versionID string) *Nr5adf73bbd5b81c09bc0db026Params {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the 5adf73bbd5b81c09bc0db026 params
func (o *Nr5adf73bbd5b81c09bc0db026Params) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *Nr5adf73bbd5b81c09bc0db026Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param patternId
	if err := r.SetPathParam("patternId", o.PatternID); err != nil {
		return err
	}

	// path param versionId
	if err := r.SetPathParam("versionId", o.VersionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
