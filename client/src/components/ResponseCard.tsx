import React from "react";
import { GiftResponse } from "../types";

const ResponseCard: React.FC<GiftResponse> = ({
                                                  GiftCollection,
                                                  CustomMessage,
                                              }: GiftResponse) => {

    return (
        <div className="flex flex-col bg-slate-100 px-4 py-2">
            <h2 className="font-bold text-md">Response:</h2>
            <p>Message: {CustomMessage}</p>
            <p>Collection Sent: {GiftCollection.CollectionName}</p>

        </div>
    );
};

export default ResponseCard;