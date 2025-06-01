<script lang="ts">
    // import { models } from "wailsjs/go/models";

    export let card;
</script>

<div class="card">
    <img alt={card.Name} src={card.ImagePath}/>
    <div class="card_count">{card.Count}</div>
</div>


<style>
    img {
        margin: 5px;
        width: 200px;
        border-radius: 10px;
        transition: transform .1s;
    }

    img:hover {
        transform: scale(1.5);
    }

    .card {
        display: inline-block;
    }
</style>