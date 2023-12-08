export interface Gift {
  ID: number;
  Name: string;
  Price: number;
  Link: string;
  Occasion: string;
  Description: string;
  Demographic: string;
  GiftCollections: GiftCollection[];
  ImageLink: string;
  Category: string[];
}

export interface GiftRequest {
  ID: number;
  CustomerID: number;
  GifteeID: number;
  Giftee: Giftee | null;
  GiftResponseID: number | null;
  Occasion: string[];
  BudgetMax: number;
  BudgetMin: number;
  GiftResponse: GiftResponse | null;
  DateNeeded: string;
  Comment: string;
  CreatedAt: string;
}
export interface GiftRequestProps {
  CustomerID: number;
  GifteeID: number;
  Occasion: string[];
  BudgetMax: number;
  BudgetMin: number;
  DateNeeded: Date;
  Comment: string;
}

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
  Password: string;
}

export interface Customer {
  ID: number;
  UserId: number;
}

export interface Giftee {
  ID: number;
  GifteeName: string;
  CustomerID: number;
  Gender: string;
  CustomerRelationship: string;
  Age: number;
  Colors: string[];
  Interests: string[];
  GiftRequests: GiftRequest[];
}
export interface Admin {
  ID: number;
  UserId: number;
}

export type FilterKey = keyof Filters;

export interface Filters {
  minPrice: number;
  maxPrice: number;
  occasion: string;
  demographic: string;
  category: string;
}
