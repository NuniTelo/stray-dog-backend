package models

import (
	"encoding/json"
	"time"
)

type Animal struct {
	Id                *int8     `db:"id" json:"id"`
	Origin            string    `db:"origin" json:"origin"`
	Gender            string    `db:"gender" json:"gender"`
	Age               int64     `db:"age" json:"age"`
	Chipped           bool      `db:"chipped" json:"chipped"`
	ChipNumber        *string   `db:"chip_number" json:"chip_number"`
	ChipDate          *string   `db:"chip_date" json:"chip_date"`
	ChipPosition      *string   `db:"chip_position" json:"chip_position"`
	ParvoVaccine      *string   `db:"parvo_vaccine" json:"parvo_vaccine"`
	DistemperVaccine  *string   `db:"distemper_vaccine" json:"distemper_vaccine"`
	PolyvalentVaccine *string   `db:"polyvalent_vaccine" json:"polyvalent_vaccine"`
	RabiesVaccine     *string   `db:"rabies_vaccine" json:"rabies_vaccine"`
	InShelter         bool      `db:"in_shelter" json:"in_shelter"`
	SterilizationDate *string   `db:"sterilization_date" json:"sterilization_date"`
	Breed             string    `db:"breed" json:"breed"`
	IsAlive           bool      `db:"is_alive" json:"is_alive"`
	DeathDate         *string   `db:"death_date" json:"death_date"`
	DeathCause        *string   `db:"death_cause" json:"death_cause"`
	Images            *[]string `json:"images"`
}

type GetAnimalListStruct struct {
	ID                int             `db:"id" json:"id"`
	Origin            string          `db:"origin" json:"origin"`
	Gender            string          `db:"gender" json:"gender"`
	Age               int64           `db:"age" json:"age"`
	ChipPosition      *string         `db:"chip_position" json:"chip_position"`
	PolyvalentVaccine *string         `db:"polyvalent_vaccine" json:"polyvalent_vaccine"`
	ChipDate          *time.Time      `db:"chip_date" json:"chip_date"`
	Name              *string         `db:"username" json:"username"`
	UserId            *int            `db:"user_id" json:"user_id"`
	Chipped           *bool           `db:"chipped" json:"chipped"`
	ChipNumber        *string         `db:"chip_number" json:"chip_number"`
	ParvoVaccine      *string         `db:"parvo_vaccine" json:"parvo_vaccine"`
	DistemperVaccine  *string         `db:"distemper_vaccine" json:"distemper_vaccine"`
	RabiesVaccine     *string         `db:"rabies_vaccine" json:"rabies_vaccine"`
	InShelter         bool            `db:"in_shelter" json:"in_shelter"`
	SterilizationDate *string         `db:"sterilization_date" json:"sterilization_date"`
	Breed             string          `db:"breed" json:"breed"`
	IsAlive           bool            `db:"is_alive" json:"is_alive"`
	DeathDate         *string         `db:"death_date" json:"death_date"`
	DeathCause        *string         `db:"death_cause" json:"death_cause"`
	Images            json.RawMessage `db:"images" json:"images"`
}

type AnimalCheckIfExists struct {
	ID     string          `db:"id" json:"id"`
	Images json.RawMessage `db:"images" json:"images"`
}
