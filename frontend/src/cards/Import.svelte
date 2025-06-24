<script type="ts">
    import {AllSets, AllCardsSet} from '../../wailsjs/go/main/App.js';
    import Card from '../Card.svelte';
    import {EventsOn} from '../../wailsjs/runtime/runtime';
    // import * as OtherA from '../../wailsjs/go/models.js';

    let cards = new Array();
    function handleSubmit() {
        cards = new Array();
        AllCardsSet("fin");
    }

    EventsOn('get_card', (card) => {
        cards.push(card);
        cards = cards; // update is triggered on assignment
    })
</script>

<div>Import cards</div>

{#await AllSets() then sets}
<div class="set_list">
    <input list="sets">
    <datalist id="sets">
    {#each sets as set}
    <option value="{set.Code.toUpperCase()} ({set.Name})">{set.Name}</option>
    {/each}
    </datalist> 
</div>
{/await}
<button on:click={handleSubmit}> Submit </button>

{#if cards}
<div class="card_list">
{#each cards as card}
    <Card card={card}/>
{/each}
</div>
{/if}