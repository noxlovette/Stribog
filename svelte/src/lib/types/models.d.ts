export interface Forge {
	id: string;
	title: string;
	description?: string;
	ownderId: string;
	createdAt: string;
	updatedAt: string;
}

export interface ForgeAccess {
	id: string;
	forgeId: string;
	userId: string;
	accessRole: string;
	addedBy: string;
	createdAt: string;
	updatedAt: string;
}

export interface ApiKey {
	id: string;
	forgeId: string;
	title: string;
	isActive: boolean;
	createdAt: string;
	lastUsedAt: string;
}

export interface Spark {
	id: string;
	forgeId: string;
	title: string;
	markdown: string;
	ownerId: string;
	createdAt: string;
	updatedAt: string;
}

export interface Toast {
	message: string | null;
	type: 'success' | 'error' | 'info' | null;
}

export interface User {
	name: string | null;
	username: string | null;
	role: string | null;
	email: string | null;
	sub: string | null;
	[key: string]: string | undefined;
}
