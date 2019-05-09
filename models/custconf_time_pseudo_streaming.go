// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfTimePseudoStreaming The HTTP pseudo-streaming policy enables Flash based video players to support seeking to random locations within a MP4 or FLV file without having to download the entire video. Flash players such as Flowplayer and JWPlayer can be configured to send a query string parameter that indicates to the server the time offset the user advanced the video to. Typically a query string parameter called "start" is used by these players.
// swagger:model custconfTimePseudoStreaming
type CustconfTimePseudoStreaming struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// This is used by the API to perform conflict checking
	ID string `json:"id,omitempty"`

	// The name of the query string parameter that indicates to the CDN the end time of the video that should be returned.
	JumpToTimeEndParam string `json:"jumpToTimeEndParam,omitempty"`

	// The name of the query string parameter that indicates to the CDN the specific time interval of the video that is being requested.
	JumpToTimeStartParam string `json:"jumpToTimeStartParam,omitempty"`
}

// Validate validates this custconf time pseudo streaming
func (m *CustconfTimePseudoStreaming) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfTimePseudoStreaming) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfTimePseudoStreaming) UnmarshalBinary(b []byte) error {
	var res CustconfTimePseudoStreaming
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
