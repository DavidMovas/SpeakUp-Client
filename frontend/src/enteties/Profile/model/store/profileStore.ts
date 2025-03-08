import { create } from 'zustand';
import { ProfileStore } from "@/enteties/Profile/model/types/profile.ts";


const useProfileStore = create<ProfileStore>((set) => ({
    data: undefined,
    loading: false,
    error: undefined,
    saveProfile: (profile)=> set({data: profile}),
    setError: (error) => set({error: error}),
    reset: () => set({data: undefined, loading: false, error: undefined})
}))

export default useProfileStore;