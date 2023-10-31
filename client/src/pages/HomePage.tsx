import NavBar from "../components/NavBarUpdated";

const HomePage = () => {
  return (
    <div className="bg-gray-100 h-screen text-white flex flex-col">
      <div className="ml-0">
        <NavBar />
      </div>
      <div className="w-full bg-gray-300 text-center py-9">
        <h1 className="text-2xl text-black">Essential Gifts</h1>
        <h1 className="text-sm text-black mt-4">Handpicked by Cait</h1>
      </div>
      <div className="flex flex-col items-center justify-center my-8">
        <div className="flex"></div>
      </div>
    </div>
  );
};

export default HomePage;


