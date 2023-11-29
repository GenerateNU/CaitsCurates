import NavBar from "../components/NavBarUpdated";
import CollectionItem from "../components/CollectionItemUpdated";
import SearchBar from "../components/SearchBar";
import GiftSortNavBar from "../components/GiftSortNavBar";
import UpdatedGiftItem from "../components/UpdatedGiftItem";

import openGift from "../images/openGift.svg";
import closedGift from "../images/closedGift.svg";
import caitsLogo from "../images/caitsCuratesLogo.svg";

import { useMockGifts } from "../Context/MockGiftsContext";
import { useState } from "react";

const HomePage = () => {
  const [giftPageStart, setGiftPageStart] = useState(0);
  const [giftPageEnd, setGiftPageEnd] = useState(16);
  const [pageNumber, setPageNumber] = useState(1);

  const { gifts } = useMockGifts();
  const { collections } = useMockGifts();

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
          <h1 className="font-serif text-4xl text-black font-medium">Budget-Friendly Bliss</h1>
          <h1 className="text-base text-black mt-4">
            Thoughtful gift-giving doesn't have to break the bank. That's why we've curated a delightful collection of affordable gifts that are sure to bring smiles.
          </h1>
        </div>
        <img src={closedGift} alt="closed gift" />
      </div>
      <div className="flex flex-col items-center my-8">
        <div className="overflow-x-auto w-full">
          <div className="flex space-x-4">
            {collections.map((collection, index) => (
              <CollectionItem key={index} name={collection.name} gifts={collection.gifts} />
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
          {gifts.slice(giftPageStart, giftPageEnd).map((gift, index) => (
            <div key={index} className="flex items-center w-80 justify-center mx-4">
              <UpdatedGiftItem name={gift.name} price={gift.price} />
            </div>
          ))}
        </div>
        <div className="relative mx-auto mt-10 flex flex-row justify-center -mx-2 w-full">
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

