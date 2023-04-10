/** @type {import('tailwindcss').Config} */
export default {
  // important: true,
  mode: 'jit',
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      spacing: {
        26: '6.5rem'
      },
      colors: {
        info: '#2d8cf0'
      },
      backgroundColor: {
        'blue-info': '#2d8cf0'
      }
    }
  },
  variants: {
    extend: {}
  },
  plugins: [require('@tailwindcss/line-clamp')]
}
