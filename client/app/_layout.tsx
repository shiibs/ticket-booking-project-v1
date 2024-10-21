import { StatusBar } from "expo-status-bar";
import { Slot } from "expo-router";

export default function RootLayout() {
  return (
    <>
      <StatusBar style={"dark"} />
      {/* Authentication provider */}
      <Slot />
      {/* Authentication provider */}
    </>
  );
}
