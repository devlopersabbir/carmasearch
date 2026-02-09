"use client";

import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Target } from "lucide-react";
import { useCompareModelStore } from "@/components/providers/models/compare-model.store";
import { VisuallyHidden } from "radix-ui";
import UrlInput from "@/components/compares/url-input";
import OptionToggle from "./option-toggle";
import AdvanceOptions from "./advance-options";
import SimpleCompareButton from "./simple-compare-button";
import DisplayError from "./display-error";
import CompareSearchResult from "./compare-search-result";

export default function CompareModal() {
  const { isCompareModalOpen, setIsCompareModalOpen } = useCompareModelStore();

  return (
    <Dialog open={isCompareModalOpen} onOpenChange={setIsCompareModalOpen}>
      <DialogContent
        className="sm:max-w-6xl max-h-[95vh] overflow-y-auto bg-black/90 backdrop-blur-xl border border-white/10"
        title="Compare"
      >
        <DialogHeader className="border-b border-white/10 pb-6">
          <VisuallyHidden.Root>
            <DialogTitle className="text-2xl font-bold text-white flex items-center gap-3">
              <Target className="h-6 w-6 text-primary" />
              Vehicle Comparison
            </DialogTitle>
          </VisuallyHidden.Root>
        </DialogHeader>

        <div className="space-y-8">
          {/* URL Input Section */}
          <UrlInput />

          {/* Advanced Options Toggle */}
          <OptionToggle />

          {/* Advanced Options */}
          <AdvanceOptions />

          {/* Simple Compare Button */}
          <SimpleCompareButton />

          {/* Error Display */}
          <DisplayError />

          {/* Search Results */}
          <CompareSearchResult />
        </div>
      </DialogContent>
    </Dialog>
  );
}
