import { env } from "@ws-tanstack/env/server";
import cors from "cors";
import express from "express";
import { createServer } from "http";
import { WebSocketServer } from "ws";
import type { Username, AuthKey } from "./types/types";

const app = express();
const wss = new WebSocketServer({ noServer: true });
const server = createServer(app);

app.use(
  cors({
    origin: env.ALLOWED_ORIGINS,
    methods: ["GET", "POST", "OPTIONS"],
  }),
);

app.use(express.json());

const AUTHORIZED_USERS_TO_KEYS = new Map<Username, AuthKey>();
const AUTHORIZED_KEYS_TO_USERS = new Map<AuthKey, Username>();

app.post("/auth", (req, res) => {
  const username = req.body?.username as Username | undefined;

  if (!username) {
    res.status(400).json({
      message: "Please provide username to authenticate.",
    });
    return;
  }

  const authKey =
    AUTHORIZED_USERS_TO_KEYS.get(username) ?? (crypto.randomUUID() as AuthKey);

  AUTHORIZED_USERS_TO_KEYS.set(username, authKey);
  AUTHORIZED_KEYS_TO_USERS.set(authKey, username);

  res.status(200).json({
    username,
    key: authKey,
  });
});

app.get("/me", (req, res) => {
  const authKey = req.headers.authorization?.replace("Bearer ", "");

  if (!isValidAuthKey(authKey)) {
    res.status(401).json({
      message: "Unauthorized. Please authenticate.",
    });
    return;
  }

  const username = AUTHORIZED_KEYS_TO_USERS.get(authKey);

  res.status(200).json({
    username,
  });
});

server.on("upgrade", (req, socket, head) => {
  const origin = req.headers.origin;

  if (!isValidOrigin(origin)) {
    socket.write("HTTP/1.1 403 Forbidden\r\n\r\n");
    socket.destroy();
    return;
  }

  const { pathname, searchParams } = new URL(
    req.url ?? "localhost",
    "http://localhost",
  );

  const token = searchParams.get("token");

  if (!isValidAuthKey(token)) {
    socket.write("HTTP/1.1 401 Unauthorized. Please authenticate.\r\n\r\n");
    socket.destroy();
    return;
  }

  const username = AUTHORIZED_KEYS_TO_USERS.get(token);

  if (!username) {
    socket.destroy();
    return;
  }

  if (pathname === "/ws") {
    wss.handleUpgrade(req, socket, head, (ws) => {
      ws.username = username;
      wss.emit("connection", ws, req);
    });
  } else {
    socket.destroy();
  }
});

wss.on("connection", (client) => {
  client.isAlive = true;

  // Heartbeat. Response from client ensures it is still connected.
  client.on("pong", () => (client.isAlive = true));
});

server.listen(3000, () => {
  console.log("Server is running on http://localhost:3000");

  const intervalId = setupWSSPing(wss);

  process.once("SIGINT", () => shutdownWSS(wss, intervalId));
  process.once("SIGTERM", () => shutdownWSS(wss, intervalId));
});

function isValidOrigin(origin: string | undefined | null): boolean {
  return !!origin && env.ALLOWED_ORIGINS.includes(origin);
}

function isValidAuthKey(key: string | undefined | null): key is AuthKey {
  return !!key && AUTHORIZED_KEYS_TO_USERS.has(key as AuthKey);
}

function setupWSSPing(wss: WebSocketServer) {
  return setInterval(() => {
    for (const client of wss.clients) {
      if (!client.isAlive) {
        client.terminate();
        continue;
      }

      client.isAlive = false;
      client.ping();
    }
  }, env.PING_INTERVAL);
}

function shutdownWSS(wss: WebSocketServer, pingIntervalId: NodeJS.Timeout) {
  clearInterval(pingIntervalId);

  wss.close();

  wss.clients.forEach((c) => {
    c.close(1001, "WSS server shutting down...");

    // close() above waits for client to acknowledge the shutdown
    // here we force close the conn. if clients don't respond for whatever reason
    setTimeout(() => c.readyState === c.CLOSED || c.terminate(), 5_000);
  });
}
