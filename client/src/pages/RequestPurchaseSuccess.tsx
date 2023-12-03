
import AccountSideBar from "../components/AccountSideBar";
import SampleOrderSummary from "../components/SampleOrderSummary";
import Navbar from "../components/Home/NavBarUpdated";

const RequestPurchaseSuccess = () => {

    const submitRequestButtonStyle = "px-4 py-2 text-FFF9F4 text-s bg-273F2A rounded-md";
    const startShoppingButtonStyle = "px-4 py-2 text-273F2A text-s bg-FFF9F4 rounded-md";

    return (
        <div className="flex flex-col h-screen bg-FFF9F4">
        <Navbar />
  
        <div className="flex flex-row h-screen">
          <AccountSideBar />
          <div className="flex flex-col items-center flex-grow text-4xl" style={{ fontFamily: 'The Seasons', marginLeft:"20px" }}>
          <div>
            <div style={{ textAlign: 'left' }}>
              <div style={{ marginBottom: '100px', marginTop: "100px" }}>Gifting</div>
              <div className="text-A65A5A" style={{ marginBottom: '40px' }}>Thank You!</div>
              <div style={{width: '835px', fontSize: '24px', marginBottom: '40px'}}> 
                We have received your purchase, and you will receive an order confirmation email shortly to 
                <span style={{ color: '#BC2C36' }}> johndoe@gmail.com</span>. 
                In the meantime, submit a request to Cait, or explore her curated gift selections for any occasion!</div>
            </div>
            <div className="flex flex-row space-x-2">
            <button
                className={`${submitRequestButtonStyle}`}
                style={{ width: "412px", height: "50px", fontSize: "18px" }}
              >
                Submit a Request
              </button>
              <button
                className={`${startShoppingButtonStyle}`}
                style={{ width: "412px", height: "50px", fontSize: "18px", border: '2px solid #273F2A', }}
              >
                Start Shopping
              </button>
              </div>
              <div style={{marginTop: '50px'}}>
              <SampleOrderSummary />
              </div>
            </div>
            </div>
          </div>
        </div>
      );
};

export default RequestPurchaseSuccess;
