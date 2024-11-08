import Navbar from "../components/Home/NavBarUpdated.tsx";
import LoginForm from "../components/Login/LoginForm.tsx";
import { loginFormData } from "../data/formData";

export type LoginInputs = {
    Email: string;
    Password: string;
};

const LoginPage = () => {
    const initialState: LoginInputs = {
        Email: "",
        Password: ""
    }

    const handleOnSubmit = (formData: LoginInputs) => {
        console.log(formData);
        // call endpoint with form data
    }

    return (
        <div className="bg-eggshell h-screen">
            <Navbar />
            <div className="flex justify-center">
                <LoginForm
                    formData={loginFormData}
                    initialState={initialState}
                    onSubmit={handleOnSubmit}
                />
            </div>
        </div>
    );
};

export default LoginPage;
