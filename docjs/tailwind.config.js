module.exports = {
    purge: ['./public/*.html', './src/**/*.{vue,js,ts,jsx,tsx}', '../**/*.go'],
    darkMode: false, // or 'media' or 'class'
    theme: {
        extend: {},
        fontFamily: {
            sans: ['Helvetica Neue', 'Helvetica', 'Arial', 'sans-serif'],
            mono: ['Menlo', 'Monaco', 'Consolas', 'Courier New', 'monospace'],
        }
    },
    variants: {
        extend: {},
    },
    plugins: [],
}