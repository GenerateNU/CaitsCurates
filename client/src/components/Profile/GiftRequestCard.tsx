import Navbar from "../Home/NavBarUpdated";

const GiftRequestCard = () => {
  return (
    <div>
      <Navbar />
      <div>
        <h1>Gift Suggestions</h1>
        {/* table */}
      </div>
      <div className="flex flex-row justify-between">
        <div>
          <h1>Gift Request Details</h1>
          <div>
            <p>Type of Occasion:</p>
            <p>Price Range:</p>
            <p>Date of Occasion:</p>
            <p>Deadline:</p>
            <p>Comments:</p>
          </div>
        </div>
        <div>
          <h1>Giftee Details</h1>
          <div>
            <p>Name:</p>
            <p>Gender Identity:</p>
            <p>Relation to Giftee:</p>
            <p>Age:</p>
            <p>Favorite Color(s):</p>
            <p>Interests:</p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default GiftRequestCard;
