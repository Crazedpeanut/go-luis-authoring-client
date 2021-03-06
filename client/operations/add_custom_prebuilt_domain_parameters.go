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

// NewAddCustomPrebuiltDomainParams creates a new AddCustomPrebuiltDomainParams object
// with the default values initialized.
func NewAddCustomPrebuiltDomainParams() *AddCustomPrebuiltDomainParams {
	var ()
	return &AddCustomPrebuiltDomainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddCustomPrebuiltDomainParamsWithTimeout creates a new AddCustomPrebuiltDomainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddCustomPrebuiltDomainParamsWithTimeout(timeout time.Duration) *AddCustomPrebuiltDomainParams {
	var ()
	return &AddCustomPrebuiltDomainParams{

		timeout: timeout,
	}
}

// NewAddCustomPrebuiltDomainParamsWithContext creates a new AddCustomPrebuiltDomainParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddCustomPrebuiltDomainParamsWithContext(ctx context.Context) *AddCustomPrebuiltDomainParams {
	var ()
	return &AddCustomPrebuiltDomainParams{

		Context: ctx,
	}
}

// NewAddCustomPrebuiltDomainParamsWithHTTPClient creates a new AddCustomPrebuiltDomainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddCustomPrebuiltDomainParamsWithHTTPClient(client *http.Client) *AddCustomPrebuiltDomainParams {
	var ()
	return &AddCustomPrebuiltDomainParams{
		HTTPClient: client,
	}
}

/*AddCustomPrebuiltDomainParams contains all the parameters to send to the API endpoint
for the add custom prebuilt domain operation typically these are written to a http.Request
*/
type AddCustomPrebuiltDomainParams struct {

	/*PrebuiltDomainCreateObject
	  A prebuilt domain create object containing the name and culture of the domain

	*/
	PrebuiltDomainCreateObject *models.PrebuiltDomainCreateObject

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add custom prebuilt domain params
func (o *AddCustomPrebuiltDomainParams) WithTimeout(timeout time.Duration) *AddCustomPrebuiltDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add custom prebuilt domain params
func (o *AddCustomPrebuiltDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add custom prebuilt domain params
func (o *AddCustomPrebuiltDomainParams) WithContext(ctx context.Context) *AddCustomPrebuiltDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add custom prebuilt domain params
func (o *AddCustomPrebuiltDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add custom prebuilt domain params
func (o *AddCustomPrebuiltDomainParams) WithHTTPClient(client *http.Client) *AddCustomPrebuiltDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add custom prebuilt domain params
func (o *AddCustomPrebuiltDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPrebuiltDomainCreateObject adds the prebuiltDomainCreateObject to the add custom prebuilt domain params
func (o *AddCustomPrebuiltDomainParams) WithPrebuiltDomainCreateObject(prebuiltDomainCreateObject *models.PrebuiltDomainCreateObject) *AddCustomPrebuiltDomainParams {
	o.SetPrebuiltDomainCreateObject(prebuiltDomainCreateObject)
	return o
}

// SetPrebuiltDomainCreateObject adds the prebuiltDomainCreateObject to the add custom prebuilt domain params
func (o *AddCustomPrebuiltDomainParams) SetPrebuiltDomainCreateObject(prebuiltDomainCreateObject *models.PrebuiltDomainCreateObject) {
	o.PrebuiltDomainCreateObject = prebuiltDomainCreateObject
}

// WriteToRequest writes these params to a swagger request
func (o *AddCustomPrebuiltDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PrebuiltDomainCreateObject != nil {
		if err := r.SetBodyParam(o.PrebuiltDomainCreateObject); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
