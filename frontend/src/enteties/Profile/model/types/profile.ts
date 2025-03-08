import { CustomError } from "@/enteties/Error/model/types/error.ts";

export interface Profile {
    id: string
    email: string
    username: string
    fullName: string
    bio: string | null
    avatarUrl?: string
    createdAt: Date
}

export interface ProfileStore {
    data: Profile | null;
    loading: boolean;
    error: CustomError | null;
    saveProfile: (profile: Profile) => void;
    setError: (error: CustomError) => void;
    reset: () => void;
}