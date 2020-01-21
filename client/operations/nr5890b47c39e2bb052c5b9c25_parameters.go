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

// NewNr5890b47c39e2bb052c5b9c25Params creates a new Nr5890b47c39e2bb052c5b9c25Params object
// with the default values initialized.
func NewNr5890b47c39e2bb052c5b9c25Params() *Nr5890b47c39e2bb052c5b9c25Params {
	var ()
	return &Nr5890b47c39e2bb052c5b9c25Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewNr5890b47c39e2bb052c5b9c25ParamsWithTimeout creates a new Nr5890b47c39e2bb052c5b9c25Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr5890b47c39e2bb052c5b9c25ParamsWithTimeout(timeout time.Duration) *Nr5890b47c39e2bb052c5b9c25Params {
	var ()
	return &Nr5890b47c39e2bb052c5b9c25Params{

		timeout: timeout,
	}
}

// NewNr5890b47c39e2bb052c5b9c25ParamsWithContext creates a new Nr5890b47c39e2bb052c5b9c25Params object
// with the default values initialized, and the ability to set a context for a request
func NewNr5890b47c39e2bb052c5b9c25ParamsWithContext(ctx context.Context) *Nr5890b47c39e2bb052c5b9c25Params {
	var ()
	return &Nr5890b47c39e2bb052c5b9c25Params{

		Context: ctx,
	}
}

// NewNr5890b47c39e2bb052c5b9c25ParamsWithHTTPClient creates a new Nr5890b47c39e2bb052c5b9c25Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr5890b47c39e2bb052c5b9c25ParamsWithHTTPClient(client *http.Client) *Nr5890b47c39e2bb052c5b9c25Params {
	var ()
	return &Nr5890b47c39e2bb052c5b9c25Params{
		HTTPClient: client,
	}
}

/*Nr5890b47c39e2bb052c5b9c25Params contains all the parameters to send to the API endpoint
for the 5890b47c39e2bb052c5b9c25 operation typically these are written to a http.Request
*/
type Nr5890b47c39e2bb052c5b9c25Params struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*CEntityID
	  Format - guid. The composite entity extractor ID.

	*/
	CEntityID string
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the 5890b47c39e2bb052c5b9c25 params
func (o *Nr5890b47c39e2bb052c5b9c25Params) WithTimeout(timeout time.Duration) *Nr5890b47c39e2bb052c5b9c25Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 5890b47c39e2bb052c5b9c25 params
func (o *Nr5890b47c39e2bb052c5b9c25Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 5890b47c39e2bb052c5b9c25 params
func (o *Nr5890b47c39e2bb052c5b9c25Params) WithContext(ctx context.Context) *Nr5890b47c39e2bb052c5b9c25Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 5890b47c39e2bb052c5b9c25 params
func (o *Nr5890b47c39e2bb052c5b9c25Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 5890b47c39e2bb052c5b9c25 params
func (o *Nr5890b47c39e2bb052c5b9c25Params) WithHTTPClient(client *http.Client) *Nr5890b47c39e2bb052c5b9c25Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 5890b47c39e2bb052c5b9c25 params
func (o *Nr5890b47c39e2bb052c5b9c25Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the 5890b47c39e2bb052c5b9c25 params
func (o *Nr5890b47c39e2bb052c5b9c25Params) WithAppID(appID string) *Nr5890b47c39e2bb052c5b9c25Params {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the 5890b47c39e2bb052c5b9c25 params
func (o *Nr5890b47c39e2bb052c5b9c25Params) SetAppID(appID string) {
	o.AppID = appID
}

// WithCEntityID adds the cEntityID to the 5890b47c39e2bb052c5b9c25 params
func (o *Nr5890b47c39e2bb052c5b9c25Params) WithCEntityID(cEntityID string) *Nr5890b47c39e2bb052c5b9c25Params {
	o.SetCEntityID(cEntityID)
	return o
}

// SetCEntityID adds the cEntityId to the 5890b47c39e2bb052c5b9c25 params
func (o *Nr5890b47c39e2bb052c5b9c25Params) SetCEntityID(cEntityID string) {
	o.CEntityID = cEntityID
}

// WithVersionID adds the versionID to the 5890b47c39e2bb052c5b9c25 params
func (o *Nr5890b47c39e2bb052c5b9c25Params) WithVersionID(versionID string) *Nr5890b47c39e2bb052c5b9c25Params {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the 5890b47c39e2bb052c5b9c25 params
func (o *Nr5890b47c39e2bb052c5b9c25Params) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *Nr5890b47c39e2bb052c5b9c25Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	// path param cEntityId
	if err := r.SetPathParam("cEntityId", o.CEntityID); err != nil {
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
