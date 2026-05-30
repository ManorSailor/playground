import "ws";
import type { Username } from "./types";

declare module "ws" {
	interface WebSocket {
		isAlive: boolean;
		username: Username;
	}
}
