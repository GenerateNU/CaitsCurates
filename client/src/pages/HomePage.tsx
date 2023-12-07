import NavBar from "../components/Home/NavBarUpdated.tsx";
import CollectionItem from "../components/Home/CollectionItemUpdated.tsx";
import SearchBar from "../components/Home/SearchBar.tsx";
import GiftSortNavBar from "../components/Nav/GiftSortNavBar.tsx";
import UpdatedGiftItem from "../components/Home/UpdatedGiftItem.tsx";
import axios from "axios";
import { useEffect, useState } from "react";
import { Gift, GiftCollection, Filters } from "../types.tsx";

import openGift from "../images/openGift.svg";
import closedGift from "../images/closedGift.svg";
import caitsLogo from "../images/caitsCuratesLogo.svg";
import pink_banner from "../images/pink_banner.svg";

const HomePage = () => {
  const exampleGifts = [
    {
      ID: 1,
      Name: "Customized Jewelry",
      Price: 100,
      Link: "https://example.com/customized-jewelry",
      Occasion: "Anniversary",
      Description: "A personalized piece of jewelry to celebrate your special day.",
      Demographic: "Adults",
      GiftCollections: [],
      Category: ["Jewelry", "Personalized"]
    },
    {
      ID: 2,
      Name: "Tech Gadgets Set",
      Price: 150,
      Link: "https://example.com/tech-gadgets-set",
      Occasion: "Birthday",
      Description: "A bundle of cutting-edge tech gadgets for the tech enthusiast in your life.",
      Demographic: "Tech Enthusiasts",
      GiftCollections: [],
      Category: ["Tech", "Gadgets"]
    },
    {
      ID: 3,
      Name: "Spa Day Experience",
      Price: 80,
      Link: "https://example.com/spa-day-experience",
      Occasion: "Relaxation",
      Description: "Treat your loved one to a rejuvenating spa day experience.",
      Demographic: "All Ages",
      GiftCollections: [],
      Category: ["Wellness", "Experience"]
    },
    {
      ID: 4,
      Name: "Cooking Class Voucher",
      Price: 60,
      Link: "https://example.com/cooking-class-voucher",
      Occasion: "Cooking Enthusiast",
      Description: "A voucher for a fun and educational cooking class.",
      Demographic: "Cooking Enthusiasts",
      GiftCollections: [],
      Category: ["Experience", "Cooking"]
    },
    {
      ID: 5,
      Name: "Book Lover's Subscription Box",
      Price: 30,
      Link: "https://example.com/book-lovers-subscription",
      Occasion: "Bookworm's Delight",
      Description: "A monthly subscription box filled with curated books and literary goodies.",
      Demographic: "Book Lovers",
      GiftCollections: [],
      Category: ["Books", "Subscription"]
    },
    {
      ID: 6,
      Name: "Customized Jewelry",
      Price: 100,
      Link: "https://example.com/customized-jewelry",
      Occasion: "Anniversary",
      Description: "A personalized piece of jewelry to celebrate your special day.",
      Demographic: "Adults",
      GiftCollections: [],
      Category: ["Jewelry"]
    },
    {
      ID: 7,
      Name: "Tech Gadgets Set",
      Price: 150,
      Link: "https://example.com/tech-gadgets-set",
      Occasion: "Birthday",
      Description: "A bundle of cutting-edge tech gadgets for the tech enthusiast in your life.",
      Demographic: "Tech Enthusiasts",
      GiftCollections: [],
      Category: ["Tech", "Gadgets"]
    },
    {
      ID: 8,
      Name: "Spa Day Experience",
      Price: 80,
      Link: "https://example.com/spa-day-experience",
      Occasion: "Relaxation",
      Description: "Treat your loved one to a rejuvenating spa day experience.",
      Demographic: "All Ages",
      GiftCollections: [],
      Category: ["Wellness", "Experience"]
    },
    {
      ID: 9,
      Name: "Cooking Class Voucher",
      Price: 60,
      Link: "https://example.com/cooking-class-voucher",
      Occasion: "Cooking Enthusiast",
      Description: "A voucher for a fun and educational cooking class.",
      Demographic: "Cooking Enthusiasts",
      GiftCollections: [],
      Category: ["Experience", "Cooking"]
    },
    {
      ID: 10,
      Name: "Book Lover's Subscription Box",
      Price: 30,
      Link: "https://example.com/book-lovers-subscription",
      Occasion: "Bookworm's Delight",
      Description: "A monthly subscription box filled with curated books and literary goodies.",
      Demographic: "Book Lovers",
      GiftCollections: [],
      Category: ["Books", "Subscription"]
    },
    {
      ID: 11,
      Name: "Customized Jewelry",
      Price: 100,
      Link: "https://example.com/customized-jewelry",
      Occasion: "Anniversary",
      Description: "A personalized piece of jewelry to celebrate your special day.",
      Demographic: "Adults",
      GiftCollections: [],
      Category: ["Jewelry"]
    },
    {
      ID: 12,
      Name: "Tech Gadgets Set",
      Price: 150,
      Link: "https://example.com/tech-gadgets-set",
      Occasion: "Birthday",
      Description: "A bundle of cutting-edge tech gadgets for the tech enthusiast in your life.",
      Demographic: "Tech Enthusiasts",
      GiftCollections: [],
      Category: ["Tech", "Gadgets"]
    },
    {
      ID: 13,
      Name: "Spa Day Experience",
      Price: 80,
      Link: "https://example.com/spa-day-experience",
      Occasion: "Relaxation",
      Description: "Treat your loved one to a rejuvenating spa day experience.",
      Demographic: "All Ages",
      GiftCollections: [],
      Category: ["Wellness", "Experience"]
    },
    {
      ID: 14,
      Name: "Cooking Class Voucher",
      Price: 60,
      Link: "https://example.com/cooking-class-voucher",
      Occasion: "Cooking Enthusiast",
      Description: "A voucher for a fun and educational cooking class.",
      Demographic: "Cooking Enthusiasts",
      GiftCollections: [],
      Category: ["Experience", "Cooking"]
    },
    {
      ID: 15,
      Name: "Book Lover's Subscription Box",
      Price: 30,
      Link: "https://example.com/book-lovers-subscription",
      Occasion: "Bookworm's Delight",
      Description: "A monthly subscription box filled with curated books and literary goodies.",
      Demographic: "Book Lovers",
      GiftCollections: [],
      Category: ["Books", "Subscription"]
    },
    {
      ID: 16,
      Name: "Customized Jewelry",
      Price: 100,
      Link: "https://example.com/customized-jewelry",
      Occasion: "Anniversary",
      Description: "A personalized piece of jewelry to celebrate your special day.",
      Demographic: "Adults",
      GiftCollections: [],
      Category: ["Jewelry"]
    },
    {
      ID: 17,
      Name: "Tech Gadgets Set",
      Price: 150,
      Link: "https://example.com/tech-gadgets-set",
      Occasion: "Birthday",
      Description: "A bundle of cutting-edge tech gadgets for the tech enthusiast in your life.",
      Demographic: "Tech Enthusiasts",
      GiftCollections: [],
      Category: ["Tech", "Gadgets"]
    },
  ];

  const exampleCustomer = {
    ID: 1,
    UserId: 1,
  };

  const exampleGiftCollection = {
    ID: 1,
    CustomerID: 1,
    Customer: exampleCustomer,
    CollectionName: "Default",
    Gifts: exampleGifts,
  };

  const customerID = 1;
  const [collections, setCollections] = useState<GiftCollection[]>([]);

  const [displayCollection, setDisplayCollection] = useState<GiftCollection>(
    exampleGiftCollection
  );

  const [searchTerm, setSearchTerm] = useState("");
  const [currentFilters, setCurrentFilters] = useState<Filters>({
    minPrice: 0,
    maxPrice: 1000,
    occasion: "",
    demographic: "",
    category: "",
  });

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
    getGifts();
  }, []);

  useEffect(() => {
    getGifts();
  }, [currentFilters, displayCollection, searchTerm]);

  const getCollection = async () => {
    try {
      const response = await axios.get(`/api/collections/${customerID}`);
      setCollections(response.data);
      return response.data;
    } catch (error) {
      console.error("An error occurred while fetching the collection:", error);
    }
  };

  const getGifts = async () => {
    try {
      const response = await axios.get(`/api/search/${displayCollection.ID}`, {
        params: {
          q: searchTerm,
          minPrice: currentFilters.minPrice,
          maxPrice: currentFilters.maxPrice,
          occasion: currentFilters.occasion,
          demographic: currentFilters.demographic,
          category: currentFilters.category,
        },
      });
      setGifts(response.data);
    } catch (error) {
      console.error("An error occurred while fetching gifts:", error);
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
        const currentCollection = updatedCollection.find((collection: GiftCollection) => collection.ID === displayCollection.ID) ?? displayCollection;
        setDisplayCollection(currentCollection);
      }
    } catch (error) {
      console.error('An error occured while favoriting a gift:', error)
    }
  }

  const handleSearchChange = (e: string) => {
    setSearchTerm(e);
  };

  return (
    <div className="bg-eggshell h-fit text-white flex flex-col">
      <div className="ml-0">
        <NavBar />
        <SearchBar updateHomePage={handleSearchChange} />
      </div>
      
      <img src={pink_banner} className="w-full" />
      
      <div className="flex flex-col items-center my-8 h-screen">
        <div className="overflow-x-auto w-full">
          <div className="flex space-x-4">
            {collections.map((collection, index) => (
              <div
                key={index}
                className={`cursor-pointer ${collection === displayCollection ? "font-bold" : ""
                  }`}
                onClick={() => setDisplayCollection(collection)}
              >
                <CollectionItem
                  key={index}
                  name={collection.CollectionName}
                  gifts={collection.Gifts}
                />
              </div>
            ))}
          </div>
        </div>
        <div className=" w-2/3">
          <GiftSortNavBar
            currentFilters={currentFilters}
            setCurrentFilters={setCurrentFilters}
          />
        </div>

        <div className="flex flex-row items-start flex-wrap mx-auto w-2/3">
          {gifts.slice(giftPageStart, giftPageEnd).map((gift, index) => {
            const isSaved = gift.GiftCollections.some((collection: GiftCollection) => collection.CollectionName === "Favorites" && collection.CustomerID === customerID)
            return (
              <div key={index} className="flex items-center w-80 justify-center mx-4">
                <UpdatedGiftItem gift={gift} isSaved={isSaved} onFavoriteClick={handleFavoriteClick} />
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

