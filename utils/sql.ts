import {neon} from "@neondatabase/serverless";

const DATABASE_URL = process.env["DATABASE_URL"];

if (DATABASE_URL === undefined) {
    throw Error("Missing environmental variable DATABASE_URL");
}

export default neon(DATABASE_URL);