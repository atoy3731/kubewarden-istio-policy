// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// EventSource EventSource contains information for an event.
//
// swagger:model EventSource
type EventSource struct {

	// Component from which the event is generated.
	Component string `json:"component,omitempty"`

	// Node name on which the event is generated.
	Host string `json:"host,omitempty"`
}
