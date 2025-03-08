import React from "react";
import { LoginForm } from "@/components/forms/login/login-form.tsx";
import ThemeToggle from "@/enteties/Theme/ui/ThemeToggle.tsx";

const LoginPage: React.FC = () => {
    return (
        <div className="min-h-screen flex flex-col items-center justify-center bg-muted p-6 md:p-10">
            <ThemeToggle className="absolute bottom-2 left-2"/>
            <div className="w-full max-w-sm md:max-w-2xl">
                <LoginForm/>
            </div>
        </div>
    )
}

export default LoginPage;