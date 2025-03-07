import { LoginForm } from "@/components/login-form.tsx";

function App() {
    return (
        <div className="light">
            <div className="flex min-h-svh flex-col items-center justify-center bg-muted p-6 md:p-10">
                <div className="w-full h-max max-w-sm md:max-w-3xl">
                    <LoginForm/>
                </div>
            </div>
        </div>
    )
}

export default App
