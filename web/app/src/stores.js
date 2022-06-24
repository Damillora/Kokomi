import { writable } from "svelte/store";

let storedToken = "";
export const token = writable(storedToken);
token.subscribe(value => {
    if (value != null) {
        storedToken = value;
    }
});

let storedRefreshToken = localStorage.getItem("refreshToken") ?? "";

export const refreshToken = writable(storedRefreshToken);
refreshToken.subscribe(value => {
    if (value != null) {
        localStorage.setItem("refreshToken", value);
    }
});
