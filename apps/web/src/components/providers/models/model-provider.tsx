"use client";
import React from "react";
import { MobileMenu } from "@/components/shared/header/mobile.menu";
import CompareModal from "@/components/shared/models/compare/compare-model";

export default function ModelProvider() {
  return (
    <React.Fragment>
      <CompareModal />
      <MobileMenu />
      {/*
      <AuthModal
        isOpen={isAuthModalOpen}
        onClose={() => setIsAuthModalOpen(false)}
        mode={authMode}
      /> */}
    </React.Fragment>
  );
}
