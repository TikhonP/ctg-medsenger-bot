templ:
	@templ generate -watch -proxy=http://localhost:8080

tailwind:
	@tailwindcss -i view/css/input.css -o public/styles.css --watch

