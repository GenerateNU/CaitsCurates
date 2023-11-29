import NavBar from "../components/Home/NavBarUpdated.tsx";
import CollectionItem from "../components/Home/CollectionItemUpdated.tsx";
import SearchBar from "../components/Home/SearchBar.tsx";
import GiftSortNavBar from "../components/Nav/GiftSortNavBar.tsx";
import UpdatedGiftItem from "../components/Home/UpdatedGiftItem.tsx";
import axios from "axios";
import { useEffect, useState } from "react";
import { Gift, GiftCollection, Filters } from "../types.tsx";

const HomePage = () => {
  const exampleGifts = [
    {
      ID: 1,
      Name: "Customized Jewelry",
      Price: 100,
      Link: "https://example.com/customized-jewelry",
      Occasion: "Anniversary",
      Description:
        "A personalized piece of jewelry to celebrate your special day.",
      Demographic: "Adults",
      GiftCollections: [],
      Category: ["Jewelry", "Personalized"],
    },
    {
      ID: 2,
      Name: "Tech Gadgets Set",
      Price: 150,
      Link: "https://example.com/tech-gadgets-set",
      Occasion: "Birthday",
      Description:
        "A bundle of cutting-edge tech gadgets for the tech enthusiast in your life.",
      Demographic: "Tech Enthusiasts",
      GiftCollections: [],
      Category: ["Tech", "Gadgets"],
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
      Category: ["Wellness", "Experience"],
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
      Category: ["Experience", "Cooking"],
    },
    {
      ID: 5,
      Name: "Book Lover's Subscription Box",
      Price: 30,
      Link: "https://example.com/book-lovers-subscription",
      Occasion: "Bookworm's Delight",
      Description:
        "A monthly subscription box filled with curated books and literary goodies.",
      Demographic: "Book Lovers",
      GiftCollections: [],
      Category: ["Books", "Subscription"],
    },
  ];

  const exampleCustomer = {
    ID: 1,
    UserId: 1,
  };

  const exampleGiftCollection = {
    ID: 1,
    CustomerId: 1,
    Customer: exampleCustomer,
    CollectionName: "Default",
    Gifts: exampleGifts,
  };

  const customerID = 1;
  const [collections, setCollections] = useState<GiftCollection[]>([]);
  const [displayCollection, setDisplayCollection] = useState<GiftCollection>(
    exampleGiftCollection
  );
  const [gifts, setGifts] = useState(exampleGifts);
  const [searchTerm, setSearchTerm] = useState("");
  const [currentFilters, setCurrentFilters] = useState<Filters>({
    minPrice: 0,
    maxPrice: 1000,
    occasion: "",
    demographic: "",
    category: "",
  });

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

  const handleSearchChange = (e: string) => {
    setSearchTerm(e);
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
    <div className="bg-gray-100 h-full text-white flex flex-col">
      <div className="ml-0">
        <NavBar />
        <SearchBar updateHomePage={handleSearchChange} />
      </div>
      <div className="w-full bg-gray-300 text-center py-9">
        <h1 className="text-2xl text-black font-seasons">Essential Gifts</h1>
        <h1 className="text-sm text-black font-proxima">Handpicked by Cait</h1>
      </div>
      <div className="flex flex-col items-center my-8">
        <div className="overflow-x-auto w-full">
          <div className="flex space-x-4">
            {collections.map((collection, index) => (
              <div
                key={index}
                className={`cursor-pointer ${
                  collection === displayCollection ? "font-bold" : ""
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
        <div className=" w-1000">
          <GiftSortNavBar
            currentFilters={currentFilters}
            setCurrentFilters={setCurrentFilters}
          />
        </div>

        <div
          className="overflow-y-auto"
          style={{ maxHeight: "290px", maxWidth: "1000px" }}>
          <div className="flex flex-wrap justify-between gap-4">
            {displayCollection.Gifts.map((gift, index) => {
                const isSaved = gift.GiftCollections === null ? false : gift.GiftCollections.some((collection) => collection.CollectionName === "Favorites" && collection.CustomerID === customerID)
                return(
                <div key={index}>
                <UpdatedGiftItem name={gift.Name} price={gift.Price} isSaved={isSaved} onFavoriteClick={handleFavoriteClick}  gift={gift}/>
              </div>
            )
            })}
          </div>
        </div>
      </div>
    </div>
  );
};

export default HomePage;

