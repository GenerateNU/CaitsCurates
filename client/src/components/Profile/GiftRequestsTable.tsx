import { useEffect, useState } from "react";
import { styled } from "@mui/material/styles";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import axios from "axios";

const StyledTableRow = styled(TableRow)(() => ({
  "&:last-child td, &:last-child th": {
    border: 0,
  },
}));

interface GiftResponse {
  id: number;
  giftCollectionID: number;
  customMessage: string;
}

interface GiftRequest {
  id: number;
  giftResponseID: number | null;
  recipientName: string;
  giftResponse: GiftResponse | null;
  dateSubmitted: string;
}

export default function GiftRequestsTable() {
  const [selectedRequest, setSelectedRequest] = useState<GiftRequest | null>(
    null
  );
  const [requests, setRequests] = useState<GiftRequest[]>([]);

  const customerID = 1;

  const handleClick = (req: GiftRequest) => {
    console.log(req.recipientName);
    setSelectedRequest(req);
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

  return (
    <div>
      <div>
        selected component:
        <p>{selectedRequest?.recipientName}</p>
      </div>
      <TableContainer className="bg-eggshell cursor-pointer">
        <Table aria-label="customized table">
          <TableHead>
            <TableRow>
              <TableCell>Order Number</TableCell>
              <TableCell>Name</TableCell>
              <TableCell>Date Submitted</TableCell>
              <TableCell>Status</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {requests.map((row) => (
              <StyledTableRow
                key={row.id}
                onClick={() => handleClick(row)}
                className={`${
                  row.giftResponseID === null ? "bg-eggshell" : "bg-linen"
                }`}
              >
                <TableCell component="th" scope="row">
                  {row.id}
                </TableCell>
                <TableCell>{row.recipientName}</TableCell>
                <TableCell>{row.dateSubmitted}</TableCell>
                <TableCell>
                  {row.giftResponseID ? "Complete" : "In Progress"}
                </TableCell>
              </StyledTableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </div>
  );
}
