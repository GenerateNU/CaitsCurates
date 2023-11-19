import { useState } from "react";
import { useMockGifts } from "../Context/MockGiftsContext";

interface GiftSortNavbarProps {
  items?: {
    name: string;
    price: number;
  }[];
  sortGifts?: () => void;
}

const GiftSortNavbar = () => {
  const [activeButton, setActiveButton] = useState("");
  const [dropdownOpen, setDropdownOpen] = useState(false);
  const [selectedFeature, setSelectedFeature] = useState("Featured");

  const { gifts, setGifts } = useMockGifts();

  const handleButtonClick = (buttonName: any) => {
    setActiveButton(buttonName);
  };

  const handleDropdownToggle = () => {
    setDropdownOpen(!dropdownOpen);
  };

  const handleFeatureSelect = (feature: any) => {
    setSelectedFeature(feature);
    setDropdownOpen(false);
    if (feature === "Price") {
      setGifts([...gifts].sort((a, b) => a.price - b.price));
    } else {
      setGifts([...gifts].sort((a, b) => a.name.localeCompare(b.name)));
    }
  };

  // const buttonStyle = "px-4 py-2 text-black text-xs rounded-md bg-gray-100 border-gray-400";
  const buttonStyle = "px-4 text-[#BC2C36] font-bold";
  const activeButtonStyle = "bg-gray-400";

  return (
    <div className="flex justify-between p-2">
      <div className="flex space-x-2 ">
        <button
          className={`${buttonStyle} ${activeButton === "Shop" ? activeButtonStyle : ""
            }`}
          onClick={() => handleButtonClick("Shop")}
        >
          Shop All Gifts
        </button>
      </div>

      <div className="flex items-center">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth="1.5"
          stroke="black"
          className="w-4 h-4"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M12 3c2.755 0 5.455.232 8.083.678.533.09.917.556.917 1.096v1.044a2.25 2.25 0 01-.659 1.591l-5.432 5.432a2.25 2.25 0 00-.659 1.591v2.927a2.25 2.25 0 01-1.244 2.013L9.75 21v-6.568a2.25 2.25 0 00-.659-1.591L3.659 7.409A2.25 2.25 0 013 5.818V4.774c0-.54.384-1.006.917-1.096A48.32 48.32 0 0112 3z"
          />
        </svg>
        <div className="relative flex flex-row">
          <p className="ml-2 text-base font-normal text-black">
            Sort By:
          </p>
          <button
            className={`${buttonStyle} ${activeButton === "SignUp" ? activeButtonStyle : ""
              }`}
            onClick={handleDropdownToggle}
          >
            {selectedFeature}
          </button>

          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M7.41 8.58008L12 13.1701L16.59 8.58008L18 10.0001L12 16.0001L6 10.0001L7.41 8.58008Z" fill="#32211E" />
          </svg>

          {dropdownOpen && (
            <div className="absolute mt-2 right-0 border border-gray-300 rounded-md bg-gray-100 text-black text-sm shadow-lg z-10">
              <ul>
                <li
                  onClick={() => handleFeatureSelect("Featured")}
                  className="cursor-pointer px-4 py-2 hover:bg-gray-300"
                >
                  Featured
                </li>
                <li
                  onClick={() => handleFeatureSelect("Price")}
                  className="cursor-pointer px-4 py-2 hover:bg-gray-300"
                >
                  Price
                </li>
                <li
                  onClick={() => handleFeatureSelect("Occasion")}
                  className="cursor-pointer px-4 py-2 hover:bg-gray-300"
                >
                  Occassion
                </li>
                {/* Add more options here */}
              </ul>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default GiftSortNavbar;
