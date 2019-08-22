import { cube } from "./math.js";
import printMe from "./print.js";
import "./style.css";

function component() {
  const element = document.createElement("pre");
  const btn = document.createElement("button");

  element.innerHTML = [
    "hello",
    "webpack",
    "5 cubed is equal to " + cube(5)
  ].join("\n\n");

  btn.innerHTML = "Click me and check the console!";
  btn.onclick = printMe;
  element.appendChild(btn);

  return element;
}

let element = component();
document.body.appendChild(element);

if (module.hot) {
  module.hot.accept("./print", function() {
    console.log("Accepting the updated printMe module!");
    document.body.removeChild(element);
    element = component();
    document.body.appendChild(element);
  });
}
