import { create } from "zustand";

export type CtaButtonsStore = {
  handleCompareClick: () => void;
};

export const useCtaButtonsStore = create<CtaButtonsStore>((set) => ({
  handleCompareClick: () => set({}),
}));
