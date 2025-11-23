import { env as dyn } from '$env/dynamic/private';

export const env = {
	API_URL: dyn.API_URL,
	STATIC_URL: dyn.STATIC_URL
};
