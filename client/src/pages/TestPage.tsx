import  { useEffect } from 'react';
import { useAdmin } from '../Context/AdminContext'; // adjust the import based on your file structure

export default function TestPage() {  const { collections, fetchGiftCollections } = useAdmin();

    useEffect(() => {
        fetchGiftCollections();
    }, []);
    console.log(collections)
    return (
        <div>
            {collections.map((collection, index) => (
                <div key={index}>
                    <h3>{collection.CollectionName}</h3>
                </div>
            ))}
        </div>
    );
};

