import {match, P} from "ts-pattern";
import sql from "../../utils/sql.ts";
import {createUnexpectedServerResponseErrorWithCode, crudWrapper, ErrorWithCode} from "../../utils/utils.ts";

export const GET = crudWrapper<number>(async (request) => {

    const urlSearchParams = new URLSearchParams(new URL(request.url).search);

    const params: {
        email: string,
        password_hashed: string,
    } | ErrorWithCode = match({
        email: urlSearchParams.get("email"),
        password_hashed: urlSearchParams.get("password_hashed")
    })
        .with({
            email: P.string,
            password_hashed: P.string
        }, (it) => it)
        .otherwise(() => new ErrorWithCode("Bad or incomplete request", 400));

    if (params instanceof ErrorWithCode) {
        return params;
    }

    const res: {
        password_hashed: string,
        id: number,
    } | ErrorWithCode = match(await sql`SELECT id, password_hashed FROM users WHERE email = ${params.email}`)
        .with([{password_hashed: P.string, id: P.number}], ([it]) => it)
        .otherwise(createUnexpectedServerResponseErrorWithCode)

    if (res instanceof ErrorWithCode) {
        return res;
    }

    return res.password_hashed === params.password_hashed ? res.id : -1;
});