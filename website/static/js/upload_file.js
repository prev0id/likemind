const inputElement = document.getElementById("upload_new_profile_image_input");
const buttonElement = document.getElementById(
  "upload_new_profile_image_button",
);
const submitElement = document.getElementById(
  "upload_new_profile_image_submit",
);
const dropboxElement = document.getElementById(
  "upload_new_profile_image_dropbox",
);
const previewElement = document.getElementById(
  "upload_new_profile_image_preview",
);
const svgElement = document.getElementById("upload_new_profile_image_svg");

inputElement.addEventListener(
  "change",
  () => {
    handleFiles(inputElement.files);
  },
  false,
);

buttonElement.addEventListener(
  "click",
  () => {
    inputElement.click();
  },
  false,
);

dropboxElement.addEventListener(
  "dragenter",
  (e) => {
    e.stopPropagation();
    e.preventDefault();
  },
  false,
);

dropboxElement.addEventListener(
  "dragover",
  (e) => {
    e.stopPropagation();
    e.preventDefault();
  },
  false,
);

dropboxElement.addEventListener(
  "drop",
  (e) => {
    e.stopPropagation();
    e.preventDefault();

    const dt = e.dataTransfer;
    const files = dt.files;

    handleFiles(files);
  },
  false,
);

function handleFiles(files) {
  const selectedElement = document.getElementById(
    "upload_new_profile_image_selected",
  );

  if (!files || files.length == 0) {
    selectedElement.textContent = "None";
    return;
  }

  const image = files[0];

  selectedElement.textContent = `${image.name} (${getFileSize(image)})`;

  previewElement.file = image;
  previewElement.classList.remove("hidden");
  svgElement.classList.add("hidden");
  submitElement.removeAttribute("disabled");

  const reader = new FileReader();
  reader.onload = (e) => {
    previewElement.src = e.target.result;
  };
  reader.readAsDataURL(image);
}

function getFileSize(file) {
  const units = ["B", "KiB", "MiB", "GiB"];

  const exponent = Math.min(
    Math.floor(Math.log(file.size) / Math.log(1024)),
    units.length - 1,
  );

  const approx = file.size / 1024 ** exponent;
  return exponent === 0
    ? `${file.size} bytes`
    : `${approx.toFixed(3)} ${units[exponent]}`;
}
