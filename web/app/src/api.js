import { token, refreshToken } from "./stores";
import axios from "axios";

let userToken = "";
token.subscribe((value) => {
    userToken = value;
});

let url = window.BASE_URL;

export async function login({ username, password }) {
    const endpoint = url + "/api/auth/login";
    const response = await axios({
        url: endpoint,
        method: "POST",
        data: JSON.stringify({
            username,
            password,
        }),
    })
    token.set(response.data.token);
    refreshToken.set(response.data.refreshToken);
    return response.data;
}
export async function refresh({ refresh }) {
    const endpoint = url + "/api/auth/refresh";
    const response = await axios({
        url: endpoint,
        method: "POST",
        data: JSON.stringify({
            refreshToken: refresh,
        }),
    })
    token.set(response.data.token);
    refreshToken.set(response.data.refreshToken);
    return response.data;
}


export async function register({ email, username, password }) {
    const endpoint = url + "/api/user/register";
    const response = await axios({
        url: endpoint,
        method: "POST",
        data: JSON.stringify({
            email,
            username,
            password,
        }),
    })
    return response.data;
}

export async function getUserProfile() {
    const endpoint = url + "/api/user/profile";
    const response = await axios({
        url: endpoint,
        method: "GET",
        headers: {
            'Authorization': 'Bearer ' + userToken,
        },
        withCredentials: true,
    });
    return response.data;
}
