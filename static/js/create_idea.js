let mainForm = document.getElementById("MainForm");
let editorForm = document.getElementById("EditorForm");

//MainMode();

let imageUpload = document.getElementById("logoUpload");
let imgPreview = document.getElementById("previewImg");

imageUpload.addEventListener("change", function() {
    const files = imageUpload.files[0];
    if (files) {
      const fileReader = new FileReader();
      fileReader.readAsDataURL(files);
      fileReader.addEventListener("load", function () {
        imgPreview.style.display = "block";
        imgPreview.setAttribute('src', this.result);
      });    
    }
});

let headerTextInput = document.getElementById("headerInput");
let headerText = document.getElementById("header");

headerTextInput.addEventListener("keyup", function() {
  headerText.innerHTML = headerTextInput.value;
});

let descTextInput = document.getElementById("descriptionInput");
let descText = document.getElementById("description");

descTextInput.addEventListener("keyup", function() {
  descText.innerHTML = descTextInput.value;
})

let authorTextInput = document.getElementById("authorInput");
let authorText = document.getElementById("author");

authorTextInput.addEventListener("keyup", function() {
  authorText.innerHTML = authorTextInput.value;
})

//let editorButton = document.getElementById("EditorButton");

function EditorMode() {
  mainForm.style.visibility = "hidden";
  editorForm.style.visibility = "visible";
  editorForm.style.top= "-50rem";
}
function MainMode() {
  mainForm.style.visibility = "visible";
  editorForm.style.visibility = "hidden";
  editorForm.style.top="0";
}

/*
editorButton.addEventListener("click", function() {
  EditorMode();
});
*/

let markdownInput = document.getElementById("markdownInput");
let markdownPreview = document.getElementById("markdownPreview");

markdownInput.addEventListener("keyup", function() {
  let md = new Remarkable();
  markdownPreview.innerHTML = md.render(markdownInput.value);
  console.log(md.render('# tada!'));
  console.log(markdownInput);
});