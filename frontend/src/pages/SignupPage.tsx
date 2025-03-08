import React from "react";
import { SignUpForm } from "@/components/forms/signup/sign-up-form.tsx";
import ThemeToggle from "@/enteties/Theme/ui/ThemeToggle.tsx";

const SignUpPage: React.FC = () => {
    return (
        <div className="min-h-screen flex flex-col items-center justify-center bg-muted p-6 md:p-10">
            <ThemeToggle className="absolute bottom-2 left-2"/>
            <div className="w-full h-max max-w-sm md:max-w-2xl">
                <SignUpForm/>
            </div>
        </div>
    )
}

export default SignUpPage;