import adapter from '@sveltejs/adapter-auto';
import preprocess from 'svelte-preprocess';
import { resolve } from "path"

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	preprocess: preprocess(),

	kit: {
		adapter: adapter(),

		// hydrate the <div id="svelte"> element in src/app.html
		target: '#svelte',
		vite: {
			resolve: {
				alias: {
					$utils: resolve("./src/utils"),
					$modules: resolve("./src/modules"),
					$lib: resolve("./src/lib"),
					$stores: resolve("./src/stores"),
				}
			}
		}
	}
};

export default config;
