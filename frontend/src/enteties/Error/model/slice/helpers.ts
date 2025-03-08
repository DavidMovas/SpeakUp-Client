import { Code, CustomError } from "@/enteties/Error/model/types/error.ts";

export function formatError(error: string): CustomError {
    const regex = /code = (\w+) desc = (.+)/;
    const match = RegExp(regex).exec(error);

    if (match) {
        return {
            name: "CustomError",
            code: formatCode(match[1]),
            message: match[2],
        }
    }

    return {
        name: "Error",
        code: Code.INTERNAL,
        message: "An error occurred while formatting the error message",
    }
}

function formatCode(code: any): Code {
    switch (code) {
        case 'Unknown':
            return Code.UNKNOWN;
        case 'OK':
            return Code.OK;
        case 'Internal':
            return Code.INTERNAL;
        case 'Canceled':
            return Code.CANCELED;
        case 'InvalidArgument':
            return Code.INVALID_ARGUMENT;
        case 'PermissionDenied':
            return Code.PERMISSION_DENIED;
        case 'Unauthenticated':
            return Code.UNAUTHENTICATED;
        case 'FailedPrecondition':
            return Code.FAILED_PRECONDITION;
        case 'NotFound':
            return Code.NOT_FOUND;
        case 'AlreadyExists':
            return Code.ALREADY_EXISTS;
        case 'OutOfRange':
            return Code.OUT_OF_RANGE;
        case 'Unimplemented':
            return Code.UNIMPLEMENTED;
        case 'Unavailable':
            return Code.UNAVAILABLE;
        default:
            return Code.INTERNAL;
    }
}