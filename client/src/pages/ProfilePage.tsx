import  { useState } from 'react';
import RequestPurchasingPage from "../components/Requests/RequestPurchasing.tsx";
import GiftRequests from "../components/Requests/GiftRequests.tsx";
import AccountSideBar from "../components/AccountSideBar.tsx";
import Navbar from "../components/Home/NavBarUpdated.tsx";


const ProfilePage = () => {
    const [activeComponent, setActiveComponent] = useState('purchaseGifts');

    return (
        <div style={{ display: 'flex', flexDirection: 'column', height: '100vh' }}>
            <div className="sticky top-0 z-10 bg-eggshell">
            <Navbar />
            </div>

            <div style={{ display: 'flex', flex: 1 }}>
                <div>
                    <AccountSideBar activeComponent={activeComponent} setActiveComponent={setActiveComponent} />
                </div>
                <div style={{ flex: 1 }}>
                    {activeComponent === 'purchaseGifts' && <RequestPurchasingPage />}
                    {activeComponent === 'requestHistory' && <GiftRequests />}
                </div>
            </div>
        </div>
    );}

    export default ProfilePage;