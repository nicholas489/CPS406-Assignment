import {defineStore} from 'pinia';

type severity = "warn" | "success" | "info" | "error" | "secondary" | "contrast" | undefined;
export const useAuthStore = defineStore('auth', {
    state: (): {
        isAuthenticated: boolean,
        email: string,
        initialized: boolean,
        isCoach: boolean,
        id: number,
        toasts: [severity, string, string][],
        amountOwed: number
    } => ({
        isAuthenticated: false,
        email: '',
        initialized: false,
        isCoach: false,
        id: 0,
        toasts: [],
        amountOwed: 0
    }),
    actions: {
        async initAuth(): Promise<void> {
            if (this.initialized) return; // Prevent re-initialization
            const response = await fetch('/api/auth/session', {method: 'POST'});
            switch (response.status) {
                case 200:
                    const res = await response.json();
                    this.email = res.email;
                    this.isCoach = res.privileges.coach;
                    this.isAuthenticated = true;
                    this.id = res.id;
                    await this.refreshOwed()
                    console.log(
                        `${this.email}, ${this.isCoach}, ${this.id}`
                    )
                    break;
                case 401:
                case 404:
                    this.isAuthenticated = false;
                    break;
                default:
                    console.error('Unexpected response status:', response.status);
            }
            this.initialized = true; // Mark as initialized

        },
        async login(formData: { email: string, password: string }, isCoach: boolean): Promise<[boolean, string]> {
            const response = await fetch(`/api/login/${isCoach ? 'coach' : 'user'}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formData)
            });
            const body = await response.json();
            if (response.ok) {
                this.email = body.email;
                this.isAuthenticated = true;
                this.isCoach = isCoach;
                this.id = body.id;
                await this.refreshOwed();
                return [true, ''];
            } else {
                return [false, await body.error];
            }
        },
        async logout(): Promise<void> {
            const response = await fetch('/api/logout', {method: 'POST'});
            if (response.ok) {
                this.isAuthenticated = false;
                this.email = '';
                this.isCoach = false;
                this.id = 0;
            } else {
                console.error('Failed to logout:', response.status);
            }
        },
        async busyWaitTillInitialized(): Promise<void> {
            console.log(this.initialized)
            console.log(this.id)
            while (!this.initialized) {
                console.log(this.initialized)
            }
            return
        },
        async pushToast(severity: severity, title: string, message: string) {
            this.toasts.push([severity, title, message])
        },
        async clearToast() {
            this.toasts = [];
        },
        async refreshOwed() {
            if (this.isCoach) {
                this.amountOwed = (await (await fetch(`/api/coach/owed/${this.id}`)).json()).owed;
            } else {
                this.amountOwed = (await (await fetch(`/api/user/${this.id}`)).json()).balance;
            }
        }
    },
});