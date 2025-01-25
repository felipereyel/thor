all: statics templ

statics:
	curl https://cdn.tailwindcss.com/3.4.16 --output internal/embeded/statics/tailwind.js
	curl https://unpkg.com/htmx.org@1.9.12/dist/htmx.min.js --output internal/embeded/statics/htmx.js
	curl https://unpkg.com/htmx.org@1.9.12/dist/ext/ws.js --output internal/embeded/statics/htmx-ws.js

templ:
	templ generate
