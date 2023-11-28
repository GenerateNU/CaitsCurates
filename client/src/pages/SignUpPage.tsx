import LoginForm from "../components/Login/LoginForm.tsx";
import { LoginInputs } from "./LoginPage";
import { signUpFormData } from "../data/formData";


type SignUpInputs = LoginInputs & {
    FirstName: string;
    LastName: string;
}

const SignUpPage = () => {
    const initialState: SignUpInputs = {
        FirstName: "",
        LastName: "",
        Email: "",
        Password: ""
    };

    const handleOnSubmit = (formData: SignUpInputs) => {
        console.log(formData);
        // call endpoint with form data
    }

    return (
        <div>
            {/* add navbar */}
            <div className="flex">
                <LoginForm
                    formData={signUpFormData}
                    initialState={initialState}
                    onSubmit={handleOnSubmit}
                />
                {/* replace with image */}
                <div className="w-[512px] bg-slate-700 lg:w-1/2" />
            </div>
        </div>
    );
}

export default SignUpPage;
