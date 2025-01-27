/** @type {import('tailwindcss').Config} */
export default {
  content: [
      "./src/**/*.{html,css,vue}",
      "./index.html",
  ],
  theme: {
    extend: {
        colors: {
            primary: "#D0C3F1",
            stress: "rgb(85,26,139)",
            "neutral-750": "#2a2a2a",
            "highlight1": "#1ac8db",
            "highlight2": "#99dfec",
        }
    },
  },
  plugins: [],
}

