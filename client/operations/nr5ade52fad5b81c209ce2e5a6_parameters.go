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

// NewNr5ade52fad5b81c209ce2e5a6Params creates a new Nr5ade52fad5b81c209ce2e5a6Params object
// with the default values initialized.
func NewNr5ade52fad5b81c209ce2e5a6Params() *Nr5ade52fad5b81c209ce2e5a6Params {
	var ()
	return &Nr5ade52fad5b81c209ce2e5a6Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewNr5ade52fad5b81c209ce2e5a6ParamsWithTimeout creates a new Nr5ade52fad5b81c209ce2e5a6Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr5ade52fad5b81c209ce2e5a6ParamsWithTimeout(timeout time.Duration) *Nr5ade52fad5b81c209ce2e5a6Params {
	var ()
	return &Nr5ade52fad5b81c209ce2e5a6Params{

		timeout: timeout,
	}
}

// NewNr5ade52fad5b81c209ce2e5a6ParamsWithContext creates a new Nr5ade52fad5b81c209ce2e5a6Params object
// with the default values initialized, and the ability to set a context for a request
func NewNr5ade52fad5b81c209ce2e5a6ParamsWithContext(ctx context.Context) *Nr5ade52fad5b81c209ce2e5a6Params {
	var ()
	return &Nr5ade52fad5b81c209ce2e5a6Params{

		Context: ctx,
	}
}

// NewNr5ade52fad5b81c209ce2e5a6ParamsWithHTTPClient creates a new Nr5ade52fad5b81c209ce2e5a6Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr5ade52fad5b81c209ce2e5a6ParamsWithHTTPClient(client *http.Client) *Nr5ade52fad5b81c209ce2e5a6Params {
	var ()
	return &Nr5ade52fad5b81c209ce2e5a6Params{
		HTTPClient: client,
	}
}

/*Nr5ade52fad5b81c209ce2e5a6Params contains all the parameters to send to the API endpoint
for the 5ade52fad5b81c209ce2e5a6 operation typically these are written to a http.Request
*/
type Nr5ade52fad5b81c209ce2e5a6Params struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*EntityID
	  Format - guid. The custom prebuilt domain entity ID.

	*/
	EntityID string
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the 5ade52fad5b81c209ce2e5a6 params
func (o *Nr5ade52fad5b81c209ce2e5a6Params) WithTimeout(timeout time.Duration) *Nr5ade52fad5b81c209ce2e5a6Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 5ade52fad5b81c209ce2e5a6 params
func (o *Nr5ade52fad5b81c209ce2e5a6Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 5ade52fad5b81c209ce2e5a6 params
func (o *Nr5ade52fad5b81c209ce2e5a6Params) WithContext(ctx context.Context) *Nr5ade52fad5b81c209ce2e5a6Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 5ade52fad5b81c209ce2e5a6 params
func (o *Nr5ade52fad5b81c209ce2e5a6Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 5ade52fad5b81c209ce2e5a6 params
func (o *Nr5ade52fad5b81c209ce2e5a6Params) WithHTTPClient(client *http.Client) *Nr5ade52fad5b81c209ce2e5a6Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 5ade52fad5b81c209ce2e5a6 params
func (o *Nr5ade52fad5b81c209ce2e5a6Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the 5ade52fad5b81c209ce2e5a6 params
func (o *Nr5ade52fad5b81c209ce2e5a6Params) WithAppID(appID string) *Nr5ade52fad5b81c209ce2e5a6Params {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the 5ade52fad5b81c209ce2e5a6 params
func (o *Nr5ade52fad5b81c209ce2e5a6Params) SetAppID(appID string) {
	o.AppID = appID
}

// WithEntityID adds the entityID to the 5ade52fad5b81c209ce2e5a6 params
func (o *Nr5ade52fad5b81c209ce2e5a6Params) WithEntityID(entityID string) *Nr5ade52fad5b81c209ce2e5a6Params {
	o.SetEntityID(entityID)
	return o
}

// SetEntityID adds the entityId to the 5ade52fad5b81c209ce2e5a6 params
func (o *Nr5ade52fad5b81c209ce2e5a6Params) SetEntityID(entityID string) {
	o.EntityID = entityID
}

// WithVersionID adds the versionID to the 5ade52fad5b81c209ce2e5a6 params
func (o *Nr5ade52fad5b81c209ce2e5a6Params) WithVersionID(versionID string) *Nr5ade52fad5b81c209ce2e5a6Params {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the 5ade52fad5b81c209ce2e5a6 params
func (o *Nr5ade52fad5b81c209ce2e5a6Params) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *Nr5ade52fad5b81c209ce2e5a6Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	// path param entityId
	if err := r.SetPathParam("entityId", o.EntityID); err != nil {
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
