<script>
    import Collection from "../cards/Collection.svelte";
    import Import from "../cards/Import.svelte";

    let component = Collection;
</script>

<div class="app">
    <div class="menu">
        <input bind:group={component} type="radio" value={Collection} id="collection"/><label for="collection">Collection</label>
        <input bind:group={component} type="radio" value={Import} id="import"/><label for="import">Import</label>
    </div>
    
    <div class="menu_content">
        <svelte:component this={component}/>
    </div>
</div>

<style>
    .app {
        display: flex;
    }

    .menu {
        background-color: green;
        height: 100%;
        padding: 10px;

        /* put the menu on the left side */
        flex-grow: 1;

        /* disply inner items as a column */
        display: flex;
        flex-direction: column;
    }

    .menu_content {
        width: 100%;
    }

    input {
        /* hide the radio input button so we only have the labels */
        display: none;
    }

    input[type="radio"]:checked + label {
        background: red;
    }

</style>