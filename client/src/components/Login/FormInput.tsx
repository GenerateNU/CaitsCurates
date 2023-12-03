import React, { useState } from "react";
import ValidationAlert from "./ValidationAlert";
import { passwordValidation } from "../../data/formData";
import EyeOff from "../../images/eye_off.svg";
import EyeOn from "../../images/eye_on.svg";

export type FormInputProps = React.InputHTMLAttributes<HTMLInputElement> & {
    label: string;
    className?: string;
    type: string;
    includeValidation?: boolean;
};

const FormInput = ({label, className, type, value, includeValidation = false, ...rest}: FormInputProps) => {
    const [icon, setIcon] = useState(EyeOn);

    const toggleShowPassword = () => {
        const input = document.getElementById(label) as HTMLInputElement;
        input.type = input.type === "password" ? "text" : "password";
        setIcon(icon === EyeOn ? EyeOff : EyeOn);
    }

    return (
        <div className={`flex flex-col ${className}`}>
            <label htmlFor={label} className="font-seasons font-bold font-base"> {label} </label>
            <div>
                <input
                    id={label}
                    className="border-brown border-solid border-[1.5px] bg-pearl rounded-md w-full text-base text-coffee font-proxima px-2 p-1"
                    type={type}
                    {...rest}
                />
                {type === "password" &&
                <>
                    <img
                        className="inline absolute ml-4 mt-2 w-5 h-5"
                        src={icon}
                        onClick={toggleShowPassword}
                    />
                    {includeValidation && value !== "" && passwordValidation.map((validation) => (
                        <ValidationAlert
                        key={validation.validationText}
                        /* will always be a string from form state*/
                        fieldText={value as string}
                        validationText={validation.validationText}
                        validationRule={validation.validationRule}
                        className="mt-1"
                    />
                    ))}
                </>
                }
            </div>
        </div>
    );
};

export default FormInput;
