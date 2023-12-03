import { useState } from "react";

const Navbar = () => {
  const [activeButton, setActiveButton] = useState("");

  const handleButtonClick = (buttonName: any) => {
    setActiveButton(buttonName);
  };

  const buttonStyle =
    "px-4 py-2 text-black text-xs rounded-md border-gray-400";
  const activeButtonStyle = "bg-gray-400";

  return (
    <div className="bg-gray-100 p-2">
      <div className="flex" style={{ maxWidth: "1440px" }}>
        <div className="flex space-x-2">
          <button
            className={`${buttonStyle} ${
              activeButton === "Shop" ? activeButtonStyle : ""
            }`}
            onClick={() => handleButtonClick("Shop")}
          >
            Shop
          </button>
          <button
            className={`${buttonStyle} ${
              activeButton === "Request" ? activeButtonStyle : ""
            }`}
            onClick={() => handleButtonClick("Request")}
          >
            Request
          </button>
          <button
            className={`${buttonStyle} ${
              activeButton === "More" ? activeButtonStyle : ""
            }`}
            onClick={() => handleButtonClick("More")}
          >
            More
          </button>
        </div>
        <div className="ml-auto flex items-center">
  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="black" className="w-4 h-4" style={{ marginRight: '10px' }}>
    <path fillRule="evenodd" d="M10.5 3.75a6.75 6.75 0 100 13.5 6.75 6.75 0 000-13.5zM2.25 10.5a8.25 8.25 0 1114.59 5.28l4.69 4.69a.75.75 0 11-1.06 1.06l-4.69-4.69A8.25 8.25 0 012.25 10.5z" clipRule="evenodd" />
  </svg>

  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="black" className="w-4 h-4" style={{ marginRight: '10px' }}>
    <path d="M11.645 20.91l-.007-.003-.022-.012a15.247 15.247 0 01-.383-.218 25.18 25.18 0 01-4.244-3.17C4.688 15.36 2.25 12.174 2.25 8.25 2.25 5.322 4.714 3 7.688 3A5.5 5.5 0 0112 5.052 5.5 5.5 0 0116.313 3c2.973 0 5.437 2.322 5.437 5.25 0 3.925-2.438 7.111-4.739 9.256a25.175 25.175 0 01-4.244 3.17 15.247 15.247 0 01-.383.219l-.022.012-.007.004-.003.001a.752.752 0 01-.704 0l-.003-.001z" />
  </svg>
          <button
            className={`${buttonStyle} ${
              activeButton === "LogIn" ? activeButtonStyle : ""
            }`}
            onClick={() => handleButtonClick("LogIn")}
          >
            Log In
          </button>
          <button
            className={`${buttonStyle} ${
              activeButton === "SignUp" ? activeButtonStyle : ""
            }`}
            onClick={() => handleButtonClick("SignUp")}
          >
            Sign Up
          </button>
        </div>
      </div>
    </div>
  );
};

export default Navbar;
