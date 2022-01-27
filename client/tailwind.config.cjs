module.exports = {
  content: ['./public/index.html', './src/**/*.svelte'],
  theme: {
    extend: {
      colors: {
        brand: {
          500: "#43ffaf"
        }
      },
      animation: {
        caret: "caret 1s cubic-bezier(.87,0,.13,1) infinite",
      },
      keyframes: {
        caret: {
          "0%, 100%": {
            opacity: 1,
          },
          "50%": {
            opacity: 0,
          },
        },
      },
    },
  },
  plugins: [],
}
