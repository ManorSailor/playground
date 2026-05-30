import { Link } from "@tanstack/react-router";

import type { PropsWithChildren } from "react";
import { ModeToggle } from "./mode-toggle";

export default function Header({ children }: PropsWithChildren) {
  const links = [{ to: "/", label: "Home" }] as const;

  return (
    <div>
      <div className="flex flex-row items-center justify-between px-2 py-1">
        <nav className="flex gap-4 text-lg">
          {links.map(({ to, label }) => {
            return (
              <Link key={to} to={to}>
                {label}
              </Link>
            );
          })}
        </nav>

        <div className="flex items-center gap-2">
          {children}
          <ModeToggle />
        </div>
      </div>
      <hr />
    </div>
  );
}
