import { create } from 'zustand';
import { ProfileStore } from "@/enteties/Profile/model/types/profile.ts";

const useProfileStore = create<ProfileStore>((set) => ({
    data: null,
    loading: false,
    error: null,
    saveProfile: (profile)=> set({data: profile}),
    setError: (error) => set({error: error}),
    reset: () => set({data: null, loading: false, error: null})
}))

export default useProfileStore;