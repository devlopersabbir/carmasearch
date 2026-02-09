"use client";
import { useAdvanceOptionStore } from "./_store/advance-option.store";

export default function DisplayError() {
  const { error } = useAdvanceOptionStore();
  return (
    error && (
      <div className="bg-red-500/20 border border-red-500/30 rounded-lg p-4">
        <div className="flex items-center">
          <div className="text-red-400 text-sm font-medium">{error}</div>
        </div>
      </div>
    )
  );
}
