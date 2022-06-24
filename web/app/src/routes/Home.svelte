<script>
  import Menu from "../components/Menu.svelte";
  import { getAutocompleteItems } from "../autocompleter";
  import { navigate } from "svelte-routing";

  let isActive = false;
  let isWidgetActive = false;

  let searchTerm = "";
  let autocompleteItems = [];
  let selectedItemIndex;
  $: {
    if (searchTerm) {
      isActive = true;
      getAutocompleteItems(searchTerm).then((items) => {
        autocompleteItems = items;
        if (items.length > 0) {
          selectedItemIndex = 0;
        } else {
          selectedItemIndex = undefined;
        }
      });
    } else {
      isActive = false;
      selectedItemIndex = undefined;
    }
  }
  const selectItem = (i) => {
    selectedItemIndex = i;
  };
  const executeTerm = (term) => {
    switch (term.type) {
      case "Google":
        navigate(term.url);
        break;
    }
  };
  let doSearch = () => {
    if (selectedItemIndex) {
      executeTerm(autocompleteItems[selectedItemIndex]);
    } else {
      if (autocompleteItems.length > 0) {
        executeTerm(autocompleteItems[0]);
      }
    }
  };
</script>

<Menu {isWidgetActive} />

<main>
  <section class="search-section" class:inactive={isWidgetActive}>
    <h1 class="search-section__search-header">
      <strong>start</strong>&nbsp;<span>here</span>
    </h1>
    <form class="search-section__search" on:submit|preventDefault={doSearch}>
      <input
        class="search-section__search-box"
        type="text"
        name="search"
        id="search-query"
        bind:value={searchTerm}
      />
      <button type="submit">
        <span class="material-icons search-section__search-icon">
          search
        </span></button
      >
      <div class="search-section__autocomplete" class:active={isActive}>
        {#each autocompleteItems as item, i}
          <a
            href={item.url}
            class="search-section__autocomplete-item"
            class:selected={item.url ==
              autocompleteItems[selectedItemIndex].url}
            on:mouseenter={() => selectItem(i)}
          >
            <span class="material-icons search-section__autocomplete-icon"
              >{item.icon}</span
            >
            <span class="search-section__autocomplete-type">{item.type}</span>
            <span class="search-section__autocomplete-text">{item.text}</span>
          </a>
        {/each}
      </div>
    </form>
  </section>
  <section class="widget-section">
    <div
      class="widget-collapse"
      on:click={() => (isWidgetActive = !isWidgetActive)}
    >
      {#if isWidgetActive}
        <span class="material-icons">expand_more</span>
      {:else}
        <span class="material-icons">expand_less</span>
      {/if}
    </div>
    <div class="widget-area" class:active={isWidgetActive} />
  </section>
</main>

<style lang="scss">
</style>
