export interface Gift {
    Name: string;
    Price: number;
    Link: string;
    Description: string;
    Demographic: string;
    GiftCollections: GiftCollection[];
}

export interface GiftRequest {
    CustomerId: number;
    GiftResponseId: number | null;
    RecipientName: string;
    RecipientAge: number;
    Occasion: string[];
    RecipientInterests: string[];
    BudgetMax: number;
    BudgetMin: number;
    GiftResponse: GiftResponse | null;
    DateNeeded: Date;
}

export interface GiftCollection {
    CustomerId: number | null;
    Customer: Customer;
    CollectionName: string;
    Gifts: Gift[];
}

export interface GiftResponse {
    GiftCollection: GiftCollection;
    GiftCollectionId: number;
    CustomMessage: string;
}

export interface User {
    Email: string;
    FirstName: string;
    LastName: string;
    Password: string;
}

export interface Customer {
    UserId: number;
}

export interface Admin {
    UserId: number;
}