import type { PageServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ cookies }) => {
	cookies.delete('hifi', { path: '/' });

	throw redirect(303, '/signin');
};
