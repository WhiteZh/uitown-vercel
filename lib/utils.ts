import {match, P} from "ts-pattern";

export type ErrorResponseType = {
    error: string;
}

export class ErrorWithCode {
    public err: Error;
    public code: number;

    public constructor(err: Error | string, code: number) {
        this.err = match(err)
            .with(P.select(P.string), msg => Error(msg))
            .with(P.select(P._), it => it)
            .exhaustive()

        this.code = code;
    }

    public createErrorResponse(): Response {
        return Response.json({
            error: this.err.message,
        } satisfies ErrorResponseType, {
            status: this.code,
        });
    }
}

export function crudWrapper<T>(method: (request: Request) => Promise<T | ErrorWithCode>): (request: Request) => Promise<Response> {
    return async (request) => {
        const result: ErrorWithCode | T = await method(request);
        if (result instanceof ErrorWithCode) {
            console.error(result.err)
            return result.createErrorResponse();
        }
        return Response.json(result);
    }
}

export const wrapBadResponse = (resBody: unknown): Error => match(resBody)
    .with({
        error: P.select(P.string)
    }, err => Error(err))
    .otherwise(() => Error("Bad request"));

export const createUnexpectedServerResponseError = () => Error("Unexpected response from server");

export const createUnexpectedServerResponseErrorWithCode = () => new ErrorWithCode(createUnexpectedServerResponseError(), 500);