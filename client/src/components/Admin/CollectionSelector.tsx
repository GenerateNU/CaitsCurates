import {useEffect} from "react";
import Select from "react-select";
import { useAdmin } from '../../Context/AdminContext.tsx';

interface CollectionSelectorProps {
    selectedOption: Number | null;
    handleOptionChange: (newValue: Number | null) => void;
}

const CollectionSelector: React.FC<CollectionSelectorProps> = ({
                                                                   selectedOption,
                                                                   handleOptionChange
                                                               }) => {
    const { collections, fetchGiftCollections } = useAdmin();

    useEffect(() => {
        fetchGiftCollections();
    }, []);


    console.log("Collections:", collections); // Log to see the fetched collections


    // Mapping collections to options
    const options = collections.map((collection) => ({
        value: collection.ID,
        label: collection.CollectionName,
    }));

    return (
        <Select
        value={options.find(option => option.value === selectedOption)}
        onChange={(option) => handleOptionChange(option ? option.value : null)}
        options={options}
       />
    );
}
export default CollectionSelector;


