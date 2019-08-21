import analyze from "rollup-plugin-analyzer";
import less from "rollup-plugin-less";

export default {
  input: "index.js",
  output: {
    file: "./dist/bundle.rollup.js",
    format: "cjs"
  },
  plugins: [analyze(), less()]
};
