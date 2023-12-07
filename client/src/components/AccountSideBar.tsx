import {  useState} from 'react';
type AccountSideBarProps = {
    activeComponent: string;
    setActiveComponent: React.Dispatch<React.SetStateAction<string>>;
};

const AccountSideBar: React.FC<AccountSideBarProps> = ({ activeComponent, setActiveComponent }) => {
  const [isGiftingOpen, setIsGiftingOpen] = useState(false);
  const [isBudgetOpen, setIsBudgetOpen] = useState(false);
  const [isCalendarOpen, setIsCalendarOpen] = useState(false);
  const toggleGifting = () => {
    setIsGiftingOpen(!isGiftingOpen);
  };

  const toggleBudget = () => {
    setIsBudgetOpen(!isBudgetOpen);
  };

  const toggleCalendar = () => {
    setIsCalendarOpen(!isCalendarOpen);
  };

  return (
    <div className="flex flex-col bg-FBF2EB" style={{ width: '366px', height: "100vh"}}>
      <div style={{ fontFamily: 'The Seasons', fontSize: '36px', marginLeft: '40px', marginTop: "50px" }}>
        Account
        <div className="flex flex-row items-center" style={{ fontFamily: 'The Seasons', fontSize: '24px', marginTop: '50px' }}>
          <svg width="48" height="48" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
            <circle cx="24" cy="24" r="24" fill="#DFB2AA"/>
            <path d="M24 16C25.0609 16 26.0783 16.4214 26.8284 17.1716C27.5786 17.9217 28 18.9391 28 20C28 21.0609 27.5786 22.0783 26.8284 22.8284C26.0783 23.5786 25.0609 24 24 24C22.9391 24 21.9217 23.5786 21.1716 22.8284C20.4214 22.0783 20 21.0609 20 20C20 18.9391 20.4214 17.9217 21.1716 17.1716C21.9217 16.4214 22.9391 16 24 16ZM24 26C28.42 26 32 27.79 32 30V32H16V30C16 27.79 19.58 26 24 26Z" fill="#32211E"/>
          </svg>
          <span style={{ marginLeft: '18px' }}>John Doe</span>
        </div>
        <div className="flex flex-row items-center" style={{ fontFamily: 'The Seasons', fontSize: '24px', marginLeft: '20px', marginTop: '40px' }}>
          <svg width="21" height="20" viewBox="0 0 21 20" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M18.5801 4.60547H16.4001C16.5101 4.29547 16.5801 3.95547 16.5801 3.60547C16.5801 1.94547 15.2401 0.605469 13.5801 0.605469C12.5301 0.605469 11.6201 1.14547 11.0801 1.95547L10.5801 2.62547L10.0801 1.94547C9.54008 1.14547 8.63008 0.605469 7.58008 0.605469C5.92008 0.605469 4.58008 1.94547 4.58008 3.60547C4.58008 3.95547 4.65008 4.29547 4.76008 4.60547H2.58008C1.47008 4.60547 0.590078 5.49547 0.590078 6.60547L0.580078 17.6055C0.580078 18.7155 1.47008 19.6055 2.58008 19.6055H18.5801C19.6901 19.6055 20.5801 18.7155 20.5801 17.6055V6.60547C20.5801 5.49547 19.6901 4.60547 18.5801 4.60547ZM13.5801 2.60547C14.1301 2.60547 14.5801 3.05547 14.5801 3.60547C14.5801 4.15547 14.1301 4.60547 13.5801 4.60547C13.0301 4.60547 12.5801 4.15547 12.5801 3.60547C12.5801 3.05547 13.0301 2.60547 13.5801 2.60547ZM7.58008 2.60547C8.13008 2.60547 8.58008 3.05547 8.58008 3.60547C8.58008 4.15547 8.13008 4.60547 7.58008 4.60547C7.03008 4.60547 6.58008 4.15547 6.58008 3.60547C6.58008 3.05547 7.03008 2.60547 7.58008 2.60547ZM18.5801 17.6055H2.58008V15.6055H18.5801V17.6055ZM18.5801 12.6055H2.58008V6.60547H7.66008L5.58008 9.43547L7.20008 10.6055L9.58008 7.36547L10.5801 6.00547L11.5801 7.36547L13.9601 10.6055L15.5801 9.43547L13.5001 6.60547H18.5801V12.6055Z" fill="#A65A5A"/>
          </svg>
          <button
            style={{
              background: 'none',
              border: 'none',
              cursor: 'pointer',
              fontFamily: 'The Seasons',
              fontSize: '24px',
              marginLeft: '20px',
            }}
            onClick={toggleGifting}
          >
            Gifting
          </button>
        </div>
        {isGiftingOpen && (
          <div
            style={{
              background: '#FBF2EB',
              padding: '10px',
              fontSize: '16px',
            }}
          >
            <div style={{marginTop: "10px"}}>
            <button
                style={{
                    background: activeComponent == "purchaseGifts" ? '#DFB2AA' : 'none',
                    width: "265px",
                    borderLeft: activeComponent == "purchaseGifts" ? '6px solid #A65A5A' : 'none',
                    fontWeight: activeComponent == "purchaseGifts" ? 'bold' : 'normal',
                 }}
                onClick={() => setActiveComponent("purchaseGifts")} >
                Purchase Gift Requests
          </button>
                </div>
            <div style={{marginTop: "10px"}}><button
                style={{
                    background: activeComponent == "requestHistory" ? '#DFB2AA' : 'none',
                    width: "265px",
                    borderLeft: activeComponent == "requestHistory" ? '6px solid #A65A5A' : 'none',
                    fontWeight: activeComponent == "requestHistory" ? 'bold' : 'normal',
                    paddingRight: '20px',
                 }}
                onClick={() => setActiveComponent("requestHistory")} >
                Gift Request History
          </button></div>
            <div style={{marginTop: "10px"}}>
            <button
                style={{
                    background: activeComponent == "giftees" ? '#DFB2AA' : 'none',
                    width: "265px",
                    borderLeft: activeComponent == "giftees" ? '6px solid #A65A5A' : 'none',
                    fontWeight: activeComponent == "giftees" ? 'bold' : 'normal',
                    paddingRight: '105px',
                 }}
                onClick={() => setActiveComponent("giftees")} >
                Giftees
          </button>
            </div>
          </div>
        )}
        <div className="flex flex-row items-center" style={{ fontFamily: 'The Seasons', fontSize: '24px', marginLeft: '20px', marginTop: '40px' }}>
          <svg width="11" height="19" viewBox="0 0 11 19" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M5.78444 8.33889C3.38833 7.71611 2.61778 7.07222 2.61778 6.06944C2.61778 4.91889 3.68389 4.11667 5.46778 4.11667C7.34667 4.11667 8.04333 5.01389 8.10667 6.33333H10.4394C10.3656 4.51778 9.25722 2.85 7.05111 2.31167V0H3.88444V2.28C1.83667 2.72333 0.19 4.05333 0.19 6.09056C0.19 8.52889 2.20611 9.74278 5.15111 10.45C7.79 11.0833 8.31778 12.0122 8.31778 12.9939C8.31778 13.7222 7.80056 14.8833 5.46778 14.8833C3.29333 14.8833 2.43833 13.9122 2.32222 12.6667H0C0.126667 14.9783 1.85778 16.2767 3.88444 16.7094V19H7.05111V16.7306C9.10944 16.34 10.7456 15.1472 10.7456 12.9833C10.7456 9.98556 8.18056 8.96167 5.78444 8.33889Z" fill="#32211E"/>
          </svg>
          <button
            style={{
              background: 'none',
              border: 'none',
              cursor: 'pointer',
              fontFamily: 'The Seasons',
              fontSize: '24px',
              marginLeft: '30px',
            }}
            onClick={toggleBudget}
          >
            Budgets
          </button>
        </div>
        {isBudgetOpen && (
          <div
            style={{
              background: '#FBF2EB',
              padding: '10px',
              fontSize: '16px',
            }}
          >
            {/*Add dropdown types here*/}
          </div>
        )}
        <div className="flex flex-row items-center" style={{ fontFamily: 'The Seasons', fontSize: '24px', marginLeft: '20px', marginTop: '40px' }}>
          <svg width="18" height="19" viewBox="0 0 18 19" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M15.5455 1.72727H14.6818V0H12.9545V1.72727H4.31818V0H2.59091V1.72727H1.72727C0.777273 1.72727 0 2.50455 0 3.45455V17.2727C0 18.2227 0.777273 19 1.72727 19H15.5455C16.4955 19 17.2727 18.2227 17.2727 17.2727V3.45455C17.2727 2.50455 16.4955 1.72727 15.5455 1.72727ZM15.5455 17.2727H1.72727V6.04545H15.5455V17.2727Z" fill="#32211E"/>
          </svg>
          <button
            style={{
              background: 'none',
              border: 'none',
              cursor: 'pointer',
              fontFamily: 'The Seasons',
              fontSize: '24px',
              marginLeft: '20px',
            }}
            onClick={toggleCalendar}
          >
            Calendar
          </button>
        </div>
        {isCalendarOpen && (
          <div
            style={{
              background: '#FBF2EB',
              padding: '10px',
              fontSize: '16px',
            }}
          >
            {/*Add dropdown types here*/}
          </div>
        )}
      </div>
    </div>
  );
};

export default AccountSideBar;