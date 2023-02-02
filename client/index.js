window.onload = function() {
  let jokeContainer = document.getElementById("joke-container");
  let memeContainer = document.getElementById("meme-container");
  let jokeTitleEl = document.getElementById("joke-title");
  let jokeBodyEl = document.getElementById("joke-body");

  let requestJokeBtn = document.getElementById("request-joke");
  let requestMemeBtn = document.getElementById("request-meme");

  requestJokeBtn.addEventListener("click", async function() {
    try {
      jokeContainer.style.display = "block";
      memeContainer.style.display = "none";
      let response = await fetch("http://localhost:5050/");
      if (response.ok) {
        jokeJSON = await response.json();
        jokeTitleEl.innerText = jokeJSON.title;
        jokeBodyEl.innerText = jokeJSON.body;
      }
    } catch (error) {
      //console.error(error)
      jokeTitleEl.innerText = "Unfortunately joke could not be retrieved. Try again later."
      jokeBodyEl.innerText = "Maybe try and click for meme?"
    }
  });

  requestMemeBtn.addEventListener("click", async function() {
    try {
      jokeContainer.style.display = "none";
      memeContainer.style.display = "block";
      let response = await fetch("http://localhost:5050/meme");
      if (response.ok) {
        memeJSON = await response.json();
        meme = JSON.parse(memeJSON.meme);
        if(meme.url) {
          memeContainer.innerHTML = `<img src=${meme.url}>`
        } else {
          memeContainer.innerHTML = "<p>Unfortunately memes not available</p>";
        }
      }
    } catch (error) {
      //console.error(error)
      memeContainer.innerHTML = "<p>Unfortunately memes not available</p>"
    }
  });

}