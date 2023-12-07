import { FormData } from "../components/Login/LoginForm.tsx";
import { ValidationAlertProps } from "../components/Login/ValidationAlert.tsx"

export const loginFormData: FormData = {
    title: "Log In",
    subText: "Log in to continue your gift giving journey!",
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
    subText: "Create an account to begin your gift-giving journey!",
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
            type: "password" ,
            includeValidation: true
        }
    ],
    additionalText: "Already have an account? ",
    linkedText: "Log in",
    link: "/login",
    buttonText: "Sign Up"
};

export const passwordValidation: Omit<ValidationAlertProps, "fieldText">[] = [
    {
        validationText: "Contains uppercase letters",
        validationRule: new RegExp("[A-Z]")
    },
    {
        validationText: "Contains lowercase letters",
        validationRule: new RegExp("[a-z]")
    },
    {
        validationText: "Contains 8 characters",
        validationRule: new RegExp(".{8,}")
    },
    {
        validationText: "Contains a number or symbol",
        validationRule: new RegExp("\\d|\\W")
    }

];