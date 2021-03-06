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

// NewImportVersionJSONParams creates a new ImportVersionJSONParams object
// with the default values initialized.
func NewImportVersionJSONParams() *ImportVersionJSONParams {
	var ()
	return &ImportVersionJSONParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImportVersionJSONParamsWithTimeout creates a new ImportVersionJSONParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImportVersionJSONParamsWithTimeout(timeout time.Duration) *ImportVersionJSONParams {
	var ()
	return &ImportVersionJSONParams{

		timeout: timeout,
	}
}

// NewImportVersionJSONParamsWithContext creates a new ImportVersionJSONParams object
// with the default values initialized, and the ability to set a context for a request
func NewImportVersionJSONParamsWithContext(ctx context.Context) *ImportVersionJSONParams {
	var ()
	return &ImportVersionJSONParams{

		Context: ctx,
	}
}

// NewImportVersionJSONParamsWithHTTPClient creates a new ImportVersionJSONParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImportVersionJSONParamsWithHTTPClient(client *http.Client) *ImportVersionJSONParams {
	var ()
	return &ImportVersionJSONParams{
		HTTPClient: client,
	}
}

/*ImportVersionJSONParams contains all the parameters to send to the API endpoint
for the import version Json operation typically these are written to a http.Request
*/
type ImportVersionJSONParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*JSONApp
	  A JSON representing the LUIS application structure.

	*/
	JSONApp *models.JSONApp
	/*VersionID
	  The imported versionId.

	*/
	VersionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the import version Json params
func (o *ImportVersionJSONParams) WithTimeout(timeout time.Duration) *ImportVersionJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import version Json params
func (o *ImportVersionJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import version Json params
func (o *ImportVersionJSONParams) WithContext(ctx context.Context) *ImportVersionJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import version Json params
func (o *ImportVersionJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import version Json params
func (o *ImportVersionJSONParams) WithHTTPClient(client *http.Client) *ImportVersionJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import version Json params
func (o *ImportVersionJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the import version Json params
func (o *ImportVersionJSONParams) WithAppID(appID string) *ImportVersionJSONParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the import version Json params
func (o *ImportVersionJSONParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithJSONApp adds the jSONApp to the import version Json params
func (o *ImportVersionJSONParams) WithJSONApp(jSONApp *models.JSONApp) *ImportVersionJSONParams {
	o.SetJSONApp(jSONApp)
	return o
}

// SetJSONApp adds the jSONApp to the import version Json params
func (o *ImportVersionJSONParams) SetJSONApp(jSONApp *models.JSONApp) {
	o.JSONApp = jSONApp
}

// WithVersionID adds the versionID to the import version Json params
func (o *ImportVersionJSONParams) WithVersionID(versionID *string) *ImportVersionJSONParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the import version Json params
func (o *ImportVersionJSONParams) SetVersionID(versionID *string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *ImportVersionJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if o.JSONApp != nil {
		if err := r.SetBodyParam(o.JSONApp); err != nil {
			return err
		}
	}

	if o.VersionID != nil {

		// query param versionId
		var qrVersionID string
		if o.VersionID != nil {
			qrVersionID = *o.VersionID
		}
		qVersionID := qrVersionID
		if qVersionID != "" {
			if err := r.SetQueryParam("versionId", qVersionID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
