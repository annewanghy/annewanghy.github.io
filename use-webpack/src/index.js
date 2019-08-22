import _ from "lodash";
import "./style.css";
import MyImage from "./my-image.png";

function component() {
  const element = document.createElement("div");

  element.innerHTML = _.join(["hello", "webpack"], "");

  element.classList.add("hello");

  const myImage = new Image();
  myImage.src = MyImage;
  element.appendChild(myImage);

  return element;
}

document.body.appendChild(component());
