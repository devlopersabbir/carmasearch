import { mileageFormatter } from "@/constants";
import { currencyFormatter } from "@/constants/currency";

// Format deal score for display
export function formatDealScore(rawScore?: number | null): {
  text: string;
  class: string;
  isGood: boolean;
} {
  if (typeof rawScore !== "number" || Number.isNaN(rawScore)) {
    return { text: "Fair Price", class: "text-gray-600", isGood: true };
  }

  let normalized = rawScore;

  if (normalized >= 0 && normalized <= 1) {
    normalized = (normalized - 0.5) * 2;
  }

  normalized = Math.max(-1, Math.min(1, normalized));

  if (Math.abs(normalized) < 0.01) {
    return { text: "Fair Price", class: "text-gray-600", isGood: true };
  }

  const percentage = Math.abs(normalized * 100).toFixed(1);

  if (normalized > 0) {
    return {
      text: `Good Deal (+${percentage}%)`,
      class: "text-green-600",
      isGood: true,
    };
  }

  return {
    text: `Overpriced (${percentage}%)`,
    class: "text-red-600",
    isGood: false,
  };
}

export function formatPrice(price?: number | null): string {
  if (price === null || price === undefined || Number.isNaN(price)) {
    return "N/A";
  }

  return currencyFormatter.format(price);
}

// Format mileage for display
export function formatMileage(mileage: number): string {
  if (mileage === null || mileage === undefined || Number.isNaN(mileage)) {
    return "N/A";
  }
  return `${mileageFormatter.format(mileage)} km`;
}
