import { Elements } from '@stripe/react-stripe-js';
import { loadStripe } from '@stripe/stripe-js';
import {
  useStripe,
  useElements,
  CardElement
} from '@stripe/react-stripe-js';

const stripePromise = loadStripe('pk_test_51OAJv1LvSQCInHXR4xvIIOKSghGMZ52FONHQXJefxwlF7E7BqC2aFdoyHZYOh9PEcgxj5vyUlbwofAnRgkgaBurR00tdvcArW1');



const handleSubmit = async (event: any) => {
  event.preventDefault();
  
}

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
