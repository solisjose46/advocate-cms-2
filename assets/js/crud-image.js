let imageNameTarget = document.getElementById("image-src");
let imageAltTarget = document.getElementById("image-alt");
let imageTarget = document.getElementById("image");
let imageIdTarget = document.getElementById("image-id");

const galleryItems = document.querySelectorAll(".image-gallery-item");

function swapInputForImg(imgSrc, imgAlt) {
    let newImg = document.createElement("img");
    newImg.src = imgSrc;
    newImg.alt = imgAlt;

    let parent = imageTarget.parentElement;

    newImg.setAttribute("id", "image");

    imageTarget.remove();
    imageTarget = newImg;
    parent.appendChild(newImg)

}


function galleryItemClickHandler(e) {
    let item = e.currentTarget;

    let imageName = item.querySelector(".image-src-value").innerText;
    let imageAlt = item.querySelector(".image-alt-value").innerText;
    let imageId = item.id.split("-")[1];

    let imgSrc = item.querySelector("img").src;

    swapInputForImg(imgSrc, imageAlt);

    imageNameTarget.value = imageName;
    imageAltTarget.value = imageAlt;
    imageIdTarget.value = imageId;
}

galleryItems.forEach(function(element) {
    element.addEventListener("click", galleryItemClickHandler);
});
