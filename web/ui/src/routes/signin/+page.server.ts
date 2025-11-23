import { superValidate, message } from 'sveltekit-superforms';
import { zod4 } from 'sveltekit-superforms/adapters';
import { fail, type Actions } from '@sveltejs/kit';
import { formSchema } from '$lib/types/auth';


import type { PageServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';
import { env } from '../../env/server';

export const load: PageServerLoad = async (event) => {
	// Redirect authenticated users away from the signin page
	const { locals, url } = event;

	if (locals.user) {
		const redirectTo = url.searchParams.get('redirect') || '/connect';
		throw redirect(303, redirectTo);
	}

	const form = await superValidate(event, zod4(formSchema));
	return { form };
};

export const actions: Actions = {
	default: async (e) => {
		const { cookies, url } = e;

		const form = await superValidate(e, zod4(formSchema));
		if (!form.valid) return fail(400, { form });
		const res = await e.fetch(`${env.API_URL}/v1/signin`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(form.data)
		});

		if (!res.ok) {
			form.valid = false;
			form.errors.username = ['Invalid username or password'];
			form.errors.password = ['Invalid username or password'];
			return fail(400, { form });
		}

		const { token, maxAge } = await res.json();

		cookies.set('hifi', token, {
			path: '/',
			httpOnly: true,
			secure: true,
			sameSite: 'strict',
			maxAge
		});

		const redirectTo = url.searchParams.get('redirect') || '/connect';

		return message(form, redirect(303, redirectTo));
	}
};
