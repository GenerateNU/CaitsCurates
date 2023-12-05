import NavBar from "../components/NavBarUpdated";
import CollectionItem from "../components/CollectionItemUpdated";
import SearchBar from "../components/SearchBar";
import GiftSortNavBar from "../components/GiftSortNavBar";
import UpdatedGiftItem from "../components/UpdatedGiftItem";
import axios from "axios";
import { useEffect, useState } from "react";
import { Gift, GiftCollection } from "../types.tsx";

import openGift from "../images/openGift.svg";
import closedGift from "../images/closedGift.svg";
import caitsLogo from "../images/caitsCuratesLogo.svg";
import { exampleGifts } from "./mocks/MockGifts";

const HomePage = () => {

  const exampleCustomer = {
    ID: 1,
    UserId: 1,
  }

  const exampleGiftCollection = {
    ID: 1,
    CustomerID: 1,
    Customer: exampleCustomer,
    CollectionName: 'Default',
    Gifts: exampleGifts,
  }

  const customerID = 1;
  const [collections, setCollections] = useState<GiftCollection[]>([]);
  const [displayCollection, setDisplayCollection] = useState<GiftCollection>(exampleGiftCollection);

  const [giftPageStart, setGiftPageStart] = useState(0);
  const [giftPageEnd, setGiftPageEnd] = useState(16);
  const [pageNumber, setPageNumber] = useState(1);

  const [gifts, setGifts] = useState<Gift[]>(exampleGifts);

  const handleNext = () => {
    if (gifts.length / 16 <= pageNumber) return;
    setGiftPageStart(giftPageStart + 16);
    setGiftPageEnd(giftPageEnd + 16);
    setPageNumber(pageNumber + 1);
  };

  const handlePrevious = () => {
    if (giftPageStart === 0) return;
    setGiftPageStart(giftPageStart - 16);
    setGiftPageEnd(giftPageEnd - 16);
    setPageNumber(pageNumber - 1);
  };

  useEffect(() => {
    getCollection();
  }, []);

  const getCollection = async (): Promise<GiftCollection[] | undefined> => {
    try {
      const response = await axios.get(`/api/collections/${customerID}`);
      setCollections(response.data);
      return response.data;
    } catch (error) {
      console.error('An error occurred while fetching the collection:', error);
    }
  };

  const handleFavoriteClick = async (gift: Gift, isSaved: boolean) => {
    const baseUrl = isSaved ? "/api/removeCustomerGiftCollection" : "/api/addCustomerGiftCollection"
    try {
      await axios.post(`${baseUrl}/Favorites/${customerID}`, gift)
      // refetch customer gift collections
      const updatedCollection = await getCollection();

      if (updatedCollection) {
        // on success set state for currently displayed collection
        const currentCollection = updatedCollection.find((collection) => collection.ID === displayCollection.ID) ?? displayCollection;
        setDisplayCollection(currentCollection);
      }
    } catch (error) {
      console.error('An error occured while favoriting a gift:', error)
    }
  }

  return (
    <div className="bg-[#FFF9F4] h-full text-white flex flex-col">

      <div className="ml-0">

        <div className="flex justify-center w-full absolute z-0 px-8 py-4">
          <img src={caitsLogo} alt="caits-logo.svg" className="mx-auto" />
        </div>

        <NavBar />
        <SearchBar />
      </div>

      <div className="flex flex-row w-full bg-[#F4E6DC] text-center py-20 justify-evenly">
        <img src={openGift} alt="open gift" />
        <div className="flex flex-col items-center w-1/3">
          <h1 className="text-2xl text-black">Essential Gifts</h1>
          <h1 className="text-sm text-black mt-4">Handpicked by Cait</h1>
        </div>
        <img src={closedGift} alt="closed gift" />
      </div>
      <div className="flex flex-col items-center my-8">
        <div className="overflow-x-auto w-full">
          <div className="flex space-x-4">
            {collections.map((collection, index) => (
              <div
                className={`cursor-pointer ${collection.ID === displayCollection.ID ? 'font-bold' : ''}`}
                onClick={() => setDisplayCollection(collection)}>
                <CollectionItem key={index} name={collection.CollectionName} gifts={collection.Gifts} />
              </div>
            ))}
          </div>
        </div>
        <div className=" w-2/3">
          <h1 className="text-2xl text-black pl-2 pt-5 font-serif">
            {gifts.length != 1 ? `${gifts.length} Items` : `${gifts.length} Item`}
          </h1>
          <GiftSortNavBar />
        </div>

        <div className="flex flex-row items-start flex-wrap mx-auto w-2/3">
          {displayCollection.Gifts.slice(giftPageStart, giftPageEnd).map((gift, index) => {
            const isSaved = gift.GiftCollections === null ? false : gift.GiftCollections.some(
              (collection) => collection.CollectionName === "Favorites" && collection.CustomerID === customerID)
            return (
              <div key={index} className="flex items-center w-80 justify-center mx-4">
                <UpdatedGiftItem key={gift.ID} gift={gift} isSaved={isSaved} onFavoriteClick={handleFavoriteClick} />
              </div>
            )
          })}
        </div>

        <div className="relative mx-auto mt-10 flex flex-row justify-center -mx-2 w-full">
          <button
            className=" text-gray-800 font-bold rounded-l px-2"
            onClick={handlePrevious}
          >
            {`<`}
          </button>
          <p className="text-gray-800 font-bold">{pageNumber}</p>
          <button
            className="text-gray-800 font-bold rounded-r px-2"
            onClick={handleNext}
          >
            {`>`}
          </button>
        </div>

      </div>
    </div>
  );
};

export default HomePage;

