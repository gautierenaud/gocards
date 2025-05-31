<script>
  import {AllCards} from '../wailsjs/go/main/App.js'
</script>

<main>
  {#await AllCards() then cards}
  {#each cards as card}
    <li><img src={card.ImagePath} alt={card.Name}/></li>
  {/each}
  {/await}
</main>

<style>
</style>
