"use client";

import { formatDealScore, formatMileage, formatPrice } from "@/lib/format";
import { useAdvanceOptionStore } from "./_store/advance-option.store";
import { Card } from "@/components/ui/card";
import { Car, ExternalLink } from "lucide-react";
import { Badge } from "@/components/ui/badge";

export default function SearchResultDisplay() {
  const { searchResults } = useAdvanceOptionStore();

  if (!searchResults?.query_vehicle) return null;
  const { results } = searchResults;
  // const comparablesToDisplay = searchResults?.comparables || [];

  return (
    <div className="space-y-4">
      {/* TODO: actual result will be display here... */}
      {results &&
        results.map((vehicle, index) => {
          // const dealScore = formatDealScore(vehicle.deal_score);
          // const similarityScore =
          //   typeof vehicle.similarity_score === "number"
          //     ? vehicle.similarity_score
          //     : typeof vehicle.score === "number"
          //       ? vehicle.score
          //       : undefined;
          // const finalScore =
          //   typeof vehicle.final_score === "number"
          //     ? vehicle.final_score
          //     : similarityScore;
          // const preferenceScore =
          //   typeof vehicle.preference_score === "number"
          //     ? vehicle.preference_score
          //     : undefined;
          // const matchPercentage =
          //   finalScore !== undefined ? (finalScore * 100).toFixed(1) : null;
          // const similarityPercentage =
          //   similarityScore !== undefined
          //     ? Math.round(similarityScore * 100)
          //     : null;
          // const preferencePercentage =
          //   preferenceScore !== undefined
          //     ? Math.round(preferenceScore * 100)
          //     : null;
          // let dealPercent: number | null = null;
          // if (typeof vehicle.deal_score === "number") {
          //   let normalized = vehicle.deal_score;
          //   if (normalized >= 0 && normalized <= 1) {
          //     normalized = (normalized - 0.5) * 100;
          //   } else {
          //     normalized = normalized * 100;
          //   }
          //   dealPercent = Math.round(normalized);
          // }
          // const filterLevel = vehicle.ranking_details?.filter_level;
          return (
            <Card
              key={vehicle.unique_id}
              className="p-4 hover:shadow-lg transition-all duration-300 bg-black/40 backdrop-blur-sm border border-white/10 hover:border-white/20"
            >
              <div className="flex gap-4">
                <div className="flex-shrink-0">
                  {vehicle.images && vehicle.images.length > 0 ? (
                    <img
                      src={vehicle.images[0]}
                      alt={`${vehicle.make} ${vehicle.model}`}
                      className="w-28 h-20 object-cover rounded-lg"
                      loading="lazy"
                      // onError={(e) => {
                      //   e.currentTarget.style.display = "none";
                      //   e.currentTarget && e.currentTarget?.nextElementSibling && e.currentTarget?.nextElementSibling.style.display =
                      //     "flex";
                      // }}
                    />
                  ) : null}
                  <div
                    className="w-28 h-20 bg-white/10 rounded-lg flex items-center justify-center"
                    style={{
                      display:
                        vehicle.images && vehicle.images.length > 0
                          ? "none"
                          : "flex",
                    }}
                  >
                    <Car className="h-8 w-8 text-white/40" />
                  </div>
                </div>

                <div className="flex-1 min-w-0">
                  <div className="flex items-start justify-between">
                    <div className="flex-1">
                      <h4 className="font-semibold text-white truncate text-lg">
                        {vehicle.make} {vehicle.model}
                      </h4>
                      <p className="text-sm text-white/70 mb-2">
                        {vehicle.production_year} • {vehicle.fuel_type} •{" "}
                        {vehicle.transmission_type}
                      </p>
                      <p className="text-sm text-white/70 mb-2">
                        {vehicle.body_type} •{" "}
                        {formatMileage(vehicle.mileage_km)}
                      </p>
                      <p className="text-sm text-white/70">
                        {vehicle.exterior_color || "N/A"} exterior •{" "}
                        {vehicle.upholstery_color ||
                          vehicle.interior_color ||
                          "N/A"}{" "}
                        interior
                      </p>
                      {vehicle.power_kw && (
                        <p className="text-sm text-white/70">
                          {vehicle.power_kw} kW
                        </p>
                      )}
                    </div>

                    <div className="text-right ml-4">
                      <div className="flex items-center gap-2 mb-2">
                        <Badge
                          variant={index === 0 ? "default" : "secondary"}
                          className={
                            index === 0
                              ? "bg-primary/20 text-primary border-primary/30"
                              : "bg-white/10 text-white border-white/20"
                          }
                        >
                          #{index + 1}
                        </Badge>
                        {/* {filterLevel ? (
                          <Badge
                            variant="outline"
                            className="border-white/10 bg-transparent text-white/70"
                          >
                            L{filterLevel}
                          </Badge>
                        ) : null}
                        <div className="flex items-center gap-1">
                          <Star className="h-3 w-3 text-primary" />
                          <span className="text-xs text-white/60">
                            {matchPercentage ?? "N/A"}% match
                          </span>
                        </div> */}
                      </div>

                      <div className="mb-2">
                        <div className="text-xl font-bold text-primary">
                          {formatPrice(vehicle.price)}
                        </div>
                        <div className="text-sm text-white/60">
                          {/* Fair: {formatPrice(vehicle.price_hat)} */}
                          Fair: {formatPrice(2043)}
                        </div>
                        {/* <div
                          className={`text-xs font-medium ${dealScore.class}`}
                        >
                          {dealScore.text}
                        </div> */}
                        {/* <div className="text-xs text-white/50 space-x-2">
                          {similarityPercentage !== null && (
                            <span>Sim {similarityPercentage}%</span>
                          )}
                          {dealPercent !== null && (
                            <span>
                              Deal {dealPercent > 0 ? "+" : ""}
                              {dealPercent}%
                            </span>
                          )}
                          {preferencePercentage !== null &&
                            preferencePercentage > 0 && (
                              <span>Pref {preferencePercentage}%</span>
                            )}
                        </div> */}
                        {/* {typeof vehicle.savings === "number" &&
                          vehicle.savings > 0 && (
                            <div className="text-xs text-emerald-400">
                              Saves {formatPrice(vehicle.savings)}
                            </div>
                          )} */}
                      </div>

                      <div>
                        <a
                          href={vehicle.listing_url}
                          target="_blank"
                          rel="noopener noreferrer"
                          className="inline-flex items-center gap-1 text-xs text-primary hover:text-primary/80 transition-colors"
                        >
                          View Listing
                          <ExternalLink className="h-3 w-3" />
                        </a>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </Card>
          );
        })}
    </div>
  );
}
