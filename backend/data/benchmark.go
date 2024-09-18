package data

import (
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Benchmark struct {
	ID                                       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name                                     string             `json:"name" form:"name" binding:"required"`
	GeographicLocation                       string             `json:"geographicLocation" form:"geographicLocation" binding:"required"`
	TotalProjectCostP90                      float64            `json:"totalProjectCostP90" form:"totalProjectCostP90" binding:"required"`
	TotalConstructionCostPerLaneKm           float64            `json:"totalConstructionCostPerLaneKm" form:"totalConstructionCostPerLaneKm" binding:"required"`
	CubicMetreRateForEarthworksPerM3         float64            `json:"cubicMetreRateForEarthworksPerM3" form:"cubicMetreRateForEarthworksPerM3" binding:"required"`
	SquareMetreRateForPavementPerBridgePerM2 float64            `json:"squareMetreRateForPavementPerBridgePerM2" form:"squareMetreRateForPavementPerBridgePerM2" binding:"required"`
}

func (b *Benchmark) Validate() (*ValidationErrorMap, error) {
	errors := make(ValidationErrorMap)

	if strings.TrimSpace(b.Name) == "" {
		errors["name"] = "benchmark name is required"
	}

	if strings.TrimSpace(b.GeographicLocation) == "" {
		errors["geographicLocation"] = "geographic location is required"
	}

	if b.TotalProjectCostP90 < 0 {
		errors["totalProjectCostP90"] = "total project cost p90 must be greater or equals to 0"
	}

	if b.TotalConstructionCostPerLaneKm < 0 {
		errors["totalConstructionCostPerLaneKm"] = "total construction cost per lane km must be greater or equals to 0"
	}

	if b.CubicMetreRateForEarthworksPerM3 < 0 {
		errors["cubicMetreRateForEarthworksPerM3"] = "cubic metre rate for earthworks per m3 must be greater or equals to 0"
	}

	if b.SquareMetreRateForPavementPerBridgePerM2 < 0 {
		errors["squareMetreRateForPavementPerBridgePerM2"] = "square metre rate for pavement per bridge per m2 must be greater or equals to 0"
	}

	if len(errors) > 0 {
		return &errors, fmt.Errorf("validation error")
	}

	return nil, nil
}
