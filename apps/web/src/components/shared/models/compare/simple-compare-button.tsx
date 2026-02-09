import { useOptionToggleStore } from "./_store/option-toggle.store";
import { Loader2 } from "lucide-react";
import { useAdvanceOptionStore } from "./_store/advance-option.store";
import { Button } from "@/components/ui/button";

export default function SimpleCompareButton() {
  const { showAdvanced } = useOptionToggleStore();
  const { isSearching, handleSaveAndCompare } = useAdvanceOptionStore();

  return (
    !showAdvanced && (
      <div className="flex justify-center">
        <Button
          onClick={handleSaveAndCompare}
          disabled={isSearching}
          size="lg"
          className="bg-white/20 backdrop-blur-sm border border-white/20 hover:bg-white/30 text-white font-medium rounded-2xl px-8 py-3 text-lg"
        >
          {isSearching ? (
            <>
              <Loader2 className="mr-2 h-5 w-5 animate-spin" />
              Analyzing Vehicle...
            </>
          ) : (
            "Compare with CARMA AI"
          )}
        </Button>
      </div>
    )
  );
}
