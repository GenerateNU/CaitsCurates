import { useState } from 'react';
import CollectionItem from '../components/CollectionItem';
import EditForm from '../components/CollectionForm';

const CollectionsPage = () => {
  const [collections, setCollections] = useState([
    {
      id: 1,
      name: 'Birthday Gifts',
      gifts: ['Toy car', 'Art supplies', 'Book', 'Candy'],
    },
    {
      id: 2,
      name: 'Christmas Gifts',
      gifts: ['Sweater', 'Toys', 'Cookies'],
    },
    {
        id: 3,
        name: 'Birthday Gifts',
        gifts: ['Toy car', 'Art supplies', 'Book', 'Candy'],
      },
      {
        id: 4,
        name: 'Christmas Gifts',
        gifts: ['Sweater', 'Toys', 'Cookies'],
      },
      {
        id: 5,
        name: 'Birthday Gifts',
        gifts: ['Toy car', 'Art supplies', 'Book', 'Candy'],
      },
      {
        id: 6,
        name: 'Christmas Gifts',
        gifts: ['Sweater', 'Toys', 'Cookies'],
      },
      {
        id: 7,
        name: 'Birthday Gifts',
        gifts: ['Toy car', 'Art supplies', 'Book', 'Candy'],
      },
      {
        id: 8,
        name: 'Christmas Gifts',
        gifts: ['Sweater', 'Toys', 'Cookies'],
      },
      {
        id: 9,
        name: 'Birthday Gifts',
        gifts: ['Toy car', 'Art supplies', 'Book', 'Candy'],
      },
      {
        id: 10,
        name: 'Christmas Gifts',
        gifts: ['Sweater', 'Toys', 'Cookies'],
      },
      {
        id: 11,
        name: 'Birthday Gifts',
        gifts: ['Toy car', 'Art supplies', 'Book', 'Candy'],
      },
      {
        id: 12,
        name: 'Christmas Gifts',
        gifts: ['Sweater', 'Toys', 'Cookies'],
      },
      {
        id: 13,
        name: 'Birthday Gifts',
        gifts: ['Toy car', 'Art supplies', 'Book', 'Candy'],
      },
      {
        id: 14,
        name: 'Christmas Gifts',
        gifts: ['Sweater', 'Toys', 'Cookies'],
      },
      
  ]);

  const [showEditForm, setShowEditForm] = useState(false);
  const [editCollectionId, setEditCollectionId] = useState(null);

  const handleEditCollection = (collectionId) => {
    setEditCollectionId(collectionId);
    setShowEditForm(true);
  };

  const handleSaveCollection = (updatedCollection) => {
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

  const handleDeleteCollection = (collectionId) => {
    setCollections((prevCollections) => prevCollections.filter((collection) => collection.id !== collectionId));
  };

  return (
    <div className="min-h-screen flex items-center justify-center">
      <div className="app" style={{ overflowX: "auto"}}>
        <div className="flex">
          {collections.map((collection) => (
            <div key={collection.id} className="m-4 flex-shrink-0 ">
              <CollectionItem name={collection.name} gifts={collection.gifts} />
              <div className="mt-2">
                <button onClick={() => handleEditCollection(collection.id)} className="m-2">
                  Edit
                </button>
                <button onClick={() => handleDeleteCollection(collection.id)} className="m-2 text-red-500">
                  Delete
                </button>
              </div>
            </div>
          ))}
        </div>
      </div>
      {showEditForm && (
        <div className="items-center justify-center">
          <EditForm
            collection={collections.find((c) => c.id === editCollectionId)}
            onSave={handleSaveCollection}
            onClose={handleCloseEditForm}
          />
        </div>
      )}
    </div>
  );
};

export default CollectionsPage;
