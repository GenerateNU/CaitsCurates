import { Gift } from "../../types.tsx";

type GiftItemProps = {
  gift: Gift;
  isSaved: boolean;
  onFavoriteClick: (gift: Gift, isSaved: boolean) => void;
};

function UpdatedGiftItem({ gift, isSaved, onFavoriteClick }: GiftItemProps) {
  return (
    <div className="w-[15rem] flex-col bg-eggshell flex-start">
        <div className="bg-gray-100 h-[15rem] w-[15rem] mx-auto mb-2 relative">
            {gift.ImageLink ? (
                <img src={gift.ImageLink} alt={gift.Name} className="w-full h-full object-cover" />
            ) : (
                <div className="w-full h-full flex items-center justify-center">
                    {/* You can add content or text here if needed */}
                </div>
            )}
        </div>
      <div className="relative">
        <div className="text-base font-proxima text-espresso">{gift.Occasion}</div>
        <div className=" w-[90%] text-2xl font-seasons text-espresso font-bold">{gift.Name}</div>
        <div className="text-xs text-espresso font-proxima">${gift.Price}</div>
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill={isSaved ? "red" : "none"}
          viewBox="0 0 24 24"
          strokeWidth="1.5"
          stroke="black"
          className="w-6 h-6 absolute top-0 right-0"
          onClick={() => onFavoriteClick(gift, isSaved)}
        >
            {gift.ImageLink && (
                <img src={gift.ImageLink} alt={gift.Name} style={{ width: '100px', height: '100px' }} />
            )}
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12z"
          />
        </svg>
      </div>
    </div>
  );
}

export default UpdatedGiftItem;
