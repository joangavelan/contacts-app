/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./web/templates/**/*.html"],
  theme: {
    extend: {}
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      {
        dark: {
          primary: "#3F8EFC",
          "base-100": "#0e131f",
          success: "#ADD7F6",
          warning: "#B3001B"
        }
      }
    ]
  }
}
