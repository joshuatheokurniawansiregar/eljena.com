document.addEventListener("DOMContentLoaded", async function () {
  getCarousel();
});

async function getCarousel() {
  const url = "http://localhost:3000/api/v1/carousel";
  const carouselImage = document.getElementById("carousel-image");

  const carouselItemNameTag = document.getElementById("carousel-item-name");
  const itemNameArray = [];

  const carouselSubCategory = document.getElementById("carousel-sub-category");
  const subCategoryArray = [];

  const carouselDescription = document.getElementById("carousel-description");
  const descriptionArray = [];

  const carouselTags = document.getElementById("carousel-tags");
  let tagsArray = [];

  try {
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }
    const json = await response.json();
    json.forEach((value) => {
      const { imageUrl, itemName, description, subCategory, tags } = value;

      const imagesElement = document.createElement("img");
      imagesElement.src = imageUrl;
      imagesElement.classList = "w-full min-h-full";
      carouselImage.appendChild(imagesElement);

      itemNameArray.push(itemName);

      subCategoryArray.push(subCategory);

      descriptionArray.push(description);

      tagsArray.push(tags);
    });
  } catch (error) {}

  carouselItemNameTag.textContent = itemNameArray[0];
  carouselSubCategory.textContent = subCategoryArray[0];
  carouselDescription.textContent = descriptionArray[0];

  tagsArray[0].forEach((each) => {
    const liElementsForTags = document.createElement("li");
    const aTest = document.createElement("a");
    aTest.href = "#";
    aTest.classList = "hover:text-red-600";
    aTest.innerHTML = `#${each}`;
    liElementsForTags.appendChild(aTest);
    liElementsForTags.classList = "mx-2 hover:text-red-600";
    carouselTags.appendChild(liElementsForTags);
  });

  let index = 0;
  let images = document.querySelectorAll("#carousel-image img");
  let carouselIndicatorContainer = document.querySelector(
    "#carousel-indicator-container"
  );

  images.forEach((_, i) => {
    const carouselIndicatorContainerChilds = document.createElement("div");
    carouselIndicatorContainerChilds.classList = `transition-all w-2 h-2 rounded-full carousel-indicator-container-childs ${
      i == index ? " bg-yellow-500" : "bg-opacity-50 bg-white"
    }`;
    carouselIndicatorContainer.appendChild(carouselIndicatorContainerChilds);
  });

  const prevSlideButton = document.getElementById("prevSlideButton");
  const nextSlideButton = document.getElementById("nextSlideButton");

  prevSlideButton.addEventListener("click", function () {
    index = index == 0 ? images.length - 1 : index - 1;
    updateCarousel();
  });
  nextSlideButton.addEventListener("click", function () {
    index = index == images.length - 1 ? 0 : index + 1;
    updateCarousel();
  });

  function updateCarousel() {
    document.getElementById("carousel-image").style.transform = `translateX(-${
      index * 100
    }%)`;
    carouselItemNameTag.textContent = itemNameArray[index];
    carouselSubCategory.textContent = subCategoryArray[index];
    carouselDescription.textContent = descriptionArray[index];
    carouselTags.innerHTML = "";
    tagsArray[index].forEach((each) => {
      const liElementsForTags = document.createElement("li");
      const aTest = document.createElement("a");
      aTest.href = "#";
      aTest.classList = "hover:text-red-600";
      aTest.innerHTML = `#${each}`;
      liElementsForTags.appendChild(aTest);
      liElementsForTags.classList = "mx-2 hover:text-red-600";
      carouselTags.appendChild(liElementsForTags);
    });

    document
      .querySelectorAll(".carousel-indicator-container-childs")
      .forEach((value, i) => {
        value.classList = `transition-all w-2 h-2 rounded-full carousel-indicator-container-childs ${
          i == index ? " bg-yellow-500" : "bg-opacity-50 bg-white"
        }`;
      });
  }
}
