import { useState } from "react";
import Select from "react-select";

const options = [
    { value: "Taylor's gifts", label: "Taylor's gifts" },
    { value: "Harry's gifts", label: "Harry's gifts" },
];

export type SelectValueType = {
    [key: string]: string;
} | null;

export default function CollectionSelector() {
    const [selectedOption, setSelectedOption] = useState<SelectValueType>(null);

    const handleOption = (selection: SelectValueType) => {
        setSelectedOption(selection);
    };

    return (
        <div className="App">
            <Select
                defaultValue={selectedOption}
                onChange={handleOption}
                options={options}
            />
        </div>
    );
}