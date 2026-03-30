customElements.define(
  "interlocutr-comments",
  class extends HTMLElement {
    connectedCallback() {
      const shadow = this.attachShadow({ mode: "open" });

      const api = this.getAttribute("api");
      const site = this.getAttribute("site");
      const page = this.getAttribute("page");
      const url = `${api}/${site}/${page}/comments`;

      shadow.innerHTML += `
      <link href="./output.css" rel="stylesheet">

      <section class="mb-12">
        <h3 class="text-2xl font-bold mb-6">Leave a comment</h3>
        <form action="#" method="POST" class="space-y-4" id="comments-form">
            <input id="comments-author" name="author" required type="text" placeholder="Your Name" class="w-1/2 px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:outline-none">
            <textarea id="comments-text" name="text" required rows="4" placeholder="Share your thoughts..." class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:outline-none"></textarea>
            <button id="comments-submit-button" type="submit" class="bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-6 rounded-lg transition duration-200">
                Post Comment
            </button>
        </form>
      </section>

      <section>
        <h3 class="text-2xl font-bold mb-6">Comments (<span id="comments-count"></span>)</h3>
        
        <div class="space-y-8" id="comments">
            Loading...
        </div>
      </section>
    `;

      const form = shadow.querySelector("#comments-form");
      const commentsText = shadow.querySelector("#comments-text");
      const submitButton = shadow.querySelector("#comments-submit-button");

      form.addEventListener("submit", async (e) => {
        e.preventDefault();

        submitButton.disabled = true;

        const response = await fetch(url, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(Object.fromEntries(new FormData(form))),
        });

        const result = await response.json();
        console.info('Comment added', result);

        submitButton.disabled = false;
        if (response.ok) {
          form.reset();
          renderComments();
        }
      });

      commentsText.addEventListener("keydown", (e) => {
        if (e.key === "Enter" && (e.ctrlKey || e.metaKey)) {
          form.requestSubmit();
        }
      });

      function e(str) {
        return str
          .replaceAll("&", "&amp;")
          .replaceAll("<", "&lt;")
          .replaceAll(">", "&gt;");
      }

      function renderComments() {
        fetch(url)
          .then((r) => r.json())
          .then((data) => {
            shadow.querySelector("#comments").innerHTML = data
              .map((c) => commentTemplate(c))
              .join("");
            shadow.querySelector("#comments-count").textContent = data.length;
          });
      }

      function commentTemplate(c) {
        const initials = e(c.author)
          .toUpperCase()
          .split(" ")
          .map((word) => word[0])
          .slice(0, 2)
          .join("");
        const date = new Date(c.created_at);

        return `
        <div class="flex gap-4">
          <div class="w-10 h-10 bg-blue-100 rounded-full shrink-0 flex items-center justify-center font-bold text-blue-600">
            ${initials}
          </div>
          <div>
            <div class="flex items-center gap-2 mb-1">
              <span class="font-semibold text-sm">${e(c.author)}</span>
              <span class="text-xs text-gray-400">${date.toLocaleDateString()} ${date.toLocaleTimeString()}</span>
            </div>
            <p class="text-gray-600 italic">${e(c.text)}</p>
          </div>
        </div>
      `;
      }

      renderComments();
    }
  },
);
