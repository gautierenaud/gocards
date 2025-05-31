<script>
  import {AllCards} from '../wailsjs/go/main/App.js'
  import Card from './Card.svelte';
</script>

<main>
  {#await AllCards() then cards}
  <div class="card_list">
    {#each cards as card}
      <Card card={card}/>
    {/each}
  </div>
  {/await}
</main>

<style>
  .card_list {
    margin-top: 100px;
    margin-bottom: 100px;
  }
</style>
