import { z } from "zod";

export const UserSchema = z.object({
  id: z.string(),
  type: z.union([z.literal("member"), z.literal("client")]),
  fullName: z.string(),
  email: z.string(),
  password: z.string(),
  avatar: z.string(),
  bio: z.string(),
  organisation: z.string(),
  clientRole: z.string(),
  lastAccess: z.coerce.date().optional()
});

export type User = z.infer<typeof UserSchema>;

export const EmptyUser: User = {
  id: "",
  type: "member",
  fullName: "",
  email: "",
  password: "",
  avatar: "",
  bio: "",
  organisation: "",
  clientRole: "",
  lastAccess: undefined
};
