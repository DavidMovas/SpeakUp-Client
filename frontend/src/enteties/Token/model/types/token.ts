import { CustomError } from "@/enteties/Error/model/types/error.ts";

export interface AccessTokenState {
    accessToken: string | null;
    isLoading: boolean;
    error: CustomError | null;
    saveToken: (token: string) => void;
    reset: () => void;
}