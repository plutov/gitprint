import type { Config } from "tailwindcss";

const config: Config = {
  darkMode: "selector",
  content: ["./index.html", "./src/**/*.{vue,ts}"],
  theme: {
    extend: {},
  },
  plugins: [],
};
export default config;
