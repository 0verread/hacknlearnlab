/** @type {import('tailwindcss').Config} */

const sharedConfig = require("@monoture/config");

module.exports = {
  ...sharedConfig,
  content: ["./**/*.{js,ts,jsx,tsx,mdx}"],
};