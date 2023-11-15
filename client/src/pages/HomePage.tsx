import NavBar from "../components/NavBarUpdated";
import CollectionItem from "../components/CollectionItemUpdated";
import SearchBar from "../components/SearchBar";
import GiftSortNavBar from "../components/GiftSortNavBar";
import UpdatedGiftItem from "../components/UpdatedGiftItem";

import openGift from "../images/openGift.svg";
import closedGift from "../images/closedGift.svg";

import { useState } from "react";

const HomePage = () => {
  const [giftPageStart, setGiftPageStart] = useState(0);
  const [giftPageEnd, setGiftPageEnd] = useState(16);
  const [pageNumber, setPageNumber] = useState(1);


  const collections = [
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

  const [gifts, setGifts] = useState([
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

  const sortGifts = () => {
    if (gifts && setGifts) {
      const sortedItems = gifts.sort((a: any, b: any) => a.price - b.price);
      setGifts(sortedItems);
      console.log("sorted");
    }
  };

  return (
    <div className="bg-[#FFF9F4] h-full text-white flex flex-col">
      <div className="ml-0">
        <NavBar />
        <SearchBar />
      </div>
      <div className="flex flex-row w-full bg-[#F4E6DC] text-center py-20">
        <img src={openGift} alt="open gift" className="mx-auto" />
        <div className="flex flex-col items-center w-1/3">
          <h1 className="text-2xl text-black">Budget-Friendly Bliss</h1>
          <h1 className="text-sm text-black mt-4">
            Thoughtful gift-giving doesn't have to break the bank. That's why we've curated a delightful collection of affordable gifts that are sure to bring smiles.
          </h1>
        </div>
        <img src={closedGift} alt="closed gift" className="mx-auto" />
      </div>
      <div className="flex flex-col items-center my-8">
        <div className="overflow-x-auto w-full">
          <div className="flex space-x-4">
            {collections.map((collection, index) => (
              <CollectionItem key={index} name={collection.name} gifts={collection.gifts} />
            ))}
          </div>
        </div>
        <div className=" w-4/5">
          <h1 className="text-2xl text-black">{`${gifts.length} Items`}</h1>
          <GiftSortNavBar sortGifts={sortGifts} />
        </div>
        <div className="flex flex-wrap -mx-2 w-4/5  gap-5">
          {gifts.slice(giftPageStart, giftPageEnd).map((gift, index) => (
            <div key={index} className="flex items-center w-80">
              <UpdatedGiftItem name={gift.name} price={gift.price} />
            </div>
          ))}
        </div>
        <div className="flex flex-row justify-center -mx-2 w-full">
          <button
            className=" text-gray-800 font-bold rounded-l px-2"
            onClick={handlePrevious}
          >
            {`<`}
          </button>
          <p className="text-gray-800 font-bold">
            {pageNumber}
          </p>
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

