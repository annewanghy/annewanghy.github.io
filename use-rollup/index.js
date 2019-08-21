import { add } from "./math";
import "./style.less";

console.log(add(1, 2));
console.log("hello world");

import Swiper from "swiper";

var mySwiper = new Swiper(".swiper-container", {
  // Optional parameters
  direction: "horizontal",
  loop: true,

  // If we need pagination
  pagination: {
    el: ".swiper-pagination"
  },

  // Navigation arrows
  navigation: {
    nextEl: ".swiper-button-next",
    prevEl: ".swiper-button-prev"
  },

  // And if we need scrollbar
  scrollbar: {
    el: ".swiper-scrollbar"
  }
});

console.log("mySwiper", mySwiper);
