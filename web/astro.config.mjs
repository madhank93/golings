// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';
import sitemap from '@astrojs/sitemap';

// Served from the custom domain at the root.
// (For GitHub project pages instead, set SITE to the github.io origin and BASE to '/golings'.)
const SITE = 'https://golings.madhan.app';
const BASE = '/';
const DESCRIPTION =
	'Learn Go the rustlings way — 106 hands-on exercises you fix one at a time, from variables to concurrency, current through Go 1.26.';

export default defineConfig({
	site: SITE,
	base: BASE,
	integrations: [
		sitemap(),
		starlight({
			title: 'golings',
			description: DESCRIPTION,
			customCss: ['./src/styles/hero.css'],
			social: [
				{ icon: 'github', label: 'GitHub', href: 'https://github.com/madhank93/golings' },
			],
			// SEO / social-share metadata applied to every page.
			head: [
				{ tag: 'meta', attrs: { property: 'og:type', content: 'website' } },
				{ tag: 'meta', attrs: { property: 'og:site_name', content: 'golings' } },
				{ tag: 'meta', attrs: { property: 'og:image', content: new URL(`${BASE}favicon.svg`, SITE).href } },
				{ tag: 'meta', attrs: { name: 'twitter:card', content: 'summary' } },
				{ tag: 'meta', attrs: { name: 'twitter:image', content: new URL(`${BASE}favicon.svg`, SITE).href } },
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
