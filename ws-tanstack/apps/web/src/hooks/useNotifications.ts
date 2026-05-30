import { type UseQueryOptions, useQuery } from "@tanstack/react-query";
import { env } from "@ws-tanstack/env/web";

type NotificationResponse = {
  id: string;
  message: string;
  receivedAt: number;
};

async function getNotifications(): Promise<NotificationResponse[]> {
  const res = await fetch(`${env.VITE_SERVER_URL}/notifications`, {
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

function useNotifications(
  config?: Omit<
    UseQueryOptions<unknown, Error, NotificationResponse>,
    "queryKey"
  >,
) {
  return useQuery({
    ...config,
    queryKey: ["notifications"],
    queryFn: getNotifications,
    retry: false,
  });
}

export { useNotifications };
