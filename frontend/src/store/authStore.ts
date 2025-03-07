import {create} from 'zustand';

interface AuthState {
    accessToken: string | null;
    setAccessToken: (accessToken: string) => void;
    clearAccessToken: () => void;
}

const useAuthStore = create<AuthState>((set) => ({
    accessToken: null,
    setAccessToken: (accessToken) => set((state) => ({...state, accessToken })),
    clearAccessToken: () => set((state) => ({...state, accessToken: null })),
}))

export default useAuthStore;

