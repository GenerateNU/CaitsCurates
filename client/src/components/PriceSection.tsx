import React, { useState } from "react";

const priceRanges = [25, 50, 100, 500];

type PriceSectionProps = {
  title: string;
  updateFilters: (min: number, max: number) => void;
};

const PriceSection: React.FC<PriceSectionProps> = ({
  title,
  updateFilters,
}) => {
  const [sectionOpen, setSectionOpen] = useState(false);
  const [minPrice, setMinPrice] = useState(0);
  const [maxPrice, setMaxPrice] = useState(1000);

  const toggle = () => {
    setSectionOpen(!sectionOpen);
  };

  const updatePrices = (min: number, max: number) => {
    setMinPrice(min);
    setMaxPrice(max);
    updateFilters(min, max);
  };

  return (
    <>
      <hr className="h-px my-2 bg-gray-500 border-0" />
      <div className="flex justify-between cursor-pointer" onClick={toggle}>
        <h2 className="text-lg font-bold my-2">{title}</h2>
        <div className="my-2" onClick={toggle}>
          {sectionOpen ? "-" : "+"}
        </div>
      </div>
      {sectionOpen && (
        <div>
          {priceRanges.map((price) => (
            <h2
              key={price}
              onClick={() => updatePrices(0, price)}
              className={`mb-2 cursor-pointer ${
                price === maxPrice ? "font-bold" : ""
              }`}
            >
              Under ${price}
            </h2>
          ))}

          <div className="flex">
            <div className="w-32">
              <label className="w-24">
                Min price:
                <input
                  className="w-full border"
                  name="minPrice"
                  type="number"
                  value={minPrice}
                  onChange={(e) =>
                    updatePrices(parseFloat(e.target.value), maxPrice)
                  }
                />
              </label>
            </div>
            <div className="w-32">
              <label className="w-24">
                Max price:
                <input
                  className="w-full border"
                  name="maxPrice"
                  type="number"
                  value={maxPrice}
                  onChange={(e) =>
                    updatePrices(minPrice, parseFloat(e.target.value))
                  }
                />
              </label>
            </div>
          </div>
        </div>
      )}
    </>
  );
};

export default PriceSection;
