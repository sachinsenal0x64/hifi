import type { Handle } from '@sveltejs/kit';
import { dev } from '$app/environment';
import { env } from './env/server';

export const handle: Handle = async ({ event, resolve }) => {
	// Special dev route
	if (dev && event.url.pathname === '/.well-known/appspecific/com.chrome.devtools.json') {
		return new Response(null, { status: 404 });
	}

	// --- Auth check ---
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
			event.cookies.delete('hifi', { path: '/' });
			event.locals.user = null;
		}
	} else {
		event.locals.user = null;
	}

	const pathname = event.url.pathname;

	const protectedRoutes = ['/connect', '/deactivate'];

	const isProtected = protectedRoutes.some(
		(route) => pathname === route || pathname.startsWith(route + '/')
	);

	// Redirect only if user is not logged in AND route is protected
	if (!event.locals.user && isProtected) {
		const redirectUrl = `/signin?redirect=${encodeURIComponent(pathname + event.url.search)}`;
		return Response.redirect(new URL(redirectUrl, event.url), 303);
	}

	const response = await resolve(event);

	response.headers.set(
		'Content-Security-Policy',
		`form-action 'self'; frame-ancestors 'self'; base-uri 'self'; upgrade-insecure-requests; object-src 'none';`
	);

	return response;
};
