export interface Gift {
    ID : number
    Name: string;
    Price: number;
    Link: string;
    Description: string;
    Demographic: string;
    GiftCollections: GiftCollection[];
}

export interface GiftRequest {
    ID : number
    CustomerId: number;
    GiftResponseId: number | null;
    RecipientName: string;
    RecipientAge: number;
    Occasion: string[];
    RecipientInterests: string[];
    BudgetMax: number;
    BudgetMin: number;
    GiftResponse: GiftResponse | null;
    DateNeeded: string;
}

export interface GiftCollection {
    ID : number
    CustomerId: number | null;
    Customer: Customer;
    CollectionName: string;
    Gifts: Gift[];
}

export interface GiftResponse {
    ID : number
    GiftCollection: GiftCollection;
    GiftCollectionId: number;
    CustomMessage: string;
}

export interface User {
    ID : number
    Email: string;
    FirstName: string;
    LastName: string;
    Password: string;
}

export interface Customer {
    ID : number
    UserId: number;
}

export interface Admin {
    ID : number
    UserId: number;
}