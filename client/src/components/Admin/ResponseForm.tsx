import CollectionSelector from "./CollectionSelector.tsx";
import axios from "axios";
import React, {useState} from "react";

interface ResponseFormProps {
    RequestID:  number; // adjust the type as necessary
}
const ResponseForm: React.FC<ResponseFormProps> = ({ RequestID }) => {
    const [selectedOption, setSelectedOption] = useState<Number | null>(null);
    const [customMessage, setCustomMessage] = useState("");

    const handleSubmit = async (event: { preventDefault: () => void; }) => {
        event.preventDefault();
        // @ts-ignore
        if (!selectedOption) {
            alert("Please select a gift collection.");
            return;
        }
        // @ts-ignore
        const payload  = {
            GiftCollectionId: selectedOption,
            CustomMessage: customMessage
        };

        try {
            const response = await axios.post("/api/addGiftResponse", payload);
            await axios.put("/api/requests", {ID: RequestID, giftResponseID : response.data.ID});
        } catch (error) {
            alert("An error occurred while saving the gift response.");
        }
    };

    return (
        <form onSubmit={handleSubmit} className="flex flex-col justify-between h-full mt-4">
            <div>
                <label className="block text-gray-700 text-sm font-bold mb-2">
                    Select a gift collection:
                    <CollectionSelector
                        selectedOption={selectedOption}
                        handleOptionChange={setSelectedOption}
                    />
                </label>
                <label className="block text-gray-700 text-sm font-bold mb-2">
                    Custom message:
                    <input
                        className="border rounded w-full py-2 px-3 text-gray-700"
                        name="message"
                        type="text"
                        value={customMessage}
                        onChange={(e) => setCustomMessage(e.target.value)}
                    />
                </label>
            </div>
            <button type="submit" className="bg-blue-600 ml-8 px-4 py-2 h-10 text-white rounded-md self-end">
                Submit
            </button>
        </form>
    );
}
export default ResponseForm;