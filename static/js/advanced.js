let markdownPreview = document.getElementById("markdownPreview");
let md = new Remarkable();
markdownPreview.innerHTML = md.render(markdownPreview.innerHTML);