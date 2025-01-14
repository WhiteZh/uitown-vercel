// export const cssCategories: string[] = [
//     "button",
//     "checkbox",
//     "toggle switch",
//     "card",
//     "loader",
//     "input",
//     "transition",
//     "special effect"
// ];

// export const jsCategories = [
//
// ];

import {notifications, user} from "@/globs";
import {getUserById} from "@/api";

export const iframeContent = (html: string, css: string) => `
  <!DOCTYPE html>
  <html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Embedded Content</title>
    <style>
      body {
        display: flex;
        justify-content: center;
        align-items: center;
        margin: 0;
        background-color: #2b2a2a;
        height: 100vh;
        width: 100vw;
      }
      ${css}
    </style>
  </head>
  <body>
    <div>
      ${html}
    </div>
  </body>
  </html>
`;

export const shadowContent = (html: string, css: string) => `
<div style="position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%);" id="the-id-of-the-shadow-root">
    <style>${css}</style>
    ${html}
</div>
`;

export function isOfType(o: unknown, properties: Record<string, (x: unknown) => boolean>, optional?: Record<string, (x: unknown) => boolean>): boolean {
    if (!(typeof o === "object" && o !== null)) {
        return false;
    }

    for (let key of Object.keys(properties)) {
        if (!(key in o && properties[key]((o as Record<string, unknown>)[key]))) {
            return false;
        }
    }

    if (optional !== undefined) {
        for (let key of Object.keys(optional)) {
            if (key in o && !optional[key]((o as Record<string, unknown>)[key])) {
                return false;
            }
        }
    }

    return true;
}

export type User = {
    id: number,
    name: string,
    email: string,
    password_hashed: string,
    aboutme: string,
    icon: string | null,
};

export type Notification = {
    message: string,
    color?: string
};

export const cssCategories = [
    "button",
    "checkbox",
    "toggle switch",
    "loader",
    "card",
    "input",
    "transition",
    "special effect",
] as const;
export type CSSCategory = typeof cssCategories[number];
export const isCSSCategory = (o: unknown): o is CSSCategory => typeof o === "string" && (cssCategories as any as string[]).includes(o)

export type CSSStyle = {
    id: number,
    name: string,
    viewed_time: number,
    author_id: number,
    html: string,
    css: string,
    category: CSSCategory
};

export async function updateUser() {
    if (user.value === undefined) return;

    let res = await getUserById(user.value.id, user.value.password_hashed);
    if (res instanceof Error) {
        notifications.push({message: res.message});
    } else {
        user.value = res;
    }
}