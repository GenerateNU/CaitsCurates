import { useEffect, useState } from "react";
import axios from "axios";
import { GiftRequest } from "../../types";

type GiftRequestsProps = {
  selectRow: React.Dispatch<React.SetStateAction<GiftRequest | null>>;
};

const GiftRequestsTable: React.FC<GiftRequestsProps> = ({ selectRow }) => {
  const [requests, setRequests] = useState<GiftRequest[]>([]);

  const customerID = 1;

  const handleClick = (req: GiftRequest) => {
    console.log(req);
    selectRow(req);
  };

  useEffect(() => {
    getRequests();
  }, []);

  const getRequests = async () => {
    try {
      const response = await axios.get(`/api/requests/${customerID}`);
      setRequests(response.data);
      console.log(response);
    } catch (error) {
      console.error("An error occurred while fetching the gifts:", error);
    }
  };

  const formatDate = (dateString: string) => {
    var date = new Date(dateString);

    return new Intl.DateTimeFormat("en-US", {
      year: "numeric",
      month: "2-digit",
      day: "2-digit",
    }).format(date);
  };


  return (
    <div>
      <h1 className="mt-2 font-seasons text-3xl text-espresso">Gifting</h1>

      <h2 className="mt-6 font-seasons text-2xl text-deeppink">
        Gift Request History
      </h2>

      <table className="mt-4 font-proxima bg-eggshell border border-gray-300 text-winered">
        <thead>
          <tr>
            <th className="py-2 px-4 font-semibold text-left ">Order Number</th>
            <th className="py-2 px-4 font-semibold text-left ">Name</th>
            <th className="py-2 px-4 font-semibold text-left">
              Date Submitted
            </th>
            <th className="py-2 px-4 font-semibold text-left">Status</th>
          </tr>
        </thead>
        <tbody className="text-espresso">
          {requests.map((item) => (
            <tr
              key={item.ID}
              className={`cursor-pointer  ${
                item.GiftResponseID != null ? "bg-linen" : "bg-eggshell"
              }`}
              onClick={() => handleClick(item)}
            >
              <td className="py-2 px-4 border-b border-gray-300">{item.ID}</td>
              <td className="py-2 px-4 border-b border-gray-300">
                {item.Giftee.GifteeName}
              </td>
              <td className="py-2 px-4 border-b border-gray-300">
                {formatDate(item.CreatedAt)}
              </td>
              <td className="py-2 px-4 border-b border-gray-300">
                {item.GiftResponseID != null ? "Complete" : "In Progress"}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default GiftRequestsTable;
