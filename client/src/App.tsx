
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import HomePage from "./pages/HomePage";
import RequestsPage from "./pages/RequestsPage";
import CollectionPage from './pages/CollectionsPage';
import TestPage from "./pages/TestPage.tsx";
import {AdminProvider} from "./Context/AdminContext.tsx";

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
                    path="/test/"
                    element={
                        <AdminProvider>
                            <TestPage/>
                        </AdminProvider>}
                />
            </Routes>
        </Router>
    );
}

export default App;