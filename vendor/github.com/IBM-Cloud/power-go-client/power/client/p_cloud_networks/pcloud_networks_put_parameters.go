// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

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

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudNetworksPutParams creates a new PcloudNetworksPutParams object
// with the default values initialized.
func NewPcloudNetworksPutParams() *PcloudNetworksPutParams {
	var ()
	return &PcloudNetworksPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudNetworksPutParamsWithTimeout creates a new PcloudNetworksPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudNetworksPutParamsWithTimeout(timeout time.Duration) *PcloudNetworksPutParams {
	var ()
	return &PcloudNetworksPutParams{

		timeout: timeout,
	}
}

// NewPcloudNetworksPutParamsWithContext creates a new PcloudNetworksPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudNetworksPutParamsWithContext(ctx context.Context) *PcloudNetworksPutParams {
	var ()
	return &PcloudNetworksPutParams{

		Context: ctx,
	}
}

// NewPcloudNetworksPutParamsWithHTTPClient creates a new PcloudNetworksPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudNetworksPutParamsWithHTTPClient(client *http.Client) *PcloudNetworksPutParams {
	var ()
	return &PcloudNetworksPutParams{
		HTTPClient: client,
	}
}

/*PcloudNetworksPutParams contains all the parameters to send to the API endpoint
for the pcloud networks put operation typically these are written to a http.Request
*/
type PcloudNetworksPutParams struct {

	/*Body
	  Parameters to update a Network

	*/
	Body *models.NetworkUpdate
	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud networks put params
func (o *PcloudNetworksPutParams) WithTimeout(timeout time.Duration) *PcloudNetworksPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud networks put params
func (o *PcloudNetworksPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud networks put params
func (o *PcloudNetworksPutParams) WithContext(ctx context.Context) *PcloudNetworksPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud networks put params
func (o *PcloudNetworksPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud networks put params
func (o *PcloudNetworksPutParams) WithHTTPClient(client *http.Client) *PcloudNetworksPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud networks put params
func (o *PcloudNetworksPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud networks put params
func (o *PcloudNetworksPutParams) WithBody(body *models.NetworkUpdate) *PcloudNetworksPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud networks put params
func (o *PcloudNetworksPutParams) SetBody(body *models.NetworkUpdate) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud networks put params
func (o *PcloudNetworksPutParams) WithCloudInstanceID(cloudInstanceID string) *PcloudNetworksPutParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud networks put params
func (o *PcloudNetworksPutParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithNetworkID adds the networkID to the pcloud networks put params
func (o *PcloudNetworksPutParams) WithNetworkID(networkID string) *PcloudNetworksPutParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the pcloud networks put params
func (o *PcloudNetworksPutParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudNetworksPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
