import { type UseQueryOptions, useQuery } from "@tanstack/react-query";
import { env } from "@ws-tanstack/env/web";

type Username = string;

type MeResponse = {
  username: Username;
};

async function getMyInfo(): Promise<MeResponse> {
  const res = await fetch(`${env.VITE_SERVER_URL}/me`, {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  });

  if (!res.ok) {
    const error = await res.json();

    throw new Error(error.message);
  }

  return res.json();
}

function useMe(
  config?: Omit<UseQueryOptions<unknown, Error, MeResponse>, "queryKey">,
) {
  return useQuery({
    ...config,
    queryKey: ["user"],
    queryFn: getMyInfo,
    retry: false,
  });
}

export { useMe };
