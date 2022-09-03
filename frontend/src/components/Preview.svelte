<script lang="ts">
  import ms from "ms";

  export let videoDetails;

  const {
    title,
    shortDescription: description,
    keywords,
    author: channel,
    channelId,
    viewCount: views,
    lengthSeconds: duration,
    thumbnail: { thumbnails },
  } = videoDetails;

  export let urls = [];
  let play = false;
  let qualityIndex = 0;

  function playVideo(_, index = 0) {
    console.log(index, "index");
    play = true;
    qualityIndex = index;
  }

  console.log(urls);
</script>

<figure>
  <div class="video-preview">
    <img src={thumbnails.at(-1).url} alt={title + " video thumbnail"} />
    {#if play}
      <video src={urls[qualityIndex].url} controls autoplay>
        <track
          kind="captions"
          src="https://www.youtube.com/api/timedtext?lang=en&v={videoDetails.videoId}"
        />
      </video>
    {/if}
    <div class="overlay">
      <div class="quality">
        quality:
        {#each urls as { qualityLabel }, i}
          <span
            class:active={i == qualityIndex}
            on:click={(_) => {
              playVideo(_, i);
            }}
          >
            {qualityLabel}
          </span>
        {/each}
      </div>
      {#if !play}
        <i class="ai-play" style:font-size="3em" on:click={playVideo} />
      {/if}
    </div>
  </div>
  <figcaption>
    <div class="top-bar">
      <div class="left-side">
        <h3>{title}</h3>
        <a href="https://www.youtube.com/channel/{channelId}">{channel}</a>
      </div>
      <div class="right-side">
        <span>
          <i class="ai-clock" />
          {ms(ms(duration +"s"))}
        </span>
        <span>
          <i class="ai-eye-open" />
          {views}
        </span>
      </div>
    </div>
    <hr />
    <p>{description}</p>
    <br />
    <h3>Keywords</h3>
    <hr />
    <div class="keywords">
      {#each keywords as keyword}
        <span class="keyword">{keyword}</span>
      {/each}
    </div>
  </figcaption>
</figure>

<style>
  figure {
    position: relative;
    display: grid;
    grid-template-columns: 1fr;
    padding: 2em;
    background-color: var(--bg-color);
    /* -webkit-backdrop-filter: blur(0.5em);
    backdrop-filter: blur(0.5em); */
    border-radius: 0.75em;
    max-width: 90vw;
    margin: auto;

    box-shadow: 0px 10px 15px -3px rgba(0, 0, 0, 0.1),
      0px 10px 15px -3px rgba(0, 0, 0, 0.1),
      0px 10px 15px -3px rgba(0, 0, 0, 0.1);
  }

  .video-preview {
    box-shadow: 0px 10px 15px -3px rgba(0, 0, 0, 0.1);
    position: relative;
    transform:
    scale(0.9)
    perspective(75em)
    rotateX(18deg);
  transform-style: preserve-3d;
  transition: transform 0.4s ease;
  }
  
  .video-preview:where(:hover, :has(video)) {
    transform: scale(1);
  }

  .video-preview:has(video) > img {
    opacity: 0.75;
  }

  img,
  video {
    border-radius: 0.5em;
    width: 100%;
    object-fit: cover;
  }

  video,
  .overlay {
    position: absolute;
    inset: 0;
    z-index: 10;
  }

  .overlay {
    display: grid;
    place-items: center;
    opacity: 0.75;
    pointer-events: none;
  }

  .overlay > * {
    pointer-events: all;
  }

  .quality {
    position: absolute;
    top: -1em;
    right: -1em;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.5em;
    font-size: 0.75em;
    padding: 0.5em;
    background-color: rgba(102, 51, 153, 0.75);
    -webkit-backdrop-filter: blur(0.5em);
    backdrop-filter: blur(0.5em);
    border-radius: 0.5em;
    z-index: 10;
  }

  .quality span {
    color: var(--text-color);
    cursor: pointer;
  }

  .quality span.active {
    color: red;
  }

  .top-bar {
    display: flex;
    align-items: center;
    gap: 1em;
  }

  .top-bar > div {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
  }

  .top-bar span {
    display: flex;
    gap: 0.25em;
    align-items: center;
  }

  a {
    color: var(--text-color);
  }

  figcaption {
    padding: 1em;
  }

  .top-bar {
    display: flex;
    justify-content: space-between;
  }

  p {
    white-space: pre-wrap;
  }

  .keywords {
    display: flex;
    flex-wrap: wrap;
    gap: 0.5em;
  }

  .keyword {
    padding: 0.5em 1em;
    border-radius: 0.5em;
    background-color: var(--bg-color);
    color: var(--text-color);
    font-size: 0.75em;
    margin-right: 0.5em;
    margin-bottom: 0.5em;

    box-shadow: 0px 10px 15px -3px rgba(0, 0, 0, 0.1);
  }
</style>
