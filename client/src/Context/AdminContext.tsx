import React, {createContext, useState, useContext, ReactNode} from 'react';
import {Gift, GiftCollection} from "../types.tsx";
import axios from "axios";

// Defining the shape of the context
interface AdminContextProps {
    gifts: Gift[];
    setGifts: React.Dispatch<React.SetStateAction<Gift[]>>;
    collections: GiftCollection[];
    setCollections: React.Dispatch<React.SetStateAction<GiftCollection[]>>;
    fetchGifts: () => void;
    fetchGiftCollections: () => void;


}

// Creating the context with a default value
const AdminContext = createContext<AdminContextProps | null>(null);

interface AdminProviderProps {
    children: ReactNode;
}

export const useAdmin = () => {
    const context = useContext(AdminContext);
    if (!context) {
        throw new Error('useAdmin must be used within an AdminProvider');
    }
    return context;
};

export function AdminProvider({ children }: AdminProviderProps) {
    const [gifts, setGifts] = useState<Gift[]>([]);
    const [collections, setCollections] = useState<GiftCollection[]>([]);

    const fetchGifts = async () => {
        try {
            const response = await axios.get('/api/gifts');
            setGifts(response.data);
        } catch (error) {
            console.error('An error occurred while fetching the gifts:', error);
        }
    };
    const fetchGiftCollections = async () => {
        try {
            const response = await axios.get('/api/collections');
            setCollections(response.data);
        } catch (error) {
            console.error('An error occurred while fetching the gifts:', error);
        }
    };

   // The function fetchGifts will run when the component mounts
    // Fetch data and other admin-specific logic

    return (
        <AdminContext.Provider value={{ gifts, setGifts, collections, setCollections, fetchGifts, fetchGiftCollections }}>
            {children}
        </AdminContext.Provider>
    );
}
