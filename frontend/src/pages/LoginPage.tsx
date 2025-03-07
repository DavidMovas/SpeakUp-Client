import React from "react";
import { LoginForm } from "@/components/forms/login/login-form.tsx";

const LoginPage: React.FC = () => {
    return (
        <div className="flex min-h-svh flex-col items-center justify-center bg-muted p-6 md:p-10">
            <div className="w-full h-max max-w-sm md:max-w-3xl">
                <LoginForm/>
            </div>
        </div>
    )
}

export default LoginPage;