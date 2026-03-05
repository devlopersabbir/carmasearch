import { Badge } from "@/components/ui/badge";
import { Card } from "@/components/ui/card";
import { formatPrice, formatMileage, formatDealScore } from "@/lib/format";
import {
  Star,
  Target,
  Car,
  ExternalLink,
  Calendar,
  Gauge,
  Euro,
  TrendingUp,
  Palette,
  ChevronRight,
} from "lucide-react";
import { useAdvanceOptionStore } from "./_store/advance-option.store";
import { useOptionToggleStore } from "./_store/option-toggle.store";
import SearchResultDisplay from "./search-result-display";

// IMPORTANT_TODO: Need to be implement scoring system on client side
export default function CompareSearchResult() {
  const { searchResults } = useAdvanceOptionStore();
  const { isDescriptionExpanded, setIsDescriptionExpanded } =
    useOptionToggleStore();

  if (!searchResults?.query_vehicle) return null;
  const { query_vehicle } = searchResults;

  return query_vehicle ? (
    <div className="space-y-8">
      {/* Original Vehicle */}
      <div className="space-y-4">
        <div className="flex items-center justify-between">
          <h3 className="text-xl font-semibold text-white flex items-center gap-2">
            <Target className="h-5 w-5 text-primary" />
            Selected Vehicle
          </h3>
          <Badge
            variant="default"
            className="bg-primary/20 text-primary border-primary/30"
          >
            Original
          </Badge>
        </div>

        <Card className="p-6 bg-gradient-to-br from-blue-500/10 to-purple-500/10 border border-blue-500/20 backdrop-blur-sm">
          <div className="flex gap-6">
            {/* Left side - Image */}
            <div className="flex-shrink-0">
              {query_vehicle.images && query_vehicle.images.length && (
                <img
                  src={query_vehicle.images[0]}
                  alt={`${query_vehicle.make} ${query_vehicle.model}`}
                  className="w-80 h-48 object-cover rounded-xl shadow-lg"
                  loading="lazy"
                  decoding="async"
                  onError={(e) => {
                    e.currentTarget.style.display = "none";
                  }}
                />
              )}
            </div>

            {/* Right side - Content */}
            <div className="flex-1 min-w-0">
              <h4 className="text-2xl font-bold text-white mb-4">
                {query_vehicle.make} {query_vehicle.model}
              </h4>

              {/* Quick Snapshot */}
              <div className="bg-black/40 backdrop-blur-sm rounded-xl p-4 mb-4 border border-white/10">
                <h5 className="font-semibold text-white mb-3 flex items-center gap-2">
                  <Star className="h-4 w-4 text-primary" />
                  Quick Snapshot
                </h5>
                <div className="grid grid-cols-3 gap-3 text-sm">
                  <div className="flex items-center gap-2">
                    <Calendar className="h-4 w-4 text-white/60" />
                    <div>
                      <span className="text-white/60">First Registration:</span>
                      <div className="font-medium text-white">
                        {query_vehicle.first_registration ||
                          query_vehicle.first_registration_date ||
                          "N/A"}
                      </div>
                    </div>
                  </div>
                  <div className="flex items-center gap-2">
                    <Gauge className="h-4 w-4 text-white/60" />
                    <div>
                      <span className="text-white/60">Mileage:</span>
                      <div className="font-medium text-white">
                        {formatMileage(query_vehicle.mileage_km)}
                      </div>
                    </div>
                  </div>
                  <div className="flex items-center gap-2">
                    <Euro className="h-4 w-4 text-white/60" />
                    <div>
                      <span className="text-white/60">Price:</span>
                      <div className="font-medium text-white text-lg">
                        {formatPrice(query_vehicle.price)}
                      </div>
                    </div>
                  </div>
                  <div className="flex items-center gap-2">
                    <TrendingUp className="h-4 w-4 text-white/60" />
                    <div>
                      <span className="text-white/60">Deal Score:</span>
                      <div
                        className={`font-medium ${formatDealScore(1).class}`}
                      >
                        {/* TODO: Implement deal score */}
                        {formatDealScore(1).text}
                      </div>
                    </div>
                  </div>
                  <div className="flex items-center gap-2">
                    <Palette className="h-4 w-4 text-white/60" />
                    <div>
                      <span className="text-white/60">Exterior Color:</span>
                      <div className="font-medium text-white">
                        {query_vehicle.exterior_color || "N/A"}
                      </div>
                    </div>
                  </div>
                  <div className="flex items-center gap-2">
                    <Car className="h-4 w-4 text-white/60" />
                    <div>
                      <span className="text-white/60">Interior:</span>
                      <div className="font-medium text-white">
                        {query_vehicle.upholstery_color ||
                          query_vehicle.interior_color ||
                          "N/A"}
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              {/* Technical Details */}
              <div className="grid grid-cols-2 gap-4 text-sm text-white/80 mb-4">
                <p>
                  <strong className="text-white">Fuel Type:</strong>{" "}
                  {/* {query_vehicle.fuel_group} */}
                  {query_vehicle.fuel_type}
                </p>
                <p>
                  <strong className="text-white">Transmission:</strong>{" "}
                  {/* {query_vehicle.transmission_group} */}
                  {query_vehicle.transmission_type}
                </p>
                <p>
                  <strong className="text-white">Body Type:</strong>{" "}
                  {/* {query_vehicle.body_group} */}
                  {query_vehicle.body_type}
                </p>
                {query_vehicle.power_kw && (
                  <p>
                    <strong className="text-white">Power:</strong>{" "}
                    {query_vehicle.power_kw} kW
                  </p>
                )}
                <p>
                  <strong className="text-white">Predicted Fair Price:</strong>{" "}
                  {/* {formatPrice(query_vehicle.price_hat)} */}
                  {formatPrice(query_vehicle.price)}
                </p>
              </div>

              {/* Expandable Description */}
              {/* {query_vehicle.description && ( */}
              {query_vehicle.subtitle && (
                <div className="mb-4">
                  <button
                    onClick={() =>
                      setIsDescriptionExpanded(!isDescriptionExpanded)
                    }
                    className="flex items-center gap-2 text-white font-medium hover:text-primary transition-colors mb-2"
                  >
                    Description:
                    <ChevronRight
                      className={`h-4 w-4 transition-transform ${isDescriptionExpanded ? "rotate-90" : ""}`}
                    />
                  </button>
                  <div
                    className={`overflow-hidden transition-all duration-300 ${
                      isDescriptionExpanded
                        ? "max-h-96 opacity-100"
                        : "max-h-0 opacity-0"
                    }`}
                  >
                    <p className="text-sm text-white/70 leading-relaxed">
                      {/* {searchResults.vehicle.description} */}
                      {query_vehicle.subtitle}
                    </p>
                  </div>
                </div>
              )}

              <a
                href={query_vehicle.listing_url}
                target="_blank"
                rel="noopener noreferrer"
                className="inline-flex items-center gap-2 text-primary hover:text-primary/80 font-medium transition-colors"
              >
                View on AutoScout24
                <ExternalLink className="h-4 w-4" />
              </a>
            </div>
          </div>
        </Card>
      </div>

      {/* Comparable Vehicles */}
      <div className="space-y-4">
        <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-2">
          <h3 className="text-xl font-semibold text-white flex items-center gap-2">
            <Star className="h-5 w-5 text-primary" />
            Similar Vehicles Found
          </h3>
          <div className="flex items-center gap-2">
            <Badge
              variant="secondary"
              className="bg-white/10 text-white border-white/20"
            >
              {/* {TODO: need to show here the list of result} */}
              20 results
            </Badge>
            {/* TODO: need to show here the list of result */}
            {/* {searchResults.metadata && (
              <span className="text-xs text-white/50">
                Requested top {searchResults.metadata.requested_top}, ranked{" "}
                {searchResults.metadata.returned} of{" "}
                {searchResults.metadata.total_candidates} candidates
              </span>
            )} */}
          </div>
        </div>

        <SearchResultDisplay />
      </div>
    </div>
  ) : (
    <h1>no data found</h1>
  );
}
