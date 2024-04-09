import {defineStore} from 'pinia';

export const useAuthStore = defineStore('auth', {
    state: () => ({
        isAuthenticated: false,
        username: '',
        baid: 0,
        initialized: false,
        isCoach: false,
    }),
    actions: {
        async initAuth(): Promise<void> {
            if (this.initialized) return; // Prevent re-initialization
            this.initialized = true; // Mark as initialized

            const response = await fetch('/api/auth/session', {method: 'POST'});
            switch (response.status) {
                case 200:
                    const res = await response.json();
                    this.username = res.Username;
                    this.baid = res.Baid;
                    this.isCoach = res.privileges.coach;
                    this.isAuthenticated = true;
                    break;
                case 401:
                    this.isAuthenticated = false;
                    break;
                default:
                    console.error('Unexpected response status:', response.status);
            }
        },
        async login(formData: {email: string, password: string}, isCoach: boolean): Promise<[boolean, string]> {
            const response = await fetch(`/api/login/${isCoach ? 'coach' : 'user'}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formData)
            });
            const body = await response.json();
            if (response.ok) {
                this.username = body.Username;
                this.baid = body.Baid;
                this.isAuthenticated = true;
                this.isCoach = isCoach;
                return [true, ''];
            } else {
                return [false, await body.error];
            }
        },
        async logout(): Promise<void> {
            const response = await fetch('/api/logout', {method: 'POST'});
            if (response.ok) {
                this.isAuthenticated = false;
                this.username = '';
                this.baid = 0;
            } else {
                console.error('Failed to logout:', response.status);
            }
        },
    },
});