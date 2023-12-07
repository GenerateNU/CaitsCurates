import GiftItem from '../components/Admin/GiftItem.tsx'
import Navbar from "../components/Admin/Navbar.tsx";
import {useAdmin} from "../Context/AdminContext.tsx";
import {useEffect, useState} from "react";
import GiftForm from "../components/Admin/GiftForm.tsx";
import {Gift} from "../types.tsx";
import axios from "axios";




const GiftManagementPage = () => {
    const { gifts, fetchGifts } = useAdmin();
    const [editingGift, setEditingGift] = useState<Gift | null>(null);
    const [showForm, setShowForm] = useState(false);

    useEffect(() => {
        fetchGifts();
    }, []);

    const handleGiftChange = () => {
        setShowForm(false);
        fetchGifts();
    };

    const handleEditingGift = (gift: Gift | null) => {
        setEditingGift(gift);
        setShowForm(true);
    };
    const handleDelete = async (gift: Gift) => {
        if (gift.ID) {
            try {
                await axios.delete(`/api/gifts/${gift.ID}`);
            } catch (error) {
                console.error("Error deleting gift:", error);
            }
            fetchGifts();
        }
    };


    return (
        <div>
            <Navbar />
            <div className="min-h-screen px-48 py-8">
                <h2 className="font-bold text-2xl mb-2">Manage Gifts</h2>
                <p>View, create, and update gifts here. </p>
                <div className="my-4">
            <button
                className="bg-blue-500 text-white p-2 rounded"
                onClick={() => handleEditingGift(null)}
            >
                + Add Gift
            </button>
                </div>
                <div className="grid grid-cols-4 gap-x-10">
                {gifts.map((gift) => (
                    <div key={gift.ID} className="m-2">
                        <GiftItem
                            gift={gift}
                            onEditClick={() => handleEditingGift(gift)}
                            onDeleteClick={() => handleDelete(gift)}
                        />
                    </div>
            ))}</div>

            {showForm &&
                <div className="fixed z-10 inset-0 overflow-y-auto" aria-labelledby="modal-title" role="dialog" aria-modal="true">
                    <div className="flex items-center justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
                        <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" aria-hidden="true"></div>
                        <span className="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">&#8203;</span>
                        <div className="inline-block align-middle bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full">
                            <GiftForm
                                {...(editingGift ? { initialGift: editingGift } : {})}
                                mode={editingGift ? 'edit' : 'add'}
                                onGiftChange={handleGiftChange}
                            />
                        </div>
                    </div>
                </div>
            }
        </div>
        </div>
    );
};

export default GiftManagementPage;
