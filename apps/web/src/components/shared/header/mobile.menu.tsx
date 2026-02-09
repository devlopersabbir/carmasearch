"use client";

import { Dialog, DialogContent } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { getMenuItems } from "./_data/menu-items";
import { useHeaderStore } from "./header.store";

export function MobileMenu() {
  const { isMobileMenuOpen, setIsMobileMenuOpen } = useHeaderStore();
  const menuItems = getMenuItems();
  return (
    <Dialog open={isMobileMenuOpen} onOpenChange={setIsMobileMenuOpen}>
      <DialogContent className="sm:max-w-sm p-0 gap-0 bg-black/40 backdrop-blur-xl border border-white/10 rounded-[32px] shadow-2xl">
        <div className="flex items-center justify-center p-6 border-b border-white/10">
          <div className="flex items-center gap-3">
            <div className="w-8 h-8 rounded-full bg-primary/20 flex items-center justify-center border border-primary/30">
              <div className="w-5 h-5 rounded-full border border-primary" />
            </div>
            <span className="font-bold text-xl text-white">CARMA</span>
          </div>
        </div>

        <div className="p-6 space-y-2">
          {menuItems.map((item) => (
            <Button
              key={item.label}
              variant="ghost"
              className="w-full justify-start gap-4 h-14 text-white/80 hover:text-white hover:bg-white/10 rounded-2xl transition-all duration-300 transform hover:scale-[1.02] hover:shadow-lg active:scale-[0.98]"
              onClick={() => {
                item.action();
                setIsMobileMenuOpen(false);
              }}
            >
              <item.icon className="h-5 w-5" />
              <span className="font-medium">{item.label}</span>
            </Button>
          ))}
        </div>

        <Separator className="bg-white/10" />

        <div className="p-6">
          <div className="text-xs text-white/40 text-center font-medium">
            CARMA v1.0.0
          </div>
        </div>
      </DialogContent>
    </Dialog>
  );
}
