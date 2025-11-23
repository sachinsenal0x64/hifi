declare global {
	namespace App {
		interface Locals {
			user: {
				username: string;
				password: string;
			} | null;
		}
	}
}

export {};
