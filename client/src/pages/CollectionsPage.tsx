import {useEffect, useState} from "react";
import CollectionItem from "../components/Admin/CollectionItem.tsx";
import EditForm from "../components/Admin/CollectionForm.tsx";
import Navbar from "../components/Admin/Navbar.tsx";
import {GiftCollection} from "../types.tsx";
import {useAdmin} from "../Context/AdminContext.tsx";
import axios from "axios";





const CollectionsPage = () => {
  const { collections, fetchGiftCollections } = useAdmin();
  const { gifts, fetchGifts } = useAdmin();
  const [showEditForm, setShowEditForm] = useState<boolean>(false);
  const [editCollectionId, setEditCollectionId] = useState<number | null>(null);

  useEffect(() => {
    fetchGiftCollections();
    fetchGifts();
  }, []);
  const handleCreateCollection =  () => {

    // Add the new collection to the state
    setEditCollectionId(0);
    setShowEditForm(true);
  };



  const handleEditCollection = (collectionId: number) => {
    setEditCollectionId(collectionId);
    setShowEditForm(true);
    fetchGiftCollections();
  };

  const handleSaveCollection = async (updatedCollection: GiftCollection) => {
    if (updatedCollection.ID != undefined) {
      await axios.put("/api/updateGiftCollection", updatedCollection);
    } else {
      await axios.post("/api/addGiftCollection", updatedCollection);
    }
    console.log(updatedCollection)
    setShowEditForm(false);
    fetchGiftCollections();
  };

  const handleCloseEditForm = () => {
    setEditCollectionId(null);
    setShowEditForm(false);
    fetchGiftCollections();
  };

  const handleDeleteCollection = async (collectionId: number) => {
    await axios.delete(`/api/deleteGiftCollection/${collectionId}`);
    fetchGiftCollections();
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
            <div key={collection.ID} className="m-2">
              <CollectionItem name={collection.CollectionName} gifts={collection.Gifts} />
              <div className="mt-0">
                <button
                  onClick={() => handleEditCollection(collection.ID)}
                  className="m-2"
                >
                  Edit
                </button>
                <button
                  onClick={() => handleDeleteCollection(collection.ID)}
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
                collection={collections.find((c) => c.ID === editCollectionId)!}
                allGifts={gifts}
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
