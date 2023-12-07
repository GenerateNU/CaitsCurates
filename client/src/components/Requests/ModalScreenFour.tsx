// ModalScreenFour.tsx
import React, { useState } from 'react';
import ModalScreenThree from './ModalScreenThree.tsx';
import ModalScreenFive from './ModalScreenFive.tsx';
import ModalX from '../../images/modal_x_out.svg';
import OccasionDropdown from '../Admin/OccasionDropdown.tsx';
import DateTextBox from '../Admin/DateTextBox.tsx';
import BudgetSlider from '../Admin/BudgetSlider.tsx';
import {GiftRequestProps} from '../../types.tsx'

interface ModalScreenFourProps {
  isOpen: boolean;
  onClose: () => void;
  children: React.ReactNode;
  gifteeID: number;
}

const ModalScreenFour: React.FC<ModalScreenFourProps> = ({ isOpen, onClose, children, gifteeID}) => {
  const [showModalThree, setShowModalThree] = useState(false);
  const [showModalFive, setShowModalFive] = useState(false);
  const [giftRequest, setGiftRequest] = useState<GiftRequestProps>({
    BudgetMax: 0,
    BudgetMin: 0,
    Comment: "",
    CustomerID: 1,
    DateNeeded: new Date(),
    GifteeID: 0,
    Occasion: []
  });

  const handleOccasionChange = (selectedOccasion: string) => {
    setGiftRequest((prevOccasion) => ({ ...prevOccasion, Occasion: [selectedOccasion]}));
  }

  const handleDateChange = (selectedDate: string) => {
    let date = new Date(selectedDate)
    setGiftRequest((prevDate) => ({ ...prevDate, DateNeeded: date}));
  }

  const handleBudgetMaxChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selectedBudgetMax = Math.max(0, e.target.valueAsNumber);
    setGiftRequest((prevBudgetMax) => ({ ...prevBudgetMax, BudgetMax: selectedBudgetMax}));
  }

  const handleBudgetMinChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selectedBudgetMin = Math.max(0, e.target.valueAsNumber);
    setGiftRequest((prevBudgetMin) => ({ ...prevBudgetMin, BudgetMin: selectedBudgetMin}));
  }




  const handleNextClick = () => {
    console.log('Giftee id:', gifteeID);

    setShowModalFive(true);
  };

  const handleBackClick = () => {
    setShowModalThree(true);
  };

  const overlayStyles = isOpen
      ? 'fixed top-0 left-0 w-full h-full flex justify-center items-center z-50'
      : 'hidden';

  const contentStyles = isOpen
      ? 'bg-FFF9F4 rounded-md shadow-md flex flex-col px-50 items-center'
      : 'hidden';

  const buttonStyle = "px-4 py-2 ml-40 text-FFF9F4 text-s bg-273F2A rounded-md";
  const activeButtonStyle = "bg-273F2A";

  return (
      <div className={overlayStyles} onClick={onClose}>
        <div className={contentStyles} onClick={(e) => e.stopPropagation()}>
          {showModalThree ? (
              <ModalScreenThree isOpen={isOpen} onClose={onClose} children={children} />
          ) : (
              <>
                {showModalFive ? (
                    <ModalScreenFive isOpen={isOpen} onClose={onClose} children={children} gifteeID={gifteeID}
                                     giftRequest={giftRequest}  />
                ) : (
                    <>
                      <div className="flex ml-10 items-center mt-10">
                        <button className=''
                                onClick={() => handleBackClick()}
                        >
                          Back
                        </button>
                        <div className="mr-auto ml-80">
                          <img
                              src={ModalX}
                              alt="caits-logo.svg"
                              onClick={onClose}
                              style={{ height: "20px", width: "20px", marginLeft: "880px", marginRight: "20px"}}
                          />
                        </div>
                      </div>
                      <div>
                        <h1 className="font-bold mb-10 text-3xl text-center items-center w-full font-seasons">Giftee Preferences</h1>

                        <div className= "flex flex-row">
                          <div className="flex flex-col">
                            <h2 className='text-2xl text-A65A5A font-seasons ml-4 font-bold leading-7'>Type of Occasion</h2>
                            <OccasionDropdown
                                onChange={handleOccasionChange}
                            />

                            <h2 className='text-2xl text-A65A5A font-seasons ml-4 mt-6 font-bold leading-7'>Price Range</h2>
                            <BudgetSlider maxBudget={giftRequest.BudgetMax} minBudget={giftRequest.BudgetMin}/>
                            <div className='flex flex-row justify-between mt-3 ml-4'>
                              <div>
                                <span className='text-2xl font-bold leading-7 text-B79D94 font-seasons mr-1'>$</span>
                                <input
                                    type="number"
                                    style={{
                                      color: '#B79D94',
                                      height: '42px',
                                      width: '84px',
                                      border: '1.5px solid #B79D94',
                                      borderRadius: '5px',
                                      borderWidth: '1.5px',
                                      textAlign: 'center',
                                    }}
                                    className="w-full px-4 py-2 rounded-m pl-2"
                                    value={giftRequest.BudgetMin}
                                    placeholder='0'
                                    onChange={handleBudgetMinChange}
                                />
                              </div>
                              <span className='text-base'>to</span>
                              <div>
                                <span className='text-2xl font-bold leading-7 text-B79D94 font-seasons mr-1'>$</span>
                                <input
                                    type="number"
                                    style={{
                                      color: '#B79D94',
                                      height: '42px',
                                      width: '84px',
                                      border: '1.5px solid #B79D94',
                                      borderRadius: '5px',
                                      borderWidth: '1.5px',
                                      textAlign: 'center',
                                    }}
                                    className="w-full px-4 py-2 rounded-m pl-2"
                                    value={giftRequest.BudgetMax}
                                    placeholder='0'
                                    onChange={handleBudgetMaxChange}
                                />
                              </div>
                            </div>
                          </div>
                          <div className='flex flex-col ml-28'>
                            <h2 className='text-2xl text-A65A5A font-seasons ml-4 font-bold leading-7'>Deadline</h2>
                            <h3 className='text-base text-32211E font-proxima leading-5 mt-3 ml-4'>When do you need suggestions by?</h3>
                            <p className='text-8D7A73 mt-1.5 text-base leading-5 ml-4'><b>Note:</b> I cannot guarantee gift suggestions within <br/> one week of needed date.</p>
                            <div className='ml-4 mt-3'>
                              <DateTextBox
                                  onChange={handleDateChange}
                                  width='359px'
                              />
                            </div>
                          </div>
                        </div>
                        <div className='mb-10 flex justify-end'>
                          <button
                              className={`${buttonStyle} ${activeButtonStyle}`}
                              onClick={() => handleNextClick()}
                              style={{ width: "170px", height: "50px"}}
                          >
                            Next
                          </button>
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

export default ModalScreenFour;