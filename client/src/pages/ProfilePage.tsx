import  { useState } from 'react';
import RequestPurchasingPage from "./RequestPurchasingPage.tsx";
import GiftRequestsPage from "./GiftRequestsPage.tsx";
import AccountSideBar from "../components/AccountSideBar.tsx";
import Navbar from "../components/Home/NavBarUpdated.tsx";


const ProfilePage = () => {
    const [activeComponent, setActiveComponent] = useState('purchaseGifts');

    return (
        <div style={{ display: 'flex', flexDirection: 'column', height: '100vh' }}>
            <Navbar />

            <div style={{ display: 'flex', flex: 1 }}>
                <div>
                    <AccountSideBar activeComponent={activeComponent} setActiveComponent={setActiveComponent} />
                </div>
                <div style={{ flex: 1 }}>
                    {activeComponent === 'purchaseGifts' && <RequestPurchasingPage />}
                    {activeComponent === 'requestHistory' && <GiftRequestsPage />}
                </div>
            </div>
        </div>
    );}

    export default ProfilePage;