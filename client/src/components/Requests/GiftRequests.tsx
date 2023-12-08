import { useState } from "react";
import GiftRequestCard from "../Profile/GiftRequestCard.tsx";
import GiftRequestsTable from "../Profile/GiftRequestsTable.tsx";
import { GiftRequest } from "../../types.tsx";

const GiftRequests = () => {
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

export default GiftRequests;
