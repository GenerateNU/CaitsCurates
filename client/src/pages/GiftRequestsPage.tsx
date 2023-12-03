import { useState } from "react";
import GiftRequestCard from "../components/Profile/GiftRequestCard";
import GiftRequests from "../components/Profile/GiftRequestsTable";
import { GiftRequest } from "../types";

const GiftRequestsPage = () => {
  const [selectedRequest, setSelectedRequest] = useState<GiftRequest | null>(
    null
  );

  {
    selectedRequest ? <GiftRequestCard /> : <GiftRequests />;
  }

  return <div></div>;
};

export default GiftRequestsPage;
