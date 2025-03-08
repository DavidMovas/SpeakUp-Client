import React, { useState } from "react";
import ThemeToggle from "@/enteties/Theme/ui/ThemeToggle.tsx";
import { AnimatePresence } from "framer-motion";
import { LoginForm } from "@/components/forms/login/login-form.tsx";
import { SignUpForm } from "@/components/forms/signup/sign-up-form.tsx";

export interface AuthPageProps {
    className?: string
    setLogin: (isLogin: boolean) => void;
}

const AuthPage: React.FC = () => {
    const [isLogin, setIsLogin] = useState(true);

    return (
        <div className="min-h-screen flex flex-col items-center justify-center bg-muted p-6 md:p-10">
            <ThemeToggle className="absolute bottom-2 left-2"/>
            <div className="w-full h-max max-w-sm md:max-w-2xl">
                <AnimatePresence mode={"wait"}>
                    {isLogin && <LoginForm key={"login"} setLogin={setIsLogin} />}
                    {!isLogin && <SignUpForm key={"signup"} setLogin={setIsLogin} />}
                </AnimatePresence>
            </div>
        </div>
    )
}

export default AuthPage;