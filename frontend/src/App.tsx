import { HashRouter, Route, Routes } from 'react-router-dom';
import { lazy, Suspense } from "react";
import { Version } from "@/components/ui/Version/Version.tsx";
import { Toaster } from "sonner";

const version = "0.0.1-alpha"

const LoginPage = lazy(() => import('./pages/LoginPage'))
const SigUpPage = lazy(() => import('./pages/SignupPage'))

function App() {
    return (
        <HashRouter basename={"/"}>
            <Version version={version}/>
            <Toaster />
            <Suspense fallback={<div>Loading...</div>}>
                <Routes>
                    <Route path="/" element={<LoginPage />}></Route>
                    <Route path="/signup" element={<SigUpPage />}></Route>
                </Routes>
            </Suspense>
        </HashRouter>
    )
}

export default App
