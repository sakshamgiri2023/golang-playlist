
const express = require('express');
const cors = require('cors');
const bodyParser = require('body-parser');

const app = express();
const PORT = 3000;


app.use(cors()); // Enable CORS for all routes
app.use(bodyParser.json()); // Parse JSON bodies


let users = [
  { id: 1, name: 'Alice', email: 'alice@example.com' },
  { id: 2, name: 'Bob', email: 'bob@example.com' },
  { id: 3, name: 'Charlie', email: 'charlie@example.com' }
];

let products = [
  { id: 101, name: 'Laptop', price: 999.99 },
  { id: 102, name: 'Phone', price: 699.99 },
  { id: 103, name: 'Tablet', price: 399.99 }
];


app.get('/', (req, res) => {
  res.send('Welcome to the Node.js/Express API Server!');
});


app.get('/api/users', (req, res) => {
  res.json(users);
});

app.get('/api/users/:id', (req, res) => {
  const user = users.find(u => u.id === parseInt(req.params.id));
  if (!user) return res.status(404).send('User not found');
  res.json(user);
});

app.post('/api/users', (req, res) => {
  const newUser = {
    id: users.length + 1,
    name: req.body.name,
    email: req.body.email
  };
  users.push(newUser);
  res.status(201).json(newUser);
});


app.get('/api/products', (req, res) => {
  res.json(products);
});

app.get('/api/products/:id', (req, res) => {
  const product = products.find(p => p.id === parseInt(req.params.id));
  if (!product) return res.status(404).send('Product not found');
  res.json(product);
});

app.use((req, res) => {
  res.status(404).send('Route not found');
});


app.listen(PORT, () => {
  console.log(`Server running on http://localhost:${PORT}`);
});