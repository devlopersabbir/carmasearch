import { create } from "zustand";
import { CompareVehicles } from "../../actions/compare-vehicles";
import { useAdvanceOptionStore } from "@/components/shared/models/compare/_store/advance-option.store";

export type CtaButtonsStore = {
  handleCompareClick: () => void;
};

export const useCtaButtonsStore = create<CtaButtonsStore>((set) => ({
  handleCompareClick: async () => {
    const query = useAdvanceOptionStore();
    try {
      const res = await CompareVehicles({
        listing_url: query.vehicleUrl,
        ...query,
      });
      if (res) {
        query.setSearchResults(res);
      }
    } catch (err) {
      query.setError(err as string);
    } finally {
      query.setIsSearching(false);
    }
  },
}));
