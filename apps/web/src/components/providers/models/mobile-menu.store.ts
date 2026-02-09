import { create } from "zustand";

export type MobileMenuStore = {
  isMobileMenuOpen: boolean;
  setIsMobileMenuOpen: (open: boolean) => void;
};

export const useMobileMenuStore = create<MobileMenuStore>((set) => ({
  isMobileMenuOpen: false,
  setIsMobileMenuOpen: (open: boolean) => set({ isMobileMenuOpen: open }),
}));
