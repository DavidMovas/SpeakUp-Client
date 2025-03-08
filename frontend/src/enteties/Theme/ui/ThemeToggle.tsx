import React from "react";
import useThemeStore from "@/enteties/Theme/store/themeStore.ts";
import { Button } from "@/components/ui/Button/button.tsx";


function ThemeToggle ({className}:React.ComponentProps<"div">) {
    const {isDark, toggleTheme} = useThemeStore();

   return (
        <Button className={className} variant="outline" onClick={toggleTheme}>
            {isDark ? '🌙' : '☀️'}
        </Button>
   )
}

export default ThemeToggle;