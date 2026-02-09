import { create } from "zustand";

type Mode = "login" | "signup";
export type AuthModelStore = {
  isAuthModalOpen: boolean;
  setIsAuthModalOpen: (open: boolean) => void;
  mode: Mode;
  activeTab: Mode;
  setActiveTab: (tab: Mode) => void;
};

export const useAuthModelStore = create<AuthModelStore>((set) => ({
  isAuthModalOpen: false,
  mode: "login",
  activeTab: "login",
  setIsAuthModalOpen: (open: boolean) => set({ isAuthModalOpen: open }),
  setMode: (mode: Mode) => set({ mode }),
  setActiveTab: (tab: Mode) => set({ activeTab: tab }),
}));
