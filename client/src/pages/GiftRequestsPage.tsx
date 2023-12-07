import { useState } from "react";
import GiftRequestCard from "../components/Profile/GiftRequestCard";
import GiftRequestsTable from "../components/Profile/GiftRequestsTable";
import { GiftRequest } from "../types";
import Navbar from "../components/Home/NavBarUpdated";

const GiftRequestsPage = () => {
  const [selectedRequest, setSelectedRequest] = useState<GiftRequest | null>(
    null
  );

  return (
    <div className=" bg-eggshell min-h-screen">
      <Navbar />

      <div className="mx-12">
        {selectedRequest ? (
          <GiftRequestCard
            selectedRequest={selectedRequest}
            clearSelected={setSelectedRequest}
          />
        ) : (
          <GiftRequestsTable selectRow={setSelectedRequest} />
        )}
      </div>
    </div>
  );
};

export default GiftRequestsPage;
