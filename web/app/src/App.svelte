<script>
	import { Router, Link, Route } from "svelte-routing";

	import { refresh } from "./api";
    import { isTokenExpired } from "./login-check.js";
	import { token, refreshToken } from "./stores";

	import Home from "./routes/Home.svelte";
	import Login from "./routes/Auth/Login.svelte";

	export let url = "";
	let baseURL = window.BASE_URL;

	let userRefreshToken = "";
	refreshToken.subscribe((value) => {
		userRefreshToken = value;
	});

	let loggedIn = false;
	token.subscribe(async (value) => {
		loggedIn = !isTokenExpired(value);

		if (loggedIn === false) {
			// Try refreshing token
			const result = await refresh({ refresh: userRefreshToken });
		}
	});
</script>

<Router {url}>
	<div>
		<Route path="/" component={Home} />
		<Route path="/login" component={Login} />
	</div>
</Router>

<style global lang="scss">
	@import "./style/main.scss";
</style>
