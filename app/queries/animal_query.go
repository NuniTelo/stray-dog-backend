package queries

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"stray-dogs/app/models"

	"github.com/jmoiron/sqlx"
)

/*
This will be used to call the database
*/
type AnimalQueries struct {
	*sqlx.DB
}

type Images []struct {
	URL string `json:"url"`
}

/*
Query to get a list of animals using pagination.
*/
func (q *AnimalQueries) GetAnimals(page int, limit int) ([]models.GetAnimalListStruct, error) {
	offset := (page - 1) * limit
	animals := []models.GetAnimalListStruct{}
	query := `
	SELECT 
		a.id,
		a.origin,
		a.gender,
		a.age,
		a.chipped,
		a.chip_number,
		a.chip_date,
		a.chip_position,
		a.parvo_vaccine,
		a.distemper_vaccine,
		a.rabies_vaccine,
		a.polyvalent_vaccine,
		a.in_shelter,
		a.sterilization_date,
		a.breed,
		a.is_alive,
		a.death_date,
		a.death_cause,
		a.chip_date,
		JSON_AGG(JSON_BUILD_OBJECT('url', ai.url)) AS images 
	FROM "public"."animal" a
	LEFT JOIN "public"."animal_images" ai on ai.animal_id = a.id
	GROUP BY a.id 
	ORDER BY a.id DESC
	LIMIT $1 
	OFFSET $2`
	err := q.Select(&animals, query, limit, offset)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return animals, nil
}

/*
Query to get a list of animals created by user.
*/
func (q *AnimalQueries) GetAnimalsCreatedByUser(page int, limit int, userId *int8) ([]models.GetAnimalListStruct, error) {
	offset := (page - 1) * limit
	animals := []models.GetAnimalListStruct{}
	query := `
	SELECT 
		a.id,
		a.origin,
		a.gender,
		a.age,
		a.chip_position,
		a.polyvalent_vaccine,
		a.chip_date,
		JSON_AGG(JSON_BUILD_OBJECT('url', ai.url)) AS images 
	FROM "public"."animal" a 
	LEFT JOIN "public"."animal_images" ai on ai.animal_id = a.id 
	WHERE a.user_id = $3 
	ORDER BY id DESC 
	LIMIT $1 
	OFFSET $2`
	err := q.Select(&animals, query, limit, offset, userId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return animals, nil
}

/*
Insert a new animal to the database.
*/
func (q *AnimalQueries) CreateAnimalQuery(a *models.Animal, id *int8) error {
	query := `INSERT INTO "public"."animal" 
		(
			"origin",
			"gender", 
			"age", 
			"chipped", 
			"chip_number", 
			"chip_date", 
			"chip_position", 
			"parvo_vaccine", 
			"distemper_vaccine", 
			"polyvalent_vaccine", 
			"rabies_vaccine", 
			"in_shelter", 
			"sterilization_date", 
			"breed", 
			"is_alive", 
			"death_date", 
			"death_cause",
			"user_id" 
		) 
			VALUES 
			(
				$1,
				$2, 
				$3, 
				$4, 
				$5, 
				$6, 
				$7, 
				$8, 
				$9, 
				$10,
				$11, 
				$12, 
				$13, 
				$14, 
				$15, 
				$16, 
				$17,
				$18
			) RETURNING id`

	var animalId int8

	err := q.QueryRow(query,
		a.Origin,
		a.Gender,
		a.Age,
		a.Chipped,
		a.ChipNumber,
		a.ChipDate,
		a.ChipPosition,
		a.ParvoVaccine,
		a.DistemperVaccine,
		a.PolyvalentVaccine,
		a.RabiesVaccine,
		a.InShelter,
		a.SterilizationDate,
		a.Breed,
		a.IsAlive,
		a.DeathDate,
		a.DeathCause,
		id,
	).Scan(&animalId)

	if err != nil {
		return err
	}

	if a.Images != nil {
		tx := q.MustBegin()
		for _, url := range *a.Images {
			tx.MustExec(`INSERT INTO "public"."animal_images" ("animal_id", "url") VALUES ($1, $2)`, animalId, url)
		}
		tx.Commit()
	}
	return nil
}

/*
Update the animal by using id.
*/
func (q *AnimalQueries) UpdateAnimal(id int, a *models.Animal, userId *int8) error {
	animalExist := models.AnimalCheckIfExists{}

	queryAnimalExists := `
		SELECT 
			a.id,
			JSON_AGG(JSON_BUILD_OBJECT('url', ai.url)) AS images 
		FROM "public"."animal" a 
		LEFT JOIN "public"."animal_images" ai on ai.animal_id = a.id 
		WHERE a.id=$1 AND a.user_id=$2
		GROUP BY a.id 
		LIMIT 1`

	animalResults := q.Get(&animalExist, queryAnimalExists, id, userId)

	switch animalResults {
	case sql.ErrNoRows:
		return errors.New("this animal does not exist")
	}

	var imageDifferences []string

	if a.Images != nil {
		var images Images
		if err := json.Unmarshal(animalExist.Images, &images); err != nil {
			fmt.Println(err)
			return err
		}
		var imagesArray []string
		for _, v := range images {
			imagesArray = append(imagesArray, v.URL)
		}
		imageDifferences = difference(*a.Images, imagesArray)
	}

	query := `
	UPDATE "public"."animal"
		SET    
			origin = $2,
			gender = $3,
			age = $4,
			chipped = $5,
			chip_number = $6,
			chip_date = $7,
			chip_position = $8,
			parvo_vaccine = $9,
			distemper_vaccine = $10,
			polyvalent_vaccine = $11,
			rabies_vaccine = $12,
			in_shelter = $13,
			sterilization_date = $14,
			breed = $15,
			is_alive = $16,
			death_date = $17,
			death_cause = $18 
		WHERE id = $1`

	_, err := q.Exec(query, id,
		a.Origin,
		a.Gender,
		a.Age,
		a.Chipped,
		a.ChipNumber,
		a.ChipDate,
		a.ChipPosition,
		a.ParvoVaccine,
		a.DistemperVaccine,
		a.PolyvalentVaccine,
		a.RabiesVaccine,
		a.InShelter,
		a.SterilizationDate,
		a.Breed,
		a.IsAlive,
		a.DeathDate,
		a.DeathCause)
	if err != nil {
		return err
	}

	if len(imageDifferences) > 0 {
		tx := q.MustBegin()
		for _, url := range imageDifferences {
			tx.MustExec(`INSERT INTO "public"."animal_images" ("animal_id", "url") VALUES ($1, $2)`, id, url)
		}
		tx.Commit()
	}

	return nil
}

/*
Delete the animal by using id.
*/
func (q *AnimalQueries) DeleteAnimal(id int, userId *int8) error {
	animalExist := models.AnimalCheckIfExists{}

	queryAnimalExists := `SELECT id FROM "public"."animal" WHERE id=$1 AND user_id=$2 LIMIT 1`
	animalResults := q.Get(&animalExist, queryAnimalExists, id, userId)

	switch animalResults {
	case sql.ErrNoRows:
		return errors.New("no animal found")
	}

	queryDeletePhotos := `DELETE FROM "public"."animal_images" WHERE animal_id = $1`
	_, errPhotoDelete := q.Exec(queryDeletePhotos, id)
	if errPhotoDelete != nil {
		return errPhotoDelete
	}

	query := `DELETE FROM "public"."animal" WHERE id = $1`
	_, err := q.Exec(query, id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
