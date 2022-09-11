<script lang="ts">
  import { router } from "tinro";

  import { Dialog } from "../../wailsjs/go/main/App.js";

  import { WindowToggleMaximise } from "../../wailsjs/runtime/runtime.js";
  export let url = "https://www.youtube.com/watch?v=GV3HUDMQ-F8";
  let newURL = url;

  async function updateURL() {
    if (newURL == "")
      await Dialog(
        "Invalid URL",
        "Please enter a valid YouTube URL",
        "warning"
      );
    else {
      url = newURL;
    }
  }

  export let downloadProgress: string;
</script>

<nav on:dblclick|self={WindowToggleMaximise}>
  <span>
    <i
      class="ai-arrow-left"
      on:click={(_) => {
        router.goto("/");
      }}
      style:font-size="1.5em"
    />
  </span>
  <form on:submit|preventDefault={updateURL}>
    <input type="url" bind:value={newURL} />
    <button type="submit">
      <i class="ai-send" style:font-size="1.5em" />
    </button>
  </form>
  <span id="download-progress">
    {downloadProgress}
  </span>
</nav>

<style>
  nav {
    position: fixed;
    z-index: 100;
    inset: 0 0 auto 0;
    display: grid;
    grid-template-columns: 1fr auto 1fr;
    align-items: center;
    text-align: center;
    padding: 0.5em;
    font-size: 0.75em;
    background-color: var(--bg-color);
    color: var(--text-color);
    -webkit-backdrop-filter: blur(0.5em);
    backdrop-filter: blur(0.5em);
    box-shadow: 0px 10px 15px -3px rgba(0, 0, 0, 0.1);
    transform: translateZ(75em);

    --wails-draggable: drag;
  }

  nav > form {
    width: fit-content;
    color: inherit;

    display: flex;
    gap: 1em;
  }

  nav > form > * {
    --wails-draggable: no-drag;
    padding: 0.5em 1em;
    color: inherit;
    background-color: var(--bg-color);
    -webkit-backdrop-filter: blur(0.25em) brightness(2);
    backdrop-filter: blur(0.25em);
    outline: none;
    border: none;
    border-radius: 50vmax;

    box-shadow: 0 0 0.25em 0.25em rgba(0, 0, 0, 0.25);
  }

  nav > form > input {
    width: 50vw;
  }

  nav > form > button {
    cursor: pointer;
    aspect-ratio: 1;
  }
</style>
