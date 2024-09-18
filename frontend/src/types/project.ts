import { z } from "zod";

const MemberSchema = z.object({
  id: z.string(),
  fullName: z.string().optional(),
  email: z.string().optional(),
  avatar: z.string().optional(),
  bio: z.string().optional()
});

export type Member = z.infer<typeof MemberSchema>;

const ProjectTeamSchema = z.object({
  projectLead: MemberSchema,
  teamMembers: z.array(MemberSchema)
});
export type ProjectTeam = z.infer<typeof ProjectTeamSchema>;

const ScopeOfEngagementSchema = z.object({
  quantification: z.boolean(),
  costEstimation: z.boolean(),
  probabilisticRiskAssessment: z.boolean(),
  basisOfEstimateReport: z.boolean(),
  numberOfMilestones: z.number(),
  estimatedCompletionDate: z.coerce.date(),
  remainsAccessibleForNDays: z.number()
});
export type ScopeOfEngagement = z.infer<typeof ScopeOfEngagementSchema>;

const ProgressToDateSchema = z.object({
  quantification: z.number(),
  costEstimation: z.number(),
  probabilisticRiskAssessment: z.number(),
  basisOfEstimateReport: z.number(),
  numberOfMilestones: z.number()
});
export type ProgressToDate = z.infer<typeof ProgressToDateSchema>;

// Assuming primitive.DateTime will be treated as a string in ISO format for simplicity.
const AnticipatedCompletionDateSchema = z.object({
  quantification: z.coerce.date(),
  costEstimation: z.coerce.date(),
  probabilisticRiskAssessment: z.coerce.date(),
  basisOfEstimateReport: z.coerce.date()
});
export type AnticipatedCompletionDate = z.infer<
  typeof AnticipatedCompletionDateSchema
>;

const CommercialInformationSchema = z.object({
  ciContractedValue: z.number(),
  ciAccrualToDate: z.number(),
  approvedVariationToDate: z.number()
});
export type CommercialInformation = z.infer<typeof CommercialInformationSchema>;

const OptionOutturnCostSchema = z.object({
  p90: z.number(),
  p50: z.number(),
  base: z.number()
});
export type OptionOutturnCost = z.infer<typeof OptionOutturnCostSchema>;

const TotalProjectCostPerMilestoneSchema = z.object({
  levelOfDesign: z.string(),
  date: z.coerce.date(), // Assuming date as string in ISO format
  baseValue: z.number(),
  p90OutturnCost: z.number(),
  p90RiskContingency: z.number(),
  p50OutturnCost: z.number(),
  p50RiskContingency: z.number(),
  currentMilstone: z.boolean()
});
export type TotalProjectCostPerMilestone = z.infer<
  typeof TotalProjectCostPerMilestoneSchema
>;

const KeyCostDriverSchema = z.object({
  driver: z.string(),
  cost: z.number()
});
export type KeyCostDriver = z.infer<typeof KeyCostDriverSchema>;

const KeyRiskSchema = z.object({
  description: z.string(),
  score: z.string(),
  trend: z.string()
});
export type KeyRisk = z.infer<typeof KeyRiskSchema>;

const DesignPackageSchema = z.object({
  description: z.string(),
  milestones: z.array(z.string())
});
export type DesignPackage = z.infer<typeof DesignPackageSchema>;

const DesignPackagesSchema = z.object({
  levelOfDesign: DesignPackageSchema,
  packages: z.array(DesignPackageSchema)
});

export type DesignPackages = z.infer<typeof DesignPackagesSchema>;

const PackageSchema = z.object({
  description: z.string(),
  progress: z.number(),
  secondStageQA: z.boolean(),
  finalQAReview: z.boolean(),
  submitted: z.boolean()
});
export type Package = z.infer<typeof PackageSchema>;

const ProjectBenchmarkSchema = z.object({
  benchmarkId: z.string(),
  displayProjectName: z.boolean()
});

export type ProjectBenchmark = z.infer<typeof ProjectBenchmarkSchema>;

const Benchmarking = z.object({
  enableGeographicLocation: z.boolean(),
  enableTotalConstructionCost: z.boolean(),
  enableSquareMetreRateForPavement: z.boolean(),
  enableTotalProjectCost: z.boolean(),
  enableCubicMetreRateForEarthworks: z.boolean(),
  geographicLocation: z.string(),
  totalConstructionCostPerLaneKm: z.number(),
  squareMetreRateForPavementPerBridgePerM2: z.number(),
  totalProjectCostP90: z.number(),
  cubicMetreRateForEarthworksPerM3: z.number(),
  benchmarks: z.array(ProjectBenchmarkSchema).or(z.null())
});
export type Benchmarking = z.infer<typeof Benchmarking>;

const MetricsSchema = z.object({
  progressToDate: ProgressToDateSchema,
  anticipatedCompletionDate: AnticipatedCompletionDateSchema,
  commercialInformation: CommercialInformationSchema,
  optionOutturnCosts: z.array(OptionOutturnCostSchema).or(z.null()),
  totalProjectCostPerMilestone: z
    .array(TotalProjectCostPerMilestoneSchema)
    .or(z.null()),
  keyCostDriversBaseValue: z.number(),
  keyCostDrivers: z.array(KeyCostDriverSchema).or(z.null()),
  keyRisks: z.array(KeyRiskSchema).or(z.null()),
  valueManagementOpportunities: z.array(z.string()).or(z.null()),
  designPackages: DesignPackagesSchema.or(z.null()),
  packages: z.array(z.array(PackageSchema)).or(z.null()),
  benchmarking: Benchmarking
});
export type Metrics = z.infer<typeof MetricsSchema>;

// Client struct {
//   FullName string `json:"fullName" form:"fullName" binding:"required"`
//   Email    string `json:"email" form:"email" binding:"required"`
// } `json:"client" form:"client" binding:"required"`

// Assuming primitive.ObjectID is represented as a string.
export const ProjectSchema = z.object({
  id: z.string(),
  name: z.string(),
  client: z.string(),
  region: z.string(),
  ciProjectNumber: z.string(),
  clientProjectNumber: z.string(),
  clientRepresentative: z.object({
    id: z.string(),
    fullName: z.string().optional(),
    email: z.string().optional(),
    clientRole: z.string().optional()
  }),
  team: ProjectTeamSchema,
  scope: ScopeOfEngagementSchema,
  startDate: z.coerce.date(), // Assuming date as string in ISO format
  completionDate: z.coerce.date(), // Assuming date as string in ISO format
  status: z.string(),
  metrics: MetricsSchema
});
export type Project = z.infer<typeof ProjectSchema>;

export const EmptyProject: Project = {
  id: "",
  name: "",
  client: "",
  region: "",
  ciProjectNumber: "",
  clientProjectNumber: "",
  clientRepresentative: {
    id: ""
  },
  team: {
    projectLead: { id: "" },
    teamMembers: [{ id: "" }, { id: "" }]
  },
  scope: {
    quantification: false,
    costEstimation: false,
    probabilisticRiskAssessment: false,
    basisOfEstimateReport: false,
    numberOfMilestones: 0,
    estimatedCompletionDate: null as any,
    remainsAccessibleForNDays: 0
  },
  startDate: null as any,
  completionDate: null as any,
  status: "",
  metrics: {
    progressToDate: {
      quantification: 0,
      costEstimation: 0,
      probabilisticRiskAssessment: 0,
      basisOfEstimateReport: 0,
      numberOfMilestones: 0
    },
    anticipatedCompletionDate: {
      quantification: new Date(),
      costEstimation: new Date(),
      probabilisticRiskAssessment: new Date(),
      basisOfEstimateReport: new Date()
    },
    commercialInformation: {
      ciContractedValue: 0,
      ciAccrualToDate: 0,
      approvedVariationToDate: 0
    },
    optionOutturnCosts: null,
    totalProjectCostPerMilestone: null,
    keyCostDriversBaseValue: 0,
    keyCostDrivers: null,
    keyRisks: null,
    valueManagementOpportunities: null,
    designPackages: null,
    packages: null,
    benchmarking: {
      enableGeographicLocation: false,
      enableTotalConstructionCost: false,
      enableSquareMetreRateForPavement: false,
      enableTotalProjectCost: false,
      enableCubicMetreRateForEarthworks: false,
      geographicLocation: "",
      totalConstructionCostPerLaneKm: 0,
      squareMetreRateForPavementPerBridgePerM2: 0,
      totalProjectCostP90: 0,
      cubicMetreRateForEarthworksPerM3: 0,
      benchmarks: null
    }
  }
};
