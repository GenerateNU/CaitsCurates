import { Gift } from "../../types.tsx";

type CollectionItemProps = {
  name: string;
  gifts: Gift[];
};

function CollectionItem({ name }: CollectionItemProps) {
  const circleStyle = {
    width: "100px", 
    height: "100px",
    borderRadius: "50%",
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    justifyContent: "center",
    margin: "10px",
  };



  return (
    <div>
      <div className="bg-beige" style={circleStyle}>
      </div>
      <h2 className="text-sm text-black text-center" style={{marginBottom: "15px"}}>{name}</h2>
    </div>
  );
}

export default CollectionItem;

