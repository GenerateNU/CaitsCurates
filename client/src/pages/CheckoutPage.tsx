import {FormEvent, useState} from "react";
import axios from "axios";


const CheckoutPage = () => {
    const [isLoading, setLoading] = useState(false);

    const handleCheckout = async (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setLoading(true);

        try {
            // Use axios to send a POST request
            const response = await axios.post('/api/create-checkout-session');

            window.location = response.data.url;

        } catch (error) {
            console.error('Error during checkout:', error);
        } finally {
            setLoading(false);
        }
    }
    return (
        <section>
            <div className="product">
                <img src="https://i.imgur.com/EHyR2nP.png" alt="The cover of Stubborn Attachments" />
                <div className="description">
                    <h3>Stubborn Attachments</h3>
                    <h5>$20.00</h5>
                </div>
            </div>
            <form onSubmit={handleCheckout}>
                <button type="submit" id="checkout-button" disabled={isLoading}>
                    {isLoading ? 'Processing...' : 'Checkout'}
                </button>
            </form>
        </section>
    );
}

export default CheckoutPage;
