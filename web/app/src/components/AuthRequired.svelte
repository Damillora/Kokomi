<script>
    import { refresh } from "../api.js";
    import { navigate } from "svelte-routing";
    import { onMount } from "svelte";
    import { isTokenExpired } from "../login-check.js";
    import { refreshToken, token } from "../stores";
    
    let loggedIn = false;
    token.subscribe((value) => {
        loggedIn = !isTokenExpired(value);
    })
    onMount(async () => {
        const apiToken = getApiToken();
        let loggedIn = !isTokenExpired(apiToken);
        if (loggedIn === false) {
            // Try refreshing token
            const result = await refresh(getRefreshToken());

            if(!result.token) {
                token.set("");
                refreshToken.set("");
                navigate("/");
            }
        }
    });
</script>
