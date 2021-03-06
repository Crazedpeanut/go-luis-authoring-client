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

	"github.com/crazedpeanut/go-luis-authoring-client/models"
)

// NewUpdateApplicationParams creates a new UpdateApplicationParams object
// with the default values initialized.
func NewUpdateApplicationParams() *UpdateApplicationParams {
	var ()
	return &UpdateApplicationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateApplicationParamsWithTimeout creates a new UpdateApplicationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateApplicationParamsWithTimeout(timeout time.Duration) *UpdateApplicationParams {
	var ()
	return &UpdateApplicationParams{

		timeout: timeout,
	}
}

// NewUpdateApplicationParamsWithContext creates a new UpdateApplicationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateApplicationParamsWithContext(ctx context.Context) *UpdateApplicationParams {
	var ()
	return &UpdateApplicationParams{

		Context: ctx,
	}
}

// NewUpdateApplicationParamsWithHTTPClient creates a new UpdateApplicationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateApplicationParamsWithHTTPClient(client *http.Client) *UpdateApplicationParams {
	var ()
	return &UpdateApplicationParams{
		HTTPClient: client,
	}
}

/*UpdateApplicationParams contains all the parameters to send to the API endpoint
for the update application operation typically these are written to a http.Request
*/
type UpdateApplicationParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*ApplicationUpdateObject
	  A JSON object containing Name and Description of the application.

	*/
	ApplicationUpdateObject *models.ApplicationUpdateObject

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update application params
func (o *UpdateApplicationParams) WithTimeout(timeout time.Duration) *UpdateApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update application params
func (o *UpdateApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update application params
func (o *UpdateApplicationParams) WithContext(ctx context.Context) *UpdateApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update application params
func (o *UpdateApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update application params
func (o *UpdateApplicationParams) WithHTTPClient(client *http.Client) *UpdateApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update application params
func (o *UpdateApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the update application params
func (o *UpdateApplicationParams) WithAppID(appID string) *UpdateApplicationParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the update application params
func (o *UpdateApplicationParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithApplicationUpdateObject adds the applicationUpdateObject to the update application params
func (o *UpdateApplicationParams) WithApplicationUpdateObject(applicationUpdateObject *models.ApplicationUpdateObject) *UpdateApplicationParams {
	o.SetApplicationUpdateObject(applicationUpdateObject)
	return o
}

// SetApplicationUpdateObject adds the applicationUpdateObject to the update application params
func (o *UpdateApplicationParams) SetApplicationUpdateObject(applicationUpdateObject *models.ApplicationUpdateObject) {
	o.ApplicationUpdateObject = applicationUpdateObject
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if o.ApplicationUpdateObject != nil {
		if err := r.SetBodyParam(o.ApplicationUpdateObject); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
