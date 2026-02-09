import { create } from "zustand";

export type CompareModelStore = {
  isCompareModalOpen: boolean;
  setIsCompareModalOpen: (open: boolean) => void;
};

export const useCompareModelStore = create<CompareModelStore>((set) => ({
  isCompareModalOpen: false,
  setIsCompareModalOpen: (open: boolean) => set({ isCompareModalOpen: open }),
}));
