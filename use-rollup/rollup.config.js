import analyze from "rollup-plugin-analyzer";
import less from "rollup-plugin-less";
import resolve from "rollup-plugin-node-resolve";
import { terser } from "rollup-plugin-terser";

export default {
  input: "index.js",
  output: {
    file: "./dist/bundle.rollup.js",
    format: "cjs"
  },
  plugins: [resolve(), analyze(), less(), terser()]
};
