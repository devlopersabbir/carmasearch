"use client";
import { LinkIcon } from "lucide-react";
import { Label } from "../ui/label";
import { Input } from "../ui/input";
import { useAdvanceOptionStore } from "@/components/shared/models/compare/_store/advance-option.store";

export default function UrlInput() {
  const { vehicleUrl, isSearching, setVehicleUrl } = useAdvanceOptionStore();
  return (
    <div className="space-y-4">
      <Label
        htmlFor="vehicle-url"
        className="text-white/90 text-sm font-medium"
      >
        Vehicle Listing URL
      </Label>
      <div className="relative">
        <LinkIcon className="absolute left-4 top-1/2 transform -translate-y-1/2 h-4 w-4 text-white/40" />
        <Input
          id="vehicle-url"
          type="url"
          placeholder="Paste AutoScout24 vehicle listing URL (e.g., https://www.autoscout24.de/angebote/...)"
          value={vehicleUrl}
          onChange={(e) => setVehicleUrl(e.target.value)}
          className="pl-12 bg-black/40 backdrop-blur-sm border border-white/10 rounded-2xl h-14 text-white placeholder:text-white/40 focus:border-white/30 focus:ring-0"
          disabled={isSearching}
        />
      </div>
    </div>
  );
}
