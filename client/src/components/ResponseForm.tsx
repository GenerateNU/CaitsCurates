import CollectionSelector from "./CollectionSelector";

const ResponseForm = () => {
    return (
        <div className="flex flex-col justify-between h-full mt-4">
            <div>
                <label className="block text-gray-700 text-sm font-bold mb-2">
                    Select a gift collection:
                    <CollectionSelector />
                </label>
                <label className="block text-gray-700 text-sm font-bold mb-2">
                    Custom message:
                    <input
                        className="border rounded w-full py-2 px-3 text-gray-700"
                        name="message"
                        type="text"
                    />
                </label>
            </div>
            <button className="bg-blue-600 ml-8 px-4 py-2 h-10 text-white rounded-md self-end">
                Submit
            </button>
        </div>
    );
};

export default ResponseForm;