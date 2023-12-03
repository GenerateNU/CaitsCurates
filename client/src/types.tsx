export interface Gift {
  ID: number;
  Name: string;
  Price: number;
  Link: string;
  Occasion: string;
  Description: string;
  Demographic: string;
  GiftCollections: GiftCollection[];
  Category: string[];
}

export interface GiftRequest {
  ID: number;
  CustomerID: number;
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

export type AuthenticatedUser = {
  type: 'Admin' | 'Customer';
  details: User;
  additionalInfo: Customer | Admin;
};


export interface GiftCollection {
  ID: number;
  CustomerID: number | null;
  Customer: Customer;
  CollectionName: string;
  Gifts: Gift[];
}

export interface GiftResponse {
  ID: number;
  GiftCollection: GiftCollection;
  GiftCollectionId: number;
  CustomMessage: string;
}

export interface User {
  ID: number;
  Email: string;
  FirstName: string;
  LastName: string;
}

export interface Customer {
  ID: number;
  UserId: number;
}

export interface Admin {
  ID: number;
  UserId: number;
}
