package data

import (
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Member struct {
	ID       string `json:"id"`
	FullName string `json:"fullName,omitempty"`
	Email    string `json:"email,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Bio      string `json:"bio,omitempty"`
}
type ProjectTeam struct {
	ProjectLead Member   `json:"projectLead" form:"projectLead" binding:"required"`
	TeamMembers []Member `json:"teamMembers" form:"teamMembers" binding:"required"`
}

type ScopeOfEngagement struct {
	Quantification              bool               `json:"quantification" form:"quantification"`
	CostEstimation              bool               `json:"costEstimation" form:"costEstimation"`
	ProbabilisticRiskAssessment bool               `json:"probabilisticRiskAssessment" form:"probabilisticRiskAssessment"`
	BasisOfEstimateReport       bool               `json:"basisOfEstimateReport" form:"basisOfEstimateReport"`
	NumberOfMilestones          int                `json:"numberOfMilestones" form:"numberOfMilestones" binding:"required"`
	EstimatedCompletionDate     primitive.DateTime `json:"estimatedCompletionDate" form:"estimatedCompletionDate" binding:"required"`
	RemainsAccessibleForNDays   int                `json:"remainsAccessibleForNDays" form:"remainsAccessibleForNDays" binding:"required"`
}

type ProgressToDate struct {
	Quantification              int `json:"quantification"`
	CostEstimation              int `json:"costEstimation"`
	ProbabilisticRiskAssessment int `json:"probabilisticRiskAssessment"`
	BasisOfEstimateReport       int `json:"basisOfEstimateReport"`
	NumberOfMilestones          int `json:"numberOfMilestones"`
}

type AnticipatedCompletionDate struct {
	Quantification              primitive.DateTime `json:"quantification"`
	CostEstimation              primitive.DateTime `json:"costEstimation"`
	ProbabilisticRiskAssessment primitive.DateTime `json:"probabilisticRiskAssessment"`
	BasisOfEstimateReport       primitive.DateTime `json:"basisOfEstimateReport"`
}

type CommercialInformation struct {
	CIContractedValue float64 `json:"ciContractedValue"`
	CIAccrualToDate   float64 `json:"ciAccrualToDate"`
	// CIRemainingFee          int `json:"ciRemainingFee"`
	ApprovedVariationToDate float64 `json:"approvedVariationToDate"`
}

type OptionOutturnCost struct {
	P90  float64 `json:"p90"`
	P50  float64 `json:"p50"`
	Base float64 `json:"base"`
}

type TotalProjectCostPerMilestone struct {
	LevelOfDesign      string             `json:"levelOfDesign"`
	Date               primitive.DateTime `json:"date"`
	BaseValue          float64            `json:"baseValue"`
	P90OutturnCost     float64            `json:"p90OutturnCost"`
	P90RiskContingency float64            `json:"p90RiskContingency"`
	P50OutturnCost     float64            `json:"p50OutturnCost"`
	P50RiskContingency float64            `json:"p50RiskContingency"`
	CurrentMilstone    bool               `json:"currentMilstone"`
}

type KeyCostDriver struct {
	Driver string  `json:"driver"`
	Cost   float64 `json:"cost"`
}

type KeyRisk struct {
	Description string `json:"description"`
	Score       string `json:"score"`
	Trend       string `json:"trend"`
}

type DesignPackage struct {
	Description string   `json:"description"`
	Milestones  []string `json:"milestones"`
}

type DesignPackages struct {
	LevelOfDesign DesignPackage   `json:"levelOfDesign"`
	Packages      []DesignPackage `json:"packages"`
}

type Package struct {
	Description   string `json:"description"`
	Progress      int    `json:"progress"`
	SecondStageQA bool   `json:"secondStageQA"`
	FinalQAReview bool   `json:"finalQAReview"`
	Submitted     bool   `json:"submitted"`
}

type ProjectBenchmark struct {
	BenchmarkID        primitive.ObjectID `json:"benchmarkId"`
	DisplayProjectName bool               `json:"displayProjectName"`
}

type Benchmarking struct {
	EnableGeographicLocation                 bool               `json:"enableGeographicLocation"`
	EnableTotalConstructionCost              bool               `json:"enableTotalConstructionCost"`
	EnableSquareMetreRateForPavement         bool               `json:"enableSquareMetreRateForPavement"`
	EnableTotalProjectCost                   bool               `json:"enableTotalProjectCost"`
	EnableCubicMetreRateForEarthworks        bool               `json:"enableCubicMetreRateForEarthworks"`
	GeographicLocation                       string             `json:"geographicLocation" form:"geographicLocation" binding:"required"`
	TotalProjectCostP90                      float64            `json:"totalProjectCostP90" form:"totalProjectCostP90" binding:"required"`
	TotalConstructionCostPerLaneKm           float64            `json:"totalConstructionCostPerLaneKm" form:"totalConstructionCostPerLaneKm" binding:"required"`
	CubicMetreRateForEarthworksPerM3         float64            `json:"cubicMetreRateForEarthworksPerM3" form:"cubicMetreRateForEarthworksPerM3" binding:"required"`
	SquareMetreRateForPavementPerBridgePerM2 float64            `json:"squareMetreRateForPavementPerBridgePerM2" form:"squareMetreRateForPavementPerBridgePerM2" binding:"required"`
	Benchmarks                               []ProjectBenchmark `json:"benchmarks"`
}

type Metrics struct {
	ProgressToDate               ProgressToDate                 `json:"progressToDate"`
	AnticipatedCompletionDate    AnticipatedCompletionDate      `json:"anticipatedCompletionDate"`
	CommercialInformation        CommercialInformation          `json:"commercialInformation"`
	OptionOutturnCosts           []OptionOutturnCost            `json:"optionOutturnCosts"`
	TotalProjectCostPerMilestone []TotalProjectCostPerMilestone `json:"totalProjectCostPerMilestone"`
	KeyCostDriversBaseValue      float64                        `json:"keyCostDriversBaseValue"`
	KeyCostDrivers               []KeyCostDriver                `json:"keyCostDrivers"`
	KeyRisks                     []KeyRisk                      `json:"keyRisks"`
	ValueManagementOpportunities []string                       `json:"valueManagementOpportunities"`
	DesignPackages               *DesignPackages                `json:"designPackages"`
	Packages                     [][]Package                    `json:"packages"`
	Benchmarking                 Benchmarking                   `json:"benchmarking"`
}

func (m *Metrics) Validate() (*ValidationErrorMap, error) {
	errors := make(ValidationErrorMap)

	for i, item := range m.Benchmarking.Benchmarks {
		if item.BenchmarkID.IsZero() {
			errors[fmt.Sprintf("benchmark-%d", i)] = "benchmark is required"
		}
	}

	if len(errors) > 0 {
		return &errors, fmt.Errorf("validation error")
	}

	return nil, nil
}

type Project struct {
	ID                   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name                 string             `json:"name" form:"name" binding:"required"`
	Client               string             `json:"client" form:"client" binding:"required"`
	Region               string             `json:"region" form:"region" binding:"required"`
	CIProjectNumber      string             `json:"ciProjectNumber" form:"ciProjectNumber" binding:"required"`
	ClientProjectNumber  string             `json:"clientProjectNumber" form:"clientProjectNumber" binding:"required"`
	ClientRepresentative struct {
		ID         string `json:"id"`
		FullName   string `json:"fullName,omitempty"`
		ClientRole string `json:"clientRole,omitempty"`
	} `json:"clientRepresentative" form:"clientRepresentative" binding:"required"`
	Team           ProjectTeam        `json:"team"`
	Scope          ScopeOfEngagement  `json:"scope"`
	StartDate      primitive.DateTime `json:"startDate" form:"startDate"`
	CompletionDate primitive.DateTime `json:"completionDate" form:"completionDate"`
	Status         string             `json:"status" form:"status"`
	Metrics        Metrics            `json:"metrics"`
}

func (p *Project) Validate() (*ValidationErrorMap, error) {
	errors := make(ValidationErrorMap)

	if strings.TrimSpace(p.Name) == "" {
		errors["name"] = "project name is required"
	}

	if strings.TrimSpace(p.Client) == "" {
		errors["client"] = "client is required"
	}

	if strings.TrimSpace(p.Region) == "" {
		errors["region"] = "region is required"
	}

	if strings.TrimSpace(p.CIProjectNumber) == "" {
		errors["ciProjectNumber"] = "CI project number is required"
	}

	if strings.TrimSpace(p.ClientProjectNumber) == "" {
		errors["clientProjectNumber"] = "client project number is required"
	}

	if strings.TrimSpace(p.ClientRepresentative.ID) == "" {
		errors["clientRepresentative"] = "client representative is required"
	}

	if strings.TrimSpace(p.Team.ProjectLead.ID) == "" {
		errors["projectLead"] = "project lead is required"
	}

	if p.Scope.RemainsAccessibleForNDays == 0 {
		errors["remainsAccessibleForNDays"] = "remains accessible for n days is required"
	}

	if p.Scope.EstimatedCompletionDate == 0 {
		errors["estimatedCompletionDate"] = "estimated completion date is required"
	}

	if p.StartDate == 0 {
		errors["startDate"] = "start date is required"
	}

	fmt.Println("dates", p.StartDate, p.CompletionDate, p.Scope.EstimatedCompletionDate)

	if len(errors) > 0 {
		return &errors, fmt.Errorf("validation error")
	}

	return nil, nil
}
