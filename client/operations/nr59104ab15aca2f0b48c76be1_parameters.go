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

// NewNr59104ab15aca2f0b48c76be1Params creates a new Nr59104ab15aca2f0b48c76be1Params object
// with the default values initialized.
func NewNr59104ab15aca2f0b48c76be1Params() *Nr59104ab15aca2f0b48c76be1Params {
	var ()
	return &Nr59104ab15aca2f0b48c76be1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewNr59104ab15aca2f0b48c76be1ParamsWithTimeout creates a new Nr59104ab15aca2f0b48c76be1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr59104ab15aca2f0b48c76be1ParamsWithTimeout(timeout time.Duration) *Nr59104ab15aca2f0b48c76be1Params {
	var ()
	return &Nr59104ab15aca2f0b48c76be1Params{

		timeout: timeout,
	}
}

// NewNr59104ab15aca2f0b48c76be1ParamsWithContext creates a new Nr59104ab15aca2f0b48c76be1Params object
// with the default values initialized, and the ability to set a context for a request
func NewNr59104ab15aca2f0b48c76be1ParamsWithContext(ctx context.Context) *Nr59104ab15aca2f0b48c76be1Params {
	var ()
	return &Nr59104ab15aca2f0b48c76be1Params{

		Context: ctx,
	}
}

// NewNr59104ab15aca2f0b48c76be1ParamsWithHTTPClient creates a new Nr59104ab15aca2f0b48c76be1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr59104ab15aca2f0b48c76be1ParamsWithHTTPClient(client *http.Client) *Nr59104ab15aca2f0b48c76be1Params {
	var ()
	return &Nr59104ab15aca2f0b48c76be1Params{
		HTTPClient: client,
	}
}

/*Nr59104ab15aca2f0b48c76be1Params contains all the parameters to send to the API endpoint
for the 59104ab15aca2f0b48c76be1 operation typically these are written to a http.Request
*/
type Nr59104ab15aca2f0b48c76be1Params struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the 59104ab15aca2f0b48c76be1 params
func (o *Nr59104ab15aca2f0b48c76be1Params) WithTimeout(timeout time.Duration) *Nr59104ab15aca2f0b48c76be1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 59104ab15aca2f0b48c76be1 params
func (o *Nr59104ab15aca2f0b48c76be1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 59104ab15aca2f0b48c76be1 params
func (o *Nr59104ab15aca2f0b48c76be1Params) WithContext(ctx context.Context) *Nr59104ab15aca2f0b48c76be1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 59104ab15aca2f0b48c76be1 params
func (o *Nr59104ab15aca2f0b48c76be1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 59104ab15aca2f0b48c76be1 params
func (o *Nr59104ab15aca2f0b48c76be1Params) WithHTTPClient(client *http.Client) *Nr59104ab15aca2f0b48c76be1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 59104ab15aca2f0b48c76be1 params
func (o *Nr59104ab15aca2f0b48c76be1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the 59104ab15aca2f0b48c76be1 params
func (o *Nr59104ab15aca2f0b48c76be1Params) WithAppID(appID string) *Nr59104ab15aca2f0b48c76be1Params {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the 59104ab15aca2f0b48c76be1 params
func (o *Nr59104ab15aca2f0b48c76be1Params) SetAppID(appID string) {
	o.AppID = appID
}

// WithVersionID adds the versionID to the 59104ab15aca2f0b48c76be1 params
func (o *Nr59104ab15aca2f0b48c76be1Params) WithVersionID(versionID string) *Nr59104ab15aca2f0b48c76be1Params {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the 59104ab15aca2f0b48c76be1 params
func (o *Nr59104ab15aca2f0b48c76be1Params) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *Nr59104ab15aca2f0b48c76be1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
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
