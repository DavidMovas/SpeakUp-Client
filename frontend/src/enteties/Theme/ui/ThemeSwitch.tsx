import React from "react";
import useThemeStore from "@/enteties/Theme/store/themeStore.ts";
import { Switch } from "@/components/ui/shadcn/switch.tsx";
import { Moon, Sun } from "lucide-react";
import { cn } from "@/lib/utils.ts";

function ThemeSwitch ({className}:React.ComponentProps<"div">) {
    const {isDark, toggleTheme} = useThemeStore();

    return (
        <div
            className={cn(className, "relative flex cursor-default select-none items-center gap-2 rounded-sm py-1.5 text-sm outline-none transition-colors focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50 [&_svg]:pointer-events-none [&_svg]:size-4 [&_svg]:shrink-0")}>
            <div className={"flex gap-2"}>
                {!isDark ? <Sun/> : <Moon/>}
                Theme
            </div>
            <div className="ml-auto">
                <Switch onClick={toggleTheme}/>
            </div>
        </div>
    )
}

export default ThemeSwitch;