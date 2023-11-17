import { useState } from "react";

interface SearchBarProps {
  updateHomePage: (e: string) => void;
}
const SearchBar: React.FC<SearchBarProps> = ({ updateHomePage }) => {
  const [searchText, setSearchText] = useState("");

  const handleInputChange = (e: any) => {
    setSearchText(e.target.value);
    updateHomePage(e.target.value);
  };

  return (
    <div className="bg-gray-100 p-4 w-full">
      <div className="flex items-center relative">
        <input
          type="text"
          className="w-full px-4 py-2 rounded-lg border border-gray-300 pl-10"
          placeholder="Search"
          value={searchText}
          onChange={handleInputChange}
          style={{ color: "black" }}
        />
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          fill="black"
          className="w-4 h-4 absolute left-3 top-3/5"
          style={{ pointerEvents: "none" }}
        >
          <path
            fillRule="evenodd"
            d="M10.5 3.75a6.75 6.75 0 100 13.5 6.75 6.75 0 000-13.5zM2.25 10.5a8.25 8.25 0 1114.59 5.28l4.69 4.69a.75.75 0 11-1.06 1.06l-4.69-4.69A8.25 8.25 0 012.25 10.5z"
            clipRule="evenodd"
          />
        </svg>
      </div>
    </div>
  );
};

export default SearchBar;
