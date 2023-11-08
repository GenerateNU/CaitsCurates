import NavBar from "../components/NavBarUpdated";
import CollectionItem from "../components/CollectionItemUpdated"; 
import SearchBar from "../components/SearchBar";
import GiftSortNavBar from "../components/GiftSortNavBar";
import UpdatedGiftItem from "../components/UpdatedGiftItem";
import axios from "axios";
import {useEffect, useState} from "react";
import {GiftCollection} from "../types.tsx";

const HomePage = () => {
  const sampleCollections = [
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
  ];

  const gifts = [
    { name: "Gift 1", price: 20 },
    { name: "Gift 2", price: 50 },
    { name: "Gift 3", price: 30 },
    { name: "Gift 4", price: 100 },
    { name: "Gift 5", price: 10 },
    { name: "Gift 6", price: 30 },
    { name: "Gift 7", price: 55 },
    { name: "Gift 8", price: 80 },
  ]

  const customerID = 1;
  const [collections, setCollections] = useState<GiftCollection[]>([]);
  const [displayCollection, setDisplayCollection] = useState(sampleCollections[0]);

  useEffect(() => {
    getCollection();
  }, []);

  const getCollection = async ()=> {
    try {
      const response = await axios.get(`/api/collections/${customerID}`);
      setCollections(response.data);
    } catch (error) {
      console.error('An error occurred while fetching the collection:', error);
    }
  };


  return (
    <div className="bg-gray-100 h-full text-white flex flex-col">
      <div className="ml-0">
        <NavBar />
        <SearchBar />
      </div>
      <div className="w-full bg-gray-300 text-center py-9">
        <h1 className="text-2xl text-black">Essential Gifts</h1>
        <h1 className="text-sm text-black mt-4">Handpicked by Cait</h1>
      </div>
      <div className="flex flex-col items-center my-8">
        <div className="overflow-x-auto w-full">
          <div className="flex space-x-4">
            {collections.map((collection, index) => (
              <CollectionItem key={index} name={collection.name} gifts={collection.gifts} />
            ))}
          </div>
        </div>
        <div className=" w-1000">
          <GiftSortNavBar />
        </div>
  <div className="overflow-y-auto" style={{ maxHeight: '290px', maxWidth: '1000px' }}>
  <div className="flex flex-wrap -mx-2">
    {displayCollection.gifts.map((gift, index) => (
      <div key={index} className="w-1/4 px-2">
        <UpdatedGiftItem name={gift.name} price={gift.price} />
      </div>
    ))}
  </div>
</div>
      </div>
    </div>
  );
};

export default HomePage;

