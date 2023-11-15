// routes.js
const express = require('express');
const stripe = require('stripe')('sk_test_51OAJv1LvSQCInHXR5XGQY3rUMqQwZsU3RfSkWLE0uW7LC4eDJMSKnco0TbgoMjVdF7PmA5eTNAFRNi7DHvBpLUm400RX29t0NM');
const router = express.Router();

router.post('/create-checkout-session', async (req: any, res: any) => {
  try {
    const session = await stripe.checkout.sessions.create({
      payment_method_types: ['card'],
      line_items: req.body.items,
      mode: 'payment',
      success_url: 'http://localhost:5173/success',
      cancel_url: 'http://localhost:5173/cancel',
    });

    res.json({ url: session.url });
  } catch (error) {
    console.error(error);
    res.status(500).json({ error: 'Internal server error' });
  }
});

module.exports = router;
