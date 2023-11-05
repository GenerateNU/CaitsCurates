import { FormData } from "../components/LoginForm";

export const loginFormData: FormData = {
    title: "Log In",
    subText: "Log in to continue your gift giving journey",
    inputs: [
        {
            className: "mt-10",
            label: "Email",
            type: "email"
        },
        {
            className: "mt-4",
            label: "Password",
            type: "password" 
        }
    ],
    additionalText: "Don't have an account? ",
    linkedText: "Sign Up",
    link: "/signup",
    buttonText: "Log In"
};

export const signUpFormData: FormData = {
    title: "Sign Up",
    subText: "Create an account to begin youre gift-giving journey!",
    inputs: [
        {
            className: "mt-10",
            label: "First Name",
            type: "text"
        },
        {
            className: "mt-4",
            label: "Last Name",
            type: "text"
        },
        {
            className: "mt-4",
            label: "Email",
            type: "email"
        },
        {
            className: "mt-4",
            label: "Password",
            type: "password" 
        }
    ],
    additionalText: "Already have an account? ",
    linkedText: "Log in",
    link: "/login",
    buttonText: "Sign Up"
};
