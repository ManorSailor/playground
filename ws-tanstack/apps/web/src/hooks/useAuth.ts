import { type UseMutationOptions, useMutation } from "@tanstack/react-query";
import { env } from "@ws-tanstack/env/web";

type Username = string;

type AuthResponse = {
  username: Username;
  key: string;
};

async function authenticateUser(username: string): Promise<AuthResponse> {
  return fetch(`${env.VITE_SERVER_URL}/auth`, {
    method: "POST",
    body: JSON.stringify({ username }),
    headers: {
      "Content-Type": "application/json",
    },
  }).then((res) => res.json());
}

function useAuth(config?: UseMutationOptions<AuthResponse, Error, Username>) {
  return useMutation({
    ...config,
    mutationFn: authenticateUser,
    onSuccess(data, ...rest) {
      localStorage.setItem("token", data.key);
      config?.onSuccess?.(data, ...rest);
    },
    retry: false,
  });
}

export { useAuth };
