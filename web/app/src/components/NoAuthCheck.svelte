<script>
    import { isTokenExpired } from "../login-check";
    import { token } from "../stores";

    let loggedIn = false;
    token.subscribe((value) => {
        loggedIn = !isTokenExpired(value);
    });
</script>

{#if loggedIn == false}
    <slot />
{/if}
