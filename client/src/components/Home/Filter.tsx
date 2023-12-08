import React, {useState} from "react";
import {Filters} from "../../types.tsx";
import FilterSection from "./FilterSection.tsx";
import PriceSection from "../PriceSection.tsx";


const recipients = [
  "For mom",
  "For dad",
  "For partners",
  "For kids",
  "For women",
  "For men",
];
const category = [
  "Best selling",
  "Fun",
  "Gadgets",
  "Home",
  "Jewelry",
  "Kitchen & bar",
  "Warm and cozy",
  "Outdoors",
];
const occasions = [
  "Birthday",
  "Bridal",
  "Get well soon",
  "New baby",
  "Thinking of you",
  "Thank you",
];

type FilterProps = {
  isOpen: boolean;
  filterToggle: () => void;
  currentFilters: Filters;
  setCurrentFilters: React.Dispatch<React.SetStateAction<Filters>>;
};

const Filter: React.FC<FilterProps> = ({
  isOpen,
  filterToggle,
    currentFilters,
  setCurrentFilters,
}) => {
  const [resetFilter, setResetFilter] = useState(false);
  const updateOccasion = (item: string) => {
    setCurrentFilters((prevFilters: Filters) => ({
      ...prevFilters,
      occasion: item,
    }));
  };

  const updateCategory = (item: string) => {
    setCurrentFilters((prevFilters: Filters) => ({
      ...prevFilters,
      category: item,
    }));
  };

  const updateDemographic = (item: string) => {
    setCurrentFilters((prevFilters) => ({
      ...prevFilters,
      demographic: item,
    }));
  };

  const updatePrices = (min: number, max: number) => {
    setCurrentFilters((prevFilters) => ({
      ...prevFilters,
      minPrice: min,
      maxPrice: max,
    }));
  };
  const clearFilters = () => {
    setCurrentFilters(() => ({
      demographic: "",
      minPrice: 0,
      maxPrice: 1000,
      category: "",
      occasion: "",
    }));
    setResetFilter(true);
    setTimeout(() => setResetFilter(false), 0);
  };

  return (
      <div>
      {isOpen && (
        <>
          {/* Dark overlay */}
          <div className="fixed top-0 left-0 w-full h-full bg-black opacity-50 z-10"></div>
          {/* Filter modal */}
          <div className="fixed top-0 right-0 w-1/4 bg-eggshell p-8 shadow-lg z-20 text-black overflow-y-auto max-h-screen">
            <div className="flex justify-between">
              <h2 className="mb-4 text-2xl text-espresso font-bold font-seasons">Filter</h2>
              <div className="cursor-pointer" onClick={filterToggle}>
                x
              </div>
            </div>
            {/* Filter sections */}
            <FilterSection
              title="By Occasion"
              items={occasions}
              reset={resetFilter}
              currentFilters={currentFilters}
              updateFilters={updateOccasion}
            />
            <FilterSection
              title="By Recipient"
              items={recipients}
              reset={resetFilter}
              currentFilters={currentFilters}
              updateFilters={updateDemographic}
            />
            <FilterSection
              title="By Category"
              items={category}
              reset={resetFilter}
              currentFilters={currentFilters}
              updateFilters={updateCategory}
            />
            <PriceSection title="By Price"
                          reset={resetFilter}
                          currentFilters = {currentFilters}
                          updateFilters={updatePrices} />
            <button className="mt-4 bg-beige text-winered font-bold py-2 px-4"
            onClick={clearFilters}>
              Clear Filters
            </button>
          </div>
        </>
      )}
    </div>
  );
};

export default Filter;
