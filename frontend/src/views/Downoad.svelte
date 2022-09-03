<script lang="ts">
  import ms from "ms";

  import { Dialog, GetVideoDetails } from "../../wailsjs/go/main/App.js";
  import FormatsDownload from "../components/FormatsDownload.svelte";
  import NavBar from "../components/NavBar.svelte";
  import Preview from "../components/Preview.svelte";

  export let id: string;

  let downloadProgress: string = "";

  let url: string = `https://www.youtube.com/watch?v=${id}`;
</script>

<NavBar bind:url bind:downloadProgress />

<main>
  {#key url}
    {#await GetVideoDetails(url)}
      <span class="eye-loader" />
    {:then result}
      {@const { videoDetails, streamingData } = JSON.parse(result)}
      <Preview {videoDetails} urls={streamingData.formats} />
      <br />
      {@const {
        adaptiveFormats,
        expiresInSeconds: expiry,
        formats,
      } = streamingData}
      <span class="expiry">
        Links will expire in {ms(ms(expiry + "s"))}
      </span>
      <FormatsDownload
        title={videoDetails.title}
        {formats}
        bind:downloadProgress
      />
      <FormatsDownload
        title={videoDetails.title}
        formats={adaptiveFormats}
        bind:downloadProgress
      />
    {:catch error}
      {Dialog("Error", error, "error")}
    {/await}
  {/key}
</main>

<style>
  main {
    margin-top: 5em;
    display: grid;
    place-items: center;
    min-height: 100vh;
    min-width: 100vw;
  }

  .expiry {
    font-size: 0.75em;
    color: var(--text-color);
    padding: 0.5em;
    background-color: var(--bg-color);
    -webkit-backdrop-filter: blur(0.5em);
    backdrop-filter: blur(0.5em);
    border-radius: 0.5em;
  }

  /* Loaders */
  .eye-loader {
    position: relative;
    width: 78px;
    height: 78px;
    border-radius: 50%;
    box-sizing: border-box;
    background: #fff;
    border: 8px solid #131a1d;
    overflow: hidden;
    box-sizing: border-box;
  }
  .eye-loader::after {
    content: "";
    position: absolute;
    left: 0;
    top: -50%;
    width: 100%;
    height: 100%;
    background: #263238;
    z-index: 5;
    border-bottom: 8px solid #131a1d;
    box-sizing: border-box;
    animation: eyeShade 3s infinite;
  }
  .eye-loader::before {
    content: "";
    position: absolute;
    left: 20px;
    bottom: 15px;
    width: 32px;
    z-index: 2;
    height: 32px;
    background: #111;
    border-radius: 50%;
    animation: eyeMove 3s infinite;
  }
  @keyframes eyeShade {
    0% {
      transform: translateY(0);
    }
    20% {
      transform: translateY(5px);
    }
    40%,
    50% {
      transform: translateY(-5px);
    }
    60% {
      transform: translateY(-8px);
    }
    75% {
      transform: translateY(5px);
    }
    100% {
      transform: translateY(10px);
    }
  }
  @keyframes eyeMove {
    0% {
      transform: translate(0, 0);
    }
    20% {
      transform: translate(0px, 5px);
    }
    40%,
    50% {
      transform: translate(0px, -5px);
    }
    60% {
      transform: translate(-10px, -5px);
    }
    75% {
      transform: translate(-20px, 5px);
    }
    100% {
      transform: translate(0, 10px);
    }
  }
</style>
