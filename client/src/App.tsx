
import {
  BrowserRouter as Router,
  Route,
  Routes,
} from 'react-router-dom';
import GiftManagementPage from './pages/GiftManagementPage';
import HomePage from './pages/HomePage';

function App() {
  return (
  <Router>
    <Routes>
      <Route path = '/' element={<HomePage/>} />
      <Route path = '/giftManagement' element={<GiftManagementPage/>} />
    </Routes>
  </Router>
  );
}

export default App;