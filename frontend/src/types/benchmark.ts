import { z } from "zod";

export const BenchmarkSchema = z.object({
  id: z.string(),
  name: z.string(),
  geographicLocation: z.string(),
  totalProjectCostP90: z.number(),
  totalConstructionCostPerLaneKm: z.number(),
  cubicMetreRateForEarthworksPerM3: z.number(),
  squareMetreRateForPavementPerBridgePerM2: z.number()
});

export type Benchmark = z.infer<typeof BenchmarkSchema>;
