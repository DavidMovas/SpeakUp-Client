export enum Code {
    UNKNOWN = "Unknown",
    OK = "OK",
    INTERNAL = "Internal",
    CANCELED= "Canceled",
    INVALID_ARGUMENT = "InvalidArgument",
    PERMISSION_DENIED = "PermissionDenied",
    UNAUTHENTICATED = "Unauthenticated",
    FAILED_PRECONDITION = "FailedPrecondition",
    NOT_FOUND = "NotFound",
    ALREADY_EXISTS = "AlreadyExists",
    OUT_OF_RANGE = "OutOfRange",
    UNIMPLEMENTED = "Unimplemented",
    UNAVAILABLE = "Unavailable",
    DATA_LOSS = "DataLoss",
    ABORTED = "Aborted",
}

export interface CustomError extends Error{
    code: Code;
    message: string;
}