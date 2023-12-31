import React, {useEffect, useState} from 'react';
import OrPicture from '../../images/or_modal.svg';
import ModalScreenOne from './ModalScreenOne';
import ModalScreenThree from './ModalScreenThree';
import ModalX from '../../images/modal_x_out.svg';
import axios from "axios";
import {Giftee} from "../../types.tsx";
import ModalScreenFour from "./ModalScreenFour.tsx";


interface ModalScreenTwoProps {
  isOpen: boolean;
  onClose: () => void;
  children: React.ReactNode;
}

const ModalScreenTwo: React.FC<ModalScreenTwoProps> = ({ isOpen, onClose, children }) => {
    const [showModalFour, setShowModalFour] = useState(false);
    const [showModalThree, setShowModalThree] = useState(false);
    const [showModalOne, setShowModalOne] = useState(false);
    const [giftees, setGiftees] = useState<Giftee[]>([]);
    const [selectedGiftee, setSelectedGiftee] = useState<Giftee>()


    const overlayStyles = isOpen
        ? 'fixed top-0 left-0 w-full h-full text-black flex justify-center items-center z-50'
        : 'hidden';

    const contentStyles = isOpen
        ? 'bg-FFF9F4 rounded-md text-black shadow-md flex flex-col'
        : 'hidden';
    const getGiftees = async () => {
        try {
            const response = await axios.get(`/api/customer/1`)
            return response.data.Giftees

        } catch (error) {
            console.error('An error occured while getting giftees a gift:', error)
            return
        }
    }
    useEffect(() => {
        const fetchData = async () => {
            try {
                const result = await getGiftees(); // Replace with your actual function
                setGiftees(result);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        }
        fetchData();
    }, []);

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
        {showModalFour ? (
                <ModalScreenFour isOpen={isOpen} onClose={onClose} children={children} gifteeID={selectedGiftee?.ID} />
            ) : (
            <>
      <div className="flex flex-row ml-10 items-center mt-10">
      <button className=''
              onClick={() => handleBackClick()}
            >
              Back
            </button>
            <div className="ml-10 mr-10">
            <img
              src={ModalX}
              alt="caits-logo.svg"
              onClick={onClose}
              style={{ height: "20px", width: "20px", marginLeft: "880px"}}
            />
            </div>
      </div>
        <div className="ml-40 mr-40 text-center text-espresso">
          <div>
            <h1 className="text-black mb-10 text-3xl font-seasons">Who are you gifting for?</h1>
            <p className="text-black mb-10 text-xl font-seasons"></p>
          </div>
            <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100%', marginBottom: '8rem' }}>
                {/* Left side - My Giftees and Buttons */}
                <div style={{ flex: 1 }}>
                    <h1 style={{ color: 'espresso', marginRight:  '10rem', fontSize: '1.25rem', marginBottom: '1rem', fontFamily: 'Seasons, sans-serif', textAlign: 'left'}}>My Giftees</h1>
                    <div className="mr-5 flex flex-wrap gap-4">
                        {giftees.map((giftee, index) => (
                            <button
                                style={{
                                    padding: '8px 16px',
                                    color: '#FFF9F4',
                                    backgroundColor: '#273F2A',
                                    borderRadius: '0.375rem',
                                    textAlign: 'left',
                                    display: 'block',
                                    marginBottom: '1rem'
                                }}
                                key={index}
                                onClick={() => { setSelectedGiftee(giftee); setShowModalFour(true); }}
                            >
                                {giftee.GifteeName}
                            </button>
                        ))}
                    </div>
                </div>

                {/* Center - Image */}
                <img
                    src={OrPicture}
                    alt="caits-logo"
                    style={{ maxHeight: '100%', maxWidth: '100%', display: 'block' }}
                />

                {/* Right side - Add New Giftee Button */}
                <div style={{ flex: 1, display: 'flex', justifyContent: 'flex-end'}}>
                    <button
                        onClick={() => handleAddGifteeClick()}
                        style={{
                            padding: '8px 16px',
                            color: '#FFF9F4',
                            backgroundColor: '#273F2A',
                            borderRadius: '0.375rem',
                            textAlign: 'center',
                        }}
                    >
                        Add New Giftee
                    </button>
                </div>
            </div>
        </div>
         </>)}
            </>)}
            </>)}
      </div>
    </div>
  );
};
export default ModalScreenTwo;
