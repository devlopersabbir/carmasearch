import { ResponseVehicleData } from "@/@types";

export type AdvanceOptionState = {
  vehicleUrl: string;
  setVehicleUrl: (vehicleUrl: string) => void;

  isSearching: boolean;
  setIsSearching: (isSearching: boolean) => void;

  registrationFrom: string;
  setRegistrationFrom: (registrationFrom: string) => void;

  registrationUntil: string;
  setRegistrationUntil: (registrationUntil: string) => void;

  mileageFrom: string;
  setMileageFrom: (mileageFrom: string) => void;

  mileageUntil: string;
  setMileageUntil: (mileageUntil: string) => void;

  showExteriorColors: boolean;
  setShowExteriorColors: (showExteriorColors: boolean) => void;

  showInteriorColors: boolean;
  setShowInteriorColors: (showInteriorColors: boolean) => void;

  showInteriorMaterials: boolean;
  setShowInteriorMaterials: (showInteriorMaterials: boolean) => void;

  selectedExteriorColors: string[];
  setSelectedExteriorColors: (selectedExteriorColors: string[]) => void;

  selectedInteriorColors: string[];
  setSelectedInteriorColors: (selectedInteriorColors: string[]) => void;

  selectedInteriorMaterials: string[];
  setSelectedInteriorMaterials: (selectedInteriorMaterials: string[]) => void;

  // handle save and compare function
  handleSaveAndCompare: () => void;
  setSearchResults: <T extends ResponseVehicleData>(
    searchResults: T | null,
  ) => void;
  searchResults: ResponseVehicleData | null;
  error: string;
  setError: (error: string) => void;
};
