/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
	"./pkg/html/**/*.go",
  ],
  safelist: [
    'bg-blue-500',
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}

