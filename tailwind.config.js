module.exports = {
	content: ["internal/views/**/*.templ"],
	theme: {
		extend: {},
	},
	plugins: [require("@tailwindcss/typography"), require("daisyui")],
	daisyui: {
		themes: ["light"]
	}
}
