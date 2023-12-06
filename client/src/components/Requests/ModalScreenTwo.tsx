import React, {useState} from 'react';
import OrPicture from '../../images/or_modal.svg';
import ModalScreenOne from './ModalScreenOne';
import ModalScreenThree from './ModalScreenThree';
import ModalX from '../../images/modal_x_out.svg'

interface ModalScreenTwoProps {
  isOpen: boolean;
  onClose: () => void;
  children: React.ReactNode;
}

const ModalScreenTwo: React.FC<ModalScreenTwoProps> = ({ isOpen, onClose, children }) => {

  const [showModalThree, setShowModalThree] = useState(false);
  const [showModalOne, setShowModalOne] = useState(false);

  const buttonStyle = "px-4 py-2 ml-40 text-FFF9F4 text-s bg-273F2A rounded-md";
  const activeButtonStyle = "bg-273F2A";

  const overlayStyles = isOpen
    ? 'fixed top-0 left-0 w-full h-full text-black flex justify-center items-center z-50'
    : 'hidden';

  const contentStyles = isOpen
    ? 'bg-FFF9F4 rounded-md text-black shadow-md flex flex-col'
    : 'hidden';

    const handleAddGifteeClick = () => {
        setShowModalThree(true);
        setShowModalOne(false);
      };

    const handleBackClick = () => {
        setShowModalOne(true);
        setShowModalThree(false);
      };

  return (
    <div className={overlayStyles} onClick={onClose}>
      <div className={contentStyles} onClick={(e) => e.stopPropagation()}>
      {showModalOne ? (
          <ModalScreenOne isOpen={isOpen} onClose={onClose} children={children} />
        ) : (
            <>
            {showModalThree ? (
          <ModalScreenThree isOpen={isOpen} onClose={onClose} children={children} />
        ) : (
            <>
      <div className="flex flex-row ml-10 items-center mt-10">
      <button className=''
              onClick={() => handleBackClick()}
            >
              Back
            </button>
            <div className="mr-0 mr-auto ml-80">
            <img
              src={ModalX}
              alt="caits-logo.svg"
              onClick={onClose}
              style={{ height: "20px", width: "20px", marginLeft: "880px"}}
            />
            </div>
            </div>
        <div className="ml-40 mr-40 text-center text-black">
        
          <div>
            <h1 className="text-black mb-10 text-3xl font-seasons">Who are you gifting for?</h1>
            <p className="text-black mb-10 text-xl font-seasons"></p>
          </div>
          <div className="flex items-center ml-0">
          <div>
          <h1 className="text-black mr-80 mb-60 text-xl font-seasons">My Giftees</h1>
          </div>
            <img
              src={OrPicture}
              alt="caits-logo.svg"
              style={{ height: "100%", width: "auto" }}
            />
            <div className="mr-0">
            <button
              className={`${buttonStyle} ${activeButtonStyle}`}
              onClick={() => handleAddGifteeClick()}
              style={{ width: "265px", height: "69px" }}
            >
              Add New Giftee
            </button>
            </div>
          </div>
        </div>
        </>
        )}
         </>
        )}
      </div>
    </div>
  );
};

export default ModalScreenTwo;
