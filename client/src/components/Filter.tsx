import React from "react";
import FilterSection from "./FilterSection";

const recipients = ["For Her", "For Him", "For Mom", "For Dad"];
const interests = ["Sports", "Cooking", "Games"];
const prices = ["Under $25", "Under $50", "Under $100", "Under $500"];
const occasions = ["Birthday", "Anniversary", "Graduation"];

type FilterProps = {
  isOpen: boolean;
  filterToggle: () => void;
};

const Filter: React.FC<FilterProps> = ({ isOpen, filterToggle }) => {
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
              <div onClick={filterToggle}>x</div>
            </div>

            {/* Filter sections */}
            <FilterSection title="By Occasion" items={occasions} />
            <FilterSection title="By Recipient" items={recipients} />
            <FilterSection title="By Interest" items={interests} />
            <FilterSection title="By Price" items={prices} />
            <FilterSection title="Cait's Picks" items={[]} />
          </div>
        </>
      )}
    </div>
  );
};

export default Filter;
