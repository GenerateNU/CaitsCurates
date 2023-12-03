import LoginForm from "../components/Login/LoginForm.tsx";
import { LoginInputs } from "./LoginPage";
import { signUpFormData } from "../data/formData";
import Navbar from "../components/Home/NavBarUpdated.tsx";


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
            <Navbar />
            <div className="bg-eggshell flex justify-center pb-10">
                <LoginForm
                    formData={signUpFormData}
                    initialState={initialState}
                    onSubmit={handleOnSubmit}
                />
            </div>
        </div>
    );
}

export default SignUpPage;
