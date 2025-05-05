export interface Forge {
	id: string;
	title: string;
	description?: string;
	ownderId: string;
	createdAt: string;
	updatedAt: string;
}

export interface Collaborator {
	id: string;
	forgeId: string;
	userId: string;
	userName: string;
	userEmail: string;
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
	email: string | null;
}
