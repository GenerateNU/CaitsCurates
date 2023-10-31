import { Gift } from "../types.tsx";

type CollectionItemProps = {
  name: string;
  gifts: Gift[];
};

function CollectionItem({ name }: CollectionItemProps) {
  const circleStyle = {
    width: "100px", 
    height: "100px",
    backgroundColor: "lightgrey",
    borderRadius: "50%",
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    justifyContent: "center",
    margin: "10px",
  };



  return (
    <div>
      <div style={circleStyle}>
      </div>
      <h2 className="text-sm text-black text-center">{name}</h2>
    </div>
  );
}

export default CollectionItem;

