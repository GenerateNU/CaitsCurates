import React, { useState } from "react";
import ResponseCard from "./ResponseCard.tsx";
import ResponseForm from "./ResponseForm.tsx";
import {Giftee, GiftRequest} from "../../types.tsx";

const RequestCard: React.FC<GiftRequest> = ({
                                                ID,
                                                BudgetMin,
                                                BudgetMax,
                                                GiftResponse,
                                                DateNeeded,
                                                Comment,
                                                Giftee
                                            }: GiftRequest) => {
    const [showForm, setShowForm] = useState(false);
    return (
        <div className="flex flex-col w-full">
            <h2 className="font-bold text-lg">
                {Giftee?.GifteeName} ({new Date(DateNeeded).toLocaleDateString()})
            </h2>
            <div key={Giftee?.GifteeName} className="px-4 py-2 bg-slate-100">
                <p>Recipient: {Giftee?.GifteeName}</p>
                {!GiftResponse && (
                    <div>
                        <p>Age: {Giftee?.Age}</p>
                        <p>Interests: {Giftee?.Interests.join(", ")}</p>
                        <p>Colors: {Giftee?.Colors.join(", ")}</p>
                        <p>Relationship: {Giftee?.CustomerRelationship}</p>
                        <p>Gender: {Giftee?.Gender}</p>
                        <p>
                            Budget: ${BudgetMin} - ${BudgetMax}
                        </p>
                        <p> Comment: {Comment}</p>
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