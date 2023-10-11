type CollectionItemProps = {
    name: string;
    gifts: string[];
  };
  
  function CollectionItem({ name, gifts }: CollectionItemProps) {
    return (
      <div className="collection-item p-4 border border-black text-center w-40" style={{ height: "245px" }}>
        <h2 className="text-2xl font-bold">{name}</h2>
        <ul className="max-h-40 overflow-y-auto">
          {gifts.map((gift, index) => (
            <li key={index}>{gift}</li>
          ))}
        </ul>
      </div>
    );
  }
  
  export default CollectionItem;
  



