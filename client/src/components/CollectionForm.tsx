import  { useState, ChangeEvent, FormEvent } from 'react';
import {Gift, GiftCollection} from "../types.tsx";





type EditFormProps = {
  collection: GiftCollection;
  allGifts: Gift[];
  onSave: (collection: GiftCollection) => void;
  onClose: () => void;
};


function CollectionForm({ collection, allGifts, onSave, onClose }: EditFormProps) {
  const [editedName, setEditedName] = useState(collection?.CollectionName || "");
  const [editedGifts, setEditedGifts] = useState(collection?.Gifts || []);

  const handleNameChange = (e: ChangeEvent<HTMLInputElement>) => {
    setEditedName(e.target.value);
  };

  const handleGiftsChange = (e: ChangeEvent<HTMLInputElement>) => {
    const gift: Gift = JSON.parse(e.target.value);

    if (e.target.checked) {
      // If the checkbox is checked, add the gift to the list
      setEditedGifts(prevGifts => [...prevGifts, gift]);
    } else {
      // If unchecked, remove the gift from the list
      setEditedGifts(prevGifts => prevGifts.filter(g => g.Name !== gift.Name));
    }
  };

      const handleSubmit = async (e: FormEvent) => {
        e.preventDefault();

        const collectionData = {
          ID: collection?.ID,
          CollectionName: editedName,
          Gifts: editedGifts,
          Customer: collection?.Customer,
          CustomerID: collection?.CustomerID
        };

        try {
          onSave(collectionData)
        } catch (error) {
          console.error("An error occurred while saving the collection:", error);
        }
        onClose();
      };

      return (
          <div className="edit-form border border-black p-4 rounded-md text-center">
            <form onSubmit={handleSubmit}>
              <div className="mb-4">
                <label htmlFor="name" className="block mb-2">
                  Name:
                </label>
                <input
                    type="text"
                    id="name"
                    className="border border-blue-500 rounded-md w-64 p-2 mx-auto"
                    value={editedName}
                    onChange={handleNameChange}
                />
              </div>
              <div className="mb-4">
              <label className="block mb-2">Add/Remove Gifts:</label>
              {allGifts.map((gift, index) => (
                  <div key={index}>
                    <input
                        type="checkbox"
                        id={`gift_${index}`}
                        value={JSON.stringify(gift)}
                        checked={editedGifts.some(editedGift => editedGift.Name === gift.Name)}
                        onChange={handleGiftsChange}
                    />
                    <label htmlFor={`gift_${index}`}>{gift.Name}</label>
                  </div>
              ))}
            </div>
              <button
                  type="submit"
                  className="bg-blue-500 text-white p-2 rounded-md border border-black-500"
              >
                Save
              </button>
              <button
                  type="button"
                  className="bg-red-500 text-white p-2 rounded-md border border-black-500 ml-2"
                  onClick={onClose}
              >
                Cancel
              </button>
            </form>
          </div>
      );
    }

export default CollectionForm;
