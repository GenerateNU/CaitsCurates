import React, { useState } from 'react';
import axios from 'axios';
import {Gift} from "../../types.tsx";


const defaultGift: Gift = {
    Demographic: "", Description: "", Occasion: "", Category: [], GiftCollections: [], ID: 0, Link: "", Name: "", Price: 0
}
type Props = {
    initialGift?: Gift;
    mode: 'add' | 'edit';
    onGiftChange: () => void;
};

const GiftForm: React.FC<Props> = ({ initialGift = defaultGift, mode, onGiftChange }) => {
    const [gift, setGift] = useState<Gift>(initialGift);

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setGift(prevState => ({
            ...prevState,
            [name]: value,
        }));
    };

    // handles multiple selection input
    const handleSelectChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        const { name, options } = e.target;
        const values = [];

        for (let i = 0; i < options.length; i++) {
            if (options[i].selected) {
                values.push(options[i].value);
            }
        }

        setGift(prevState => ({
            ...prevState,
            [name]: values,
        }));
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            var sentGift = {
                ...gift,
                Price: parseFloat(gift.Price.toString())
            }
            if (mode === 'add') {
                await axios.post('/api/addGift', sentGift);
                console.log("sentGift", sentGift);
            } else if (mode === 'edit' && gift.ID) {
                await axios.put(`/api/gifts/${gift.ID}`, sentGift);
            }
            onGiftChange();
        } catch (error) {
            console.error("Error saving gift:", error);
        }
    };



    return (
        <form onSubmit={handleSubmit} className="p-6">
            <div className="flex justify-between items-center">
                <h3 className="text-lg leading-6 font-medium text-gray-900" id="modal-title">
                    {mode === 'add' ? 'Add Gift' : 'Edit Gift'}
                </h3>
                <button type="button" onClick={() => onGiftChange()}> {/* Adjusted the closing functionality */}
                    X
                </button>
            </div>

            <div className="mb-4">
                <label htmlFor="name" className="block text-sm font-medium text-gray-700">
                    Name:
                </label>
                <input
                    type="text"
                    id="name"
                    name="Name"
                    value={gift.Name}
                    onChange={handleInputChange}
                    className="mt-1 p-2 w-full border-2 border-gray-300 rounded-md"
                />
            </div>
            <div className="mb-4">
                <label htmlFor="description" className="block text-sm font-medium text-gray-700">
                    Description:
                </label>
                <input
                    type="text"
                    id="description"
                    name="Description"
                    value={gift.Description}
                    onChange={handleInputChange}
                    className="mt-1 p-2 w-full border-2 border-gray-300 rounded-md"
                />
            </div>
            <div className="mb-4">
                <label htmlFor="occasion" className="block text-sm font-medium text-gray-700">
                    Occasion:
                </label>
                <select
                    id="occasion"
                    name="Occasion"
                    value={gift.Occasion}
                    onChange={handleInputChange}
                    className="mt-1 p-2 w-full border-2 border-gray-300 rounded-md"
                >
                    <option value="Birthday">Birthday</option>
                    <option value="Bridal">Bridal</option>
                    <option value="Get well soon">Get well soon</option>
                    <option value="New baby">New baby</option>
                    <option value="Thinking of you">Thinking of you</option>
                    <option value="Thank you">Thank you</option>
                </select>
            </div>
            <div className="mb-4">
                <label htmlFor="demographic" className="block text-sm font-medium text-gray-700">
                    Demographic:
                </label>
                <select
                    id="demographic"
                    name="Demographic"
                    value={"" || gift.Demographic}
                    onChange={handleInputChange}
                    className="mt-1 p-2 w-full border-2 border-gray-300 rounded-md"
                >
                    <option value="For mom">For mom</option>
                    <option value="For dad">For dad</option>
                    <option value="For kids">For kids</option>
                    <option value="For partners">For partners</option>
                    <option value="For men">For men</option>
                    <option value="For women">For women</option>
                </select>
            </div>
            <div className="mb-4">
                <label htmlFor="category" className="block text-sm font-medium text-gray-700">
                    Category:
                </label>
                <select
                    id="category"
                    name="Category"
                    value={gift.Category}
                    onChange={handleSelectChange}
                    className="mt-1 p-2 w-full border-2 border-gray-300 rounded-md"
                    multiple
                >
                    <option value="Best selling">Best selling</option>
                    <option value="Fun">Fun</option>
                    <option value="Gadgets">Gadgets</option>
                    <option value="Home">Home</option>
                    <option value="Jewelry">Jewelry</option>
                    <option value="Kitchen & bar">Kitchen & bar</option>
                    <option value="Warm and cozy">Warm and cozy</option>
                    <option value="Outdoors">Outdoors</option>
                </select>
            </div>
            <div className="mb-4">
                <label htmlFor="price" className="block text-sm font-medium text-gray-700">
                    Price:
                </label>
                <input
                    type="number"
                    id="price"
                    name="Price"
                    value={gift.Price}
                    onChange={handleInputChange}
                    className="mt-1 p-2 w-full border-2 border-gray-300 rounded-md"
                />
            </div>
            <div className="mb-4">
                <label htmlFor="link" className="block text-sm font-medium text-gray-700">
                    Link:
                </label>
                <input
                    type="text"
                    id="link"
                    name="Link"
                    value={gift.Link}
                    onChange={handleInputChange}
                    className="mt-1 p-2 w-full border-2 border-gray-300 rounded-md"
                />
            </div>
            <button type="submit">{mode === 'add' ? 'Add' : 'Save'}</button>
        </form>
    );
};

export default GiftForm;
