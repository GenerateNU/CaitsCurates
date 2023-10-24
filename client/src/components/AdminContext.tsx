import React, { createContext, useState, useEffect, ReactNode, useContext } from 'react';
import axios from 'axios';
import { Gift, GiftCollection } from '../types';

interface AdminContextType {
  allGifts: Gift[];
  allGiftCollections: GiftCollection[];
}

const CurrentAdminContext = createContext<AdminContextType | undefined>(undefined);

export const useAdminContext = () => {
  const context = useContext(CurrentAdminContext);
  if (context === undefined) {
    throw new Error('useAdminContext must be used within an AdminContextProvider');
  }
  return context;
};

export const AdminContextProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [allGifts, setAllGifts] = useState<Gift[]>([]);
  const [allGiftCollections, setAllCollections] = useState<GiftCollection[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      const giftsResponse = await axios.get('/api/gifts');
      const giftCollectionsResponse = await axios.get('/api/collections');

      setAllGifts(giftsResponse.data);
      setAllCollections(giftCollectionsResponse.data);
    };

    fetchData();
  }, []);

  return (
    <CurrentAdminContext.Provider value={{ allGifts, allGiftCollections }}>
      {children}
    </CurrentAdminContext.Provider>
  );
};