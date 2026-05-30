import { env } from "@ws-tanstack/env/server";
import cors from "cors";
import type { UUID } from "crypto";
import express from "express";

const app = express();

app.use(
  cors({
    origin: env.CORS_ORIGIN,
    methods: ["GET", "POST", "OPTIONS"],
  }),
);

app.use(express.json());

type AuthKey = UUID & { __brand: "auth-key" };
type Username = string & { __brand: "username" };

const AUTHORIZED_USERS_TO_KEYS = new Map<Username, AuthKey>();
const AUTHORIZED_KEYS_TO_USERS = new Map<AuthKey, Username>();

app.post("/auth", (req, res) => {
  const userName = req.body.userName as Username;

  if (!userName) {
    res.status(400).json({
      message: "Please provide username to authenticate."
    })
    return;
  }

  const authKey = AUTHORIZED_USERS_TO_KEYS.get(userName) ?? crypto.randomUUID() as AuthKey;

  AUTHORIZED_USERS_TO_KEYS.set(userName, authKey);
  AUTHORIZED_KEYS_TO_USERS.set(authKey, userName);

  res.status(200).json({
    userName,
    key: authKey,
  })
});

app.get("/me", (req, res) => {
  const authKey = req.headers.authorization?.replace("Bearer ", "") as AuthKey | undefined;

  if (!authKey || !AUTHORIZED_KEYS_TO_USERS.get(authKey)) {
    res.status(401).json({
      message: "Unauthorized. Please authenticate."
    })
    return;
  }

  const userName = AUTHORIZED_KEYS_TO_USERS.get(authKey);

  res.status(200).json({
    userName,
  })
});


app.listen(3000, () => {
  console.log("Server is running on http://localhost:3000");
});
