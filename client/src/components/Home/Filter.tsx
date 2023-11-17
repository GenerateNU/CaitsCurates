import React, { useState } from "react";
import FilterSection from "./FilterSection";
import PriceSection from "./PriceSection";
import { Filters } from "../types";

const recipients = ["For Her", "For Him", "For Mom", "For Dad"];
const category = ["Sports", "Cooking", "Games"];
const prices = ["Under $25", "Under $50", "Under $100", "Under $500"];
const occasions = ["Birthday", "Anniversary", "Graduation"];

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

  return (
    <div>
      {isOpen && (
        <>
          {/* Dark overlay */}
          <div className="fixed top-0 left-0 w-full h-full bg-black opacity-50 z-10"></div>
          {/* Filter modal */}
          <div className="absolute top-0 right-0 mt-16 w-1/4 bg-white p-8 shadow-lg z-20 text-black overflow-y-auto max-h-screen">
            <div className="flex justify-between">
              <h2 className="mb-4 text-lg font-bold">Filter</h2>
              <div className="cursor-pointer" onClick={filterToggle}>
                x
              </div>
            </div>

            {/* Filter sections */}
            <FilterSection
              title="By Occasion"
              items={occasions}
              updateFilters={updateOccasion}
            />
            <FilterSection
              title="By Recipient"
              items={recipients}
              updateFilters={updateDemographic}
            />
            <FilterSection
              title="By Category"
              items={category}
              updateFilters={updateCategory}
            />
            <PriceSection title="By Price" updateFilters={updatePrices} />
            <button className="bg-gray-400 mt-6 px-4 py-2 h-10 text-white rounded-md float-right">
              Apply filters
            </button>
          </div>
        </>
      )}
    </div>
  );
};

export default Filter;
