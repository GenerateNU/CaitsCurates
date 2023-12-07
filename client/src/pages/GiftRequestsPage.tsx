import { useState } from "react";
import GiftRequestCard from "../components/Profile/GiftRequestCard";
import GiftRequestsTable from "../components/Profile/GiftRequestsTable";
import { GiftRequest } from "../types";

const GiftRequestsPage = () => {
  const [selectedRequest, setSelectedRequest] = useState<GiftRequest | null>(
    null
  );

  return (
    <div className=" bg-eggshell min-h-screen">
        <div className="flex flex-row h-screen">
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
    </div>
  );
};

export default GiftRequestsPage;
