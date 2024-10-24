import { Button } from "@/components/Button";
import { VStack } from "@/components/VStack";
import { useAuth } from "@/context/AuthContext";

export default function SettingScreen() {
  const { logout } = useAuth();

  return (
    <VStack flex={1} m={20}>
      <Button onPress={logout}>Logout</Button>
    </VStack>
  );
}
