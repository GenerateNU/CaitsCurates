/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        'seasons': ['the-seasons', 'sans-serif'],
        'proxima': ['proxima-nova', 'sans-serif']
      },
      colors: {
        'eggshell': '#fff9f4',
        'beige': '#f4e6dc',
        'pearl': '#fffefc',
        'brown': '#b79d94',
        'drkbrown': '#8d7a73',
        'coffee': '#32211e',
        'pink': '#dfb2aa',
        'dkpink': '#D08484',
        'rose': '#A65A5A',
        'petalpink': '#F9DCD7',
        'red': '#bc2c35',
        'green': '#64765b',
        'dkgreen': '#273f2a',
        'evergreen': '#305749'
      }
    },
  },
  plugins: [],
}