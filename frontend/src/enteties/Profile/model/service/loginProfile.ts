import { LoginUser } from "../../../../../wailsjs/go/handlers/UsersHandler";
import { Profile } from "@/enteties/Profile/model/types/profile.ts";

const loginProfile = async (
    email: string,
    password: string,
    saveProfile: (profile: Profile) => void,
    saveToken: (token: string) => void,
): Promise<void> => {
    try {
        const response = await LoginUser(email, password)

        if (response.accessToken) {
            saveToken(response.accessToken)
        }

        if (response.user) {
            saveProfile({
                id: response.user.id,
                email: response.user.email,
                username: response.user.username,
                fullName: response.user.fullName,
                avatarUrl: response.user.avatarUrl,
                bio: response.user.bio,
                lastLoginAt: response.user.lastLoginAt? new Date(response.user.lastLoginAt) : undefined,
                createdAt: new Date(response.user.createdAt),
                updatedAt: response.user.updatedAt? new Date(response.user.updatedAt) : undefined,
            })
        }

        return Promise.resolve()
    } catch (err) {
        return  Promise.reject(err as Error)
    }
}

export default loginProfile