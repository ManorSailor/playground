import { Button } from "@ws-tanstack/ui/components/button";
import { useAuth } from "@/hooks/useAuth";
import { useMe } from "@/hooks/useMe";

function AuthNControls() {
  const { mutate: authenticateUser, isPending: isAuthPending } = useAuth();

  const { refetch: fetchUser, isFetching: isMePending } = useMe({
    enabled: false,
  });

  const handleAuthenticate = () => authenticateUser("Sars");

  const handleMe = () => fetchUser();

  return (
    <>
      <Button
        variant="outline"
        onClick={handleAuthenticate}
        disabled={isAuthPending}
      >
        Authenticate
      </Button>

      <Button variant="outline" onClick={handleMe} disabled={isMePending}>
        Me
      </Button>
    </>
  );
}

export { AuthNControls };
