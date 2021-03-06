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

// NewGetApplicationVersionKeyParams creates a new GetApplicationVersionKeyParams object
// with the default values initialized.
func NewGetApplicationVersionKeyParams() *GetApplicationVersionKeyParams {
	var ()
	return &GetApplicationVersionKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetApplicationVersionKeyParamsWithTimeout creates a new GetApplicationVersionKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetApplicationVersionKeyParamsWithTimeout(timeout time.Duration) *GetApplicationVersionKeyParams {
	var ()
	return &GetApplicationVersionKeyParams{

		timeout: timeout,
	}
}

// NewGetApplicationVersionKeyParamsWithContext creates a new GetApplicationVersionKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetApplicationVersionKeyParamsWithContext(ctx context.Context) *GetApplicationVersionKeyParams {
	var ()
	return &GetApplicationVersionKeyParams{

		Context: ctx,
	}
}

// NewGetApplicationVersionKeyParamsWithHTTPClient creates a new GetApplicationVersionKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetApplicationVersionKeyParamsWithHTTPClient(client *http.Client) *GetApplicationVersionKeyParams {
	var ()
	return &GetApplicationVersionKeyParams{
		HTTPClient: client,
	}
}

/*GetApplicationVersionKeyParams contains all the parameters to send to the API endpoint
for the get application version key operation typically these are written to a http.Request
*/
type GetApplicationVersionKeyParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*VersionID
	  The application version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get application version key params
func (o *GetApplicationVersionKeyParams) WithTimeout(timeout time.Duration) *GetApplicationVersionKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get application version key params
func (o *GetApplicationVersionKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get application version key params
func (o *GetApplicationVersionKeyParams) WithContext(ctx context.Context) *GetApplicationVersionKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get application version key params
func (o *GetApplicationVersionKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get application version key params
func (o *GetApplicationVersionKeyParams) WithHTTPClient(client *http.Client) *GetApplicationVersionKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get application version key params
func (o *GetApplicationVersionKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get application version key params
func (o *GetApplicationVersionKeyParams) WithAppID(appID string) *GetApplicationVersionKeyParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get application version key params
func (o *GetApplicationVersionKeyParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithVersionID adds the versionID to the get application version key params
func (o *GetApplicationVersionKeyParams) WithVersionID(versionID string) *GetApplicationVersionKeyParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the get application version key params
func (o *GetApplicationVersionKeyParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetApplicationVersionKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
