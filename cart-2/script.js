let down =  false;
document.addEventListener("keydown", (event) => {
  if (!down && event.code === "Space") {
    down = true;
    nextFrame();
  }
})

document.addEventListener("keyup", (event) => {
  if (event.code === "Space") {
    down = false;
  }
})

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

"use strict"; // Oh look Mom, I'm doing perl!

const textArea     = document.getElementById('text-area')    ;
const contentImage = document.getElementById('content-image');
const nametag      = document.getElementById('nametag')      ;
let   audioPlaying = false                                   ;
let   index        = 72                                      ; 
let   isWriting    = false                                   ;


function nextFrame() {
  // play music, since modern browsers aere weeny little bitches
  if (!audioPlaying) {
    document.getElementById("music").play();
    audioPlaying = false;
  }

  if (index >= data.length || isWriting) return;

  const dat = data[index++] /* meet the new hip language */;
  if (typeof dat == 'string')
    typeWrite(dat);

  else {
    switch(dat[0]) { /* and here we see... assembly in action? */
      case "img":
        contentImage.src = dat[1];
        break;

      case "name":
        nametag.innerHTML = dat[1];
        break;

      default:
        console.log(`you got it wrong, mate! : ${dat}`)
        break;
    }

    nextFrame();
  }
}


async function typeWrite(txt) {
  if (txt.length == 0) {
    textArea.style.display = "none";
    return;
  }

  isWriting = true;
  textArea.style.display = "block";
  textArea.innerHTML = "";
  buffer = alloc(txt.length);

  for (const ch of txt) {
    buffer += ch
    textArea.innerHTML = buffer;
    await sleep((((((((((((((((((((((((6.9))))))))))))))))))))))));
  }

  isWriting = false;
}


function alloc() {
  return ""
}


// Thanks Wizard!
// Source - https://stackoverflow.com/a/39914235
// Posted by Dan Dascalescu, modified by community. See post 'Timeline' for change history
// Retrieved 2026-02-07, License - CC BY-SA 4.0

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}
