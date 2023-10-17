import RequestCard from "../components/RequestCard";
import { completeRequests, incompleteRequests } from "./mockData";
import { useState } from "react";
import Select from "react-select";

export type SelectValueType = {
    [key: string]: string;
} | null;

const options = [
    { value: "0", label: "Incomplete requests" },
    { value: "1", label: "Complete requests" },
    { value: "2", label: "All requests" },
];

export default function RequestsPage() {
    const [selectedOption, setSelectedOption] = useState<SelectValueType>({
        value: "2",
        label: "All requests",
    });

    const handleOption = (selection: SelectValueType) => {
        setSelectedOption(selection);
    };
    return (
        <div className="flex flex-col px-96 py-8">
            <h2 className="font-bold text-xl">View gift requests</h2>
            <p>Filter gift requests using the dropdown below. </p>
            <Select
                defaultValue={selectedOption}
                onChange={handleOption}
                options={options}
            />
            {(selectedOption?.value == "0" || selectedOption?.value == "2") && (
                <div className="mt-6">
                    <h2 className="font-bold text-xl text-blue-800">
                        Incomplete requests
                    </h2>
                    {incompleteRequests.map((req) => {
                        return <RequestCard  {...req} />;
                    })}
                </div>
            )}
            {(selectedOption?.value == "1" || selectedOption?.value == "2") && (
                <div className="mt-6">
                    <h2 className="font-bold text-xl text-blue-800">Complete requests</h2>
                    {completeRequests.map((req) => {
                        return <RequestCard {...req} />;
                    })}
                </div>
            )}
        </div>
    );
}