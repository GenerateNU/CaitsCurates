import React, { ReactNode, useState } from 'react';
import CaitPhoto from '../../images/modal_cait.svg';
import ModalScreenTwo from './ModalScreenTwo.tsx';

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  children: ReactNode;
}

const ModalScreenOne: React.FC<ModalProps> = ({ isOpen, onClose, children }) => {
  const [activeButton, setActiveButton] = useState("");
  const [showModalTwo, setShowModalTwo] = useState(false);

  const handleButtonClick = (buttonName: any) => {
    setActiveButton(buttonName);
    setShowModalTwo(true);
  };

  const buttonStyle = "px-4 py-2 text-FFF9F4 text-s bg-273F2A rounded-md";
  const activeButtonStyle = "bg-273F2A";

  const overlayStyles = isOpen
    ? 'fixed top-0 left-0 w-full h-full bg-black bg-opacity-50 flex justify-center items-center z-50'
    : 'hidden';

  const contentStyles = isOpen
    ? 'bg-FFF9F4 rounded-md shadow-md flex px-50 items-center'
    : 'hidden';

  return (
    <div className={overlayStyles} onClick={onClose}>
      <div className={contentStyles} onClick={(e) => e.stopPropagation()}>
        {showModalTwo ? (
          <ModalScreenTwo isOpen={isOpen} onClose={onClose} children={children} />
        ) : (
          <>
            <img
              src={CaitPhoto}
              alt="caits-logo.svg"
              className="ml-0"
              style={{ height: "100%", width: "auto" }}
            />
            <div className="ml-40 mr-40">
              <div style={{ maxWidth: "455px" }}  >
                <h1 className="text-black mb-10 text-3xl font-seasons">Let Cait Do it For You</h1>
                <h1 className="text-black mb-10 text-xl font-seasons">Submit a Gift Request to receive personalized gift recommendations for your friends and loved ones. Simply tell me a little about
                your giftee and I’ll take it from there!</h1>
                <h1 className="text-black mb-10 text-xl font-seasons">You currently have <strong>0 Gift Requests.</strong></h1>
              </div>
              <button
                className={`${buttonStyle} ${activeButton === "Purchase" ? activeButtonStyle : ""}`}
                onClick={() => handleButtonClick("Purchase")}
                style={{ width: "455px", height: "50px" }}
              >
                Purchase Gift Requests
              </button>
            </div>
          </>
        )}
      </div>
    </div>
  );
};

export default ModalScreenOne;
