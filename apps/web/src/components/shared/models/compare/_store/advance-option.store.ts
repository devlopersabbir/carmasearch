import { create } from "zustand";
import { AdvanceOptionState } from "./types";

export const useAdvanceOptionStore = create<AdvanceOptionState>((set) => ({
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
  handleSaveAndCompare: () => {
    set({ isSearching: true });
    setTimeout(() => {
      set({ isSearching: false });
    }, 2000);
  },
  setSearchResults: <T>(searchResults: T) => {
    console.log(searchResults);
    set({ searchResults });
  },
  searchResults: null,
  setError: (error: string) => set({ error }),
}));
