import { useState } from "react";

const Navbar = () => {
  const [activeButton, setActiveButton] = useState("");

  const handleButtonClick = (buttonName: any) => {
    setActiveButton(buttonName);
  };

  const buttonStyle =
    "px-4 py-2 text-[#32211E] font-bold rounded-md border-gray-400";
  const activeButtonStyle = "bg-[#F9DCD7]";

  return (
    <div className="px-8 py-4">
      <div className="flex w-full justify-between">
        <div className="flex space-x-2 z-10">
          <button
            className={`${buttonStyle} ${activeButton === "Shop" ? activeButtonStyle : ""
              }`}
            onClick={() => handleButtonClick("Shop")}
          >
            Shop
          </button>
          <button
            className={`${buttonStyle} ${activeButton === "Request" ? activeButtonStyle : ""
              }`}
            onClick={() => handleButtonClick("Request")}
          >
            Request
          </button>
          <button
            className={`${buttonStyle} ${activeButton === "More" ? activeButtonStyle : ""
              }`}
            onClick={() => handleButtonClick("More")}
          >
            More
          </button>
        </div>

        <div className="flex items-center space-x-5 z-20">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M6.5 0C8.22391 0 9.87721 0.684819 11.0962 1.90381C12.3152 3.12279 13 4.77609 13 6.5C13 8.11 12.41 9.59 11.44 10.73L11.71 11H12.5L17.5 16L16 17.5L11 12.5V11.71L10.73 11.44C9.59 12.41 8.11 13 6.5 13C4.77609 13 3.12279 12.3152 1.90381 11.0962C0.684819 9.87721 0 8.22391 0 6.5C0 4.77609 0.684819 3.12279 1.90381 1.90381C3.12279 0.684819 4.77609 0 6.5 0ZM6.5 2C4 2 2 4 2 6.5C2 9 4 11 6.5 11C9 11 11 9 11 6.5C11 4 9 2 6.5 2Z" fill="#32211E" />
          </svg>

          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M10 18.35L8.55 17.03C3.4 12.36 0 9.27 0 5.5C0 2.41 2.42 0 5.5 0C7.24 0 8.91 0.81 10 2.08C11.09 0.81 12.76 0 14.5 0C17.58 0 20 2.41 20 5.5C20 9.27 16.6 12.36 11.45 17.03L10 18.35Z" fill="#D08484" />
          </svg>

          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 0C9.06087 0 10.0783 0.421427 10.8284 1.17157C11.5786 1.92172 12 2.93913 12 4C12 5.06087 11.5786 6.07828 10.8284 6.82843C10.0783 7.57857 9.06087 8 8 8C6.93913 8 5.92172 7.57857 5.17157 6.82843C4.42143 6.07828 4 5.06087 4 4C4 2.93913 4.42143 1.92172 5.17157 1.17157C5.92172 0.421427 6.93913 0 8 0ZM8 10C12.42 10 16 11.79 16 14V16H0V14C0 11.79 3.58 10 8 10Z" fill="#32211E" />
          </svg>

        </div>
      </div>
    </div>
  );
};

export default Navbar;
