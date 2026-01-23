import { Hono } from "hono";
import { serve } from "@hono/node-server";
import { askUserQuery, type UserPrompt } from "./main.js";

const app = new Hono();

app.post("/ask", async (c) => {
  const body = await c.req.json<UserPrompt>();
  const payload = await askUserQuery(body.message);
  return c.json(payload);
});

serve(
  {
    fetch: app.fetch,
    port: +process.env.SERVER_PORT!,
  },
  (info) => {
    console.log(`Server is running on http://localhost:${info.port}`);
  },
);
