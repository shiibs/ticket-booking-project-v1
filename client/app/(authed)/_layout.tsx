import { Redirect, Stack } from "expo-router";

export default function AppLayout() {
  // check from context of user is logged in
  const isLoggedin = false;

  if (!isLoggedin) {
    return <Redirect href="/login" />;
  }

  return <Stack screenOptions={{ headerShown: false }} />;
}
