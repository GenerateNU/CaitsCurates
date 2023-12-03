import React, { ReactNode } from 'react';

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  children: ReactNode;
}

const RequestModal: React.FC<ModalProps> = ({ isOpen, onClose, children }) => {
  const overlayStyles = isOpen
    ? 'fixed top-0 left-0 w-full h-full bg-black bg-opacity-50 flex justify-center items-center z-50'
    : 'hidden';

  const contentStyles = isOpen
    ? 'bg-white p-8 rounded-md shadow-md'
    : 'hidden';

  return (
    <div className={overlayStyles} onClick={onClose}>
      <div className={contentStyles} onClick={(e) => e.stopPropagation()}>
        {children}
        <button
          className="bg-blue-500 text-white px-4 py-2 rounded-md mt-4"
          onClick={onClose}
        >
          Close
        </button>
      </div>
    </div>
  );
};

export default RequestModal;