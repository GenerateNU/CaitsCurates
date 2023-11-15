// server.js
const cors = require('cors');
const routes = require('./routes');

const app = express();
app.use(express.json());
app.use(cors());

app.use('/api', routes); // Use a specific route path, e.g., /api/create-checkout-session

const PORT = 5000;
app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
