<script>
    import {AllCards} from '../../wailsjs/go/main/App.js'
    import Card from '../Card.svelte';
</script>

{#await AllCards() then cards}
<div class="card_list">
  {#each cards as card}
    <Card card={card}/>
  {/each}
</div>
{/await}

<style>
    .card_list{
        padding-top: 50px;
    }
</style>