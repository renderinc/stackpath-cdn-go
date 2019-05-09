// Code generated by go-swagger; DO NOT EDIT.

package cdn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/renderinc/stackpath-cdn-go/models"
)

// GetCDNIpsReader is a Reader for the GetCDNIps structure.
type GetCDNIpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCDNIpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCDNIpsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetCDNIpsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCDNIpsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetCDNIpsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCDNIpsOK creates a GetCDNIpsOK with default headers values
func NewGetCDNIpsOK() *GetCDNIpsOK {
	return &GetCDNIpsOK{}
}

/*GetCDNIpsOK handles this case with default header values.

GetCDNIpsOK get c d n ips o k
*/
type GetCDNIpsOK struct {
	Payload *models.CdnGetCDNIpsResponse
}

func (o *GetCDNIpsOK) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/ips][%d] getCDNIpsOK  %+v", 200, o.Payload)
}

func (o *GetCDNIpsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CdnGetCDNIpsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCDNIpsUnauthorized creates a GetCDNIpsUnauthorized with default headers values
func NewGetCDNIpsUnauthorized() *GetCDNIpsUnauthorized {
	return &GetCDNIpsUnauthorized{}
}

/*GetCDNIpsUnauthorized handles this case with default header values.

Returned when an unauthorized request is attempted.
*/
type GetCDNIpsUnauthorized struct {
	Payload *models.APIStatus
}

func (o *GetCDNIpsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/ips][%d] getCDNIpsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCDNIpsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCDNIpsInternalServerError creates a GetCDNIpsInternalServerError with default headers values
func NewGetCDNIpsInternalServerError() *GetCDNIpsInternalServerError {
	return &GetCDNIpsInternalServerError{}
}

/*GetCDNIpsInternalServerError handles this case with default header values.

Internal server error.
*/
type GetCDNIpsInternalServerError struct {
	Payload *models.APIStatus
}

func (o *GetCDNIpsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/ips][%d] getCDNIpsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCDNIpsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCDNIpsDefault creates a GetCDNIpsDefault with default headers values
func NewGetCDNIpsDefault(code int) *GetCDNIpsDefault {
	return &GetCDNIpsDefault{
		_statusCode: code,
	}
}

/*GetCDNIpsDefault handles this case with default header values.

Default error structure.
*/
type GetCDNIpsDefault struct {
	_statusCode int

	Payload *models.APIStatus
}

// Code gets the status code for the get c d n ips default response
func (o *GetCDNIpsDefault) Code() int {
	return o._statusCode
}

func (o *GetCDNIpsDefault) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/ips][%d] GetCDNIPs default  %+v", o._statusCode, o.Payload)
}

func (o *GetCDNIpsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
