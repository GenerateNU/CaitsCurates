import { Elements } from '@stripe/react-stripe-js';
import { loadStripe } from '@stripe/stripe-js';
import {
  useStripe,
  useElements,
  CardElement
} from '@stripe/react-stripe-js';

const stripePromise = loadStripe('pk_test_51OAJv1LvSQCInHXR4xvIIOKSghGMZ52FONHQXJefxwlF7E7BqC2aFdoyHZYOh9PEcgxj5vyUlbwofAnRgkgaBurR00tdvcArW1');



const handleSubmit = async (event: any) => {

  const stripe = useStripe();
  const elements = useElements();

  event.preventDefault();
  if (!stripe || !elements) {
    // Stripe.js has not yet loaded.
    // Make sure to disable form submission until Stripe.js has loaded.
    return;
  }

  const cardElement = elements.getElement(CardElement);

  const { token, error } = await stripe.createToken(cardElement);

  if (error) {
    console.error(error);
  } else {
    // Send the token to your server for processing
    handlePayment(token);
  }
};

const handlePayment = async (token) => {
  try {
    const response = await fetch('http://localhost:5000/process-payment', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        token: token.id, // Pass the token to your server
        amount: 1000, // Adjust the amount according to your requirements
      }),
    });

    if (response.ok) {
      console.log('Payment successful!');
      // Handle success on the client side if needed
    } else {
      console.error('Payment failed!');
      // Handle failure on the client side if needed
    }
  } catch (error) {
    console.error('Error processing payment:', error);
  }
};

function PaymentForm() {
  return (
    <Elements stripe={stripePromise}>
      <form onSubmit={handleSubmit}>
        <div style={{ width: '500px' }}>
          <label>
            Card number
            <CardElement
              options={{
                style: {
                  base: {
                    fontSize: '16px',
                  },
                },
              }}
            />
          </label>
        </div>
        <button type="submit">Pay</button>
      </form>
    </Elements>
  );
}

export default PaymentForm;
