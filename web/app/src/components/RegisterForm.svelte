<script>
    import { register } from "../api.js";
    import { navigate } from "svelte-routing";

    let username = "";
    let password = "";
    let email = "";
    let error = "";

    const doRegister = async () => {
        try {
            const tokenData = await register({ email, username, password });
            navigate("/");
        } catch (error) {
            error = "We had trouble registering you";
        }
    };
</script>

<form on:submit|preventDefault={doRegister}>
    <div>
        <label for="email" class="label">Email</label>
    </div>
    <div>
        <input
            id="email"
            class="input"
            type="text"
            placeholder="Email"
            bind:value={email}
            required
        />
    </div>
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
    <div>
        <button class="button is-link">Register</button>
    </div>
</form>
