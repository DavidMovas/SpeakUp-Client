import React from "react";
import { SignUpForm } from "@/components/forms/signup/sign-up-form.tsx";

const SignUpPage: React.FC = () => {
    return (
        <div className="flex min-h-svh flex-col items-center justify-center bg-muted p-6 md:p-10">
            <div className="w-full h-max max-w-sm md:max-w-3xl">
                <SignUpForm/>
            </div>
        </div>
    )
}

export default SignUpPage;