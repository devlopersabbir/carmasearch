"use client";
import { Button } from "@/components/ui/button";
import { useCtaButtonsStore } from "./cta.buttons.store";
import { useCompareModelStore } from "@/components/providers/models/compare-model.store";
import { toast } from "sonner";

export default function CtaButtons() {
  const { handleCompareClick } = useCtaButtonsStore();
  const { setIsCompareModalOpen } = useCompareModelStore();
  return (
    <>
      <Button
        size="lg"
        onClick={() => {
          handleCompareClick();
          setIsCompareModalOpen(true);
        }}
        className="text-lg px-8 py-3"
      >
        Compare Vehicles
      </Button>
      <Button
        variant="outline"
        size="lg"
        className="text-lg px-8 py-3"
        onClick={() => {
          // TODO: first need to be check user is authenticated or not
          // if user is not authenticated then redirect to the login page
          // if user is authenticated then open the price alerts modal
          toast.info("Price Alerts", {
            description:
              "You can now set up price alerts from your account dashboard.",
          });
        }}
      >
        Price Alerts
      </Button>
    </>
  );
}
