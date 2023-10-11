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

function CollectionForm({ collection, onSave, onClose }: EditFormProps) {
  const [editedName, setEditedName] = useState(collection.name);
  const [editedGifts, setEditedGifts] = useState(collection.gifts.join(', '));

  const handleNameChange = (e: ChangeEvent<HTMLInputElement>) => {
    setEditedName(e.target.value);
  };

  const handleGiftsChange = (e: ChangeEvent<HTMLInputElement>) => {
    setEditedGifts(e.target.value);
  };

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    onSave({
      id: collection.id,
      name: editedName,
      gifts: editedGifts.split(',').map((gift) => gift.trim()),
    });
    onClose();
  };


  return (
    <div className="edit-form border border-black p-4 rounded-md text-center">
      <form onSubmit={handleSubmit}>
        <div className="mb-4">
          <label htmlFor="name" className="block mb-2">Name:</label>
          <input
            type="text"
            id="name"
            className="border border-blue-500 rounded-md w-64 p-2 mx-auto"
            value={editedName}
            onChange={handleNameChange}
          />
        </div>
        <div className="mb-4">
          <label htmlFor="gifts" className="block mb-2">Gifts:</label>
          <input
            type="text"
            id="gifts"
            className="border border-blue-500 rounded-md w-64 p-2 mx-auto"
            value={editedGifts}
            onChange={handleGiftsChange}
          />
        </div>
        <button type="submit" className="bg-blue-500 text-white p-2 rounded-md border border-black-500">
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
