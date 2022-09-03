<script lang="ts">
  import lline from "../assets/images/llline-2.svg";
  import bblury from "../assets/images/bbblurry-2.svg";
  import { router } from "tinro";

  let url;

  function start() {
    const id = new URL(url).searchParams.get("v");

    router.goto(`/${id}`);
  }
</script>

<main>
  <form on:submit|preventDefault={start}>
    <input
      type="url"
      bind:value={url}
      placeholder="https://youtube.com/?v=xxxxxxxxx or id"
    />
    <button type="submit">
      <i class="ai-send" />
    </button>
  </form>
  <img src={lline} alt="">
  <img src={bblury} alt="" style:width="100vw" style:bottom=0>
</main>

<style>
  main {
    display: grid;
    place-items: center;
    position: absolute;
    inset: 0;
    color: var(--text-color);

    --wails-draggable: drag;
  }

  form {
    width: fit-content;
    color: inherit;

    display: flex;
    gap: 1em;

    position: relative;
    overflow: visible;
  }

  form input,
  form button {
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

  input {
    width: 50vw;
  }

  input::placeholder {
    color: var(--text-color);
    opacity: 0.5;
  }

  button {
    cursor: pointer;
    aspect-ratio: 1;
  }

  img {
    position: absolute;
    z-index: -1;
    width: 50vw;
    object-position: center;
    opacity: 0.75;
    mix-blend-mode: multiply;
  }
</style>
