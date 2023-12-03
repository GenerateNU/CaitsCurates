
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import HomePage from "./pages/HomePage";
import RequestsPage from "./pages/RequestsPage";
import CollectionPage from './pages/CollectionsPage';
import {AdminProvider} from "./Context/AdminContext.tsx";
import GiftManagementPage from "./pages/GiftManagementPage.tsx";
import LoginPage  from "./pages/LoginPage.tsx";
import SignUpPage from "./pages/SignUpPage.tsx";
import RequestsPurchasingPage from "./pages/RequestPurchasingPage.tsx";
import RequestPurchaseSuccess from "./pages/RequestPurchaseSuccess.tsx";
import ReturnHomePage from "./pages/ReturnHomePage.tsx";

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<HomePage/>}/>
                <Route
                    path="/requests/"
                    element={
                        <AdminProvider>
                            <RequestsPage/>
                        </AdminProvider>}
                />
                <Route
                    path="/collections/"
                    element={
                        <AdminProvider>
                            <CollectionPage/>
                        </AdminProvider>}
                />
                <Route
                    path="/gifts/"
                    element={
                        <AdminProvider>
                            <GiftManagementPage/>
                        </AdminProvider>}
                />
                <Route
                    path="/home/"
                    element={<ReturnHomePage />}
                />
                <Route
                    path="/signup/"
                    element={<SignUpPage/>}
                />
                <Route
                    path="/purchase-requests/"
                    element={<RequestsPurchasingPage/>}
                />
                <Route
                    path="/login/"
                    element={<LoginPage/>}
                />
                 <Route
                    path="/success/"
                    element={<RequestPurchaseSuccess/>}
                />
            </Routes>
        </Router>
    );
}

export default App;