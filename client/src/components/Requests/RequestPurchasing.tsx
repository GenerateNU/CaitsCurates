import {useEffect, useState} from 'react';
import PurchaseOneCredit from "../PurchaseOneCredit.tsx";
import PurchaseThreeCredits from "../PurchaseThreeCredits.tsx";
import PurchaseFiveCredits from "../PurchaseFiveCredits.tsx";
import axios from "axios";
import RequestPurchaseSuccess from "../../pages/RequestPurchaseSuccess.tsx";
const RequestsPurchasing = () => {
  const [isLoading, setLoading] = useState(false);
  const [checkoutSuccessful, setCheckoutSuccessful] = useState(false);
  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    if (params.get('checkout') === 'success') {
      setCheckoutSuccessful(true)
    }
  }, []);
  const handleCheckout = async (NumGifts: number) => {
    setLoading(true);
    try {
      // Use axios to send a POST request
      const response = await axios.post('/api/create-checkout-session');
      window.location = response.data.url;
    } catch (error) {
      console.error('Error during checkout:', error);
    } finally {
      setLoading(false);
    }
  }
  const buttonStyle = "px-4 py-2 text-FFF9F4 text-s bg-273F2A rounded-md";
  return (
      <div>
        {checkoutSuccessful ? (
            <RequestPurchaseSuccess />
        ) : (
    <div className="flex flex-col h-screen bg-FFF9F4">
        <div className="flex flex-col items-center flex-grow text-4xl" style={{ fontFamily: 'The Seasons', marginLeft:"20px" }}>
          <div>
            <div style={{ textAlign: 'left' }}>
              <div style={{ marginBottom: '100px', marginTop: "100px" }}>Gifting</div>
              <div className="text-A65A5A" style={{ marginBottom: '40px' }}>Purchase Gift Requests</div>
            </div>
            <div className="flex flex-row space-x-4">
              <PurchaseOneCredit />
              <PurchaseThreeCredits />
              <PurchaseFiveCredits />
            </div>
            <div className="flex flex-row space-x-4" style={{ marginTop: "30px", fontFamily: "Proxima Nova", fontSize: "18px" }}>
              <button
                className={`${buttonStyle} Purchase`}
                disabled={isLoading}
                onClick={() => handleCheckout(1)}
                style={{ width: "300px", height: "50px" }}
              >
                Purchase
              </button>
              <button
                  className={`${buttonStyle} Purchase`}
                  disabled={isLoading}
                onClick={() => handleCheckout(3)}
                style={{ width: "300px", height: "50px" }}
              >
                Purchase
              </button>
              <button
                  className={`${buttonStyle} Purchase`}
                  disabled={isLoading}
                onClick={() => handleCheckout(5)}
                style={{ width: "300px", height: "50px" }}
              >
                Purchase
              </button>
            </div>
          </div>
        </div>
      </div>)}
      </div>
  );
};

export default RequestsPurchasing;
