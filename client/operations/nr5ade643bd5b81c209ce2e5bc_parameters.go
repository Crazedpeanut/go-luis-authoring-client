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

// NewNr5ade643bd5b81c209ce2e5bcParams creates a new Nr5ade643bd5b81c209ce2e5bcParams object
// with the default values initialized.
func NewNr5ade643bd5b81c209ce2e5bcParams() *Nr5ade643bd5b81c209ce2e5bcParams {
	var ()
	return &Nr5ade643bd5b81c209ce2e5bcParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNr5ade643bd5b81c209ce2e5bcParamsWithTimeout creates a new Nr5ade643bd5b81c209ce2e5bcParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr5ade643bd5b81c209ce2e5bcParamsWithTimeout(timeout time.Duration) *Nr5ade643bd5b81c209ce2e5bcParams {
	var ()
	return &Nr5ade643bd5b81c209ce2e5bcParams{

		timeout: timeout,
	}
}

// NewNr5ade643bd5b81c209ce2e5bcParamsWithContext creates a new Nr5ade643bd5b81c209ce2e5bcParams object
// with the default values initialized, and the ability to set a context for a request
func NewNr5ade643bd5b81c209ce2e5bcParamsWithContext(ctx context.Context) *Nr5ade643bd5b81c209ce2e5bcParams {
	var ()
	return &Nr5ade643bd5b81c209ce2e5bcParams{

		Context: ctx,
	}
}

// NewNr5ade643bd5b81c209ce2e5bcParamsWithHTTPClient creates a new Nr5ade643bd5b81c209ce2e5bcParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr5ade643bd5b81c209ce2e5bcParamsWithHTTPClient(client *http.Client) *Nr5ade643bd5b81c209ce2e5bcParams {
	var ()
	return &Nr5ade643bd5b81c209ce2e5bcParams{
		HTTPClient: client,
	}
}

/*Nr5ade643bd5b81c209ce2e5bcParams contains all the parameters to send to the API endpoint
for the 5ade643bd5b81c209ce2e5bc operation typically these are written to a http.Request
*/
type Nr5ade643bd5b81c209ce2e5bcParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*EntityID
	  Format - guid. The Pattern.Any entity ID.

	*/
	EntityID string
	/*RoleID
	  Format - guid. The role ID.

	*/
	RoleID string
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the 5ade643bd5b81c209ce2e5bc params
func (o *Nr5ade643bd5b81c209ce2e5bcParams) WithTimeout(timeout time.Duration) *Nr5ade643bd5b81c209ce2e5bcParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 5ade643bd5b81c209ce2e5bc params
func (o *Nr5ade643bd5b81c209ce2e5bcParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 5ade643bd5b81c209ce2e5bc params
func (o *Nr5ade643bd5b81c209ce2e5bcParams) WithContext(ctx context.Context) *Nr5ade643bd5b81c209ce2e5bcParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 5ade643bd5b81c209ce2e5bc params
func (o *Nr5ade643bd5b81c209ce2e5bcParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 5ade643bd5b81c209ce2e5bc params
func (o *Nr5ade643bd5b81c209ce2e5bcParams) WithHTTPClient(client *http.Client) *Nr5ade643bd5b81c209ce2e5bcParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 5ade643bd5b81c209ce2e5bc params
func (o *Nr5ade643bd5b81c209ce2e5bcParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the 5ade643bd5b81c209ce2e5bc params
func (o *Nr5ade643bd5b81c209ce2e5bcParams) WithAppID(appID string) *Nr5ade643bd5b81c209ce2e5bcParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the 5ade643bd5b81c209ce2e5bc params
func (o *Nr5ade643bd5b81c209ce2e5bcParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithEntityID adds the entityID to the 5ade643bd5b81c209ce2e5bc params
func (o *Nr5ade643bd5b81c209ce2e5bcParams) WithEntityID(entityID string) *Nr5ade643bd5b81c209ce2e5bcParams {
	o.SetEntityID(entityID)
	return o
}

// SetEntityID adds the entityId to the 5ade643bd5b81c209ce2e5bc params
func (o *Nr5ade643bd5b81c209ce2e5bcParams) SetEntityID(entityID string) {
	o.EntityID = entityID
}

// WithRoleID adds the roleID to the 5ade643bd5b81c209ce2e5bc params
func (o *Nr5ade643bd5b81c209ce2e5bcParams) WithRoleID(roleID string) *Nr5ade643bd5b81c209ce2e5bcParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the 5ade643bd5b81c209ce2e5bc params
func (o *Nr5ade643bd5b81c209ce2e5bcParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WithVersionID adds the versionID to the 5ade643bd5b81c209ce2e5bc params
func (o *Nr5ade643bd5b81c209ce2e5bcParams) WithVersionID(versionID string) *Nr5ade643bd5b81c209ce2e5bcParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the 5ade643bd5b81c209ce2e5bc params
func (o *Nr5ade643bd5b81c209ce2e5bcParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *Nr5ade643bd5b81c209ce2e5bcParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param roleId
	if err := r.SetPathParam("roleId", o.RoleID); err != nil {
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
