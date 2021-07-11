<script>
import axios from 'axios';

let input = "";
let selectedWallet;

async function PutWallet(input) {
    let data = await axios.put("/wallet", input).then(response => {
        return response.data;
    });
    if (data.Error != undefined) {
        alert(data.Error);
        return;
    }
    selectedWallet = data.ID;
}
</script>

{#if selectedWallet == undefined}
    <br />
    <div class="grouped">
        <input type="text" bind:value={input} placeholder="Enter New ID" />
        <!-- svelte-ignore a11y-missing-attribute -->
        <button class="button primary" on:click={async () => { await PutWallet(input); }}>Create</button>
    </div>
    <!-- svelte-ignore empty-block -->
    {#await axios.get('/wallet').then(res => res.data)}
    {:then wallets }
        {#each wallets as wallet}
        <br />
        <!-- svelte-ignore a11y-missing-attribute -->
        <div class="card">
            <header>
              <h4>{wallet}</h4>
            </header>
            <footer class="is-right">
              <a class="button primary" on:click={() => {selectedWallet=wallet;}}>Detail</a>
              <a class="button">Remove</a>
            </footer>
        </div>
        {/each}
    {/await}
{:else}
    <!-- svelte-ignore empty-block -->
    {#await axios.get(`/wallet/${selectedWallet}`).then(res => res.data)}
    {:then wallet} 
    <br />
    <!-- svelte-ignore a11y-missing-attribute -->
    <div class="card">
        <header>
          <h4>{wallet.ID}</h4>
        </header>
        <h5>Used:</h5>
        <p>{wallet.Used}</p>
        <h5>PrivateKey:</h5>
        <p>{wallet.PrivateKey}</p>
        <h5>PublicKey:</h5>
        <p>{wallet.PublicKey}</p>
        <footer class="is-right">
          <a class="button primary">Use</a>
          <a class="button" on:click={() => {selectedWallet=undefined;}}>Back</a>
        </footer>
    </div>
    {/await}
{/if}