//import Gift from "./Gift.tsx";
import {Gift} from "../types.tsx";

type CollectionItemProps = {
  name: string;
  gifts: Gift[];
};

function CollectionItem({ name, gifts }: CollectionItemProps) {
  return (
    <div className="collection-item p-4 border border-black text-center w-auto h-48">
      <h2 className="text-2xl font-bold">{name}</h2>
      <ul className="max-h-40 overflow-y-auto">
        {gifts.map((gift, index) => (
          <li key={index}>{gift.Name}</li>
        ))}
      </ul>
    </div>
  );
}

export default CollectionItem;
