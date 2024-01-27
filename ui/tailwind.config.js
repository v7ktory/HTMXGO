/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./template/**/*{html,tmpl}"],
  theme: {
    extend: {
      animation: {
        'pulse-slow': 'pulse 5s cubic-bezier(0.4, 0, 0.6, 1) infinite;' 
      }
    },
  },
  plugins: [],
}

