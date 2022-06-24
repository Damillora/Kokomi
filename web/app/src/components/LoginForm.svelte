<script>
    import { login } from "../api.js";
    import { navigate } from "svelte-routing";

    let username = "";
    let password = "";
    let error = "";

    const doLogin = async () => {
        error = "";
        try {
            const tokenData = await login({ username, password });
            navigate("/");
        } catch (e) {
            error = "We had trouble logging you in";
            return;
        }
    };
</script>

<form on:submit|preventDefault={doLogin}>
    <div>
        <label for="username" class="label">Username</label>
    </div>
    <div>
        <input
            id="username"
            class="input"
            type="text"
            placeholder="Username"
            bind:value={username}
            required
        />
    </div>
    <div>
        <label for="password" class="label">Password</label>
    </div>
    <div>
        <input
            id="password"
            class="input"
            type="password"
            placeholder="Password"
            bind:value={password}
            required
        />
    </div>
    {#if error}
        <div>
            <p class="has-text-danger">{error}</p>
        </div>
    {/if}
    <div>
        <button>Login</button>
    </div>
</form>
