package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type Datum struct {
	ID         uuid.UUID `json:"id" db:"id"`
	SeaLevelFt float32   `json:"sealevelft" db:"sealevelft"`
	WindMPH    float32   `json:"windmph" db:"windmph"`
	WindDIR    string    `json:"winddir" db:"winddir"`
	IsFlooding bool      `json:"isflooding" db:"isflooding"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (d Datum) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Data is not required by pop and may be deleted
type Data []Datum

// String is not required by pop and may be deleted
func (d Data) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *Datum) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *Datum) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *Datum) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
