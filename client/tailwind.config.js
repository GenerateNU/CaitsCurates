/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      fontFamily: {
        seasons: ["the-seasons", "sans-serif"],
        proxima: ["proxima-nova", "sans-serif"],
      },
      colors: {
        eggshell: "#fff9f4",
        linen: "#fbf2eb",
        beige: "#f4e6dc",
        pearl: "#fffefc",
        brown: "#b79d94",
        mocha: "#8d7a73",
        espresso: "#32211e",
        pink: "#dfb2aa",
        dkpink: "#D08484",
        deeppink: "#a65a5a",
        petalpink: "#F9DCD7",
        red: "#bc2c35",
        winered: "#bc2c36",
        green: "#64765b",
        dkgreen: "#273f2a",
        evergreen: "#305749",
      },
    },
  },
  plugins: [],
};