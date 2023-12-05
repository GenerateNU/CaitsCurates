import React, {useState} from 'react';
import ModalScreenTwo from './ModalScreenTwo';
import ModalScreenFour from './ModalScreenFour';
import ModalX from '../../images/modal_x_out.svg'
import TextBox from '../Admin/TextBox'
import AgeDropdown from '../Admin/AgeDropdown';
import GenderDropdown from '../Admin/GenderDropdown';
import ColorDropdown from '../Admin/ColorDropdown';

interface ModalScreenThreeProps {
  isOpen: boolean;
  onClose: () => void;
  children: React.ReactNode;
}

const ModalScreenThree: React.FC<ModalScreenThreeProps> = ({ isOpen, onClose, children }) => {
    const [nameNickname, setNameNickname] = useState<string>('');

    const handleAgeChange = (selectedAge: number) => {
        console.log(`Selected age: ${selectedAge}`);

      };

      const handleGenderChange = (selectedGender: String) => {
        console.log(`Selected Gender: ${selectedGender}`);

      };

      const handleColorChange = (selectedColors: String[]) => {
        console.log(`Selected Colors: ${selectedColors}`);

      };

      const handleNextClick = () => {
        setNameNickname(nameNickname);
        setShowModalFour(true);
      };

  const [showModalTwo, setShowModalTwo] = useState(false);
  const [showModalFour, setShowModalFour] = useState(false);

  const buttonStyle = "px-4 py-2 ml-40 text-FFF9F4 text-s bg-273F2A rounded-md";
  const activeButtonStyle = "bg-273F2A";

  const overlayStyles = isOpen
    ? 'fixed top-0 left-0 w-full h-full flex justify-center items-center z-50'
    : 'hidden';

  const contentStyles = isOpen
    ? 'bg-FFF9F4 rounded-md shadow-md flex flex-col px-50 items-center'
    : 'hidden';

    const handleBackClick = (buttonName: any) => {
        setShowModalTwo(true);
      };

    return (
        <div className={overlayStyles} onClick={onClose}>
          <div className={contentStyles} onClick={(e) => e.stopPropagation()}>
          {showModalTwo ? (
              <ModalScreenTwo isOpen={isOpen} onClose={onClose} children={children} />
            ) : (
                <> 
          <div className="flex ml-10 items-center mt-10">
          <button className=''
                  onClick={() => handleBackClick("Back")}
                >
                  Back
                </button>
                <div className="mr-0 mr-auto ml-80">
                <img
                  src={ModalX}
                  alt="caits-logo.svg"
                  onClick={onClose}
                  style={{ height: "20px", width: "20px", marginLeft: "880px", marginRight: "20px"}}
                />
                </div>
                </div>
                <div>
                <h1 className="mb-10 text-3xl text-center items-center w-full font-seasons">Add New Giftee</h1>
               
                <div className= "flex flex-row">
                <TextBox 
                placeholder="Name or Nickname"
               
              
                />
                <ColorDropdown onChange={handleColorChange}/>
                </div>
                <div className= "flex flex-row ">
                 <TextBox placeholder="Relationship to You" />
                 <TextBox placeholder="Type to Add Interests" />
                </div>
                <div className='flew-col'>
                    <AgeDropdown onChange={handleAgeChange}/>
            
                </div>
                <div >
                <GenderDropdown onChange={handleGenderChange}/>
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
          </div>
        </div>
      );
    };

export default ModalScreenThree;
