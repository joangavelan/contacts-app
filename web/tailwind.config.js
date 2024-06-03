/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./templates/**/*.html'],
  theme: {
    extend: {}
  },
  plugins: [require('daisyui')],
  daisyui: {
    themes: [
      {
        dark: {
          primary: '#3F8EFC',
          'base-100': '#0e131f',
          success: '#ADD7F6',
          error: '#B3001B'
        }
      }
    ]
  }
}
