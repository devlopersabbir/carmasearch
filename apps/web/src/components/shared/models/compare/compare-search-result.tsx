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

export default function CompareSearchResult() {
  const { searchResults } = useAdvanceOptionStore();
  const { isDescriptionExpanded, setIsDescriptionExpanded } =
    useOptionToggleStore();
  return (
    searchResults && (
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
                {searchResults.vehicle.images &&
                  searchResults.vehicle.images.length > 0 && (
                    <img
                      src={searchResults.vehicle.images[0]}
                      alt={`${searchResults.vehicle.make} ${searchResults.vehicle.model}`}
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
                  {searchResults.vehicle.make} {searchResults.vehicle.model}
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
                        <span className="text-white/60">
                          First Registration:
                        </span>
                        <div className="font-medium text-white">
                          {searchResults.vehicle.year}
                        </div>
                      </div>
                    </div>
                    <div className="flex items-center gap-2">
                      <Gauge className="h-4 w-4 text-white/60" />
                      <div>
                        <span className="text-white/60">Mileage:</span>
                        <div className="font-medium text-white">
                          {formatMileage(searchResults.vehicle.mileage_km)}
                        </div>
                      </div>
                    </div>
                    <div className="flex items-center gap-2">
                      <Euro className="h-4 w-4 text-white/60" />
                      <div>
                        <span className="text-white/60">Price:</span>
                        <div className="font-medium text-white text-lg">
                          {formatPrice(searchResults.vehicle.price_eur)}
                        </div>
                      </div>
                    </div>
                    <div className="flex items-center gap-2">
                      <TrendingUp className="h-4 w-4 text-white/60" />
                      <div>
                        <span className="text-white/60">Deal Score:</span>
                        <div
                          className={`font-medium ${formatDealScore(searchResults.vehicle.deal_score).class}`}
                        >
                          {
                            formatDealScore(searchResults.vehicle.deal_score)
                              .text
                          }
                        </div>
                      </div>
                    </div>
                    <div className="flex items-center gap-2">
                      <Palette className="h-4 w-4 text-white/60" />
                      <div>
                        <span className="text-white/60">Exterior Color:</span>
                        <div className="font-medium text-white">
                          {searchResults.vehicle.exterior_color || "N/A"}
                        </div>
                      </div>
                    </div>
                    <div className="flex items-center gap-2">
                      <Car className="h-4 w-4 text-white/60" />
                      <div>
                        <span className="text-white/60">Interior:</span>
                        <div className="font-medium text-white">
                          {searchResults.vehicle.upholstery_color ||
                            searchResults.vehicle.interior_color ||
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
                    {searchResults.vehicle.fuel_group}
                  </p>
                  <p>
                    <strong className="text-white">Transmission:</strong>{" "}
                    {searchResults.vehicle.transmission_group}
                  </p>
                  <p>
                    <strong className="text-white">Body Type:</strong>{" "}
                    {searchResults.vehicle.body_group}
                  </p>
                  {searchResults.vehicle.power_kw && (
                    <p>
                      <strong className="text-white">Power:</strong>{" "}
                      {searchResults.vehicle.power_kw} kW
                    </p>
                  )}
                  <p>
                    <strong className="text-white">
                      Predicted Fair Price:
                    </strong>{" "}
                    {formatPrice(searchResults.vehicle.price_hat)}
                  </p>
                </div>

                {/* Expandable Description */}
                {searchResults.vehicle.description && (
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
                        {searchResults.vehicle.description}
                      </p>
                    </div>
                  </div>
                )}

                <a
                  href={searchResults.vehicle.url}
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
              {searchResults.metadata && (
                <span className="text-xs text-white/50">
                  Requested top {searchResults.metadata.requested_top}, ranked{" "}
                  {searchResults.metadata.returned} of{" "}
                  {searchResults.metadata.total_candidates} candidates
                </span>
              )}
            </div>
          </div>

          <SearchResultDisplay />
        </div>
      </div>
    )
  );
}
