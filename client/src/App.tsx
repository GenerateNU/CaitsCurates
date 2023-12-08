import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import HomePage from "./pages/HomePage";
import RequestsPage from "./pages/RequestsPage.tsx";
import CheckoutPage from "./pages/CheckoutPage.tsx";
import CollectionPage from "./pages/CollectionsPage";
import { AdminProvider } from "./Context/AdminContext.tsx";
import GiftManagementPage from "./pages/GiftManagementPage.tsx";
import LoginPage from "./pages/LoginPage.tsx";
import SignUpPage from "./pages/SignUpPage.tsx";
import RequestPurchaseSuccess from "./pages/RequestPurchaseSuccess.tsx";
import ReturnHomePage from "./pages/ReturnHomePage.tsx";
import ProfilePage from "./pages/ProfilePage.tsx";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route
          path="/requests/"
          element={
            <AdminProvider>
              <RequestsPage />
            </AdminProvider>
          }
        />
        <Route
          path="/collections/"
          element={
            <AdminProvider>
              <CollectionPage />
            </AdminProvider>
          }
        />
        <Route
          path="/checkout/"
          element={
            <AdminProvider>
              <CheckoutPage />
            </AdminProvider>
          }
        />
        <Route
          path="/gifts/"
          element={
            <AdminProvider>
              <GiftManagementPage />
            </AdminProvider>
          }
        />
        <Route path="/about/" element={<ReturnHomePage />} />
        <Route path="/signup/" element={<SignUpPage />} />
        <Route path="/login/" element={<LoginPage />} />
        <Route path="/success/" element={<RequestPurchaseSuccess />} />
        <Route path="/profile/" element={<ProfilePage />} />
      </Routes>
    </Router>
  );
}

export default App;
