import { HashRouter, Route, Routes } from 'react-router-dom';
import { lazy, Suspense } from "react";
import { Version } from "@/components/ui/Version/Version.tsx";
import { Toaster } from "sonner";
import useThemeStore from "@/enteties/Theme/store/themeStore.ts";

const version = "0.0.3-alpha"

const AuthPage = lazy(() => import('./pages/AuthPage'))
const ChatPage = lazy(() => import('./pages/ChatPage'))

function App() {
    const {isDark} = useThemeStore();

    return (
        <div className={isDark ? 'dark' : ''}>
            <HashRouter basename={"/"}>
                <Version version={version}/>
                <Toaster />
                <Suspense fallback={<div>Loading...</div>}>
                    <Routes>
                        <Route path="/auth" element={<AuthPage />}></Route>
                        <Route path="/" element={<ChatPage />}></Route>
                    </Routes>
                </Suspense>
            </HashRouter>
        </div>
    )
}

export default App
