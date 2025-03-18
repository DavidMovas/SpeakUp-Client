import React from "react";
import useThemeStore from "@/enteties/Theme/store/themeStore.ts";
import { Button } from "@/components/ui/shadcn/button.tsx";
import { Moon, Sun } from "lucide-react";


function ThemeToggle ({className}:React.ComponentProps<"div">) {
    const {isDark, toggleTheme} = useThemeStore();

   return (
        <Button className={className} onClick={toggleTheme}>
            {!isDark ? <Sun/> : <Moon/>}
        </Button>
   )
}

export default ThemeToggle;