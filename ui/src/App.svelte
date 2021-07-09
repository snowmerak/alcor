<script>
	import axios from 'axios'
import { loop_guard } from 'svelte/internal';

	let selectedMenu = 0
</script>

<ul uk-tab>
	<li class={selectedMenu==0?"uk-active":""}><a on:click={() => selectedMenu = 0} href=".">Wallet</a></li>
	<li class={selectedMenu==1?"uk-active":""}><a on:click={() => selectedMenu = 1} href=".">Clients</a></li>
	<li class={selectedMenu==2?"uk-active":""}><a on:click={() => selectedMenu = 2} href=".">Wallet</a></li>
	<li class={selectedMenu==3?"uk-active":""}><a on:click={() => selectedMenu = 3} href=".">Transactions</a></li>
    <li class={selectedMenu==4?"uk-active":""}><a on:click={() => selectedMenu = 4} href=".">Pairs</a></li>
</ul>

{#if selectedMenu == 0}
	<!-- svelte-ignore empty-block -->
	<ul uk-accordion="multiple: true">
	{#await axios.get("/wallet/get").then(res => res.data)}
	{:then accounts}
		{#each accounts as account}
		<li>
			<a class="uk-accordion-title" href=".">{account.ID}</a>
			<div class="uk-accordion-content">
				<p>PrivateKey: {account.PrivateKey}</p>
				<p>PublicKey: {account.PublicKey}</p>
			</div>
		</li>
		{/each}
	{/await}
	</ul>
{:else if selectedMenu == 1}
	<!-- svelte-ignore empty-block -->
	<ul uk-accordion="multiple: true">
		{#await axios.get("/client/get").then(res => res.data)}
		{:then clients}
			{#each clients as client}
			<li>
				<a class="uk-accordion-title" href=".">{client.ID}</a>
				<div class="uk-accordion-content">
					<p>PublicKey: {client.PublicKey}</p>
				</div>
			</li>
			{/each}
		{/await}
		</ul>
{/if}

<!-- <ul uk-accordion="multiple: true">
    <li class="uk-open">
        <a class="uk-accordion-title" href=".">Item 1</a>
        <div class="uk-accordion-content">
            <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
        </div>
    </li>
</ul> -->