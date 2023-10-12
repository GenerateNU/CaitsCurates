import React, { useState, ChangeEvent, FormEvent } from 'react';

type Collection = {
  id: number;
  name: string;
  gifts: string[];
};

type EditFormProps = {
  collection: Collection;
  onSave: (collection: Collection) => void;
  onClose: () => void;
};

const predefinedGifts = ["Gift 1", "Gift 2", "Gift 3", "Gift 4", "Gift 5"];

function CollectionForm({ collection, onSave, onClose }: EditFormProps) {
  const [editedName, setEditedName] = useState(collection.name);
  const [editedGifts, setEditedGifts] = useState(collection.gifts);

  const handleNameChange = (e: ChangeEvent<HTMLInputElement>) => {
    setEditedName(e.target.value);
  };

  const handleGiftsChange = (e: ChangeEvent<HTMLSelectElement>) => {
    const selectedOptions = Array.from(e.target.options);
    const selectedGifts = selectedOptions
      .filter((option) => option.selected)
      .map((option) => option.value);

    // Here, we concatenate the selected gifts with the existing gifts
    setEditedGifts([...editedGifts, ...selectedGifts]);
  };

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    onSave({
      id: collection.id,
      name: editedName,
      gifts: editedGifts,
    });
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
          <label htmlFor="gifts" className="block mb-2">
            Select a Gift to Add:
          </label>
          <select
            id="gifts"
            className="border border-blue-500 rounded-md w-64 p-2 mx-auto"
            multiple
            value={editedGifts}
            onChange={handleGiftsChange}
          >
            {predefinedGifts.map((gift) => (
              <option key={gift} value={gift}>
                {gift}
              </option>
            ))}
          </select>
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
