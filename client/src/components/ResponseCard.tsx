import React from "react";
import Gift from "./Gift";
import { GiftResponse } from "../types";

const ResponseCard: React.FC<GiftResponse> = ({
                                                  GiftCollection,
                                                  CustomMessage,
                                              }: GiftResponse) => {
    return (
        <div className="flex flex-col bg-slate-100 px-4 py-2">
            <h2 className="font-bold text-md">Response:</h2>
            <p>{CustomMessage}</p>
            <div>
                {GiftCollection.Gifts.map((gift) => {
                    return (
                        <div className="mt-2">
                            <Gift {...gift} />
                        </div>
                    );
                })}
            </div>
        </div>
    );
};

export default ResponseCard;