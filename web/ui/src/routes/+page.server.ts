import { superValidate, message } from 'sveltekit-superforms';
import { zod4 } from 'sveltekit-superforms/adapters';
import { fail, type Actions } from '@sveltejs/kit';
import { formSchema } from '$lib/types/auth';

import type { PageServerLoad } from './$types';
import { env } from '../env/server';

export const load: PageServerLoad = async (event) => {
	const { fetch } = event;

	const sessionUser = event.locals.user;
	const form = await superValidate(event, zod4(formSchema));

	const albumres = await fetch(`${env.STATIC_URL}/rest/fresh.view`);
	const albums = albumres.ok ? await albumres.json() : [];

	const titles = [];
	if (albums.length > 1) {
		const i = Math.floor(Math.random() * albums.length);
		let j = Math.floor(Math.random() * albums.length);

		while (j === i) {
			j = Math.floor(Math.random() * albums.length);
		}
		titles.push(albums[i].title, albums[j].title);
	}

	return {
		form,
		user: sessionUser,
		albums,
		titles
	};
};

export const actions: Actions = {
	default: async (e) => {
		const form = await superValidate(e, zod4(formSchema));
		if (!form.valid) return fail(400, { form });
		const res = await e.fetch(`${env.API_URL}/v1/signup`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(form.data)
		});

		if (!res.ok) {
			form.valid = false;
			form.errors.username = ['Invalid username'];
			return fail(400, { form });
		}

		return message(form, `Signup successful!`);
	}
};
