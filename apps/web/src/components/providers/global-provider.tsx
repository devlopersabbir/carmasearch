import { Toaster } from "sonner";
import ModelProvider from "./models/model-provider";
import React from "react";

export default function GlobalProvider() {
  return (
    <React.Fragment>
      <ModelProvider />
      <Toaster />
    </React.Fragment>
  );
}
