import { Link } from "react-router-dom";

const Navbar = () => {
  const buttonStyle =
    "bg-blue-600 px-4 py-2 h-10 text-white rounded-md self-end";

  return (
    <div className="bg-white p-4 shadow-md sticky top-0 z-50">
      <div className="HEADER-FLEX flex justify-between items-center w-full h-full px-6 laptop:px-10">
        <Link to={"/"}>
          <p className="text-4xl text-black hover:text-pink-600">
            Caits Curates
          </p>
        </Link>
        <div className="flex space-x-2">
          <button className={buttonStyle}>
            <Link to={"/gifts/"}>Manage gifts</Link>
          </button>
          <button className={buttonStyle}>
            <Link to={"/collections/"}>Manage collections</Link>
          </button>
          <button className={buttonStyle}>
            <Link to={"/requests/"}>Manage requests</Link>
          </button>
        </div>
      </div>
    </div>
  );
};

export default Navbar;
