import NavBar from "../components/Home/NavBarUpdated.tsx";
import CollectionItem from "../components/Home/CollectionItemUpdated.tsx";
import SearchBar from "../components/Home/SearchBar.tsx";
import GiftSortNavBar from "../components/Nav/GiftSortNavBar.tsx";
import UpdatedGiftItem from "../components/Home/UpdatedGiftItem.tsx";
import axios from "axios";
import { useEffect, useState } from "react";
import { Gift, GiftCollection, Filters } from "../types.tsx";
import homeGiftsImage from "../images/home_gifts_image.svg";
import Footer from "../components/Home/Footer.tsx";

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
        const currentCollection = updatedCollection.find((collection : GiftCollection) => collection.ID === displayCollection.ID) ?? displayCollection;
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
    <div className="bg-eggshell h-full text-white flex flex-col">
      <NavBar />
      <SearchBar updateHomePage={handleSearchChange} />
      <img src={homeGiftsImage}  height="50"/>
      <div className="flex flex-col px-[3vw] xl:px-[13vw] mt-20">
        <div className="grid grid-cols-2 gap-6 sm:grid-cols-3 lg:grid-cols-6">
          {collections.map((collection, index) => (
              <div
                  key={index}
                  className={"cursor-pointer"}
                  onClick={() => setDisplayCollection(collection)}
              >
                <CollectionItem
                    key={index}
                    name={collection.CollectionName}
                    collectionIndex={index}
                    selected={collection === displayCollection}
                />
              </div>
          ))}
        </div>
        <GiftSortNavBar
          currentFilters={currentFilters}
          setCurrentFilters={setCurrentFilters}
          className="mt-20"
        />
        <div className="min-h-[19rem] grid grid-cols-1 gap-y-5 justify-items-center mb-40 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
          {gifts.map((gift, index) => {
            const isSaved = gift.GiftCollections.some((collection: GiftCollection) => collection.CollectionName === "Favorites" && collection.CustomerID === customerID)

              return (
                <UpdatedGiftItem key={index} gift={gift} isSaved={isSaved} onFavoriteClick={handleFavoriteClick}/>
              )}
          )}
        </div>
      </div>
      <Footer />
    </div>
  );
};
export default HomePage;

