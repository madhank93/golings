// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';
import sitemap from '@astrojs/sitemap';

// Project pages: https://madhank93.github.io/golings
// (When the custom domain golings.madhan.app is set up, change base to '/'.)
const SITE = 'https://madhank93.github.io';
const BASE = '/golings';
const DESCRIPTION =
	'Learn Go the rustlings way — 97 hands-on exercises you fix one at a time, from variables to concurrency, current through Go 1.26.';

export default defineConfig({
	site: SITE,
	base: BASE,
	integrations: [
		sitemap(),
		starlight({
			title: 'golings',
			description: DESCRIPTION,
			social: [
				{ icon: 'github', label: 'GitHub', href: 'https://github.com/madhank93/golings' },
			],
			// SEO / social-share metadata applied to every page.
			head: [
				{ tag: 'meta', attrs: { property: 'og:type', content: 'website' } },
				{ tag: 'meta', attrs: { property: 'og:site_name', content: 'golings' } },
				{ tag: 'meta', attrs: { property: 'og:image', content: `${SITE}${BASE}/favicon.svg` } },
				{ tag: 'meta', attrs: { name: 'twitter:card', content: 'summary' } },
				{ tag: 'meta', attrs: { name: 'twitter:image', content: `${SITE}${BASE}/favicon.svg` } },
				{ tag: 'meta', attrs: { name: 'theme-color', content: '#00add8' } },
				{ tag: 'meta', attrs: { name: 'keywords', content: 'Go, Golang, rustlings, exercises, learn Go, golang tutorial, TUI' } },
			],
			sidebar: [
				{ label: 'Getting started', slug: 'getting-started' },
				{ label: 'Curriculum', slug: 'curriculum' },
				{ label: 'Topics', items: [{ autogenerate: { directory: 'curriculum/topics' } }] },
			],
		}),
	],
});
