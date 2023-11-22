import Navbar from "../components/NavBarUpdated";
import {useState} from 'react';
import PurchaseOneCredit from "../components/PurchaseOneCredit";
import PurchaseThreeCredits from "../components/PurchaseThreeCredits";
import PurchaseFiveCredits from "../components/PurchaseFiveCredits";

const RequestsPurchasingPage = () => {

    const [activeButton, setActiveButton] = useState("");

    const handleButtonClick = (buttonName: any) => {
      setActiveButton(buttonName);
    };
  
    const buttonStyle = "px-4 py-2 text-FFF9F4 text-s bg-273F2A rounded-md";
    const activeButtonStyle = "bg-273F2A";

  return (
    <div className="flex flex-col min-h-screen bg-FFF9F4">
      <Navbar />
      <div className="flex flex-col items-center justify-center flex-grow text-3xl" style={{ fontFamily: 'The Seasons'}}>
        <div>
        <div  style={{ textAlign: 'left'}}>
        <div style={{ marginBottom: '100px'}}>Gifting</div>
        <div className="text-A65A5A" style={{ marginBottom: '40px' }}>Purchase Gift Requests</div>
        </div>
        <div className="flex flex-row space-x-4">
          <PurchaseOneCredit />
          <PurchaseThreeCredits />
          <PurchaseFiveCredits />
        </div>
        <div className="flex flex-row space-x-4" style={{ marginTop: "30px", fontFamily: "Proxima Nova", fontSize: "14px"}}>
                <button
                className={`${buttonStyle} ${activeButton === "Purchase" ? activeButtonStyle : ""}`}
                onClick={() => handleButtonClick("Purchase")}
                style={{ width: "265px", height: "50px" }}
                >
                Purchase
                </button>
                <button
                className={`${buttonStyle} ${activeButton === "Purchase" ? activeButtonStyle : ""}`}
                onClick={() => handleButtonClick("Purchase")}
                style={{ width: "265px", height: "50px" }}
                >
                Purchase
                </button>
                <button
                className={`${buttonStyle} ${activeButton === "Purchase" ? activeButtonStyle : ""}`}
                onClick={() => handleButtonClick("Purchase")}
                style={{ width: "265px", height: "50px" }}
                >
                Purchase
                </button>
        </div>
        </div>
      </div>
    </div>
  );
};

export default RequestsPurchasingPage;
