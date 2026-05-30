import { createFileRoute } from "@tanstack/react-router";
import { Button } from "@ws-tanstack/ui/components/button";
import { useState } from "react";
import { AuthNControls } from "@/components/authn-controls";
import Header from "@/components/header";
import { useMe } from "@/hooks/useMe";

export const Route = createFileRoute("/")({
  component: HomeComponent,
});

function HomeComponent() {
  const [showNotifications, setShowNotifications] = useState(false);
  const { data: user, isSuccess, error } = useMe({ enabled: false });

  return (
    <div className="h-svh">
      <Header>
        <AuthNControls />
        <Button variant="outline" onClick={() => setShowNotifications(true)}>
          Show Notifications
        </Button>
      </Header>

      <div className="container mx-auto max-w-3xl px-4 py-2">
        <section className="grid gap-1 rounded-lg border p-4">
          <p className="flex gap-2">
            <span>Username:</span>
            <span>{user?.username || "Guest"}</span>
          </p>

          <p className="flex gap-2">
            <span>Authenticated:</span>
            <span>{isSuccess ? "Yes" : "No"}</span>
          </p>

          <p className="flex gap-2">
            <span>WebSocket Status:</span>
            <span>Offline</span>
          </p>

          {error && <p className="text-destructive italic">{error.message}</p>}
        </section>
      </div>
    </div>
  );
}
