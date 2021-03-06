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

// NewVersionsGetApplicationVersionSettingsParams creates a new VersionsGetApplicationVersionSettingsParams object
// with the default values initialized.
func NewVersionsGetApplicationVersionSettingsParams() *VersionsGetApplicationVersionSettingsParams {
	var ()
	return &VersionsGetApplicationVersionSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVersionsGetApplicationVersionSettingsParamsWithTimeout creates a new VersionsGetApplicationVersionSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVersionsGetApplicationVersionSettingsParamsWithTimeout(timeout time.Duration) *VersionsGetApplicationVersionSettingsParams {
	var ()
	return &VersionsGetApplicationVersionSettingsParams{

		timeout: timeout,
	}
}

// NewVersionsGetApplicationVersionSettingsParamsWithContext creates a new VersionsGetApplicationVersionSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewVersionsGetApplicationVersionSettingsParamsWithContext(ctx context.Context) *VersionsGetApplicationVersionSettingsParams {
	var ()
	return &VersionsGetApplicationVersionSettingsParams{

		Context: ctx,
	}
}

// NewVersionsGetApplicationVersionSettingsParamsWithHTTPClient creates a new VersionsGetApplicationVersionSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVersionsGetApplicationVersionSettingsParamsWithHTTPClient(client *http.Client) *VersionsGetApplicationVersionSettingsParams {
	var ()
	return &VersionsGetApplicationVersionSettingsParams{
		HTTPClient: client,
	}
}

/*VersionsGetApplicationVersionSettingsParams contains all the parameters to send to the API endpoint
for the versions get application version settings operation typically these are written to a http.Request
*/
type VersionsGetApplicationVersionSettingsParams struct {

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

// WithTimeout adds the timeout to the versions get application version settings params
func (o *VersionsGetApplicationVersionSettingsParams) WithTimeout(timeout time.Duration) *VersionsGetApplicationVersionSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the versions get application version settings params
func (o *VersionsGetApplicationVersionSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the versions get application version settings params
func (o *VersionsGetApplicationVersionSettingsParams) WithContext(ctx context.Context) *VersionsGetApplicationVersionSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the versions get application version settings params
func (o *VersionsGetApplicationVersionSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the versions get application version settings params
func (o *VersionsGetApplicationVersionSettingsParams) WithHTTPClient(client *http.Client) *VersionsGetApplicationVersionSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the versions get application version settings params
func (o *VersionsGetApplicationVersionSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the versions get application version settings params
func (o *VersionsGetApplicationVersionSettingsParams) WithAppID(appID string) *VersionsGetApplicationVersionSettingsParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the versions get application version settings params
func (o *VersionsGetApplicationVersionSettingsParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithVersionID adds the versionID to the versions get application version settings params
func (o *VersionsGetApplicationVersionSettingsParams) WithVersionID(versionID string) *VersionsGetApplicationVersionSettingsParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the versions get application version settings params
func (o *VersionsGetApplicationVersionSettingsParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *VersionsGetApplicationVersionSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
