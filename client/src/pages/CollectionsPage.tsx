import { useState } from "react";
import CollectionItem from "../components/CollectionItem";
import EditForm from "../components/CollectionForm";
import Navbar from "../components/Navbar";

type Gift = {
  name: string;
  description: string;
  price: number;
};

type Collection = {
  id: number;
  name: string;
  gifts: Gift[];
};

const predefinedGifts: Gift[] = [
  {
    name: "Gift 1",
    description: "Description of Gift 1",
    price: 10,
  },
  {
    name: "Gift 2",
    description: "Description of Gift 2",
    price: 20,
  },
  {
    name: "Gift 3",
    description: "Description of Gift 3",
    price: 30,
  },
  {
    name: "Gift 4",
    description: "Description of Gift 4",
    price: 40,
  },
  {
    name: "Gift 5",
    description: "Description of Gift 5",
    price: 50,
  },
];

const predefinedGifts2: Gift[] = [
  {
    name: "Gift 10",
    description: "Description of Gift 1",
    price: 10,
  },
  {
    name: "Gift 11",
    description: "Description of Gift 2",
    price: 20,
  },
];

const CollectionsPage = () => {
  const [collections, setCollections] = useState([
    {
      id: 1,
      name: "Birthday Gifts",
      gifts: [],
    },
    {
      id: 2,
      name: "Christmas Gifts",
      gifts: predefinedGifts,
    },
    {
      id: 3,
      name: "Birthday Gifts",
      gifts: predefinedGifts2,
    },
    {
      id: 4,
      name: "Christmas Gifts",
      gifts: predefinedGifts,
    },
    {
      id: 5,
      name: "Birthday Gifts",
      gifts: predefinedGifts,
    },
    {
      id: 6,
      name: "Christmas Gifts",
      gifts: predefinedGifts,
    },
    {
      id: 7,
      name: "Birthday Gifts",
      gifts: predefinedGifts,
    },
    {
      id: 8,
      name: "Christmas Gifts",
      gifts: predefinedGifts,
    },
    {
      id: 9,
      name: "Birthday Gifts",
      gifts: predefinedGifts,
    },
    {
      id: 10,
      name: "Christmas Gifts",
      gifts: predefinedGifts,
    },
    {
      id: 11,
      name: "Birthday Gifts",
      gifts: predefinedGifts,
    },
    {
      id: 12,
      name: "Christmas Gifts",
      gifts: predefinedGifts,
    },
    {
      id: 13,
      name: "Birthday Gifts",
      gifts: predefinedGifts,
    },
    {
      id: 14,
      name: "Christmas Gifts",
      gifts: predefinedGifts,
    },
  ]);

  const handleCreateCollection = () => {
    const newCollection: Collection = {
      id: Date.now(),
      name: "New Collection",
      gifts: [],
    };

    // Add the new collection to the state
    setCollections((prevCollections) => [...prevCollections, newCollection]);
    setEditCollectionId(newCollection.id);
    setShowEditForm(true);
  };

  const [showEditForm, setShowEditForm] = useState<boolean>(false);
  const [editCollectionId, setEditCollectionId] = useState<number | null>(null);

  const handleEditCollection = (collectionId: number) => {
    setEditCollectionId(collectionId);
    setShowEditForm(true);
  };

  const handleSaveCollection = (updatedCollection: Collection) => {
    setCollections((prevCollections) =>
      prevCollections.map((collection) =>
        collection.id === updatedCollection.id ? updatedCollection : collection
      )
    );
    setShowEditForm(false);
  };

  const handleCloseEditForm = () => {
    setEditCollectionId(null);
    setShowEditForm(false);
  };

  const handleDeleteCollection = (collectionId: number) => {
    setCollections((prevCollections) =>
      prevCollections.filter((collection) => collection.id !== collectionId)
    );
  };

  return (
    <div>
      <Navbar />
      <div className="min-h-screen px-48 py-8">
        <h2 className="font-bold text-2xl mb-2">Manage Collections</h2>
        <p>View, create, and update collections here. </p>
        <div className="my-4">
          <button
            onClick={handleCreateCollection}
            className="bg-blue-500 text-white p-2 rounded"
          >
            Create new collection
          </button>
        </div>

        {/* Collection grid */}
        <div className="grid grid-cols-3 gap-x-24">
          {collections.map((collection) => (
            <div key={collection.id} className="m-2">
              <CollectionItem name={collection.name} gifts={collection.gifts} />
              <div className="mt-0">
                <button
                  onClick={() => handleEditCollection(collection.id)}
                  className="m-2"
                >
                  Edit
                </button>
                <button
                  onClick={() => handleDeleteCollection(collection.id)}
                  className="m-2 text-red-500"
                >
                  Delete
                </button>
              </div>
            </div>
          ))}
        </div>

        {/* Modal for editing/adding collection */}
        {showEditForm && (
          <div className="fixed top-0 left-0 w-full h-full bg-black bg-opacity-50 flex items-center justify-center">
            <div className="bg-white p-4">
              <EditForm
                collection={collections.find((c) => c.id === editCollectionId)!}
                onSave={handleSaveCollection}
                onClose={handleCloseEditForm}
              />
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default CollectionsPage;
