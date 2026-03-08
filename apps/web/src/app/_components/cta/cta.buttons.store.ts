import { create } from "zustand";

export type CtaButtonsStore = {
  handleCompareClick: () => void;
  handlePriceAlertsClick: () => void;
};

export const ctaButtonsStore = create<CtaButtonsStore>((set) => ({
  handleCompareClick: () => {
    console.log("Compare clicked");
  },
  handlePriceAlertsClick: () => {
    console.log("Price alerts clicked");
  },
}));
