import { type Config } from 'tailwindcss'

export default {
	content: [
		'routes/**/*.svelte',
		'lib/**/*.svelte',
		'app.html',
	],
	theme: {
		colors: {
			white: '#fafafa',
			black: '#1a1a1a',
			gray: '#cccccc',
			grayer: '#dddddd',
		},
		fontFamily: {
			zenkaku: ['Zen Kaku Gothic New', 'sans-serif'],
		},
	},
	darkMode: 'class',
} satisfies Config
