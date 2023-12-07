import { collectionIconArray } from "../../data/collectionIcons.ts";

type CollectionItemProps = {
  name: string;
  collectionIndex: number;
  selected: boolean;
};

function CollectionItem({ name, collectionIndex, selected }: CollectionItemProps) {
  const icon = collectionIconArray[collectionIndex % collectionIconArray.length]

  return (
    <div className="flex flex-col items-center">
      <img src={icon} className="h-40 w-40" />
      <div className={`text-2xl font-seasons font-bold mt-3 ${selected ? "text-red" : "text-coffee"}`}>{name}</div>
    </div>
  );
}

export default CollectionItem;

