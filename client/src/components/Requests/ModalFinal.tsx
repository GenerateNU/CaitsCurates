import React, { ReactNode, useState } from 'react';
import CaitPhoto from '../../images/modal_cait.svg';

interface ModalProps {
    isOpen: boolean;
    onClose: () => void;
    children: ReactNode;
    date: Date;
}

const ModalScreenFinal: React.FC<ModalProps> = ({ isOpen, onClose, children , date}) => {


    const overlayStyles = isOpen
        ? 'fixed top-0 left-0 w-full h-full bg-black bg-opacity-50 flex justify-center items-center z-50'
        : 'hidden';

    const contentStyles = isOpen
        ? 'bg-FFF9F4 rounded-md shadow-md flex px-50 items-center'
        : 'hidden';

    return (
        <div className={overlayStyles} onClick={onClose}>
            <div className={contentStyles} onClick={(e) => e.stopPropagation()}>
                <img
                    src={CaitPhoto}
                    alt="caits-logo.svg"
                    className="ml-0"
                    style={{ height: "100%", width: "auto" }}
                />
                <div className="ml-40 mr-40">
                    <div style={{ maxWidth: "455px" }}  >
                        <h1 className="text-black mb-10 text-3xl font-seasons">Thank You</h1>
                        <h1 className="text-black mb-10 text-xl font-seasons">Your Gift Request has been submitted! I look forward to finding the perfect gift for your giftee. You will receive your suggestions by {date.toDateString()} In the meantime, feel free to browse the site and don’t hesitate to reach out to me through the Contact form if you have any questions, comments, or concerns</h1>
                        <h1 className="text-black mb-10 text-xl font-bold">Happy Gifting!</h1>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default ModalScreenFinal;