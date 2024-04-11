export type Event = {
        ID: number;
        CreatedAt: string;
        UpdatedAt: string;
        DeletedAt: string | null;
        name: string;
        coach_email: string;
        location: string;
        cost: number;
        Users: [number];
        event_expenses: number;
        coach_expenses: number;
};

