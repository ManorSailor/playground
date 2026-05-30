import type { UUID } from "node:crypto";

type Branded<T, TBrand> = T & { __brand: TBrand };

type AuthKey = Branded<UUID, "auth-key">;
type Username = Branded<string, "username">;

export type { AuthKey, Branded, Username };
