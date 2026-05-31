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
import { type Notification, useNotifications } from "@/hooks/useNotifications";

function Notifications() {
  const { data: notifications, isLoading } = useNotifications();

  return (
    <Card className="mx-auto max-w-1/2 rounded-xl">
      <CardHeader>
        <CardTitle>Notifications</CardTitle>

        <CardAction>
          <Button size="xs" variant="outline">
            Unread
          </Button>
        </CardAction>
      </CardHeader>

      <CardContent className="scrollbar-thin max-h-72 overflow-y-auto">
        <NotificationList notifications={notifications} />
      </CardContent>
    </Card>
  );
}

type NotificationListProps = {
  notifications?: Notification[];
};

function NotificationList({ notifications }: NotificationListProps) {
  if (!notifications?.length) {
    return (
      <Item>
        <ItemContent className="items-center">
          <ItemTitle>So Empty.</ItemTitle>
        </ItemContent>
      </Item>
    );
  }

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

export { Notifications };
