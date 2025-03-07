import {Button} from "@/components/ui/button"
import {useState} from "react";
import {RegisterUser} from "../wailsjs/go/handlers/UsersHandler";
import './app.css'

function App() {
    const [fullName, setFullName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const [result, setResult] = useState("");

    const sendInput = async () => {
        try {
            const response = await RegisterUser(fullName, email, password)
            setResult("User ID: " + response.user?.id);
        } catch (error) {
            setResult('Error');
        }
    }

    return (
        <div className="min-h-screen bg-white grid place-items-center mx-auto py-8">
        <div className="text-blue-900 text-2xl font-bold flex flex-col items-center space-y-4">
            <h1 className="text-purple-800 text-3xl font-bold">Speak Up</h1>
            <div className="input-box grid gap-4">
                <input
                    className="border-2 border-black border-accent text-zinc-900"
                    type="text"
                    placeholder="Full Name"
                    value={fullName}
                    onChange={(e) => setFullName(e.target.value)}
                />
                <input
                    className="border-2 border-black border-accent text-zinc-900"
                    type="text"
                    placeholder="Email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
                <input
                    className="border-2 border-black border-accent text-zinc-900"
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
            </div>
            <Button className="col-auto" onClick={sendInput}>Register</Button>
            <div className="result">{result}</div>
        </div>
        </div>
    )
}

export default App
