import NavBar from "../components/NavBarUpdated";
import CollectionItem from "../components/CollectionItemUpdated"; // Import your CollectionItem component

const HomePage = () => {
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

  return (
    <div className="bg-gray-100 h-screen text-white flex flex-col">
      <div className="ml-0">
        <NavBar />
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
      </div>
    </div>
  );
};

export default HomePage;











