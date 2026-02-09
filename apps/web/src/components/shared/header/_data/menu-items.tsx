import {
  Home,
  Search,
  TrendingUp,
  Bell,
  Settings,
  HelpCircle,
} from "lucide-react";
import { useRouter } from "next/navigation";

export const getMenuItems = () => {
  const router = useRouter();
  return [
    { icon: Home, label: "Home", action: () => router.push("/") },
    {
      icon: Search,
      label: "Compare Vehicles",
      action: () => router.push("/#compare-section"),
    },
    {
      icon: TrendingUp,
      label: "Portfolio",
      action: () => router.push("/portfolio"),
    },
    {
      icon: Bell,
      label: "Price Alerts",
      action: () => router.push("/#price-alerts"),
    },
    {
      icon: Settings,
      label: "Account Settings",
      action: () => router.push("/account"),
    },
    {
      icon: HelpCircle,
      label: "Help & Support",
      action: () => router.push("/help"),
    },
  ];
};
