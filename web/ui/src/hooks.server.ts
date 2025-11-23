import type { Handle } from '@sveltejs/kit';

import { dev } from '$app/environment';
import { env } from './env/server';

export const handle: Handle = async ({ event, resolve }) => {
	if (dev && event.url.pathname === '/.well-known/appspecific/com.chrome.devtools.json') {
		return new Response(null, { status: 404 });
	}

	const token = event.cookies.get('hifi');

	if (token) {
		try {
			const res = await fetch(`${env.API_URL}/v1/validate`, {
				headers: { Authorization: `Bearer ${token}` }
			});

			if (res.ok) {
				event.locals.user = await res.json();
			} else {
				event.cookies.delete('hifi', { path: '/' });
				event.locals.user = null;
			}
		} catch {
			event.locals.user = null;
			event.cookies.delete('hifi', { path: '/' });
		}
	} else {
		event.locals.user = null;
	}

	const publicRoutes = ['/', '/signin'];

	const redirectableRoutes = ['/connect', '/deactivate'];

	const isPublic = publicRoutes.some(
		(path) => event.url.pathname === path || event.url.pathname.startsWith(`${path}/`)
	);

	if (!event.locals.user && !isPublic) {
		const path = event.url.pathname;

		if (path === '/connect') {
			return Response.redirect(new URL('/signin', event.url), 303);
		}

		if (redirectableRoutes.some((r) => path === r || path.startsWith(`${r}/`))) {
			const redirectUrl = `/signin?redirect=${encodeURIComponent(path + event.url.search)}`;
			return Response.redirect(new URL(redirectUrl, event.url), 303);
		}

		return Response.redirect(new URL('/signin', event.url), 303);
	}

	const response = await resolve(event);

	response.headers.set('Access-Control-Allow-Origin', `${env.API_URL}`);
	response.headers.set(
		'Content-Security-Policy',
		`form-action 'self'; frame-ancestors 'self'; base-uri 'self'; upgrade-insecure-requests; object-src 'none';`
	);

	return response;
};
