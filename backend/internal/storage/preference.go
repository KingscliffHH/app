package storage

import (
	"context"
	"fmt"

	"github.com/Infinities-ICT-Solutions/project-dashboard/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PreferenceStorage interface {
	GetOrganisations() ([]string, error)
	AddOrganisation(organisation string) error
}

type mongoPreferenceStorage struct {
	db *mongo.Database
}

type Preference struct {
	ID    string      `json:"id" bson:"_id,omitempty"`
	Value interface{} `json:"value" form:"value" binding:"required"`
}

func NewPreferenceStorage(db *mongo.Database) *mongoPreferenceStorage {
	return &mongoPreferenceStorage{db: db}
}

func (p *mongoPreferenceStorage) GetOrganisations() ([]string, error) {
	id := "organisations"
	preference := &Preference{}
	err := p.db.Collection("preferences").FindOne(context.TODO(), bson.M{"_id": id}).Decode(preference)

	if err == mongo.ErrNoDocuments {
		// Handle or ignore the error
		fmt.Println("No document was found matching the filter.")
		return []string{}, nil
	} else if err != nil {
		return nil, err
	}

	orgs := []string{}

	if value, ok := preference.Value.(primitive.A); ok {
		for _, v := range value {
			if str, ok := v.(string); ok {
				orgs = append(orgs, str)
			} else {
				fmt.Println("An element is not a string", v)
			}
		}
	}

	return orgs, nil
}

func (p *mongoPreferenceStorage) AddOrganisation(organisation string) error {
	organisations, err := p.GetOrganisations()

	if err != nil {
		return err
	}

	if utils.Contains(organisations, organisation) {
		return nil
	}

	organisations = append(organisations, organisation)

	preference := bson.M{
		"value": organisations,
	}

	upsert := true
	_, err = p.db.Collection("preferences").UpdateOne(
		context.TODO(),
		bson.M{"_id": "organisations"},
		bson.M{"$set": preference},
		&options.UpdateOptions{Upsert: &upsert},
	)

	if err != nil {
		return err
	}

	return nil
}
