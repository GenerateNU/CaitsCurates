import React, { useState } from "react";

type FilterSectionProps = {
  title: string;
  items: string[];
  additional?: any;
};

const FilterSection: React.FC<FilterSectionProps> = ({
  title,
  items,
  additional,
}) => {
  const [sectionOpen, setSectionOpen] = useState(false);

  const toggle = () => {
    setSectionOpen(!sectionOpen);
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
          <h2 key={item} className="mb-2">
            {item}
          </h2>
        ))}
      {sectionOpen && additional}
    </>
  );
};

export default FilterSection;
