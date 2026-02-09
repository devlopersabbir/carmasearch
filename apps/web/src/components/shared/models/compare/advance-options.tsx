"use client";

import { Card } from "@/components/ui/card";
import { useOptionToggleStore } from "./_store/option-toggle.store";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useAdvanceOptionStore } from "./_store/advance-option.store";
import {
  exteriorColors,
  getColorSwatch,
  interiorColors,
  interiorMaterials,
  registrationYears,
} from "./_constants";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { ChevronDown, ChevronUp } from "lucide-react";
import { Checkbox } from "@/components/ui/checkbox";
import { useCompareModelStore } from "@/components/providers/models/compare-model.store";

export default function AdvanceOptions() {
  const { showAdvanced } = useOptionToggleStore();
  const {
    isSearching,
    handleSaveAndCompare,
    registrationFrom,
    registrationUntil,
    mileageFrom,
    mileageUntil,
    setRegistrationFrom,
    setRegistrationUntil,
    setMileageFrom,
    setMileageUntil,
    showExteriorColors,
    setShowExteriorColors,
    showInteriorColors,
    setShowInteriorColors,
    showInteriorMaterials,
    setShowInteriorMaterials,
    selectedExteriorColors,
    setSelectedExteriorColors,
    selectedInteriorColors,
    setSelectedInteriorColors,
    selectedInteriorMaterials,
    setSelectedInteriorMaterials,
  } = useAdvanceOptionStore();

  return (
    showAdvanced && (
      <Card className="p-6 space-y-6 bg-black/30 backdrop-blur-sm border border-white/10">
        <div className="space-y-6">
          {/* First Registration Range - Moved to Top */}
          <div className="space-y-3">
            <Label className="text-white/90 font-semibold">
              First Registration
            </Label>
            <div className="flex gap-4">
              <div className="flex-1">
                <Label
                  htmlFor="registration-from"
                  className="text-sm text-white/70"
                >
                  From
                </Label>
                <Select
                  value={registrationFrom}
                  onValueChange={setRegistrationFrom}
                >
                  <SelectTrigger className="bg-black/40 border-white/10 text-white">
                    <SelectValue placeholder="From" />
                  </SelectTrigger>
                  <SelectContent className="bg-black/90 border-white/10">
                    {registrationYears.map((year) => (
                      <SelectItem key={year} value={year}>
                        {year}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              </div>
              <div className="flex-1">
                <Label
                  htmlFor="registration-until"
                  className="text-sm text-white/70"
                >
                  Until
                </Label>
                <Select
                  value={registrationUntil}
                  onValueChange={setRegistrationUntil}
                >
                  <SelectTrigger className="bg-black/40 border-white/10 text-white">
                    <SelectValue placeholder="Until" />
                  </SelectTrigger>
                  <SelectContent className="bg-black/90 border-white/10">
                    {registrationYears.map((year) => (
                      <SelectItem key={year} value={year}>
                        {year}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              </div>
            </div>
          </div>

          {/* Mileage Range - Moved to Top */}
          <div className="space-y-3">
            <Label className="text-white/90 font-semibold">Mileage</Label>
            <div className="flex gap-4">
              <div className="flex-1">
                <Label htmlFor="mileage-from" className="text-sm text-white/70">
                  From
                </Label>
                <div className="relative">
                  <Input
                    id="mileage-from"
                    type="number"
                    placeholder="From"
                    value={mileageFrom}
                    onChange={(e) => setMileageFrom(e.target.value)}
                    className="bg-black/40 border-white/10 text-white pr-8"
                  />
                  <span className="absolute right-3 top-1/2 transform -translate-y-1/2 text-white/70 text-sm">
                    km
                  </span>
                </div>
              </div>
              <div className="flex-1">
                <Label
                  htmlFor="mileage-until"
                  className="text-sm text-white/70"
                >
                  Until
                </Label>
                <div className="relative">
                  <Input
                    id="mileage-until"
                    type="number"
                    placeholder="Until"
                    value={mileageUntil}
                    onChange={(e) => setMileageUntil(e.target.value)}
                    className="bg-black/40 border-white/10 text-white pr-8"
                  />
                  <span className="absolute right-3 top-1/2 transform -translate-y-1/2 text-white/70 text-sm">
                    km
                  </span>
                </div>
              </div>
            </div>
          </div>

          {/* Exterior Color - Expandable */}
          <div className="space-y-3">
            <div className="flex items-center justify-between">
              <Label className="text-white/90 font-semibold">
                Exterior Color
              </Label>
              <Button
                variant="ghost"
                size="sm"
                onClick={() => setShowExteriorColors(!showExteriorColors)}
                className="text-white/70 hover:text-white"
              >
                {showExteriorColors ? (
                  <ChevronUp className="h-4 w-4" />
                ) : (
                  <ChevronDown className="h-4 w-4" />
                )}
              </Button>
            </div>
            {showExteriorColors && (
              <div className="grid grid-cols-4 gap-2">
                {exteriorColors.map((color) => (
                  <div key={color} className="flex items-center space-x-2">
                    <Checkbox
                      id={`exterior-${color}`}
                      checked={selectedExteriorColors.includes(color)}
                      onCheckedChange={(checked) => {
                        if (checked) {
                          setSelectedExteriorColors([
                            ...selectedExteriorColors,
                            color,
                          ]);
                        } else {
                          setSelectedExteriorColors(
                            selectedExteriorColors.filter((c) => c !== color),
                          );
                        }
                      }}
                      className="border-white/20"
                    />
                    <div className="flex items-center space-x-2">
                      <div
                        className="w-4 h-4 rounded border border-white/20"
                        style={{ backgroundColor: getColorSwatch(color) }}
                      />
                      <Label
                        htmlFor={`exterior-${color}`}
                        className="text-sm text-white/80"
                      >
                        {color}
                      </Label>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </div>

          {/* Interior Color - Expandable */}
          <div className="space-y-3">
            <div className="flex items-center justify-between">
              <Label className="text-white/90 font-semibold">
                Interior Color
              </Label>
              <Button
                variant="ghost"
                size="sm"
                onClick={() => setShowInteriorColors(!showInteriorColors)}
                className="text-white/70 hover:text-white"
              >
                {showInteriorColors ? (
                  <ChevronUp className="h-4 w-4" />
                ) : (
                  <ChevronDown className="h-4 w-4" />
                )}
              </Button>
            </div>
            {showInteriorColors && (
              <div className="grid grid-cols-4 gap-2">
                {interiorColors.map((color) => (
                  <div key={color} className="flex items-center space-x-2">
                    <Checkbox
                      id={`interior-${color}`}
                      checked={selectedInteriorColors.includes(color)}
                      onCheckedChange={(checked) => {
                        if (checked) {
                          setSelectedInteriorColors([
                            ...selectedInteriorColors,
                            color,
                          ]);
                        } else {
                          setSelectedInteriorColors(
                            selectedInteriorColors.filter((c) => c !== color),
                          );
                        }
                      }}
                      className="border-white/20"
                    />
                    <div className="flex items-center space-x-2">
                      <div
                        className="w-4 h-4 rounded border border-white/20"
                        style={{ backgroundColor: getColorSwatch(color) }}
                      />
                      <Label
                        htmlFor={`interior-${color}`}
                        className="text-sm text-white/80"
                      >
                        {color}
                      </Label>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </div>

          {/* Interior Material - Expandable */}
          <div className="space-y-3">
            <div className="flex items-center justify-between">
              <Label className="text-white/90 font-semibold">
                Interior Material
              </Label>
              <Button
                variant="ghost"
                size="sm"
                onClick={() => setShowInteriorMaterials(!showInteriorMaterials)}
                className="text-white/70 hover:text-white"
              >
                {showInteriorMaterials ? (
                  <ChevronUp className="h-4 w-4" />
                ) : (
                  <ChevronDown className="h-4 w-4" />
                )}
              </Button>
            </div>
            {showInteriorMaterials && (
              <div className="grid grid-cols-3 gap-2">
                {interiorMaterials.map((material) => (
                  <div key={material} className="flex items-center space-x-2">
                    <Checkbox
                      id={`material-${material}`}
                      checked={selectedInteriorMaterials.includes(material)}
                      onCheckedChange={(checked) => {
                        if (checked) {
                          setSelectedInteriorMaterials([
                            ...selectedInteriorMaterials,
                            material,
                          ]);
                        } else {
                          setSelectedInteriorMaterials(
                            selectedInteriorMaterials.filter(
                              (m) => m !== material,
                            ),
                          );
                        }
                      }}
                      className="border-white/20"
                    />
                    <Label
                      htmlFor={`material-${material}`}
                      className="text-sm text-white/80"
                    >
                      {material}
                    </Label>
                  </div>
                ))}
              </div>
            )}
          </div>
        </div>

        <div className="flex justify-end">
          <Button
            onClick={handleSaveAndCompare}
            disabled={isSearching}
            className="bg-white/20 backdrop-blur-sm border border-white/20 hover:bg-white/30 text-white font-medium rounded-2xl px-6 py-2"
          >
            {isSearching ? "Searching..." : "Save & Compare"}
          </Button>
        </div>
      </Card>
    )
  );
}
