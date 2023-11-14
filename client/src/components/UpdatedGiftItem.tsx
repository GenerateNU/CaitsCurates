type GiftItemProps = {
  name: string;
  price: number;
};

function UpdatedGiftItem({ name, price }: GiftItemProps) {
  return (
    <div className="relative flex flex-col bg-gray-100 flex-start">
      <div className="bg-gray-200 w-40 h-40 mx-auto mb-2 relative">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth="1.5"
          stroke="black"
          className="w-6 h-6 absolute bottom-2 right-2"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12z"
          />
        </svg>
      </div>
      <div className="">
        <h2 className="text-sm text-black font-bold">{name}</h2>
        <h2 className="text-xs text-black">${price}</h2>
      </div>
    </div>
  );
}

export default UpdatedGiftItem;
