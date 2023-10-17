import { GiftRequest } from "../types";

export const completeRequests: GiftRequest[] = [
    {
        CustomerId: 1,
        GiftResponseId: null,
        RecipientName: "Alice",
        RecipientAge: 25,
        Occasion: ["Birthday", "Anniversary"],
        RecipientInterests: ["Reading", "Traveling"],
        BudgetMax: 100,
        BudgetMin: 50,
        GiftResponse: {
            GiftCollection: {
                CustomerId: 1,
                Customer: { UserId: 1 },
                CollectionName: "Books",
                Gifts: [
                    {
                        Name: "The Catcher in the Rye",
                        Price: 20,
                        Link: "https://book-link",
                        Description: "Classic novel",
                        Demographic: "Adult",
                        GiftCollections: [],
                    },
                    {
                        Name: "To Kill a Mockingbird",
                        Price: 15,
                        Link: "https://book-link",
                        Description: "Classic novel",
                        Demographic: "Adult",
                        GiftCollections: [],
                    },
                ],
            },
            GiftCollectionId: 1,
            CustomMessage: "Happy Birthday!",
        },
        DateNeeded: new Date("2023-01-15"),
    },
    {
        CustomerId: 2,
        GiftResponseId: null,
        RecipientName: "Bob",
        RecipientAge: 30,
        Occasion: ["Christmas"],
        RecipientInterests: ["Music", "Sports"],
        BudgetMax: 150,
        BudgetMin: 100,
        GiftResponse: {
            GiftCollection: {
                CustomerId: 2,
                Customer: { UserId: 2 },
                CollectionName: "Tech Gadgets",
                Gifts: [
                    {
                        Name: "Wireless Headphones",
                        Price: 80,
                        Link: "https://headphones-link",
                        Description: "High-quality sound",
                        Demographic: "Adult",
                        GiftCollections: [],
                    },
                    {
                        Name: "Smartwatch",
                        Price: 120,
                        Link: "https://smartwatch-link",
                        Description: "Fitness tracking and notifications",
                        Demographic: "Adult",
                        GiftCollections: [],
                    },
                ],
            },
            GiftCollectionId: 2,
            CustomMessage: "Merry Christmas!",
        },
        DateNeeded: new Date("2023-12-25"),
    },
    {
        CustomerId: 3,
        GiftResponseId: null,
        RecipientName: "Charlie",
        RecipientAge: 22,
        Occasion: ["Graduation"],
        RecipientInterests: ["Art", "Movies"],
        BudgetMax: 80,
        BudgetMin: 50,
        GiftResponse: {
            GiftCollection: {
                CustomerId: 3,
                Customer: { UserId: 3 },
                CollectionName: "Art Supplies",
                Gifts: [
                    {
                        Name: "Acrylic Paint Set",
                        Price: 30,
                        Link: "https://paint-set-link",
                        Description: "High-quality pigments",
                        Demographic: "Young Adult",
                        GiftCollections: [],
                    },
                    {
                        Name: "Sketchbook",
                        Price: 20,
                        Link: "https://sketchbook-link",
                        Description: "Blank pages for creative ideas",
                        Demographic: "Young Adult",
                        GiftCollections: [],
                    },
                ],
            },
            GiftCollectionId: 3,
            CustomMessage: "Congratulations on your graduation!",
        },
        DateNeeded: new Date("2023-05-20"),
    },
];

export const incompleteRequests: GiftRequest[] = [
    {
        CustomerId: 3,
        GiftResponseId: null,
        RecipientName: "Charlie",
        RecipientAge: 22,
        Occasion: ["Graduation"],
        RecipientInterests: ["Art", "Movies"],
        BudgetMax: 80,
        BudgetMin: 50,
        GiftResponse: null,
        DateNeeded: new Date("2023-05-20"),
    },
];