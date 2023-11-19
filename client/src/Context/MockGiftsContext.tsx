import React, { createContext, useState, useContext, ReactNode } from 'react';
import { Gift, GiftCollection } from "../types.tsx";

// Defining the shape of the context
interface MockGiftsContextProps {
    gifts: Gift[];
    setGifts: React.Dispatch<React.SetStateAction<Gift[]>>;
    collections: GiftCollection[];
    setCollections: React.Dispatch<React.SetStateAction<GiftCollection[]>>;
    fetchGifts: () => void;
    fetchGiftCollections: () => void;
}

// Creating the context with a default value
const MockGiftsContext = createContext<MockGiftsContextProps | null>(null);

interface MockGiftsProviderProps {
    children: ReactNode;
}

export const useMockGifts = () => {
    const context = useContext(MockGiftsContext);
    if (!context) {
        throw new Error('useMockGifts must be used within an MockGiftsProvider');
    }
    return context;
};

export function MockGiftsProvider({ children }: MockGiftsProviderProps) {
    const [gifts, setGifts] = useState<Gift[]>([
        { name: "Gift 1", price: 20 },
        { name: "Gift 2", price: 50 },
        { name: "Gift 3", price: 30 },
        { name: "Gift 4", price: 100 },
        { name: "Gift 5", price: 10 },
        { name: "Gift 6", price: 30 },
        { name: "Gift 7", price: 55 },
        { name: "Gift 8", price: 80 },
        { name: "Gift 9", price: 20 },
        { name: "Gift 10", price: 50 },
        { name: "Gift 11", price: 30 },
        { name: "Gift 12", price: 100 },
        { name: "Gift 13", price: 10 },
        { name: "Gift 14", price: 30 },
        { name: "Gift 15", price: 55 },
        { name: "Gift 16", price: 80 },
        { name: "Gift 17", price: 20 },
        { name: "Gift 18", price: 50 },
        { name: "Gift 19", price: 30 },
        { name: "Gift 20", price: 100 },
        { name: "Gift 21", price: 10 },
        { name: "Gift 22", price: 30 },
    ]);

    const [collections, setCollections] = useState<GiftCollection[]>([
        { name: "Collection 1", gifts: [] },
        { name: "Collection 2", gifts: [] },
        { name: "Collection 3", gifts: [] },
        { name: "Collection 4", gifts: [] },
        { name: "Collection 5", gifts: [] },
        { name: "Collection 6", gifts: [] },
        { name: "Collection 7", gifts: [] },
        { name: "Collection 8", gifts: [] },
        { name: "Collection 9", gifts: [] },
        { name: "Collection 10", gifts: [] },
        { name: "Collection 11", gifts: [] },
        { name: "Collection 12", gifts: [] },
        { name: "Collection 13", gifts: [] },
        { name: "Collection 14", gifts: [] },
        { name: "Collection 15", gifts: [] },
    ]);

    // The function fetchGifts will run when the component mounts
    // Fetch data and other admin-specific logic
    return (
        <MockGiftsContext.Provider value={{ gifts, setGifts, collections, setCollections }}>
            {children}
        </MockGiftsContext.Provider>
    );
}
