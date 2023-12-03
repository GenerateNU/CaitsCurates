import { useState } from "react";
import RequestModal from "../Requests/RequestModal";
import logoImage from "../../images/logo.svg"

const Navbar = () => {
  const [activeButton, setActiveButton] = useState("");
  const [isRequestModalOpen, setIsRequestModalOpen] = useState(false);

  const handleButtonClick = (buttonName: any) => {
    setActiveButton(buttonName);

    if (buttonName === "Request") {
      setIsRequestModalOpen(true);
    }
  };

  const closeRequestModal = () => {
    setIsRequestModalOpen(false);
    setActiveButton("");
  };

  const buttonStyle =
      "px-4 py-2 text-black text-s rounded-md border-gray-400";
  const activeButtonStyle = "bg-DFB2AA";

  return (
      <div className="bg-FFF9F4 p-5">
        <div className="flex">
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
              About
            </button>
          </div>
          <div className="flex items-center justify-center flex-grow">
            <img
                src={logoImage}
                alt="Logo"
                className="w-15 h-15"
            />
          </div>
          <div className="ml-auto flex items-center space-x-4">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="black" className="w-6 h-6" style={{ marginRight: '10px' }}>
              <path fill-rule="evenodd" d="M10.5 3.75a6.75 6.75 0 100 13.5 6.75 6.75 0 000-13.5zM2.25 10.5a8.25 8.25 0 1114.59 5.28l4.69 4.69a.75.75 0 11-1.06 1.06l-4.69-4.69A8.25 8.25 0 012.25 10.5z" clip-rule="evenodd" />
            </svg>

            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="pink" className="w-6 h-6" style={{ marginRight: '10px' }}>
              <path d="M11.645 20.91l-.007-.003-.022-.012a15.247 15.247 0 01-.383-.218 25.18 25.18 0 01-4.244-3.17C4.688 15.36 2.25 12.174 2.25 8.25 2.25 5.322 4.714 3 7.688 3A5.5 5.5 0 0112 5.052 5.5 5.5 0 0116.313 3c2.973 0 5.437 2.322 5.437 5.25 0 3.925-2.438 7.111-4.739 9.256a25.175 25.175 0 01-4.244 3.17 15.247 15.247 0 01-.383.219l-.022.012-.007.004-.003.001a.752.752 0 01-.704 0l-.003-.001z" />
            </svg>
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="black" className="w-6 h-6" style={{marginRight: "30px"}}>
              <path fill-rule="evenodd" d="M7.5 6a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM3.751 20.105a8.25 8.25 0 0116.498 0 .75.75 0 01-.437.695A18.683 18.683 0 0112 22.5c-2.786 0-5.433-.608-7.812-1.7a.75.75 0 01-.437-.695z" clip-rule="evenodd" />
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
          {isRequestModalOpen && (
              <RequestModal isOpen={isRequestModalOpen} onClose={closeRequestModal}>
                  {/* Add content for the RequestModal if needed */}
                  <p className="text-black">This is the modal content</p>
              </RequestModal>
          )}
      </div>
  );
};

export default Navbar;