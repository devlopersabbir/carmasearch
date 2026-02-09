import { Button } from "@/components/ui/button";
import { useOptionToggleStore } from "./_store/option-toggle.store";
import { ChevronDown } from "lucide-react";

export default function OptionToggle() {
  const { showAdvanced, setShowAdvanced } = useOptionToggleStore();
  return (
    <div className="flex items-center justify-between">
      <Button
        variant="ghost"
        onClick={() => setShowAdvanced(!showAdvanced)}
        className="text-sm text-white/60 hover:text-white p-0"
      >
        Advanced Options
        <ChevronDown
          className={`ml-1 h-4 w-4 transition-transform ${showAdvanced ? "rotate-180" : ""}`}
        />
      </Button>
    </div>
  );
}
