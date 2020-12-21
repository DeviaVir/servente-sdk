package sdk

import (
	"time"
)

// JSONResponse the data struct used to render API responses
type JSONResponse struct {
	Error     bool
	Timestamp time.Time
	Data      *JSONData
}

// JSONData represents the useful data inside the response
type JSONData struct {
	Message       string     `json:",omitempty"`
	Teams         []JSONTeam `json:",omitempty"`
	NextPageToken string     `json:",omitempty"`
	APIVersions   []JSONAPI  `json:",omitempty"`
}

// JSONTeam only required parts here are Name & ID, the rest is extra data to
// enrich a user's experience
type JSONTeam struct {
	Name         string
	ID           string
	Email        string `json:",omitempty"`
	MembersCount int64  `json:",omitempty"`
}

// JSONAPI helps Servente figure out which API versions are compatible
type JSONAPI struct {
	ID string
}
