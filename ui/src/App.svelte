<script>
import Clients from "./Clients.svelte";
import Config from "./Config.svelte";
import Wallet from "./wallet/Wallet.svelte";

let selectedMenu = 0
</script>

<!-- svelte-ignore a11y-missing-attribute -->
<nav class="tabs is-full">
    <a class={selectedMenu==0?"active":""} on:click={()=>selectedMenu=0}>Wallet</a>
    <a class={selectedMenu==1?"active":""} on:click={()=>selectedMenu=1}>Clients</a>
    <a class={selectedMenu==4?"active":""} on:click={()=>selectedMenu=4}>Config</a>
</nav>

{#if selectedMenu == 0}
<Wallet></Wallet>
{:else if selectedMenu == 1}
<Clients></Clients>
{:else if selectedMenu == 4}
<Config></Config>
{/if}

<!-- <ul uk-accordion="multiple: true">
    <li class="uk-open">
        <a class="uk-accordion-title" href=".">Item 1</a>
        <div class="uk-accordion-content">
            <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
        </div>
    </li>
</ul> -->