import { create } from "zustand";
import { AdvanceOptionState } from "./types";
import { CompareVehicles } from "@/app/(home)/actions/compare-vehicles";

export const useAdvanceOptionStore = create<AdvanceOptionState>((set, get) => ({
  vehicleUrl: "",
  isSearching: false,

  registrationFrom: "",
  registrationUntil: "",
  mileageFrom: "",
  mileageUntil: "",
  showExteriorColors: false,
  showInteriorColors: false,
  showInteriorMaterials: false,

  selectedExteriorColors: [],
  selectedInteriorColors: [],
  selectedInteriorMaterials: [],

  error: "",

  setVehicleUrl: (vehicleUrl: string) => set({ vehicleUrl }),
  setIsSearching: (isSearching: boolean) => set({ isSearching }),

  setRegistrationFrom: (registrationFrom: string) => set({ registrationFrom }),
  setRegistrationUntil: (registrationUntil: string) =>
    set({ registrationUntil }),
  setMileageFrom: (mileageFrom: string) => set({ mileageFrom }),
  setMileageUntil: (mileageUntil: string) => set({ mileageUntil }),
  setShowExteriorColors: (showExteriorColors: boolean) =>
    set({ showExteriorColors }),
  setShowInteriorColors: (showInteriorColors: boolean) =>
    set({ showInteriorColors }),
  setShowInteriorMaterials: (showInteriorMaterials: boolean) =>
    set({ showInteriorMaterials }),

  setSelectedExteriorColors: (selectedExteriorColors: string[]) =>
    set({ selectedExteriorColors }),
  setSelectedInteriorColors: (selectedInteriorColors: string[]) =>
    set({ selectedInteriorColors }),
  setSelectedInteriorMaterials: (selectedInteriorMaterials: string[]) =>
    set({ selectedInteriorMaterials }),

  // TODO: handle save and compare function here...
  // TODO: compare api call will be from here...
  handleSaveAndCompare: async () => {
    set({ isSearching: true });
    // TODO: save query to local storage
    // TODO: call api to get data
    const query = get();
    try {
      const response = await CompareVehicles({
        listing_url: query.vehicleUrl,
      });
      if (response) {
        set({ searchResults: response });
      }
    } catch (err) {
      set({ error: err as string });
    } finally {
      set({ isSearching: false });
    }
  },
  setSearchResults: (searchResults) => {
    set({ searchResults });
  },
  searchResults: null,
  setError: (error: string) => set({ error }),
}));
