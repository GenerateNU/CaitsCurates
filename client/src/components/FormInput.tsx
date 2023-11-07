import React from "react";

export type FormInputProps = React.InputHTMLAttributes<HTMLInputElement> & {
    label: string;
    className?: string;
    type: string;
};

const FormInput = ({label, className, type, ...rest}: FormInputProps) => {

    const toggleShowPassword = () => {
        const input = document.getElementById(label) as HTMLInputElement;
        input.type = input.type === "password" ? "text" : "password";
    }

    return (
        <div className={`flex flex-col ${className}`}>
            <label htmlFor={label}> {label} </label>
            <div>
                <input
                    id={label}
                    className="border-solid border-2 rounded-md w-full"
                    type={type}
                    {...rest}
                />
                {type === "password" && 
                    <button
                        className="absolute ml-4 text-sm"
                        type="button"
                        onClick={toggleShowPassword}> 
                        {"Toggle"}
                    </button>
                }
            </div>
        </div>
    );
};

export default FormInput;
