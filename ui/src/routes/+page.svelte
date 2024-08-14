<script>
	import { UserService } from '$lib/proto/outward/v1/user_service_connect';
	import { createPromiseClient } from '@connectrpc/connect';
	import { createConnectTransport } from '@connectrpc/connect-web';

	import Counter from './Counter.svelte';
	import welcome from '$lib/images/svelte-welcome.webp';
	import welcome_fallback from '$lib/images/svelte-welcome.png';
	import { GetUserListRequest } from '$lib/proto/outward/v1/user_service_pb';
	import { onMount } from 'svelte';

	onMount(() => {
		const transport = createConnectTransport({
			baseUrl: '/',
			// We pass `fetch` provided by Svelte to the Transport. The function
			// behaves the same as native fetch(), but it inherits cookies, and
			// it can make relative requests, so you don't have to specify an
			// absolute baseUrl.
			// For more information, see https://kit.svelte.dev/docs/load#making-fetch-requests
			fetch
		});
		const client = createPromiseClient(UserService, transport);

		const request = new GetUserListRequest({ pageSize: 10 });

		const response = client.getUserList(request);

		console.log(response);
	});
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<h1>
		<span class="welcome">
			<picture>
				<source srcset={welcome} type="image/webp" />
				<img src={welcome_fallback} alt="Welcome" />
			</picture>
		</span>

		to your new<br />SvelteKit app
	</h1>

	<h2>
		try editing <strong>src/routes/+page.svelte</strong>
	</h2>

	<Counter />
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
	}

	h1 {
		width: 100%;
	}

	.welcome {
		display: block;
		position: relative;
		width: 100%;
		height: 0;
		padding: 0 0 calc(100% * 495 / 2048) 0;
	}

	.welcome img {
		position: absolute;
		width: 100%;
		height: 100%;
		top: 0;
		display: block;
	}
</style>
