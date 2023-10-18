import  { useEffect } from 'react';
import { useAdmin } from '../Context/AdminContext'; // adjust the import based on your file structure

export default function TestPage() {  const { gifts, fetchGifts } = useAdmin();

    useEffect(() => {
        fetchGifts();
    }, []);

    return (
        <div>
            {gifts.map((gift, index) => (
                <div key={index}>
                    <h3>{gift.Name}</h3>
                    <p>Price: ${gift.Price}</p>
                    {/* Render other gift properties as needed */}
                </div>
            ))}
        </div>
    );
};

