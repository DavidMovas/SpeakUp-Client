import { create } from 'zustand';
import {  AccessTokenState } from "@/enteties/Token/model/types/token.ts";

const useTokenStore = create<AccessTokenState>((set) => ({
    accessToken: null,
    isLoading: false,
    error: null,
    saveToken: (token: string) => set({accessToken: token}),
    reset: () => set({ accessToken: null, isLoading: false, error: null }),
}))

export default useTokenStore;