import { useState } from "react";
import {FilterKey, Filters} from "../../types.tsx";
import Filter from "../Home/Filter.tsx";
import filterIcon from "../../images/filter.svg";
import sortIcon from "../../images/sort.svg";
import downIcon from "../../images/chevron_down.svg";



type GiftSortNavBarProps = {
  currentFilters: Filters;
  setCurrentFilters: React.Dispatch<React.SetStateAction<Filters>>;
  className?: string;
  selectedFeature: boolean,
  setSelectedFeature: React.Dispatch<React.SetStateAction<boolean>>
};

const GiftSortNavbar: React.FC<GiftSortNavBarProps> = ({
  currentFilters,
  setCurrentFilters ,
  className, selectedFeature , setSelectedFeature
}) => {
  const [dropdownOpen, setDropdownOpen] = useState(false);
  const [filterOpen, setFilterOpen] = useState(false);



  const handleDropdownToggle = () => {
    setDropdownOpen(!dropdownOpen);
  };

  const handleFilterToggle = () => {
    setFilterOpen(!filterOpen);
  };

  const handleFeatureSelect = (feature: boolean) => {
    setSelectedFeature(feature);
    setDropdownOpen(false);
  };



  const removeFilter = (key: FilterKey) => {
    const newFilters = { ...currentFilters };
    delete newFilters[key];
    setCurrentFilters(newFilters);
  };

  return (
    <div className={`bg-eggshell py-4 ${className}`}>
      <div className="flex " >
        <div className="flex space-x-2 items-center">
          <img src={filterIcon} className="h-9 w-9" onClick={handleFilterToggle}  alt={"filter"}/>
        </div>
          <Filter
            isOpen={filterOpen}
            filterToggle={handleFilterToggle}
            currentFilters={currentFilters}
            setCurrentFilters={setCurrentFilters}
          />
        <div className="flex flex-wrap items-left mr-10 gap-4">
            {Object.entries(currentFilters).map(([key, value]) => {
              if (value === "") return null;
              const filterKey = key as FilterKey;
              return (
                  <div className="flex items-center bg-beige text-red px-3 py-1 " key={key}>
                    <span className="mr-2">{`${filterKey}: ${value}`}</span>
                    <button
                        className="bg-red-500 hover:bg-red-700 text-white rounded-full h-6 w-6 flex items-center justify-center"
                        onClick={() => removeFilter(filterKey)}
                    >
                      x
                    </button>
                  </div>)})}
            </div>
        <div className="ml-auto flex items-center">
          <img src={sortIcon} className="h-6 w-6" />
          <div className="relative">
            <button
                className= ""
                onClick={handleDropdownToggle}>
              <span className="text-espresso text-2xl font-proxima"> Sort by: </span>
              <div className="text-red inline text-xl">{selectedFeature ? "Price: Low to High" : "Price: High to Low"} </div>
              <img src={downIcon} className="w-6 h-6 inline ml-" />
            </button>
            {dropdownOpen && (
                <div className="absolute mt-2 right-0  rounded-md bg-beige text-red text-sm shadow-lg z-10">
                  <ul>
                    <li
                        onClick={() => handleFeatureSelect(true)}
                        className="cursor-pointer px-4 py-2 hover:bg-gray-300"
                    >
                      Price: Low to High
                    </li>
                    <li
                        onClick={() => handleFeatureSelect(false)}
                        className="cursor-pointer px-4 py-2 hover:bg-gray-300"
                    >
                      Price: High to Low
                    </li>
                  </ul>
                </div>
            )}
          </div>
          </div>
        </div>
      </div>
  );
};

export default GiftSortNavbar;
