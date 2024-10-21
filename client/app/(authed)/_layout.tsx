import { Redirect, Stack } from "expo-router";

export default function AppLayout() {
  // check from context of user is logged in
  const isLoggedin = true;

  if (!isLoggedin) {
    return <Redirect href="/login" />;
  }

  return <Stack screenOptions={{ headerShown: false }} />;
}
