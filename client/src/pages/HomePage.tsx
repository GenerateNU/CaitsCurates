import Navbar from "../components/Navbar";
import logo from "../logo.png";

const HomePage = () => {
  return (
    <div>
      <Navbar />
      <div className="flex flex-col items-center justify-center my-8">
        <p className="text-4xl text-black hover:text-pink-600">Caits Curates</p>
        <div className="flex">
          <img src={logo} alt="Logo" width={500} />
        </div>
      </div>
    </div>
  );
};
export default HomePage;
