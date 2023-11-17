import NavBar from "../components/NavBarUpdated";
import CollectionItem from "../components/CollectionItemUpdated"; 
import SearchBar from "../components/SearchBar";
import GiftSortNavBar from "../components/GiftSortNavBar";
import UpdatedGiftItem from "../components/UpdatedGiftItem";
import axios from "axios";
import {useEffect, useState} from "react";
import {GiftCollection} from "../types.tsx";

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
    }
  ];

  const exampleCustomer = {
    ID: 1,
    UserId: 1,
  }

  const exampleGiftCollection = {
    ID: 1,
    CustomerId: 1,
    Customer: exampleCustomer,
    CollectionName: 'Default',
    Gifts: exampleGifts,
  }

  const customerID = 1;
  const [collections, setCollections] = useState<GiftCollection[]>([]);
  const [displayCollection, setDisplayCollection] = useState<GiftCollection>(exampleGiftCollection);

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
                <div
                    className={`cursor-pointer ${collection === displayCollection ? 'font-bold' : ''}`}
                    onClick={() => setDisplayCollection(collection)}>
                  <CollectionItem  key={index} name={collection.CollectionName} gifts={collection.Gifts} />
                </div>
            ))}
          </div>
        </div>
        <div className=" w-1000">
          <GiftSortNavBar />
        </div>
  <div className="overflow-y-auto" style={{ maxHeight: '290px', maxWidth: '1000px' }}>
  <div className="flex flex-wrap -mx-2">
    {displayCollection.Gifts.map((gift, index) => (
      <div key={index} className="w-1/4 px-2">
        <UpdatedGiftItem name={gift.Name} price={gift.Price} description={gift.Description}/>
      </div>
    ))}
  </div>
</div>
      </div>
    </div>
  );
};

export default HomePage;

