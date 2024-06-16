const defaultTheme = require("tailwindcss/defaultTheme");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["components/**/*.templ", "views/**/*.templ"],
  theme: {
    extend: {
      fontFamily: {
        sans: ['"Inter Variable"', ...defaultTheme.fontFamily.sans],
      },
      boxShadow: {
        inner:
          "inset 0 -0.25rem 0.25rem var(--tw-shadow-color, rgb(0 0 0 / .1))",
      },
    },
    borderRadius: {
      none: "0",
      sm: "0.25rem",
      DEFAULT: "0.25rem",
    },
  },
  plugins: [],
};
// 'Inter Variable'
