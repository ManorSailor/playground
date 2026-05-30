The goal of this project is to understand how:
- TanStack Query w/ WebSocket work alongside normal REST APIs

__Requirements:__
- Server must expose 1 WS endpoint - `/notifications` & 2 REST endpoints - `/auth` & `/me`.
- Client must allow user to authenticate, show their name & latest notifications as they arrive.

__Server Flow:__
- Client must hit `/auth` first to get a key with their `username`.
- Client must then `/me` with `key` in `Authorization` header to ensure they are authenticated.
- Client must then send an upgrade request to `ws://<ip>:<port>/notifications?token={key}` with `token` set to received key.

__Client Implementation:__
- Implement 3 buttons:
    - Authenticate - hits `/auth`, gets `key`, & sets to `localStorage`
    - Me - retrieves `key` from `localStorage`, hits `/me`, 
    - Notifications
- Hitting `me` or `notifications` should without `key` must print the _raw_ server response as-is.
- Hitting `notifications` with correct `key` should display latest notifications as they arrive.
