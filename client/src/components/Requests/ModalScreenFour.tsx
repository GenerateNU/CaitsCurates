// ModalScreenFour.tsx
import React from 'react';
import {Giftee} from '../../types.tsx'

interface ModalScreenFourProps {
  isOpen: boolean;
  onClose: () => void;
  children: React.ReactNode;
  giftee: Giftee;
}

const ModalScreenFour: React.FC<ModalScreenFourProps> = ({ isOpen, onClose, children, giftee}) => {
  const overlayStyles = isOpen
    ? 'fixed top-0 left-0 w-full h-full flex justify-center items-center z-50'
    : 'hidden';

  const contentStyles = isOpen
    ? 'bg-FFF9F4 rounded-md shadow-md flex flex-col px-50 items-center'
    : 'hidden';

  return (
    <div className={overlayStyles} onClick={onClose}>
      <div className={contentStyles} onClick={(e) => e.stopPropagation()}>
        <div className="flex ml-10 items-center mt-10">
          {/* Add any necessary elements for your header */}
        </div>
        <div>
          <h1 className="mb-10 text-3xl text-center items-center w-full font-seasons">Modal Four Title</h1>
          <p></p>
          {/* Add additional elements to display other information */}
        </div>
      </div>
    </div>
  );
};

export default ModalScreenFour;
