
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import HomePage from "./pages/HomePage";
import RequestsPage from "./pages/RequestsPage";
import CollectionPage from './pages/CollectionsPage';

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<HomePage />} />
                <Route path="/requests/" element={<RequestsPage />} />
                <Route path="/collections/" element={<CollectionPage />} />
            </Routes>
        </Router>
    );
}

export default App;