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

// NewNr590aff885aca2f09e404ec3fParams creates a new Nr590aff885aca2f09e404ec3fParams object
// with the default values initialized.
func NewNr590aff885aca2f09e404ec3fParams() *Nr590aff885aca2f09e404ec3fParams {
	var ()
	return &Nr590aff885aca2f09e404ec3fParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNr590aff885aca2f09e404ec3fParamsWithTimeout creates a new Nr590aff885aca2f09e404ec3fParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr590aff885aca2f09e404ec3fParamsWithTimeout(timeout time.Duration) *Nr590aff885aca2f09e404ec3fParams {
	var ()
	return &Nr590aff885aca2f09e404ec3fParams{

		timeout: timeout,
	}
}

// NewNr590aff885aca2f09e404ec3fParamsWithContext creates a new Nr590aff885aca2f09e404ec3fParams object
// with the default values initialized, and the ability to set a context for a request
func NewNr590aff885aca2f09e404ec3fParamsWithContext(ctx context.Context) *Nr590aff885aca2f09e404ec3fParams {
	var ()
	return &Nr590aff885aca2f09e404ec3fParams{

		Context: ctx,
	}
}

// NewNr590aff885aca2f09e404ec3fParamsWithHTTPClient creates a new Nr590aff885aca2f09e404ec3fParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr590aff885aca2f09e404ec3fParamsWithHTTPClient(client *http.Client) *Nr590aff885aca2f09e404ec3fParams {
	var ()
	return &Nr590aff885aca2f09e404ec3fParams{
		HTTPClient: client,
	}
}

/*Nr590aff885aca2f09e404ec3fParams contains all the parameters to send to the API endpoint
for the 590aff885aca2f09e404ec3f operation typically these are written to a http.Request
*/
type Nr590aff885aca2f09e404ec3fParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the 590aff885aca2f09e404ec3f params
func (o *Nr590aff885aca2f09e404ec3fParams) WithTimeout(timeout time.Duration) *Nr590aff885aca2f09e404ec3fParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 590aff885aca2f09e404ec3f params
func (o *Nr590aff885aca2f09e404ec3fParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 590aff885aca2f09e404ec3f params
func (o *Nr590aff885aca2f09e404ec3fParams) WithContext(ctx context.Context) *Nr590aff885aca2f09e404ec3fParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 590aff885aca2f09e404ec3f params
func (o *Nr590aff885aca2f09e404ec3fParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 590aff885aca2f09e404ec3f params
func (o *Nr590aff885aca2f09e404ec3fParams) WithHTTPClient(client *http.Client) *Nr590aff885aca2f09e404ec3fParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 590aff885aca2f09e404ec3f params
func (o *Nr590aff885aca2f09e404ec3fParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the 590aff885aca2f09e404ec3f params
func (o *Nr590aff885aca2f09e404ec3fParams) WithAppID(appID string) *Nr590aff885aca2f09e404ec3fParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the 590aff885aca2f09e404ec3f params
func (o *Nr590aff885aca2f09e404ec3fParams) SetAppID(appID string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *Nr590aff885aca2f09e404ec3fParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
