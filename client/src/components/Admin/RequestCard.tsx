import React, { useState } from "react";
import ResponseCard from "./ResponseCard.tsx";
import ResponseForm from "./ResponseForm.tsx";
import { GiftRequest } from "../../types.tsx";

const RequestCard: React.FC<GiftRequest> = ({
                                                ID,
                                                RecipientName,
                                                RecipientAge,
                                                RecipientInterests,
                                                BudgetMin,
                                                BudgetMax,
                                                GiftResponse,
                                                DateNeeded,
                                            }: GiftRequest) => {
    const [showForm, setShowForm] = useState(false);
    return (
        <div className="flex flex-col w-full">
            <h2 className="font-bold text-lg">
                {RecipientName} ({new Date(DateNeeded).toLocaleDateString()})
            </h2>
            <div key={RecipientName} className="px-4 py-2 bg-slate-100">
                <p>Recipient: {RecipientName}</p>
                {!GiftResponse && (
                    <div>
                        <p>Recipient age: {RecipientAge}</p>
                        <p>Recipient interests: {RecipientInterests.join(", ")}</p>
                        <p>
                            Budget: ${BudgetMin} - ${BudgetMax}
                        </p>
                        <p>Needed by: ({new Date(DateNeeded).toLocaleDateString()})</p>
                    </div>
                )}
            </div>
            <div>
                {GiftResponse && <ResponseCard {...GiftResponse} />}
                {!GiftResponse && !showForm && (
                    <button
                        className="bg-blue-600 px-4 py-2  text-white rounded-md mt-4"
                        onClick={() => setShowForm(true)}
                    >
                        Add response
                    </button>
                )}
                {showForm && <ResponseForm RequestID={ID} />}
            </div>
        </div>
    );
};

export default RequestCard