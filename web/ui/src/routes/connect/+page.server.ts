import type { PageServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';

import { superValidate, message } from 'sveltekit-superforms';
import { zod4 } from 'sveltekit-superforms/adapters';
import { updateSchema } from '$lib/types/auth';
import { fail, type Actions } from '@sveltejs/kit';
import { env } from '../../env/server';

export const load: PageServerLoad = async (event) => {
	const { locals } = event;

	if (!locals.user) {
		const redirectTo = '/signin';
		throw redirect(303, redirectTo);
	}

	const form = await superValidate(event, zod4(updateSchema));

	return {
		form,
		user: locals.user
	};
};

export const actions: Actions = {
	default: async (e) => {
		const { locals } = e;

		const form = await superValidate(e, zod4(updateSchema));
		if (!form.valid) return fail(400, { form });

		if (form.data.username === locals.user?.username) {
			form.valid = false;
			form.errors.username = ['You already have this username'];
			return fail(400, { form });
		}

		const isChangingPassword = !!form.data.password;

		if (isChangingPassword && locals.user?.password !== form.data.oldpassword) {
			form.valid = false;
			form.errors.oldpassword = ['Old password is incorrect'];
			return fail(400, { form });
		}

		const pretoken = e.cookies.get('hifi');

		const res = await e.fetch(`${env.API_URL}/v1/update`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${pretoken}`,
				'X-Username': `${locals.user?.username}`
			},
			body: JSON.stringify(form.data)
		});

		if (!res.ok) {
			form.valid = false;
			form.errors.username = ['Please try again later'];
			return fail(400, { form });
		}

		return message(form, 'Username changed successfully');
	}
};
