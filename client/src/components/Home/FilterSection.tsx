import React, { useState } from "react";

type FilterSectionProps = {
  title: string;
  items: string[];
  updateFilters: (e: string) => void;
};

const FilterSection: React.FC<FilterSectionProps> = ({
  title,
  items,
  updateFilters,
}) => {
  const [sectionOpen, setSectionOpen] = useState(false);
  const [selectedItem, setSelectedItem] = useState("");

  const toggle = () => {
    setSectionOpen(!sectionOpen);
  };

  const handleClick = (item: any) => {
    if (item !== selectedItem) {
      setSelectedItem(item);
      updateFilters(item);
    } else {
      setSelectedItem("");
      updateFilters("");
    }
  };
  return (
    <>
      <hr className="h-px my-2 bg-gray-500 border-0" />
      <div className="flex justify-between cursor-pointer" onClick={toggle}>
        <h2 className="text-lg font-bold my-2">{title}</h2>
        <div className="my-2" onClick={toggle}>
          {items.length > 0 ? (sectionOpen ? "-" : "+") : null}
        </div>
      </div>
      {sectionOpen &&
        items.map((item) => (
          <h2
            key={item}
            onClick={() => handleClick(item)}
            className={`mb-2 cursor-pointer ${
              item === selectedItem ? "font-bold" : ""
            }`}
          >
            {item}
          </h2>
        ))}
    </>
  );
};

export default FilterSection;
