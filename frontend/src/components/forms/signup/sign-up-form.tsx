import React, { useState } from "react";
import { cn } from "@/lib/utils.ts"
import { Button } from "@/components/ui/Button/button.tsx"
import { Card, CardContent } from "@/components/ui/Card/card.tsx"
import { Input } from "@/components/ui/Input/input.tsx"
import { Label } from "@/components/ui/Label/label.tsx"
import registerProfile from "@/enteties/Profile/model/service/registerProfile.ts";
import useProfileStore from "@/enteties/Profile/model/store/profileStore.ts";
import useTokenStore from "@/enteties/Token/model/store/tokenStore.ts";
import { toast } from "sonner";
import { formatError } from "@/enteties/Error/model/slice/helpers.ts";
import { Code } from "@/enteties/Error/model/types/error.ts";
import { motion } from 'framer-motion';
import { AuthPageProps } from "@/pages/AuthPage.tsx";

export function SignUpForm({
  className,
  ...props
}: React.ComponentProps<'div'> & AuthPageProps) {
  const {setLogin} = props

  const saveProfile = useProfileStore((state) => state.saveProfile);
  const saveToken = useTokenStore((state)=> state.saveToken);

  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const setLoginHandle = () => {
    setLogin(true);
  }

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      await registerProfile(name, email, password,saveProfile, saveToken)
      toast.success("Registration successful")
    } catch (err) {
      const error = formatError(err as string)
      if (error.code === Code.ALREADY_EXISTS) {
        toast.error("User with the same email already exists")
      } else {
        toast.error(`Registration failed\n ${formatError(err as string).message}`)
      }
    }
  }

  return (
      <motion.div
          initial={{ opacity: 0, y: 100 }}
          animate={{ opacity: 1, y: 0 }}
          exit={{ opacity: 0, y: -100 }}
          transition={{ duration: 0.7 }}
      >
        <div className={cn("flex flex-col gap-6", className)} {...props}>
          <Card className="overflow-hidden">
            <CardContent className="flex justify-center">
              <form className="w-7/12 p-6 md:p-8" onSubmit={onSubmit}>
                <div className="flex flex-col gap-6">
                  <div className="flex flex-col items-center text-center">
                    <h1 className="text-2xl font-bold">Welcome to SpeakUp</h1>
                    <p className="text-balance text-muted-foreground">
                      Sing up in SpeakUp
                    </p>
                  </div>
                  <div className="grid gap-2">
                    <Label htmlFor="name">Name</Label>
                    <Input
                        id="name"
                        type="name"
                        placeholder="Name"
                        required
                        onChange={(e) => {
                          setName(e.target.value)
                        }}
                    />
                  </div>
                  <div className="grid gap-2">
                    <Label htmlFor="email">Email</Label>
                    <Input
                        id="email"
                        type="email"
                        placeholder="mail@example.com"
                        required
                        onChange={(e) => {
                          setEmail(e.target.value)
                        }}
                    />
                  </div>
                  <div className="grid gap-2">
                    <div className="flex items-center">
                      <Label htmlFor="password">Password</Label>
                    </div>
                    <Input
                        id="password"
                        type="password"
                        required
                        onChange={(e) => {
                          setPassword(e.target.value)
                        }}
                    />
                  </div>
                  <Button type="submit" className="w-full">
                    Sign up
                  </Button>
                  <div
                      className="relative text-center text-sm after:absolute after:inset-0 after:top-1/2 after:z-0 after:flex after:items-center after:border-t after:border-border">
                <span className="relative z-10 bg-background px-2 text-muted-foreground">
                  Or continue with
                </span>
                  </div>
                  <div className="flex justify-center">
                    <Button variant="outline" className="w-full">
                      <div className="flex gap-4">
                        <svg className="w-6 h-6" xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="xMidYMid"
                             viewBox="0 0 256 262">
                          <path fill="#4285F4"
                                d="M255.878 133.451c0-10.734-.871-18.567-2.756-26.69H130.55v48.448h71.947c-1.45 12.04-9.283 30.172-26.69 42.356l-.244 1.622 38.755 30.023 2.685.268c24.659-22.774 38.875-56.282 38.875-96.027"/>
                          <path fill="#34A853"
                                d="M130.55 261.1c35.248 0 64.839-11.605 86.453-31.622l-41.196-31.913c-11.024 7.688-25.82 13.055-45.257 13.055-34.523 0-63.824-22.773-74.269-54.25l-1.531.13-40.298 31.187-.527 1.465C35.393 231.798 79.49 261.1 130.55 261.1"/>
                          <path fill="#FBBC05"
                                d="M56.281 156.37c-2.756-8.123-4.351-16.827-4.351-25.82 0-8.994 1.595-17.697 4.206-25.82l-.073-1.73L15.26 71.312l-1.335.635C5.077 89.644 0 109.517 0 130.55s5.077 40.905 13.925 58.602z"/>
                          <path fill="#EB4335"
                                d="M130.55 50.479c24.514 0 41.05 10.589 50.479 19.438l36.844-35.974C195.245 12.91 165.798 0 130.55 0 79.49 0 35.393 29.301 13.925 71.947l42.211 32.783c10.59-31.477 39.891-54.251 74.414-54.251"/>
                        </svg>
                        <div className="">Sign up with Google</div>
                      </div>
                    </Button>
                  </div>
                  <div className="text-center text-sm">
                    Already have account?{" "}
                    <a href="#/" onClick={setLoginHandle} className="underline underline-offset-4">
                      Login
                    </a>
                  </div>
                </div>
              </form>
            </CardContent>
          </Card>
          <div
              className="text-balance text-center text-xs text-muted-foreground [&_a]:underline [&_a]:underline-offset-4 hover:[&_a]:text-primary">
            By clicking continue, you agree to our <a href="#">Terms of Service</a>{" "}
            and <a href="#">Privacy Policy</a>.
          </div>
        </div>
      </motion.div>
  )
}
