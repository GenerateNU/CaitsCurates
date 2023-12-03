import React from "react";

type Props = {
    titleText: string;
    subText: string;
    buttonText: string;
    onClick: () => void;
    image?: string;
    imageStyle?: string;
    textStyle?: string;
    className?: string;
}
const HomePrompt = ({ 
    titleText,
    subText,
    buttonText,
    onClick,
    image,
    imageStyle,
    textStyle,
    className
}: Props) => {
    return (
        <div className={`flex flex-col ${className}`}>
            {image && <img className={`w-[22rem] h-[22rem] mb-12 ${imageStyle}`} src={image} />}
            <div className={`text-4xl font-seasons font-bold ${textStyle}`}> {titleText} </div>
            <div className="text-base font-proxima mt-4"> {subText} </div>
            <button className="btn-primary mt-12" onClick={onClick}> {buttonText} </button>
        </div>
    )
}


export default HomePrompt;
