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

// NewNr5ade5b5dd5b81c209ce2e5b4Params creates a new Nr5ade5b5dd5b81c209ce2e5b4Params object
// with the default values initialized.
func NewNr5ade5b5dd5b81c209ce2e5b4Params() *Nr5ade5b5dd5b81c209ce2e5b4Params {
	var ()
	return &Nr5ade5b5dd5b81c209ce2e5b4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewNr5ade5b5dd5b81c209ce2e5b4ParamsWithTimeout creates a new Nr5ade5b5dd5b81c209ce2e5b4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr5ade5b5dd5b81c209ce2e5b4ParamsWithTimeout(timeout time.Duration) *Nr5ade5b5dd5b81c209ce2e5b4Params {
	var ()
	return &Nr5ade5b5dd5b81c209ce2e5b4Params{

		timeout: timeout,
	}
}

// NewNr5ade5b5dd5b81c209ce2e5b4ParamsWithContext creates a new Nr5ade5b5dd5b81c209ce2e5b4Params object
// with the default values initialized, and the ability to set a context for a request
func NewNr5ade5b5dd5b81c209ce2e5b4ParamsWithContext(ctx context.Context) *Nr5ade5b5dd5b81c209ce2e5b4Params {
	var ()
	return &Nr5ade5b5dd5b81c209ce2e5b4Params{

		Context: ctx,
	}
}

// NewNr5ade5b5dd5b81c209ce2e5b4ParamsWithHTTPClient creates a new Nr5ade5b5dd5b81c209ce2e5b4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr5ade5b5dd5b81c209ce2e5b4ParamsWithHTTPClient(client *http.Client) *Nr5ade5b5dd5b81c209ce2e5b4Params {
	var ()
	return &Nr5ade5b5dd5b81c209ce2e5b4Params{
		HTTPClient: client,
	}
}

/*Nr5ade5b5dd5b81c209ce2e5b4Params contains all the parameters to send to the API endpoint
for the 5ade5b5dd5b81c209ce2e5b4 operation typically these are written to a http.Request
*/
type Nr5ade5b5dd5b81c209ce2e5b4Params struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*Body
	  A json object containing the new role name.

	*/
	Body interface{}
	/*EntityID
	  Format - guid. The closed list entity ID.

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

// WithTimeout adds the timeout to the 5ade5b5dd5b81c209ce2e5b4 params
func (o *Nr5ade5b5dd5b81c209ce2e5b4Params) WithTimeout(timeout time.Duration) *Nr5ade5b5dd5b81c209ce2e5b4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 5ade5b5dd5b81c209ce2e5b4 params
func (o *Nr5ade5b5dd5b81c209ce2e5b4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 5ade5b5dd5b81c209ce2e5b4 params
func (o *Nr5ade5b5dd5b81c209ce2e5b4Params) WithContext(ctx context.Context) *Nr5ade5b5dd5b81c209ce2e5b4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 5ade5b5dd5b81c209ce2e5b4 params
func (o *Nr5ade5b5dd5b81c209ce2e5b4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 5ade5b5dd5b81c209ce2e5b4 params
func (o *Nr5ade5b5dd5b81c209ce2e5b4Params) WithHTTPClient(client *http.Client) *Nr5ade5b5dd5b81c209ce2e5b4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 5ade5b5dd5b81c209ce2e5b4 params
func (o *Nr5ade5b5dd5b81c209ce2e5b4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the 5ade5b5dd5b81c209ce2e5b4 params
func (o *Nr5ade5b5dd5b81c209ce2e5b4Params) WithAppID(appID string) *Nr5ade5b5dd5b81c209ce2e5b4Params {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the 5ade5b5dd5b81c209ce2e5b4 params
func (o *Nr5ade5b5dd5b81c209ce2e5b4Params) SetAppID(appID string) {
	o.AppID = appID
}

// WithBody adds the body to the 5ade5b5dd5b81c209ce2e5b4 params
func (o *Nr5ade5b5dd5b81c209ce2e5b4Params) WithBody(body interface{}) *Nr5ade5b5dd5b81c209ce2e5b4Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the 5ade5b5dd5b81c209ce2e5b4 params
func (o *Nr5ade5b5dd5b81c209ce2e5b4Params) SetBody(body interface{}) {
	o.Body = body
}

// WithEntityID adds the entityID to the 5ade5b5dd5b81c209ce2e5b4 params
func (o *Nr5ade5b5dd5b81c209ce2e5b4Params) WithEntityID(entityID string) *Nr5ade5b5dd5b81c209ce2e5b4Params {
	o.SetEntityID(entityID)
	return o
}

// SetEntityID adds the entityId to the 5ade5b5dd5b81c209ce2e5b4 params
func (o *Nr5ade5b5dd5b81c209ce2e5b4Params) SetEntityID(entityID string) {
	o.EntityID = entityID
}

// WithRoleID adds the roleID to the 5ade5b5dd5b81c209ce2e5b4 params
func (o *Nr5ade5b5dd5b81c209ce2e5b4Params) WithRoleID(roleID string) *Nr5ade5b5dd5b81c209ce2e5b4Params {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the 5ade5b5dd5b81c209ce2e5b4 params
func (o *Nr5ade5b5dd5b81c209ce2e5b4Params) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WithVersionID adds the versionID to the 5ade5b5dd5b81c209ce2e5b4 params
func (o *Nr5ade5b5dd5b81c209ce2e5b4Params) WithVersionID(versionID string) *Nr5ade5b5dd5b81c209ce2e5b4Params {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the 5ade5b5dd5b81c209ce2e5b4 params
func (o *Nr5ade5b5dd5b81c209ce2e5b4Params) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *Nr5ade5b5dd5b81c209ce2e5b4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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