<script>
  import {AllCards, AllSets} from '../wailsjs/go/main/App.js'
  import Card from './Card.svelte';
  import { LogPrint, EventsOn } from '../wailsjs/runtime/runtime';
  EventsOn("card_changed", () => { LogPrint("qergjhbqergkhjbqerjhf;bqer");  });
</script>

<main>
  {#await AllCards() then cards}
  <div class="card_list">
    {#each cards as card}
      <Card card={card}/>
    {/each}
  </div>
  {/await}

  {#await AllSets() then sets}
  <div class="set_list">
    <input list="sets">
    <datalist id="sets">
      {#each sets as set}
      <option value={set}></option>
      {/each}
    </datalist> 
    <select id="sets" name="sets" size=10 multiple>
      {#each sets as set}
      <option value={set.Code}>{set.Code.toUpperCase()} ({set.Name})</option>
      {/each}
    </select>
  </div>
  {/await}
</main>

<style>
  .card_list {
    margin-top: 100px;
    margin-bottom: 100px;
  }
</style>
