import Navbar from "../components/Admin/Navbar.tsx";
import RequestCard from "../components/Admin/RequestCard.tsx";
import {useEffect, useState} from "react";
import Select, {ActionMeta} from "react-select";
import {GiftRequest} from "../types.tsx";
import axios from "axios";

export type SelectValueType = {
  [key: string]: string
  | null
};

const options = [
  { value: "0", label: "Incomplete requests" },
  { value: "1", label: "Complete requests" },
];

export default function RequestsPage() {
  const [selectedOption, setSelectedOption] = useState<SelectValueType>({
    value: "0",
    label: "Incomplete requests",
  });
  const [requests, setRequests] = useState<GiftRequest[]>([]);

  useEffect(() => {
    let url = '/api/requests'; // Base URL to fetch all requests

    if (selectedOption.value === '0') {
      url += '/incomplete';
    } else if (selectedOption.value === '1') {
      url += '/complete';
    }
    const fetchRequests = async () => {
      try {
        const response = await axios.get(url);
        setRequests(response.data);
      } catch (error) {
        console.error('An error occurred while fetching the requests:', error);
      }
    };
    fetchRequests();
  }, [selectedOption]);
  const handleOption = (selection: SelectValueType) => {
    setSelectedOption(selection);
  };

  return (
    <div>
      <Navbar />
      <div className="flex flex-col px-48 py-8">
        <h2 className="font-bold text-2xl mb-2">Manage Requests</h2>
        <p>Filter gift requests using the dropdown below. </p>
        <Select
          defaultValue={selectedOption}
          onChange={handleOption}
          options={options}
          className="mt-2"
        />
        {(selectedOption?.value == "0") && (
          <div className="mt-6">
            <h2 className="font-bold text-xl text-blue-800">
              Incomplete requests
            </h2>
            {requests.map((req) => {
              return <RequestCard key={req.ID} {...req} />;
            })}
          </div>
        )}
        {(selectedOption?.value == "1") && (
          <div className="mt-6">
            <h2 className="font-bold text-xl text-blue-800">
              Complete requests
            </h2>
            {requests.map((req) => {
              return <RequestCard key={req.ID} {...req} />;            })}
          </div>
        )}
      </div>
    </div>
  );
}
