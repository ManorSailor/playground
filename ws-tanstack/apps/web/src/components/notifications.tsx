import { Button } from "@ws-tanstack/ui/components/button";
import {
  Card,
  CardAction,
  CardContent,
  CardHeader,
  CardTitle,
} from "@ws-tanstack/ui/components/card";
import {
  Item,
  ItemContent,
  ItemDescription,
  ItemGroup,
  ItemTitle,
} from "@ws-tanstack/ui/components/item";
import { Skeleton } from "@ws-tanstack/ui/components/skeleton";
import { type Notification, useNotifications } from "@/hooks/useNotifications";

function Notifications() {
  const { data: notifications, isLoading } = useNotifications();

  return (
    <Card className="mx-auto max-w-3xl rounded-xl">
      <CardHeader>
        <CardTitle>Notifications</CardTitle>

        <CardAction>
          <Button size="xs" variant="outline">
            Unread
          </Button>
        </CardAction>
      </CardHeader>

      <CardContent className="scrollbar-thin max-h-72 overflow-y-auto">
        {isLoading && <NotificationSkeleton />}

        {!notifications?.length && !isLoading && <NoNotificationsYet />}

        {notifications?.length && (
          <NotificationList notifications={notifications} />
        )}
      </CardContent>
    </Card>
  );
}

type NotificationListProps = {
  notifications: Notification[];
};

function NotificationList({ notifications }: NotificationListProps) {
  const dateFormatter = new Intl.DateTimeFormat("en-US", {
    hour: "numeric",
    dayPeriod: "narrow",
  });

  return (
    <ItemGroup>
      {notifications.map((n) => (
        <Item key={n.id} variant="outline">
          <ItemContent>
            <ItemTitle>{n.message}</ItemTitle>
            <ItemDescription>
              {dateFormatter.format(n.receivedAt)}
            </ItemDescription>
          </ItemContent>
        </Item>
      ))}
    </ItemGroup>
  );
}

function NoNotificationsYet() {
  return (
    <Item>
      <ItemContent className="items-center">
        <ItemTitle>So Empty.</ItemTitle>
      </ItemContent>
    </Item>
  );
}

type NotificationSkeletonProps = {
  maxCount?: number;
};

function NotificationSkeleton({ maxCount = 4 }: NotificationSkeletonProps) {
  const skeletons = new Array<null>(maxCount).fill(null);

  return (
    <ItemGroup>
      {skeletons.map((_, idx) => (
        <Item key={idx} variant="outline">
          <ItemContent>
            <Skeleton className="h-3 w-48 rounded-lg bg-secondary" />
            <Skeleton className="h-3 w-24 rounded-lg bg-secondary" />
          </ItemContent>
        </Item>
      ))}
    </ItemGroup>
  );
}

export { Notifications };
