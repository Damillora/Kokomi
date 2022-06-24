<script>
    import AuthCheck from "./AuthCheck.svelte";
    import NoAuthCheck from "./NoAuthCheck.svelte";
    import LoginForm from "./LoginForm.svelte";
    import RegisterForm from "./RegisterForm.svelte";

    import { token, refreshToken } from "../stores";
    import { navigate } from "svelte-routing";

    let isVisible = false;

    export let isWidgetActive = false;

    let logOut = () => {
        refreshToken.set("");
        token.set("");
        navigate("/");
    };

    let isRegister = false;
</script>

<div
    class="menu-button"
    on:click={() => (isVisible = !isVisible)}
    class:widget-active={isWidgetActive}
>
    <span class="material-icons menu-button__icon"> settings </span>
</div>

<div class="menu" class:active={isVisible}>
    <AuthCheck>
        <div>Welcome!</div>
        <div>
            <button on:click={logOut}>Logout</button>
        </div>
    </AuthCheck>
    <NoAuthCheck>
        {#if isRegister}
            <RegisterForm />
            <a href={'#'} on:click={() => (isRegister = false)}
                >Login</a
            >
        {:else}
            <LoginForm />
            <a href={'#'} on:click={() => (isRegister = true)}
                >Register</a
            >
        {/if}
    </NoAuthCheck>
</div>
