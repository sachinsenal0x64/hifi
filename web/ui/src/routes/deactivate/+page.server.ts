import type { PageServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';

import { superValidate, message } from 'sveltekit-superforms';
import { zod4 } from 'sveltekit-superforms/adapters';
import { formSchema2 } from '$lib/types/auth';

import { fail, type Actions } from '@sveltejs/kit';
import { env } from '../../env/server';

export const load: PageServerLoad = async (event) => {
	const { locals } = event;

	if (!locals.user) {
		const redirectTo = '/signin';
		throw redirect(303, redirectTo);
	}

	const form = await superValidate(event, zod4(formSchema2));

	return {
		form,
		user: locals.user
	};
};

export const actions: Actions = {
	default: async (e) => {
		const { cookies } = e;

		const form = await superValidate(e, zod4(formSchema2));
		if (!form.valid) return fail(400, { form });

		const token = e.cookies.get('hifi');

		const res = await e.fetch(`${env.API_URL}/v1/delete`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${token}` },
			body: JSON.stringify(form.data)
		});

		if (!res.ok) {
			form.valid = false;
			form.errors.username = ['Please try again later'];
			return fail(400, { form });
		}

		const { message } = await res.json();

		cookies.delete('hifi', { path: '/' });

		const redirectTo = '/signin';

		return message(form, redirect(303, redirectTo));
	}
};
