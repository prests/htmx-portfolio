/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./internal/server/http/web/components/**/*.templ",
    "./internal/server/http/web/templates/**/*.templ",
  ],
  theme: {
    extend: {
      colors: {
        "orange-red-crystal": "#fe5f55",
        "light-coral": "#263747",
        "light-grey": "#aaaeb4",
        "custom-white": "#f7f7ff",
        "black-coral": "#2f3f4e",
      },
      fontFamily: {
        lato: ["lato", "sans-serif"],
      },
    },
  },
  plugins: [],
};
