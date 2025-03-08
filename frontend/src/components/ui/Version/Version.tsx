import { FC, memo } from "react";
import { cn } from "@/lib/utils.ts"

interface VersionProps {
    className?: string;
    version: string;
}

export const Version:FC<VersionProps> = memo((props: VersionProps) => {
    const {version, className} = props;

    return (
        <div className={cn("text-primary end-0 bottom-0 p-2 absolute", className)}>v{version}</div>
    );
});
