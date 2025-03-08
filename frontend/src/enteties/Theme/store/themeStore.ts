import {create} from 'zustand';
import { ThemeState } from "@/enteties/Theme/types/theme.ts";

const useThemeStore = create<ThemeState>((set) => ({
    isDark: false,
    toggleTheme: () => set((state) => ({isDark: !state.isDark}))
}))

export default useThemeStore;