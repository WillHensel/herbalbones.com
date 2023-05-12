/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './src/**/*.{html,js}'
  ],
  theme: {
    fontFamily: {
      'sans': ['Helvetica Neue']
    },
    extend: {
      backgroundImage: {
        'mushroom': 'url("/images/mushroombg.png")'
      },
      colors: {
        'green-light': '#BFCFAF',
        'green': '#0B293F'
      },
    },
  },
  plugins: [],
}

