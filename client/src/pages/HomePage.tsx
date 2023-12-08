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
  const customerID = 1;
  const [collections, setCollections] = useState<GiftCollection[]>();
  const [displayCollection, setDisplayCollection] = useState<GiftCollection>();
  const [gifts, setGifts] = useState([]);
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
  const [selectedFeature, setSelectedFeature] = useState(true);
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
      setDisplayCollection(response.data[0])
      return response.data;
    } catch (error) {
      console.error("An error occurred while fetching the collection:", error);
    }
  };

  const getGifts = async () => {
    if (displayCollection == undefined){
      return
    }
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
    if (displayCollection == undefined){
      return
    }
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
  const sortedGifts: Gift[] = gifts.sort((a: Gift, b: Gift) => selectedFeature ? a.Price - b.Price : b.Price - a.Price);


  return (
    <div className="bg-eggshell h-full text-white flex flex-col">
      <div className="sticky top-0 z-10 bg-eggshell">
        <NavBar />
        <SearchBar updateHomePage={handleSearchChange} />
      </div>
      <img src={homeGiftsImage}  height="50" alt="home"/>
      <div className="flex flex-col items-center px-[10vw] xl:px-[15vw] mt-20">
        <div className="grid grid-cols-2 gap-6 sm:grid-cols-3 lg:grid-cols-6">
          {collections?.map((collection, index) => (
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
        <div className="max-w-7xl ">
        <GiftSortNavBar
          currentFilters={currentFilters}
          setCurrentFilters={setCurrentFilters}
          selectedFeature={selectedFeature}
          setSelectedFeature={setSelectedFeature}
          className="mt-20"
        />
        <div className="min-h-[19rem] grid grid-cols-1 gap-10 justify-items-center mb-40 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
          {sortedGifts.map((gift: Gift, index) => {
            const isSaved = gift.GiftCollections.some((collection: GiftCollection) => collection.CollectionName === "Favorites" && collection.CustomerID === customerID)
              return (
                <UpdatedGiftItem key={index} gift={gift} isSaved={isSaved} onFavoriteClick={handleFavoriteClick}/>
              )}
          )})
        </div>
        </div>
        <div className="relative mx-auto mb-10 text-xl flex flex-row justify-center  w-full">
          <button
              className=" text-red  font-bold rounded-l px-2"
              onClick={handlePrevious}
          >
            {`<`}
          </button>
          <p className="text-espresso font-bold">{pageNumber}</p>
          <button
              className="text-red font-bold rounded-r px-2"
              onClick={handleNext}
          >
            {`>`}
          </button>
        </div>
      </div>
      <Footer />
    </div>
  );
};
export default HomePage;

