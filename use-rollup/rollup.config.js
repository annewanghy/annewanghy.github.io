import analyze from "rollup-plugin-analyzer";

export default {
  input: "index.js",
  output: {
    file: "./dist/bundle.rollup.js",
    format: "cjs"
  },
  plugins: [analyze()]
};
