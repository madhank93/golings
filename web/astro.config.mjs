// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

// Project pages: https://madhank93.github.io/golings
// (When the custom domain golings.madhan.app is set up, change base to '/'.)
export default defineConfig({
	site: 'https://madhank93.github.io',
	base: '/golings',
	integrations: [
		starlight({
			title: 'golings',
			description: 'Learn Go the rustlings way — 97 exercises, basics to advanced.',
			social: [
				{ icon: 'github', label: 'GitHub', href: 'https://github.com/madhank93/golings' },
			],
			sidebar: [
				{ label: 'Getting started', slug: 'getting-started' },
				{ label: 'Curriculum', slug: 'curriculum' },
				{ label: 'Topics', items: [{ autogenerate: { directory: 'curriculum/topics' } }] },
			],
		}),
	],
});
