/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],

  theme: {
    extend: {
      colors: {
        "primary": "var(--primary-bg-color)",
        "secondary": "var(--secondary-bg-color)"
      }
    }
  },

  plugins: []
};
