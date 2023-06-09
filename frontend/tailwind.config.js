/*
 * @Author: iRorikon
 * @Date: 2023-04-06 20:31:05
 * @FilePath: \http-file\frontend\tailwind.config.js
 */
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
  }
  // plugins: [require('@tailwindcss/line-clamp')]
}
