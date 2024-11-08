import React, { useState } from "react";
import { Link } from "react-router-dom";
import FormInput, { FormInputProps } from "./FormInput.tsx";
import { LoginInputs } from "../../pages/LoginPage.tsx";
import GoogleIcon from "../../images/google.svg";

export type FormData = {
    title: string;
    subText: string;
    inputs: FormInputProps[];
    additionalText: string;
    linkedText: string;
    link: string;
    buttonText: string;
}

type Props<T> = {
    formData: FormData;
    initialState: T;
    onSubmit: (formData: T) => void;
}

const LoginForm = <T extends LoginInputs>({formData, initialState, onSubmit}: Props<T>) => {
    const [formState, setFormState] = useState(initialState);

    const handleOnChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setFormState(prevState => ({
            ...prevState,
            [e.target.name]: e.target.value
        }));
    }

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        onSubmit(formState);
    }

    const lineStyle = "border-solid border-b-2 border-black w-1/2";

    return (
        <form
        className="w-[512px] shrink-0 flex flex-col px-14 pt-8 pr-28 lg:w-1/2"
        onSubmit={handleSubmit}
    >
        <div className="text-4xl mt-4 font-seasons font-bold"> {formData.title} </div>
        <div className="text-base text-espresso mt-4 font-proxima"> {formData.subText} </div>
        {formData.inputs.map((props, index) => {
            const [field, value] = Object.entries(formState)[index];
            return (
                <FormInput
                    key={field}
                    name={field}
                    value={value}
                    required
                    onChange={handleOnChange}
                    {...props}
                />
            )
        }
        )}
        <div className="mt-5 text-base text-espresso font-proxima">
            <span> {formData.additionalText} </span>
            <Link 
                className="underline"
                to={formData.link}
            > 
                {formData.linkedText} 
            </Link>
        </div>
        <button
            type="submit"
            className="btn-primary mt-5"
        >
            {formData.buttonText}
        </button>
        <div className="flex items-center mt-8">
            <span className={`${lineStyle}`}/>
            <span className="px-3 text-center text-base font-proxima"> {"or"} </span>
            <span className={`${lineStyle}`} />
        </div>
        <div className="flex justify-center space-x-12 mt-4">
            <div className="w-16 h-16 rounded-full bg-petalpink flex items-center justify-center">
                <img src={GoogleIcon} className="w-8 h-8" />
            </div>
        </div>
    </form>
    )
}

export default LoginForm;
