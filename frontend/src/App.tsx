import { HashRouter, Route, Routes } from 'react-router-dom';
import { lazy, Suspense } from "react";

const LoginPage = lazy(() => import('./pages/LoginPage'))
const SigUpPage = lazy(() => import('./pages/SignupPage'))

function App() {
    return (
        <HashRouter basename={"/"}>
            <div className="end-0 bottom-0 absolute">v0.0.1</div>
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
