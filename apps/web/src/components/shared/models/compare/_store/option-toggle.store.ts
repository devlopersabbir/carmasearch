import { create } from "zustand";

interface OptionToggleState {
  showAdvanced: boolean;
  setShowAdvanced: (showAdvanced: boolean) => void;

  isDescriptionExpanded: boolean;
  setIsDescriptionExpanded: (isDescriptionExpanded: boolean) => void;
}

export const useOptionToggleStore = create<OptionToggleState>((set) => ({
  showAdvanced: false,
  setShowAdvanced: (showAdvanced: boolean) => set({ showAdvanced }),

  isDescriptionExpanded: false,
  setIsDescriptionExpanded: (isDescriptionExpanded: boolean) =>
    set({ isDescriptionExpanded }),
}));
