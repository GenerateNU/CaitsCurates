import React, { ReactNode, useState } from 'react';

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  children: ReactNode;
}

const RequestModal: React.FC<ModalProps> = ({ isOpen, onClose, children }) => {
    const [gifteeName, setGifteeName] = useState("");
    const []

    
    const overlayStyles = isOpen
        ? 'fixed top-0 left-0 w-full h-full bg-black bg-opacity-50 flex justify-center items-center z-50'
        : 'hidden';

    const contentStyles = isOpen
        ? 'bg-white px-60 py-40 rounded-md shadow-md'
        : 'hidden';

    return (
        <div className={overlayStyles} onClick={onClose}>
        <div className={contentStyles} onClick={(e) => e.stopPropagation()}>
            {children}
            <button
            className="bg-blue-500 text-white px-2 py-4 rounded-md mt-4"
            onClick={onClose}
            >
            Close
            </button>
        </div>
        </div>
    );
};

export default RequestModal;