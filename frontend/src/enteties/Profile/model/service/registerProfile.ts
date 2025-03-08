import useTokenStore from "@/enteties/Token/model/store/tokenStore.ts";
import useProfileStore from "@/enteties/Profile/model/store/profileStore.ts";
import { Profile } from "@/enteties/Profile/model/types/profile.ts";
import { RegisterUser } from "../../../../../wailsjs/go/handlers/UsersHandler";
import { formatError } from "@/enteties/Error/model/slice/helpers.ts";
import { CustomError } from "@/enteties/Error/model/types/error.ts";


const registerProfile = async (
    name: string,
    email: string,
    password: string,
    saveToken: (token: string) => void,
    savaProfile: (profile: Profile) => void,
    setError: (err: CustomError) => void
    ): Promise<void> => {
    useProfileStore.getState().reset();
    useTokenStore.getState().reset();

    try {
        const response = await RegisterUser(name, email, password)

        if (response.accessToken) {
            saveToken(response.accessToken)
        }

        if (response.user) {
            const profile: Profile = {
                id: response.user.id,
                email: response.user.email,
                username: response.user.username,
                fullName: response.user.fullName,
                bio: response.user.bio,
                avatarUrl: response.user.avatarUrl,
                createdAt: new Date(response.user.createdAt),
            }

            savaProfile(profile)

            return Promise.resolve()
        }
    } catch (err) {
        const customError = formatError(err as string);

        setError(customError)

        return Promise.reject(customError)
    }
}

export default registerProfile