<script>
	import axios from 'axios'
	import { loop_guard } from 'svelte/internal';
	import { get } from 'svelte/store';

	let selectedMenu = 0
	let input = ""
</script>

<ul uk-tab>
	<li class={selectedMenu==0?"uk-active":""}><a on:click={() => selectedMenu = 0} href=".">Wallet</a></li>
	<li class={selectedMenu==1?"uk-active":""}><a on:click={() => selectedMenu = 1} href=".">Clients</a></li>
	<li class={selectedMenu==2?"uk-active":""}><a on:click={() => selectedMenu = 2} href=".">Transactions</a></li>
	<li class={selectedMenu==3?"uk-active":""}><a on:click={() => selectedMenu = 3} href=".">Pairs</a></li>
    <li class={selectedMenu==4?"uk-active":""}><a on:click={() => selectedMenu = 4} href=".">Config</a></li>
</ul>

{#if selectedMenu == 0}
	<form class="uk-grid-small" uk-grid>
	    <div class="uk-width-4-5">
	        <input class="uk-input" type="text" placeholder="ID" value={input}>
	    </div>
	    <div class="uk-width-1-5">
	        <button class="uk-button" on:click={async () => { let result = await axios.put("/wallet", input).then(res => res.data); alert(result) }}>Add</button>
	    </div>
	</form>
	<!-- svelte-ignore empty-block -->
	<ul uk-accordion="multiple: true">
	{#await axios.get("/wallet").then(res => res.data)}
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
		{#await axios.get("/client").then(res => res.data)}
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
{:else if selectedMenu == 4}
	<form class="uk-form-horizontal">
		<div class="uk-margin">
			<!-- svelte-ignore empty-block -->
			{#await axios.get("/config/higheraddress").then(res => res.data)}
			{:then address}
			<label class="uk-form-label" for="form-stacked-text">Higher Node Address</label>
			<div class="uk-inline">
				<!-- svelte-ignore a11y-missing-content -->
				<a class="uk-form-icon" href="." uk-icon="icon: pencil"></a>
				<input class="uk-input" id="form-stacked-text" type="text" value={address} placeholder="higher node address">
			</div>
			{/await}
		</div>
	</form>
{/if}

<!-- <ul uk-accordion="multiple: true">
    <li class="uk-open">
        <a class="uk-accordion-title" href=".">Item 1</a>
        <div class="uk-accordion-content">
            <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
        </div>
    </li>
</ul> -->