<script lang="ts">
  import prettyBytes from "pretty-bytes";
  import { getExtension } from "mime/lite";

  import {
    Dialog,
    DownloadFile,
    GetDownloadProgress,
    GetDownloadSize,
  } from "../../wailsjs/go/main/App";

  export let formats;

  export let downloadProgress = "";
  export let title = "video";

  async function downloadFile(url, name) {
    DownloadFile(url, name)
      .then(async () => {
        downloadProgress = "Download Complete";
        await Dialog("Download Complete", "Download complete", "info");
      })
      .catch(async (err) => {
        downloadProgress = "Download Failed";
        await Dialog("Download Failed", err, "error");
      });

    const timer = setInterval(async () => {
      const progress = await GetDownloadProgress();
      downloadProgress = `Downloading ${progress}%`;
      if (progress == 100) {
        downloadProgress = "Download Complete";
        clearInterval(timer);
      }
    }, 0.4e3);
  }
</script>

<div class="formats">
  {#each formats as { url, qualityLabel, quality, mimeType }}
    <div class="format" data-type={mimeType.split("/")[0]}>
      <span class="quality">{qualityLabel ?? quality}</span>
      {#await GetDownloadSize(url)}
        <span class="download-loader" />
      {:then size}
        <span class="size">{prettyBytes(size)}</span>
      {/await}
      <a
        href={url}
        on:click|preventDefault={() => {
          downloadFile(
            title + "-" + quality + "." + getExtension(mimeType),
            url
          );
        }}
      >
        <i class="ai-download" />
      </a>
    </div>
  {/each}
</div>

<style>
  .formats {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    justify-content: center;
    gap: 1em;
    padding: 2em 0.5em;
  }

  .format {
    display: flex;
    flex-direction: row;
    place-items: center;
    gap: 0.5em;
    padding: 0.5em;
    background-color: var(--bg-color);
    -webkit-backdrop-filter: blur(0.5em);
    backdrop-filter: blur(0.5em);
    border-radius: 0.5em;
    box-shadow: 0px 10px 15px -3px rgba(0, 0, 0, 0.1);

    position: relative;
  }

  .format::after {
    content: attr(data-type);
    display: block;
    text-align: center;
    font-size: 0.5em;
    color: var(--text-color);
    position: absolute;
    top: -0.5em;
    right: -2ch;
    padding: 0 0.5em;
    border-radius: 50vmax;
    box-shadow: 0px 10px 15px -3px rgba(0, 0, 0, 0.1);
    background-color: var(--bg-color);
    -webkit-backdrop-filter: blur(0.5em);
    backdrop-filter: blur(0.5em);
    transition: all 0.4s ease;
    --alpha: 0.5;
  }

  .format:not(:hover)::after {
    content: "";
    width: 1.5em;
    aspect-ratio: 1;
    top: -0.5em;
    right: -0.5em;
    --alpha: 0.75;
  }

  .format[data-type="video"]::after {
    --bg-color: rgba(255, 122, 0, var(--alpha));
  }

  .format[data-type="audio"]::after {
    --bg-color: rgba(0, 122, 255, var(--alpha));
  }

  a {
    color: inherit;
    text-decoration: none;
  }

  .download-loader {
    width: 0.5em;
    aspect-ratio: 1;
    border-radius: 50%;
    display: block;
    margin: 15px auto;
    position: relative;
    color: #fff;
    left: -100px;
    box-sizing: border-box;
    animation: shadowRolling 2s linear infinite;
  }

  @keyframes shadowRolling {
    0% {
      box-shadow: 0px 0 rgba(255, 255, 255, 0), 0px 0 rgba(255, 255, 255, 0),
        0px 0 rgba(255, 255, 255, 0), 0px 0 rgba(255, 255, 255, 0);
    }
    12% {
      box-shadow: 100px 0 white, 0px 0 rgba(255, 255, 255, 0),
        0px 0 rgba(255, 255, 255, 0), 0px 0 rgba(255, 255, 255, 0);
    }
    25% {
      box-shadow: 110px 0 white, 100px 0 white, 0px 0 rgba(255, 255, 255, 0),
        0px 0 rgba(255, 255, 255, 0);
    }
    36% {
      box-shadow: 120px 0 white, 110px 0 white, 100px 0 white,
        0px 0 rgba(255, 255, 255, 0);
    }
    50% {
      box-shadow: 130px 0 white, 120px 0 white, 110px 0 white, 100px 0 white;
    }
    62% {
      box-shadow: 200px 0 rgba(255, 255, 255, 0), 130px 0 white, 120px 0 white,
        110px 0 white;
    }
    75% {
      box-shadow: 200px 0 rgba(255, 255, 255, 0), 200px 0 rgba(255, 255, 255, 0),
        130px 0 white, 120px 0 white;
    }
    87% {
      box-shadow: 200px 0 rgba(255, 255, 255, 0), 200px 0 rgba(255, 255, 255, 0),
        200px 0 rgba(255, 255, 255, 0), 130px 0 white;
    }
    100% {
      box-shadow: 200px 0 rgba(255, 255, 255, 0), 200px 0 rgba(255, 255, 255, 0),
        200px 0 rgba(255, 255, 255, 0), 200px 0 rgba(255, 255, 255, 0);
    }
  }
</style>
