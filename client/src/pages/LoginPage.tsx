import LoginForm from "../components/LoginForm";
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
        <div>
            {/* add navbar */}
            <div className="flex">
                <LoginForm
                    formData={loginFormData}
                    initialState={initialState}
                    onSubmit={handleOnSubmit}
                />
                {/* replace with image */}
                <div className="w-[512px] bg-slate-700 lg:w-1/2" />
            </div>
        </div>
    );
};

export default LoginPage;
