import {create} from 'zustand';
import { models } from "../../wailsjs/go/models.ts";

interface UserStore {
    user: models.User | null;
    setUser: (user: models.User) => void,
    clearUser: () => void,
}

const useUserStore = create<UserStore>((set) => ({
    user: null,
    setUser: (user: models.User) => set((state) => ({...state, user })),
    clearUser: () => set((state) => ({...state, user: null })),
}))

export default useUserStore;