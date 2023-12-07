import { useState } from "react";
import {Filters} from "../../types.tsx";
import Filter from "../Home/Filter.tsx";
import filterIcon from "../../images/filter.svg";
import sortIcon from "../../images/sort.svg";
import downIcon from "../../images/chevron_down.svg";


type GiftSortNavBarProps = {
  currentFilters: Filters;
  setCurrentFilters: React.Dispatch<React.SetStateAction<Filters>>;
  className?: string;
};

const GiftSortNavbar: React.FC<GiftSortNavBarProps> = ({
  currentFilters,
  setCurrentFilters,
  className
}) => {
  const [activeButton, setActiveButton] = useState("");
  const [dropdownOpen, setDropdownOpen] = useState(false);
  const [filterOpen, setFilterOpen] = useState(false);
  const [selectedFeature, setSelectedFeature] = useState("Featured");

  const handleButtonClick = (buttonName: any) => {
    setActiveButton(buttonName);
  };

  const handleDropdownToggle = () => {
    setDropdownOpen(!dropdownOpen);
  };

  const handleFilterToggle = () => {
    setFilterOpen(!filterOpen);
  };

  const handleFeatureSelect = (feature: any) => {
    setSelectedFeature(feature);
    setDropdownOpen(false);
  };

  const buttonStyle =
    "px-4 py-1 text-red font-proxima";
  const activeButtonStyle = "bg-beige";

  return (
    <div className={`bg-eggshell py-4 ${className}`}>
      <div className="flex" style={{ maxWidth: "1440px" }}>
        <div className="flex space-x-2 items-center">
          <img src={filterIcon} className="h-6 w-6" onClick={handleFilterToggle} />
          <button
            className={`${buttonStyle} text-sm font-bold ${
              activeButton === "Shop" ? activeButtonStyle : "bg-eggshell"
            }`}
            onClick={() => handleButtonClick("Shop")}
          >
            Shop All Gifts
          </button>
        </div>

        <div className="ml-auto flex items-center">
          <img src={sortIcon} className="h-6 w-6" />
          <div className="relative">
            <button
              className={`${buttonStyle} text-2xl ${
                activeButton === "SignUp" ? activeButtonStyle : ""
              }`}
              onClick={handleDropdownToggle}
            >
              <span className="text-coffee text-2xl font-proxima"> Sort by: </span>
              {selectedFeature}
              <img src={downIcon} className="w-6 h-6 inline ml-6" />
            </button>
            {dropdownOpen && (
              <div className="absolute mt-2 right-0 border border-gray-300 rounded-md bg-gray-100 text-red text-sm shadow-lg z-10">
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
                    Occasion
                  </li>
                </ul>
              </div>
            )}
          </div>
          <Filter
            isOpen={filterOpen}
            filterToggle={handleFilterToggle}
            currentFilters={currentFilters}
            setCurrentFilters={setCurrentFilters}
          />
        </div>
      </div>
    </div>
  );
};

export default GiftSortNavbar;
