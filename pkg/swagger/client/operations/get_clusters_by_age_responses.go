package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/validate"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/drycc/workflow-manager/pkg/swagger/models"
)

// GetClustersByAgeReader is a Reader for the GetClustersByAge structure.
type GetClustersByAgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetClustersByAgeReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClustersByAgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetClustersByAgeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetClustersByAgeOK creates a GetClustersByAgeOK with default headers values
func NewGetClustersByAgeOK() *GetClustersByAgeOK {
	return &GetClustersByAgeOK{}
}

/*GetClustersByAgeOK handles this case with default header values.

clusters details response
*/
type GetClustersByAgeOK struct {
	Payload GetClustersByAgeOKBodyBody
}

func (o *GetClustersByAgeOK) Error() string {
	return fmt.Sprintf("[GET /v3/clusters/age][%d] getClustersByAgeOK  %+v", 200, o.Payload)
}

func (o *GetClustersByAgeOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersByAgeDefault creates a GetClustersByAgeDefault with default headers values
func NewGetClustersByAgeDefault(code int) *GetClustersByAgeDefault {
	return &GetClustersByAgeDefault{
		_statusCode: code,
	}
}

/*GetClustersByAgeDefault handles this case with default header values.

unexpected error
*/
type GetClustersByAgeDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get clusters by age default response
func (o *GetClustersByAgeDefault) Code() int {
	return o._statusCode
}

func (o *GetClustersByAgeDefault) Error() string {
	return fmt.Sprintf("[GET /v3/clusters/age][%d] getClustersByAge default  %+v", o._statusCode, o.Payload)
}

func (o *GetClustersByAgeDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetClustersByAgeOKBodyBody get clusters by age o k body body

swagger:model GetClustersByAgeOKBodyBody
*/
type GetClustersByAgeOKBodyBody struct {

	/* data

	Required: true
	*/
	Data []*models.Cluster `json:"data"`
}

// Validate validates this get clusters by age o k body body
func (o *GetClustersByAgeOKBodyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetClustersByAgeOKBodyBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getClustersByAgeOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if err := o.Data[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
