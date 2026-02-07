let audioPlaying = false;

function nextFrame() {
  if (!audioPlaying) {
    document.getElementById("music").play();
    audioPlaying = false;
  }
}
